package interfaces

import (
	"Go-Chi/models"
)

type IRepository interface {
	RCheckNonAddFriend(friends models.Friends) bool
	RAddFriend(friends models.Friends) error
	RFindFriendsOfUser(m models.Email) []string
	RFindCommonFriends(friends models.Friends)[]string
	RCheckNonFollow(subscribe models.Request) bool
	RFollowFriend(subscribe models.Request) error
	RCheckNonBlock(subscribe models.Request) bool
	RBlockFriend(subscribe models.Request) error
	//RNonBlockByEmail(recipients models.Recipients) []string
}
