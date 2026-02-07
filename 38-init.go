package main

import (
	"belajar-go-dasar/database"
	_ "belajar-go-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println("Database connection type:", database.GetDatabaseConnection())
}
