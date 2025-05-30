package service

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"log"
	"next_device/backend/models"
	"next_device/backend/repository"
	"os"
)

type NotificationService struct {
	repo            *repository.PushRepository
	vapidPublicKey  string
	vapidPrivateKey string
}

func NewNotificationService(repo *repository.PushRepository) *NotificationService {
	publicKey := os.Getenv("PUBLIC_KEY_VAPID")
	privateKey := os.Getenv("PRIVATE_KEY_VAPID")

	if publicKey == "" || privateKey == "" {
		panic("VAPID keys not set in environment variables")
	}

	return &NotificationService{
		repo:            repo,
		vapidPublicKey:  publicKey,
		vapidPrivateKey: privateKey,
	}
}

func (s *NotificationService) GetPublicKey() string {
	return s.vapidPublicKey
}

func (s *NotificationService) AddSubscription(sub *models.Subscription, userID int) {
	_ = s.repo.Save(sub, userID)
}

func (s *NotificationService) SendNotificationToAll(notif models.Notification) error {
	subs, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	payload, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	for _, sub := range subs {
		wsub := webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.Keys.P256dh,
				Auth:   sub.Keys.Auth,
			},
		}

		_, err := webpush.SendNotification(payload, &wsub, &webpush.Options{
			Subscriber:      os.Getenv("VAPID_SUBJECT"),
			VAPIDPublicKey:  s.vapidPublicKey,
			VAPIDPrivateKey: s.vapidPrivateKey,
		})

		if err != nil {
			continue
		}
	}

	return nil
}

func (s *NotificationService) SendNotificationToUser(userID int, notif models.Notification) error {
	subs, err := s.repo.GetByUserID(userID)
	if err != nil || len(subs) == 0 {
		return err
	}

	payload, _ := json.Marshal(notif)

	for _, sub := range subs {
		wsub := webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.Keys.P256dh,
				Auth:   sub.Keys.Auth,
			},
		}

		_, err := webpush.SendNotification(payload, &wsub, &webpush.Options{
			Subscriber:      os.Getenv("VAPID_SUBJECT"),
			VAPIDPublicKey:  s.vapidPublicKey,
			VAPIDPrivateKey: s.vapidPrivateKey,
		})

		if err != nil {
			log.Printf("❌ Ошибка при отправке уведомления: %v", err)
			return err
		}
	}

	fmt.Println("Отправлен")

	return nil
}
