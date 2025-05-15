package service

import (
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	"next_device/backend/models"
	"os"
	"sync"
)

type NotificationService struct {
	subscriptions   []models.Subscription
	mu              sync.Mutex
	vapidPublicKey  string
	vapidPrivateKey string
}

func NewNotificationService() *NotificationService {
	publicKey := os.Getenv("PUBLIC_KEY_VAPID")
	privateKey := os.Getenv("PRIVATE_KEY_VAPID")

	if publicKey == "" || privateKey == "" {
		panic("VAPID keys not set in environment variables")
	}

	return &NotificationService{
		subscriptions:   make([]models.Subscription, 0),
		vapidPublicKey:  publicKey,
		vapidPrivateKey: privateKey,
	}
}

func (s *NotificationService) GetPublicKey() string {
	return s.vapidPublicKey
}

func (s *NotificationService) AddSubscription(sub models.Subscription) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subscriptions = append(s.subscriptions, sub)
}

func (s *NotificationService) SendNotificationToAll(notif models.Notification) error {
	s.mu.Lock()
	subs := make([]models.Subscription, len(s.subscriptions))
	copy(subs, s.subscriptions)
	s.mu.Unlock()

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
