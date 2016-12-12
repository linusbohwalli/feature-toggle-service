package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type FeatureToggleStoreImpl struct {
	db *sql.DB
}

func NewFeatureToggleStoreImpl() *FeatureToggleStoreImpl {
	return new(FeatureToggleStoreImpl)
}

func (fs *FeatureToggleStoreImpl) Open() error {
	config := GetConfig()
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.HOST, config.Database.USER, config.Database.PASSWORD, config.Database.NAME, config.Database.PORT)
	db, err := sql.Open("postgres", dbinfo)
	if err == nil {
		fs.db = db
	}
	return err
}

func (fs *FeatureToggleStoreImpl) Close() {
	fs.db.Close()
}
