package repositories

import "go_project_structure/cockroach/entities"

type CockroachRepository interface {
  InsertCockroachData(in *entities.InsertCockroachDto) error
}