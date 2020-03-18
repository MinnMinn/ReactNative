package controllers

import (
	"Go-Chi/driver"
	"Go-Chi/models"
	"Go-Chi/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := services.FriendService(db)
	var friends models.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var block = models.Request{Requestor: friends.Friends[0], Target:friends.Friends[1]}
	var beBlock = models.Request{Requestor: friends.Friends[1], Target:friends.Friends[0]}

	var checkNonAddFriend = connectionRepo.SCheckNonAddFriend(friends)
	var checkNonBlock = connectionRepo.SCheckNonBlock(block)
	var checkNonBeBlock = connectionRepo.SCheckNonBlock(beBlock)

	if checkNonAddFriend && checkNonBlock && checkNonBeBlock{
		connectionRepo.SAddFriend(friends)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed add friend"})
	}
}

func FindFriendsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := services.FriendService(db)
	var mail models.Email
	json.NewDecoder(r.Body).Decode(&mail)
	var status = models.Friends{Success: true,Friends: connectionRepo.SFindFriendsOfUser(mail), Count: len(connectionRepo.SFindFriendsOfUser(mail))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Friends Of User"})
	}
}

func FindCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := services.FriendService(db)
	var friends models.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var status = models.Friends{Success: true, Friends: connectionRepo.SFindCommonFriends(friends), Count: len(connectionRepo.SFindCommonFriends(friends))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Common Friends"})
	}
}

func FollowFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := services.FriendService(db)
	var subscribe models.Request
	json.NewDecoder(r.Body).Decode(&subscribe)

	var checkNonFollow = connectionRepo.SCheckNonFollow(subscribe)

	if checkNonFollow{
		connectionRepo.SFollowFriend(subscribe)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed follow"})
	}
}

func BlockFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := services.FriendService(db)
	var block models.Request
	json.NewDecoder(r.Body).Decode(&block)

	var checkNonBlock = connectionRepo.SCheckNonBlock(block)
	if checkNonBlock {
		connectionRepo.SBlockFriend(block)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"failed": "You were blocked this account !!!"})
	}
}

//func receiveUpdatesFromEmail(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	db := driver.DBConn()
//	connectionRepo := repository.FriendRepo(db)
//	var recipients Model.Recipients
//	json.NewDecoder(r.Body).Decode(&recipients)
//	var emails = connectionRepo.NonBlockByEmail(recipients)
//
//
//	if checkNonFollow && checkNonBlock && checkNonBeBlock{
//		connectionRepo.Follow(subscribe)
//		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
//	} else {
//		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed follow"})
//	}
//}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}