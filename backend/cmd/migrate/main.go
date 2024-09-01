package main

import "repocon/infrastructure/postgres/migrate"

func main() {
	migrate.AutoMigrate()
}
