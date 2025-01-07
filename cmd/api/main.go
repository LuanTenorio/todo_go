package main

import (
	"github.com/LuanTenorio/todo_go/internal/config"
	"github.com/LuanTenorio/todo_go/internal/router"
)

func main() {
	config.ReadConfig()
	router.Inicialize()
}
