package main

import (
  "go_project_structure/config"
  "go_project_structure/database"
  "go_project_structure/server"
)

func main() {
  conf := config.GetConfig()
  db := database.NewPostgresDatabase(conf)
  server.NewEchoServer(conf, db).Start()
}