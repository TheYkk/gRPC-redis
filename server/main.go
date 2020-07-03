package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "server/redisimport"
)
var rdb *redis.Client
func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server started on port 50051")
	grpcServer := grpc.NewServer()

	pb.RegisterRedisImportServer(grpcServer,new(userService))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}


}

type userService struct {
	users []*pb.User
	m     sync.Mutex
}
var ctx = context.Background()

func (u userService) Import(ctx context.Context, user *pb.User) (*pb.ImportReply, error) {
	u.m.Lock()
	defer u.m.Unlock()
	//log.Print(user)

	resp := new( pb.ImportReply)
	json, err := json.Marshal(user)
	if err != nil {
		resp.Message="Error"
		return resp, err
	}

	err = rdb.Set(ctx, "users:"+user.UserName, json,0).Err()
	if err != nil {
		panic(err)
	}
	resp.Message=fmt.Sprintf("User %v succesfully added",user.UserName)

	return resp,nil
}

