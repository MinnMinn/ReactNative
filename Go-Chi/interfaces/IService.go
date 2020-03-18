package interfaces

import (
	"Go-Chi/models"
)

type IService interface {
	SCheckNonAddFriend(friends models.Friends) bool
	SAddFriend(friends models.Friends) error
	SFindFriendsOfUser(m models.Email) []string
	SFindCommonFriends(friends models.Friends)[]string
	SCheckNonFollow(subscribe models.Request) bool
	SFollowFriend(subscribe models.Request) error
	SCheckNonBlock(subscribe models.Request) bool
	SBlockFriend(subscribe models.Request) error
	//SNonBlockByEmail(recipients models.Recipients) []string
}
