package main

import (
	"fmt"
	"repocon/di"
	"repocon/infrastructure/database"
	"repocon/interface/router"
	"time"
)

func main() {
	fmt.Println("Start api server")
	db, err := database.NewDBWithRetry(5, 5*time.Second)
	if err != nil {
		fmt.Println("error starting the api server")
		return
	}
	defer database.CloseDB(db)
	e := router.NewRouter(di.User(db))
	e.Logger.Fatal(e.Start(":8080"))
}
