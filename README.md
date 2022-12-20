# Distributed Systems Project
Authors:
Harsh Apte (ha2179)
Manan Chawda (mrc9419)
Note: Open the src folder as the root of the project


A. Raft Setup on Local:

1. Clone CoreOS raft into local package
cd src/github.com
mkdir go.etcd.io
cd go.etcd.io
git clone https://github.com/etcd-io/etcd.git

2. Package Resolution from Module Graph of Go
go mod tidy

3. Install goreman
cd github.com/go.etcd.io/etcd/contrib/raftexample
go get github.com/mattn/goreman

4. Build raft Example
go build -o raftexample

4. Run goreman
goreman start

B. Run Local Services

1. Run TokenActionService
cd web/services/tokenActions/tokenActions_service
go run tokenActionService.go

2. Run UserService
cd web/services/user/user_service
go run userService.go

3. Run Post Service
cd web/services/post/post_service
go run postService.go

4. Run UI(Web) Service
cd cmd/web
go run web.go

Verify the application is running on : http://localhost:8000/login/

C. Run the test cases by running the following (as one command): go run mainTest.go registerLoginTest.go followUnfollowTest.go postTest.go
