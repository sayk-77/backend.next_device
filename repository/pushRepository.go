package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type PushRepository struct {
	db *gorm.DB
}

func NewPushRepository(db *gorm.DB) *PushRepository {
	return &PushRepository{db: db}
}

func (p *PushRepository) Save(sub *models.Subscription, userID int) error {
	pushSub := models.PushSubscription{
		Endpoint: sub.Endpoint,
		P256dh:   sub.Keys.P256dh,
		Auth:     sub.Keys.Auth,
		UserID:   userID,
	}

	var existing models.PushSubscription
	if err := p.db.Where("endpoint = ?", sub.Endpoint).First(&existing).Error; err != nil {
		return p.db.Create(&pushSub).Error
	}

	return nil
}

func (p *PushRepository) GetAll() ([]models.Subscription, error) {
	var records []models.PushSubscription
	if err := p.db.Find(&records).Error; err != nil {
		return nil, err
	}

	subs := make([]models.Subscription, len(records))
	for i, r := range records {
		subs[i] = models.Subscription{
			Endpoint: r.Endpoint,
			Keys: models.Keys{
				P256dh: r.P256dh,
				Auth:   r.Auth,
			},
		}
	}

	return subs, nil
}

func (p *PushRepository) GetByUserID(userID int) ([]models.Subscription, error) {
	var records []models.PushSubscription
	if err := p.db.Where("user_id = ?", userID).Find(&records).Error; err != nil {
		return nil, err
	}

	subs := make([]models.Subscription, len(records))
	for i, r := range records {
		subs[i] = models.Subscription{
			Endpoint: r.Endpoint,
			Keys: models.Keys{
				P256dh: r.P256dh,
				Auth:   r.Auth,
			},
		}
	}

	return subs, nil
}
