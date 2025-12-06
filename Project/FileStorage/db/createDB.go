package db

import (
	"bufio"
	"errors"
	datamodel "filestorage/db/dataModel"
	"filestorage/utils"
	"io"
	"log"
	"strconv"
	"strings"

	"fmt"
	"os"

	"github.com/gofrs/flock"
)

type DBWrapper struct {
	datamodel.DB
	datamodel.Record
}

func (db *DBWrapper) ParseLine(line string) (*datamodel.Record, error) {
	parts, err := utils.SplitEscaped(line, db.Delim)
	if err != nil {
		return nil, err
	}

	if len(parts) < 3 {
		return nil, errors.New("malformed record (not enough fields)")
	}
	id, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	if err != nil {
		return nil, err
	}
	name, err := utils.UnescapeField(parts[1], db.Delim)
	if err != nil {
		return nil, err
	}
	email, err := utils.UnescapeField(parts[2], db.Delim)
	if err != nil {
		return nil, err
	}

	return &datamodel.Record{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (db *DBWrapper) formatLine(r *datamodel.Record) string {
	return fmt.Sprintf("%d%s%s%s%s\n", r.ID, string(db.Delim), utils.EscapeField(r.Name, db.Delim), string(db.Delim), utils.EscapeField(r.Email, db.Delim))
}

func (db *DBWrapper) LoadMaxID() error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	file, err := os.Open(db.Path)
	if err != nil {
		if os.IsNotExist(err) {
			db.NextID = 1
			db.LoadMax = true
			return nil
		}
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var max int64 = 0
	lineNo := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		lineNo++
		trim := strings.TrimRight(line, "\n")
		if strings.TrimSpace(trim) != "" {
			rec, perr := db.ParseLine(trim)
			if perr != nil {
				log.Printf("loadMaxID : malformed line %d:%v (preserved)\n", lineNo, perr)
			} else {
				if rec.ID > max {
					max = rec.ID
				}
			}
		}
		if err == io.EOF {
			break
		}
	}
	db.NextID = max + 1
	db.LoadMax = true
	return nil

}

func NewDB(path string, delim string) (*datamodel.DB, error) {
	if delim == "" {
		delim = "|"
	}

	drunes := []rune(delim)

	if len(drunes) != 1 {
		return nil, errors.New("the delimiter character must be a single rune/character")
	}

	db := &DBWrapper{
		DB: datamodel.DB{
			Path:     path,
			Delim:    drunes[0],
			Perm:     0644,
			Filelock: flock.New(path + ".lock"),
		},
	}

	err := db.Filelock.Lock()
	if err != nil {
		return nil, fmt.Errorf("locking failed during New DB : %w", err)
	}

	defer func() {
		_ = db.Filelock.Unlock()
	}()

	err = db.LoadMaxID()
	if err != nil {
		return nil, err
	}

	return &db.DB, nil
}
