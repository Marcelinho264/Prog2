package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/teste")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	r.GET("/usuario", func(c *gin.Context) {
		result, err :=db.Exec("insert into ")

		if err != nil {
			c.JSON(200, gin.H{
				"message": err,
			})
		}

		c.JSON(200, gin.H{
			"message": result,
		})
	})
	r.Run()

}

