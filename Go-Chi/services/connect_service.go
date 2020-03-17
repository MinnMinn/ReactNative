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
	var block = Model.Request{Requestor:friends.Friends[0], Target:friends.Friends[1]}
	var beBlock = Model.Request{Requestor:friends.Friends[1], Target:friends.Friends[0]}

	var checkNonAddFriend = connectionRepo.CheckNonAddFriend(friends)
	var checkNonBlock = connectionRepo.CheckNonBlock(block)
	var checkNonBeBlock = connectionRepo.CheckNonBlock(beBlock)

	if checkNonAddFriend && checkNonBlock && checkNonBeBlock{
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
	var mail Model.Email
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

func FollowFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var subscribe Model.Request
	json.NewDecoder(r.Body).Decode(&subscribe)

	var checkNonFollow = connectionRepo.CheckNonFollow(subscribe)
	if checkNonFollow{
		connectionRepo.Follow(subscribe)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed follow"})
	}
}

func BlockFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var block Model.Request
	json.NewDecoder(r.Body).Decode(&block)

	var checkNonBlock = connectionRepo.CheckNonBlock(block)
	if checkNonBlock {
		connectionRepo.Block(block)
		respondwithJSON(w, http.StatusCreated, map[string]bool{"success": true})
	} else {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"failed": "You were blocked this account !!!"})
	}
}

func ReceiveUpdatesFromEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := driver.DBConn()
	connectionRepo := repository.FriendRepo(db)
	var sender Model.Sender
	json.NewDecoder(r.Body).Decode(&sender)
	var receiveUpdates []string

	var emails = connectionRepo.NonBlockByEmail(sender)
	for i := 0; i < len(emails); i++ {
		var friends = Model.Friends{Friends:[]string{sender.Sender, emails[i]}}
		var subscribe = Model.Request{Requestor:sender.Sender, Target: emails[i]}
		var checkNonAddFriend = connectionRepo.CheckNonAddFriend(friends)
		var checkNonFollow = connectionRepo.CheckNonFollow(subscribe)
		if !checkNonAddFriend || !checkNonFollow {
			receiveUpdates = append(receiveUpdates, emails[i])
		}
	}
	if len(receiveUpdates) > 0 {
		respondwithJSON(w, http.StatusCreated, Model.Recipients{Success: true, Recipients:receiveUpdates})
	}
}