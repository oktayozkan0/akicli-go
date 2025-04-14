package db

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Db struct {
	Urls struct {
		First struct {
			BaseUrl string `json:"base_url"`
		} `json:"1"`
	} `json:"urls"`
	Default struct {
		First struct {
			Token   string `json:"token"`
			Account int    `json:"account"`
		} `json:"1"`
	} `json:"_default"`
	path      string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getPath(path string) string {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path = filepath.Join(home, ".akinoncli", "db.json")
	}
	return path
}

func GetDb(path string) (*Db, error) {
	path = getPath(path)
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var db Db
	if err = json.Unmarshal(file, &db); err != nil {
		return nil, err
	}
	db.path = path
	return &db, nil
}

func (db *Db) Save() error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(db.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
