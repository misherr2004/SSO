package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3" // Драйвер SQLite
	"github.com/pressly/goose/v3"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "storage.db", "path to SQLite database file")
	flag.StringVar(&migrationsPath, "migrations-path", "migrations", "path to migrations folder")
	flag.StringVar(&migrationsTable, "migrations-table", "goose_migrations", "name of migrations table")
	flag.Parse()

	// Валидация
	if storagePath == "" {
		log.Fatal("storage-path is required")
	}
	if migrationsPath == "" {
		log.Fatal("migrations-path is required")
	}

	// Подключение к SQLite
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Указываем диалект БД (SQLite)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal(err)
	}

	// Запуск миграций
	if err := goose.Up(db, migrationsPath); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	log.Println("migrations applied successfully")
}
