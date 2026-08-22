package sqlconnect

import (
	"context"
	"encoding/json"
	"order_mgt/Internal/models"
	"order_mgt/pkg/utils"
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
	_, err = tx.ExecContext(ctx, "UPDATE files SET session_id = NULL WHERE session_id = ?", sessionID)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Unable to finalize uploaded files")
	}

	// 3. Delete files :
	_, err = tx.ExecContext(ctx, "DELETE FROM upload_sessions WHERE id = ?", sessionID)
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
