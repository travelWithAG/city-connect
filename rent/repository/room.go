package repository

import (
	"connect/rent/entity"
	"connect/rent/handler"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(room *entity.Room) error
	GetRoomByID(id uint) (*entity.Room, error)
}

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepo(db *gorm.DB) RoomRepository {
	return &roomRepo{db: db}
}

func (r *roomRepo) CreateRoom(roomInput *entity.Room) error {
	if err := r.db.Create(roomInput).Error; err != nil {
		return handler.ExceptionError()
	}
	return nil
}

func (r *roomRepo) GetRoomByID(id uint) (*entity.Room, error) {
	var room entity.Room

	if err := r.db.Preload("Rooms").First(&room, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, handler.NotFound("Room", "Not found")
		}
		return nil, handler.ExceptionError()
	}

	return &room, nil

}

