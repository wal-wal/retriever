package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"os"
	"retriever/cmd/di"
)

func main() {
	app := fiber.New()
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/WALWAL")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	injector := di.NewDependencyInjector(db)
	router := injector.PureDI()
	router.Initialize(app)

	_ = app.Listen(":8080")
}
