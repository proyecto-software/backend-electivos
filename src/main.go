// main.go

package main

import (
	"database/sql"
	"electivos-ucn/src/database"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	api := r.Group("/ucn")
	{
		public := api.Group("/public")
		{
			public.POST("/login")
			//public.POST("/data", controllers.TestEndPoint)
		}

	}
	return r
}

func main() {
	db := database.DBConnection()
	r := setupRouter(db)
	port := os.Getenv("PORT")
	r.Run(port)
}
