package main

import (
	"fmt"

	"github.com/aryansh13/go-restapi-gin/database"
)

func main() {
	// Inisialisasi koneksi database
	db, err := database.SetupDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Jalankan migrasi tabel
	err = database.RunMigrations(db)
	if err != nil {
		fmt.Println("Failed to run migrations:", err)
		return
	}

	fmt.Println("Database setup and migrations completed successfully!")
}
