package repository

import "Go-Chi/Model"

type FriendService interface {
	CheckAddFriend(friends Model.Friends) bool
	AddFriend(friends Model.Friends) error
	FindFriendsOfUser(m Model.Mail) []string
	FindCommonFriends(friends Model.Friends)[]string
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