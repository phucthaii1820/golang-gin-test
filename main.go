package main

import (
	"test/inits"
	"test/migrations"
	routers "test/router"

	"github.com/gin-gonic/gin"
)


func init() {
	inits.LoadEnvVariable()
	inits.ConnectDB()
	migrations.Migrate()
}

func main() {
	r := gin.Default()

	routers.CardRouter(r)

	r.Run()
}