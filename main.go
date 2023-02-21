package main

import (
	"database/sql"
	"fmt"
	"golang-rest-api/controllers"
	"golang-rest-api/db"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load env")
	} else {
		fmt.Println("success read env")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))
	DB, err = sql.Open("postgres", psqlInfo)
	fmt.Println(DB)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
	} else {
		fmt.Println("DB Connection Success")
	}
	db.DbMigrate(DB)
	defer DB.Close()

	router := gin.Default()
	router.GET("/transaction", controllers.GetAllTransaction)
	router.POST("/transaction", controllers.InsertTransaction)
	router.PUT("/transaction/:id", controllers.UpdateTransaction)
	router.DELETE("/transaction/:id", controllers.DeleteTransaction)
	router.GET("/account", controllers.GetAllAccount)
	router.POST("/account", controllers.InsertAccount)
	router.PUT("/account/:id", controllers.UpdateAccount)
	router.DELETE("/account/:id", controllers.DeleteAccount)

	// router.Run(":" + ("8019"))
	router.Run(":" + os.Getenv("PORT"))
}
