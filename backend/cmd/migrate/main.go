package main

import (
	"fmt"
	"repocon/infrastructure/postgres/migrate"
)

func main() {
	fmt.Println("Start Migration")
	migrate.AutoMigrate()
}
