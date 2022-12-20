/*
	Initialize the client ports as 9000, 9100 and 9200 (which is where the grpc services will start at).
	Set the dialed connection to the respective clients to be used anywhere needed.

	reference: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
*/

package modules

import (
	"src/services/post/post_pb"
	"src/services/tokenActions/tokenActions_pb"
	"src/services/user/user_pb"
	"google.golang.org/grpc"
)

type Clients struct {
	TokenClient tokenActions_pb.TokenActionsServiceClient
	PostClient  post_pb.PostsServiceClient
	UserClient  user_pb.UserServiceClient
}

type ClientPorts struct {
	AuthPort string
	PostPort string
	UserPort string
}

var clientObj Clients

var clientPortObj ClientPorts

func InitializeClientPorts() {
	clientPortObj.AuthPort = ":9000"
	clientPortObj.PostPort = ":9100"
	clientPortObj.UserPort = ":9200"
}

func DialToServers() {

	var authenticationConnection *grpc.ClientConn
	var postConnection *grpc.ClientConn
	var userConnection *grpc.ClientConn

	authenticationConnection, _ = grpc.Dial(clientPortObj.AuthPort, grpc.WithInsecure())
	postConnection, _ = grpc.Dial(clientPortObj.PostPort, grpc.WithInsecure())
	userConnection, _ = grpc.Dial(clientPortObj.UserPort, grpc.WithInsecure())

	clientObj.TokenClient = tokenActions_pb.NewTokenActionsServiceClient(authenticationConnection)
	clientObj.PostClient = post_pb.NewPostsServiceClient(postConnection)
	clientObj.UserClient = user_pb.NewUserServiceClient(userConnection)
}

func Initialize() {
	// initialize client port
	InitializeClientPorts()
	// initialize client connections
	DialToServers()
}
