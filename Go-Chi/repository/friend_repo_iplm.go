package repository

import (
	"Go-Chi/Model"
	"database/sql"
	"fmt"
	"rest-api-mysql/driver"
)

type FriendSer struct {
	Db *sql.DB
}

func FriendRepo(db *sql.DB) FriendService {
	return &FriendSer {
		Db: db,
	}
}

func (s *FriendSer) CheckAddFriend(friends Model.Friends) bool {
	connect, err := driver.DbConn().Query("select `user_id` from `connection` where `user_id` = (select `id` from `user` where `email`=?) AND `connect_id` = (select `id` from `user` where `email`=?)", friends.Friends[0], friends.Friends[1])
	catch(err)

	for connect.Next(){
		var userId sql.NullInt64
		err = connect.Scan(&userId)
		catch(err)
		fmt.Println(userId)
		if userId.Valid {
			return false
		}
	}

	defer connect.Close()
	return true
}

func (s *FriendSer) AddFriend(friends Model.Friends) error{
	user, err := driver.DbConn().Query("select `id` from `user` where `email`=?", friends.Friends[0])
	catch(err)
	connectUser, err := driver.DbConn().Query("select `id` from `user` where `email`=?", friends.Friends[1])
	catch(err)
	addFriend, err := driver.DbConn().Prepare("INSERT `connection` SET user_id=?, connect_id=?")
	catch(err)
	var user_id int
	for user.Next(){
		user.Scan(&user_id)
	}
	var connect_id int
	for connectUser.Next(){
		connectUser.Scan(&connect_id)
	}
	_, err = addFriend.Exec(user_id, connect_id)
	catch(err)
	_, err = addFriend.Exec(connect_id, user_id)
	catch(err)
	defer user.Close()
	defer connectUser.Close()
	defer addFriend.Close()
	return err
}

func (s *FriendSer) FindFriendsOfUser(m Model.Mail) []string {
	idFriends, err := driver.DbConn().Query("select `connect_id` from `connection` where `user_id` = (select `id` from `user` where `email`=?)", m.Mail)
	var id []int
	var email []string
	for idFriends.Next(){
		var connect_id int
		err = idFriends.Scan(&connect_id)
		catch(err)
		id = append(id, connect_id)
		emailFriends, err := driver.DbConn().Query("select `email` from `user` where `id`=?", connect_id)
		for emailFriends.Next(){
			var mail string
			err = emailFriends.Scan(&mail)
			catch(err)
			email = append(email, mail)
		}
	}
	return email
}

func (s *FriendSer) FindCommonFriends(friends Model.Friends)[]string{
	friendsOfUser1, err := driver.DbConn().Query("SELECT `connect_id` FROM `connection` where `user_id` = (select `id` from `user` where `email`=?)", friends.Friends[0])
	catch(err)
	friendsOfUser2, err := driver.DbConn().Query("SELECT `connect_id` FROM `connection` where `user_id` = (select `id` from `user` where `email`=?)", friends.Friends[1])
	catch(err)

	var idUser1 []int
	for friendsOfUser1.Next(){
		var idUser int
		friendsOfUser1.Scan(&idUser)
		idUser1 = append(idUser1, idUser)
	}

	var idUser2 []int
	for friendsOfUser2.Next(){
		var idUser int
		friendsOfUser2.Scan(&idUser)
		idUser2 = append(idUser2, idUser)
	}

	var commonIds []int
	for i := 0; i < len(idUser1); i++ {
		for j := 0; j <len(idUser2); j++ {
			if idUser1[i] == idUser2[j] {
				commonIds = append(commonIds, idUser1[i])
			}
		}
	}

	var email []string
	for i:= 0; i < len(commonIds); i++ {
		emailCommonFriends, err := driver.DbConn().Query("select `email` from `user` where `id`=?", commonIds[i])
		catch(err)
		for emailCommonFriends.Next(){
			var mail string
			emailCommonFriends.Scan(&mail)
			email = append(email, mail)
		}
	}
	return email
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}