package services

import (
	"Go-Chi/interfaces"
	"Go-Chi/driver"
	"Go-Chi/models"
	"Go-Chi/repositories"
	"database/sql"
)

type Service struct {
	interfaces.IService
}

func FriendService(db *sql.DB) interfaces.IService {
	return &Service {
	}
}

func (service *Service) SCheckNonAddFriend(friends models.Friends) bool {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RCheckNonAddFriend(friends)
}

func (service *Service) SAddFriend(friends models.Friends) error {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RAddFriend(friends)
}

func (service *Service) SFindFriendsOfUser(m models.Email) []string {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RFindFriendsOfUser(m)
}

func (service *Service) SFindCommonFriends(friends models.Friends) []string {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RFindCommonFriends(friends)
}

func (service *Service) SCheckNonFollow(subscribe models.Request) bool {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RCheckNonFollow(subscribe)
}

func (service *Service) SFollowFriend(subscribe models.Request) error {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RFollowFriend(subscribe)
}

func (service *Service) SCheckNonBlock(subscribe models.Request) bool {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RCheckNonBlock(subscribe)
}

func (service *Service) SBlockFriend(subscribe models.Request) error {
	db := driver.DBConn()
	connectionRepo := repositories.FriendRepository(db)
	return  connectionRepo.RBlockFriend(subscribe)
}
