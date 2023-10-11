package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/crossplane/provider-userprovider/grpc-server/proto/gen/go/userapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	userapi.UnimplementedUserServiceServer
	cache sync.Map // concurrent map to store User entities in-memory
}

func (s *server) CreateUser(ctx context.Context, user *userapi.User) (*userapi.UserResponse, error) {
	log.Printf("CreateUser request for ID: %s\n", user.Id)

	_, found := s.cache.Load(user.Id)
	if found {
		err := status.Errorf(
			codes.AlreadyExists,
			"User with ID '%s' already exists.",
			user.Id,
		)
		log.Println(err)
		return nil, err
	}
	s.cache.Store(user.Id, user)

	log.Printf("User with ID: %s created successfully\n", user.Id)
	return &userapi.UserResponse{
		Status:  "Success",
		Message: "User created successfully",
		User:    user,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *userapi.GetRequest) (*userapi.User, error) {
	log.Printf("GetUser request for ID: %s\n", req.Id)

	value, found := s.cache.Load(req.Id)
	if !found {
		err := status.Errorf(
			codes.NotFound,
			"User with ID '%s' not found.",
			req.Id,
		)
		log.Println(err)
		return nil, err
	}
	log.Printf("User with ID: %s retrieved successfully\n", req.Id)
	return value.(*userapi.User), nil
}

func (s *server) UpdateUser(ctx context.Context, user *userapi.User) (*userapi.UserResponse, error) {
	log.Printf("UpdateUser request for ID: %s\n", user.Id)

	_, found := s.cache.Load(user.Id)
	if !found {
		err := status.Errorf(
			codes.NotFound,
			"User with ID '%s' not found.",
			user.Id,
		)
		log.Println(err)
		return nil, err
	}
	s.cache.Store(user.Id, user)

	log.Printf("User with ID: %s updated successfully\n", user.Id)
	return &userapi.UserResponse{
		Status:  "Success",
		Message: "User updated successfully",
		User:    user,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *userapi.GetRequest) (*userapi.UserResponse, error) {
	log.Printf("DeleteUser request for ID: %s\n", req.Id)

	_, found := s.cache.Load(req.Id)
	if !found {
		err := status.Errorf(
			codes.NotFound,
			"User with ID '%s' not found.",
			req.Id,
		)
		log.Println(err)
		return nil, err
	}
	s.cache.Delete(req.Id)

	log.Printf("User with ID: %s deleted successfully\n", req.Id)
	return &userapi.UserResponse{
		Status:  "Success",
		Message: "User deleted successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userapi.RegisterUserServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
