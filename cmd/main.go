package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bootcamp-go/desafio-cierre-db.git/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db, err := sql.Open("mysql", "user1:secret_password@/fantasy_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	router.NewRouter(r, db).MapRoutes()

	r.Run()

}
