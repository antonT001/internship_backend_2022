package objectstorage

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"user_balance/service/internal/helpers"
	"user_balance/service/internal/logger"
)

type ObjectStorage interface {
	SaveCsvFile(path string, records [][]string) error
}

type objectStorage struct {
	baseStoragePath string
	logger          logger.Logger
}

func New(logger logger.Logger) ObjectStorage {
	baseStoragePath := helpers.GetBaseStoragePath()

	return &objectStorage{
		baseStoragePath: baseStoragePath,
		logger:          logger,
	}
}

func (o *objectStorage) SaveCsvFile(path string, records [][]string) error {
	baseDir := filepath.Dir(o.baseStoragePath + path)

	err := os.MkdirAll(baseDir, 0777)
	if err != nil {
		return err
	}

	f, err := os.Create(o.baseStoragePath + path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	w := csv.NewWriter(f)
	w.Comma = ';'
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		return fmt.Errorf("error writing csv: %v", err)
	}

	return nil
}
