package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type FeatureToggleStoreImpl struct {
	db     *sql.DB
	config *Config
}

func NewFeatureToggleStoreImpl() *FeatureToggleStoreImpl {
	return new(FeatureToggleStoreImpl)
}

func (fs *FeatureToggleStoreImpl) Open() error {
	var config Config
	if fs.config == nil {
		config = GetConfig()
	} else {
		config = *fs.config
	}
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.HOST, config.Database.USER, config.Database.PASSWORD, config.Database.NAME, config.Database.PORT)
	db, err := sql.Open("postgres", dbinfo)
	if err == nil {
		fs.db = db
	}
	return err
}

func (fs *FeatureToggleStoreImpl) SetConfig(config Config) {
	fs.config = &config
}

func (fs *FeatureToggleStoreImpl) Close() {
	fs.db.Close()
}
