package db

import (
	"api/pkg/utils/reader"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

var MyDB *sql.DB

func ConnDB() {
	var cfg = getYamlInfo("configs/configs.yaml")
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", psqlConn)
	check(err)

	err = db.Ping()
	check(err)

	MyDB = db
}

func getYamlInfo(fileName string) Config {
	fileContext := reader.ReadFile(fileName)
	var configs Config
	err := yaml.Unmarshal(fileContext, &configs)
	check(err)
	return configs
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
