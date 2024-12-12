package services

import (
  "time"
  
  "go_project_structure/cockroach/entities"
  "go_project_structure/cockroach/models"
  "go_project_structure/cockroach/repositories"
)
type cockroachServiceImpl struct {
  cockroachRepository repositories.CockroachRepository
  cockroachMessaging  repositories.CockroachMessaging
}

func NewCockroachServiceImpl(
  cockroachRepository repositories.CockroachRepository,
  cockroachMessaging repositories.CockroachMessaging,
) CockroachService {
  return &cockroachServiceImpl{
    cockroachRepository: cockroachRepository,
    cockroachMessaging:  cockroachMessaging,
  }
}

func (u *cockroachServiceImpl) CockroachDataProcessing(in *models.AddCockroachData) error {
  insertCockroachData := &entities.InsertCockroachDto{
    Amount: in.Amount,
  }

  if err := u.cockroachRepository.InsertCockroachData(insertCockroachData); err != nil {
    return err
  }

  pushCockroachData := &entities.CockroachPushNotificationDto{
    Title:        "Cockroach Detected 🪳 !!!",
    Amount:       in.Amount,
    ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
  }

  if err := u.cockroachMessaging.PushNotification(pushCockroachData); err != nil {
    return err
  }

  return nil
}