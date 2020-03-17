package repository

import "Go-Chi/Model"

type FriendService interface {
	CheckNonAddFriend(friends Model.Friends) bool
	AddFriend(friends Model.Friends) error
	FindFriendsOfUser(m Model.Email) []string
	FindCommonFriends(friends Model.Friends)[]string
	CheckNonFollow(subscribe Model.Request) bool
	Follow(subscribe Model.Request) error
	CheckNonBlock(subscribe Model.Request) bool
	Block(subscribe Model.Request) error
	NonBlockByEmail(sender Model.Sender) []string
}

//func AllPosts() ([]Model.Post) {
//	note := repository.AllPosts()
//	return note
//}
//
//func DetailPost(id string) (Model.Post) {
//	note := repository.DetailPost(id)
//	return note
//}
//
//func CreatePost(p Model.Post) {
//	repository.CreatePost(p)
//}
//
//func UpdatePost(id string){
//	repository.UpdatePost(id)
//}
//
//func DeletePost(id int){
//	repository.DeletePost(id)
//}