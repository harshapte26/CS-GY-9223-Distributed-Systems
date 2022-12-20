/*
	New gprc service started here for userService

	reference: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
*/

package main

import (
	"log"
	"net"

	"src/services/user/user_pb"
	user_methods "src/services/user/user_repository_storage"
	"google.golang.org/grpc"
)

func main() {
	log.Println("User Server started at Port 9200 ")
	post_listen, err := net.Listen("tcp", ":9200")

	if err != nil {
		log.Fatalf("Failed to listen to user server: %v", err)
	}

	userServer := grpc.NewServer()

	user_pb.RegisterUserServiceServer(userServer, &user_methods.Server{})
	user_methods.InitializeUserService()

	if err := userServer.Serve(post_listen); err != nil {
		log.Fatalf("Failed to serve user server: %v", err)
	}
}
