package main

import (
	"log"
	"os"

	"ginExample/controllers"

	"ginExample/routes"
	"ginExample/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal("err", err)
	}
	todoController := controllers.NewTodoController(db)
	r := gin.Default()
	routes.CreateRoutes(r, todoController)
	r.Run()
}
