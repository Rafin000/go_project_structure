package repositories

import "go_project_structure/cockroach/entities"

type CockroachMessaging interface {
  PushNotification(m *entities.CockroachPushNotificationDto) error
}