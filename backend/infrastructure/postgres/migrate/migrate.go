package migrate

import (
	userDomain "repocon/domain/user"
	"repocon/infrastructure/database"

	"time"
)

func AutoMigrate() {
	dbConn, err := database.NewDBWithRetry(5, 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&userDomain.User{}); err != nil {
		panic(err)
	}
}
