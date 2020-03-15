package repository

import "Go-Chi/Model"

type PostService interface {
	AllPosts() ([]Model.Post, error)
	DetailPost(id int64) (Model.Post, error)
	CreatePost(p Model.Post) (error)
	UpdatePost(p Model.Post) (error)
	DeletePost(id int64) (error)
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