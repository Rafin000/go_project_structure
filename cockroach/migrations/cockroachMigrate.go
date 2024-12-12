package main

import (
 "go_project_structure/cockroach/entities"
 "go_project_structure/config"
 "go_project_structure/database"
)

func main() {
  conf := config.GetConfig()
  db := database.NewPostgresDatabase(conf)
  cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
  db.GetDb().Migrator().CreateTable(&entities.Cockroach{})
  db.GetDb().CreateInBatches([]entities.Cockroach{
    {Amount: 1},
    {Amount: 2},
    {Amount: 2},
    {Amount: 5},
    {Amount: 3},
  }, 10)
}