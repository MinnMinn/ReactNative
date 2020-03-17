package router

import (
	"Go-Chi/services"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

var router *chi.Mux

func Post_router() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/posts", services.AllPosts)
	router.Get("/posts/{id}", services.DetailPost)
	router.Post("/posts", services.CreatePost)
	router.Put("/posts/{id}", services.UpdatePost)
	router.Delete("/posts/{id}", services.DeletePost)
	router.Post("/addfriend", services.AddFriend)
	router.Get("/findFriendOfUser", services.FindFriendsOfUser)
	router.Get("/findCommonFriends", services.FindCommonFriends)
	router.Post("/followFriend", services.FollowFriend)
	router.Post("/blockFriend", services.BlockFriend)
	router.Get("/receiveUpdatesFromEmail", services.ReceiveUpdatesFromEmail)
	http.ListenAndServe(":8005", logger())
}

func logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}