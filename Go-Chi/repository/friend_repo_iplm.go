package repository

import (
	"Go-Chi/Model"
	"Go-Chi/driver"
	"database/sql"
)

type FriendSer struct {
	Db *sql.DB
}

func FriendRepo(db *sql.DB) FriendService {
	return &FriendSer {
		Db: db,
	}
}

func (s *FriendSer) CheckNonAddFriend(friends Model.Friends) bool {
	connect, err := driver.DBConn().Query("select `user_id` from `connection` where `user_id` = (select `id` from `user` where `email`=?) AND `connect_id` = (select `id` from `user` where `email`=?)", friends.Friends[0], friends.Friends[1])
	catch(err)

	for connect.Next(){
		var userId sql.NullInt64
		err = connect.Scan(&userId)
		catch(err)
		if userId.Valid {
			return false
		}
	}

	defer connect.Close()
	return true
}

func (s *FriendSer) AddFriend(friends Model.Friends) error{
	addFriend, err := driver.DBConn().Prepare("INSERT `connection` SET user_id=(select `id` from `user` where `email`=?), connect_id=(select `id` from `user` where `email`=?)")
	catch(err)
	_, err = addFriend.Exec(friends.Friends[0], friends.Friends[1])
	catch(err)
	_, err = addFriend.Exec(friends.Friends[1], friends.Friends[0])
	catch(err)
	defer addFriend.Close()
	return err
}

func (s *FriendSer) FindFriendsOfUser(m Model.Email) []string {
	emailFriends, err := driver.DBConn().Query("select `email` from `user` where `id` in (select `connect_id` from `connection` where `user_id` = (select `id` from `user` where `email`=?))", m.Email)
	catch(err)
	var email []string
	for emailFriends.Next(){
		var e string
		err = emailFriends.Scan(&e)
		catch(err)
		email = append(email, e)
	}
	return email
}

func (s *FriendSer) FindCommonFriends(friends Model.Friends)[]string{
	commonFriends, err := driver.DBConn().Query("SELECT `email` from `user` WHERE `id` IN (SELECT `user_id` from `connection` JOIN (SELECT `id` FROM `user` where `email` in ( ?, ?)) t ON `connect_id` = `id` group by `user_id` having count(`user_id`) > 1)", friends.Friends[0], friends.Friends[1])
	catch(err)
	var email []string
	for commonFriends.Next(){
		var e string
		commonFriends.Scan(&e)
		email = append(email, e)
	}
	return email
}

func (s *FriendSer) CheckNonFollow(subscribe Model.Request) bool {
	follow, err := driver.DBConn().Query("select `user_id` from `follow` where `user_id` = (select `id` from `user` where `email`=?) AND `follow_id` = (select `id` from `user` where `email`=?)", subscribe.Requestor, subscribe.Target)
	catch(err)

	for follow.Next(){
		var userId sql.NullInt64
		err = follow.Scan(&userId)
		catch(err)
		if userId.Valid {
			return false
		}
	}

	defer follow.Close()
	return true
}

func (s *FriendSer) Follow(subscribe Model.Request) error {
	followUser, err := driver.DBConn().Prepare("INSERT `follow` SET `user_id`=(select `id` from `user` where `email`=?), follow_id=(select `id` from `user` where `email`=?)")
	catch(err)
	_, err = followUser.Exec(subscribe.Requestor, subscribe.Target)
	catch(err)
	defer followUser.Close()
	return err
}

func (s *FriendSer) CheckNonBlock(subscribe Model.Request) bool {
	block, err := driver.DBConn().Query("select `user_id` from `block` where `user_id` = (select `id` from `user` where `email`=?) AND `block_id` = (select `id` from `user` where `email`=?)", subscribe.Requestor, subscribe.Target)
	catch(err)

	for block.Next(){
		var userId sql.NullInt64
		err = block.Scan(&userId)
		catch(err)
		if userId.Valid {
			return false
		}
	}

	defer block.Close()
	return true
}

func (s *FriendSer) Block(subscribe Model.Request) error {
	blockUser, err := driver.DBConn().Prepare("INSERT `block` SET `user_id`=(select `id` from `user` where `email`=?), block_id=(select `id` from `user` where `email`=?)")
	catch(err)
	_, err = blockUser.Exec(subscribe.Requestor, subscribe.Target)
	catch(err)
	defer blockUser.Close()
	return err
}

func (s *FriendSer) NonBlockByEmail(recipients Model.Sender) []string {
	nonBlockId, err := driver.DBConn().Query("SELECT `email` FROM `user` WHERE `id` NOT IN (SELECT `block_id` from `block` join( SELECT `id` FROM `user` where `email` = ?) `u` ON `user_id` = `u`.`id`)", recipients.Sender)
	catch(err)
	var emails []string
	for nonBlockId.Next() {
		var email string
		err = nonBlockId.Scan(&email)
		catch(err)
		emails = append(emails, email)
	}
	return emails
}

func catch(err error) {
	if err != nil {
		panic(err.Error())
	}
}