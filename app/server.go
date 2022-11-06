package app

import (
	"TodoList-GO/app/utils"
	"TodoList-GO/config"
	"TodoList-GO/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(utils.MsgGetEnvErr, err)
	}

	// Environment Variables
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	// Database
	DB, err := config.InitDatabase(dbUser, dbPassword, dbPort, dbHost, dbName)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Router
	router := echo.New()
	v1 := router.Group("/v1")

	// Routes
	routes.TodoRoute(v1, DB)
	routes.SubTodoRoute(v1, DB)

	err = router.Start(":8080")
	if err != nil {
		log.Fatalf(utils.MsgRunServerErr)
	}
}
