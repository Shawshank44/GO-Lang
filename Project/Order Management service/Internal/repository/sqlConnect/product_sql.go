package sqlconnect

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"order_mgt/Internal/models"
	"order_mgt/pkg/storage"
	"order_mgt/pkg/utils"
	"os"
	"strings"
)

func CreateProductInDB(ctx context.Context, product *models.Product, sessionID string) (int64, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, utils.ErrorHandler(err, "unable to start transaction")
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	imgJSON, err := json.Marshal(&product.Images)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Invalid marshalling")
	}

	query := `INSERT INTO products (sku,name,description,category,brand,manufacturer,price,currency,stock,unit,status,created_at,spec_updated_at,inventory_updated_at,updated_by,images) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?);`

	res, err := tx.ExecContext(ctx, query,
		product.SKU,
		product.Name,
		product.Description,
		product.Category,
		product.Brand,
		product.Manufacturer,
		product.Price,
		product.Currency,
		product.Stock,
		product.Unit,
		product.Status,
		product.CreatedAt,
		product.SpecsUpdatedAt,
		product.InventoryUpdatedAt,
		product.UpdatedBy,
		imgJSON,
	)

	if err != nil {
		return 0, utils.ErrorHandler(err, "Unable to create product in database")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, utils.ErrorHandler(err, "unable to fetch last inserted id")
	}

	// 2. Finalize files :
	_, err = tx.ExecContext(ctx, `UPDATE files SET session_id = NULL WHERE session_id = ?`, sessionID)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Unable to finalize uploaded files")
	}

	// 3. Delete files :
	_, err = tx.ExecContext(ctx, `DELETE FROM upload_sessions WHERE id = ?`, sessionID)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Unable to delete upload session")
	}

	// 4. Commit transaction :
	err = tx.Commit()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Unable to commit the transaction")
	}
	return id, nil
}

func GetProductsFromDB(ctx context.Context, r *http.Request, limit, page int) ([]models.Product, int, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, 0, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	query := `SELECT sku, name, description, category, brand, manufacturer, price, currency, stock, unit, status, created_at, spec_updated_at, inventory_updated_at, updated_by, images FROM products WHERE 1=1`
	var args []interface{}
	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	// Pagination :
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, utils.ErrorHandler(err, "Unable to query the product")
	}

	defer rows.Close()

	productList := make([]models.Product, 0)

	for rows.Next() {
		var product models.Product
		var imagesJSON []byte
		err = rows.Scan(
			&product.SKU,
			&product.Name,
			&product.Description,
			&product.Category,
			&product.Brand,
			&product.Manufacturer,
			&product.Price,
			&product.Currency,
			&product.Stock,
			&product.Unit,
			&product.Status,
			&product.CreatedAt,
			&product.SpecsUpdatedAt,
			&product.InventoryUpdatedAt,
			&product.UpdatedBy,
			&imagesJSON,
		)
		if err != nil {
			return nil, 0, utils.ErrorHandler(err, "Unable to find the row")
		}
		if len(imagesJSON) > 0 {
			if err := json.Unmarshal(imagesJSON, &product.Images); err != nil {
				return nil, 0, utils.ErrorHandler(err, "Unable to parse product images")
			}
		}
		productList = append(productList, product)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, utils.ErrorHandler(err, "Unable to read product rows")
	}

	countQuery := `SELECT COUNT(*) FROM products WHERE 1=1`

	var countArgs []interface{}

	countQuery, countArgs = utils.AddFilters(r, countQuery, countArgs)

	var totalProducts int
	err = db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&totalProducts)
	if err != nil {
		return nil, 0, utils.ErrorHandler(err, "Unable to count admins")
	}

	return productList, totalProducts, nil
}

func GetProductFromDB(ctx context.Context, id int) (*models.Product, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	var product models.Product
	var imageJSON []byte

	query := `SELECT id, sku, name, description, category, brand, manufacturer, price, currency, stock, unit, status, created_at, spec_updated_at, inventory_updated_at, updated_by, images FROM products WHERE id = ?`

	err = db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.SKU,
		&product.Name,
		&product.Description,
		&product.Category,
		&product.Brand,
		&product.Manufacturer,
		&product.Price,
		&product.Currency,
		&product.Stock,
		&product.Unit,
		&product.Status,
		&product.CreatedAt,
		&product.SpecsUpdatedAt,
		&product.InventoryUpdatedAt,
		&product.UpdatedBy,
		&imageJSON,
	)

	if err == sql.ErrNoRows {
		return &models.Product{}, utils.ErrorHandler(err, "unable to find the product from db")
	}

	if err != nil {
		return nil, utils.ErrorHandler(err, "unable to find the product from db")
	}

	if len(imageJSON) > 0 {
		if err := json.Unmarshal(imageJSON, &product.Images); err != nil {
			return nil, utils.ErrorHandler(err, "Unable to parse product images")
		}
	}

	return &product, nil
}

func UpdateProductInDB(ctx context.Context, minioService *storage.MinioService, product *models.Product, id int, sessionID string) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Internal server error")
	}

	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return utils.ErrorHandler(err, "Internal tx error")
	}

	defer tx.Rollback()

	var oldImageJSON []byte
	err = tx.QueryRowContext(ctx, `SELECT images FROM products WHERE id=?`, id).Scan(&oldImageJSON)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.ErrorHandler(err, "product not found")
		}
		return utils.ErrorHandler(err, "unable to fetch product")
	}

	// Existing images stored in DB
	var oldImages []models.ProductImage

	if len(oldImageJSON) > 0 {
		if err := json.Unmarshal(oldImageJSON, &oldImages); err != nil {
			return utils.ErrorHandler(err, "Invalid old images")
		}
	}

	// New images coming from request
	newImages := product.Images

	imgContent, err := json.Marshal(newImages)
	if err != nil {
		return utils.ErrorHandler(err, "image serialization failed")
	}

	res, err := tx.ExecContext(
		ctx,
		`UPDATE products
	 SET sku=?,
	     name=?,
	     description=?,
	     category=?,
	     brand=?,
	     manufacturer=?,
	     spec_updated_at=CURRENT_TIMESTAMP,
	     updated_by=?,
	     images=?
	 WHERE id=?`,
		product.SKU,
		product.Name,
		product.Description,
		product.Category,
		product.Brand,
		product.Manufacturer,
		product.UpdatedBy,
		imgContent,
		id,
	)
	if err != nil {
		return utils.ErrorHandler(err, "unable to update product")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "unable to determine updated rows")
	}
	if rows == 0 {
		return utils.ErrorHandler(
			errors.New("product not found"),
			"product not found",
		)
	}

	if sessionID != "" {
		_, err = tx.ExecContext(ctx, `UPDATE files SET session_id = NULL WHERE session_id = ?`, sessionID)
		if err != nil {
			return utils.ErrorHandler(err, "Unable to finalize uploaded files")
		}
		_, err = tx.ExecContext(ctx, `DELETE FROM upload_sessions WHERE id = ?`, sessionID)
		if err != nil {
			return utils.ErrorHandler(err, "Unable to delete upload session")
		}
	}

	// Store all URLs that still exist after the update
	newSet := make(map[string]struct{}, len(newImages))

	for _, img := range newImages {
		if img.URL == "" {
			continue
		}

		newSet[img.URL] = struct{}{}
	}

	var objectsToDelete []string

	// Find old images that were removed
	for _, old := range oldImages {
		objName := strings.TrimPrefix(old.URL, fmt.Sprintf("http://localhost%s/product-images/", os.Getenv("MINIO_PORT")))
		if old.URL == "" {
			continue
		}
		// Image still exists, so don't delete it
		if _, exists := newSet[old.URL]; exists {
			continue
		}

		// Remove file record from DB
		_, err = tx.ExecContext(ctx, `DELETE FROM files WHERE file_path = ?`, objName)
		if err != nil {
			return utils.ErrorHandler(err, "Unable to delete the file")
		}

		objectsToDelete = append(objectsToDelete, objName)
	}

	if err := tx.Commit(); err != nil {
		return utils.ErrorHandler(err, "transaction failed")
	}

	// MinIO is external to the SQL transaction.
	// Delete objects only after the DB transaction commits.
	for _, objName := range objectsToDelete {
		if err := minioService.Delete(ctx, objName); err != nil {
			log.Printf(
				"failed to delete obsolete MinIO object %q: %v",
				objName,
				err,
			)
		}
	}

	return nil
}

func InventoryUpdateInDB(ctx context.Context, inventory *models.Inventory, id int) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "internal server error")
	}

	defer db.Close()

	query := `UPDATE products SET price=?, currency=?, stock=?, unit=?, updated_by=?, inventory_updated_at=CURRENT_TIMESTAMP, status=? WHERE id=?`

	res, err := db.ExecContext(ctx, query, inventory.Price, inventory.Currency, inventory.Stock, inventory.Unit, inventory.UpdatedBy, inventory.Status, id)
	if err != nil {
		return utils.ErrorHandler(err, "unable to update the inventory.")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "unable to determine the updated rows.")
	}

	if rows == 0 {
		return utils.ErrorHandler(errors.New("product not found"), "product not found")
	}

	return nil
}
