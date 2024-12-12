package services

import "go_project_structure/cockroach/models"

type CockroachService interface {
  CockroachDataProcessing(in *models.AddCockroachData) error
}