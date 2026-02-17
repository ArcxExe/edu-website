package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "edu.db")
    if err != nil {
        log.Fatal(err)
    }

    // Создание таблиц
    _, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            username TEXT UNIQUE,
            password TEXT
        );
        
        CREATE TABLE IF NOT EXISTS lessons (
            id INTEGER PRIMARY KEY,
            title TEXT,
            content TEXT
        );
    `)
    if err != nil {
        log.Fatal(err)
    }
}
