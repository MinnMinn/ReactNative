package services

import (
	"Go-Chi/Model"
	"Go-Chi/driver"
	"Go-Chi/repository"
	"encoding/json"
	"net/http"
)

func AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var friends Model.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var checkAddFriend = connectionRepo.CheckAddFriend(friends)
	if checkAddFriend {
		connectionRepo.AddFriend(friends)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed add friend"})
	}
}

func FindFriendsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var mail Model.Mail
	json.NewDecoder(r.Body).Decode(&mail)
	var status = Model.Friends{Success: true,Friends: connectionRepo.FindFriendsOfUser(mail), Count: len(connectionRepo.FindFriendsOfUser(mail))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Friends Of User"})
	}
}

func FindCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var friends Model.Friends
	json.NewDecoder(r.Body).Decode(&friends)
	var status = Model.Friends{Success: true, Friends: connectionRepo.FindCommonFriends(friends), Count: len(connectionRepo.FindCommonFriends(friends))}
	if len(status.Friends) >0 {
		respondwithJSON(w, http.StatusOK, status)
	} else {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Not Found Common Friends"})
	}
}