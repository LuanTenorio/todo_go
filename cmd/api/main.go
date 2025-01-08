package main

import (
	"github.com/LuanTenorio/todo_go/internal/config"
	"github.com/LuanTenorio/todo_go/internal/db"
	"github.com/LuanTenorio/todo_go/internal/router"
)

func main() {
	config.ReadConfig()
	db.Inicialize()
	router.Inicialize()
}
