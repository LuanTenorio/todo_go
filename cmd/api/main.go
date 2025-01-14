package main

import (
	"github.com/LuanTenorio/todo_go/internal/config"
	"github.com/LuanTenorio/todo_go/internal/database"
	"github.com/LuanTenorio/todo_go/internal/server"
)

func main() {
	conf := config.GetConfig()
	var db database.Database = database.NewPostgresDatabase(&conf)
	server.NewEchoServer(&conf, db).Start()
}
