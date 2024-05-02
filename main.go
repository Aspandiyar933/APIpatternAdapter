package main

import (
	"github.com/Aspandiyar933/APIpatternAdapter/api"
	"github.com/Aspandiyar933/APIpatternAdapter/config"
	"github.com/Aspandiyar933/APIpatternAdapter/store"
	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User: config.Envs.DBUser,
	}

	api := api.NewAPIServer(":3000", store)
	api.Serve()
}