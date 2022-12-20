/*
	New gprc service started here for postService

	reference: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
*/

package main

import (
	"log"
	"net"

	"src/services/post/post_pb"
	post_methods "src/services/post/post_repository_storage"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Post Server started at Port 9100 ")
	post_listen, err := net.Listen("tcp", ":9100")

	if err != nil {
		log.Fatalf("Failed to listen to post server: %v", err)
	}

	postServer := grpc.NewServer()
	post_pb.RegisterPostsServiceServer(postServer, &post_methods.Server{})
	post_methods.InitializePostService()

	if err := postServer.Serve(post_listen); err != nil {
		log.Fatalf("Failed to server post server: %v", err)
	}
}
