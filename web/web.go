package main

import (
	"log"
	"net/http"

	"src/modules"
)

func main() {
	log.Println("Connecting to server on port 8000")
	modules.Initialize()
	http.HandleFunc("/", modules.LoginFunctionality)
	http.HandleFunc("/register/", modules.RegisterFunctionality)
	http.HandleFunc("/login/", modules.LoginFunctionality)
	http.HandleFunc("/logout/", modules.LogoutFunctionality)
	http.HandleFunc("/createpost/", modules.CreatePost)
	http.HandleFunc("/findfollowers/", modules.FollowUser)
	http.HandleFunc("/unfollowuser/", modules.UnFollowUser)
	http.HandleFunc("/followers/", modules.FolloweeList)
	http.HandleFunc("/feed/", modules.GetPosts)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
}
