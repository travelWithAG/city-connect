package service

import (
	"connect/rent/entity"
	"connect/rent/repository"
	"fmt"
)

type RoomService interface {
	CreateRoom(room RoomInput) ( *entity.Room, error )
	GetRoomByID(id uint) ( *entity.Room, string )
}

type roomService struct {
	rs repository.RoomRepository
}

func NewRoomService(rs repository.RoomRepository) RoomService {
	return &roomService{rs: rs}
}

func (r *roomService) CreateRoom(dto RoomInput) (*entity.Room, error){

	room := &entity.Room{
		PropertyID: dto.PropertyID,
		RoomNumber: dto.RoomNumber,
		Features: dto.Features,
	}

	if err := r.rs.CreateRoom(room); err != nil{
		return nil, fmt.Errorf("Failed: %s", err)
	}

	return room, nil
}

func (r *roomService) GetRoomByID(id uint) (*entity.Room, string) {
	room, err := r.rs.GetRoomByID(id)
	if err != nil {
		return nil, fmt.Sprintf("%s", err)
	}

	return room, "Data Found"
}