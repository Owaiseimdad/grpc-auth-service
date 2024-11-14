// cmd/server.go
package main

import (
	"context"
	"fmt"
	"grpc-auth-service/api"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	api.UnimplementedAuthServiceServer
}

// Implement the Authenticate method with error handling and logging
func (s *server) Authenticate(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	log.Printf("Received authentication request for username: %s", req.GetUsername()) // Log incoming request

	// Simple validation
	if req.GetUsername() == "" || req.GetPassword() == "" {
		log.Println("Invalid request: username or password is empty") // Log error
		return nil, status.Error(codes.InvalidArgument, "Username and password cannot be empty")
	}

	// Simulate credential validation
	if req.GetUsername() != "user1" || req.GetPassword() != "password123" {
		log.Printf("Failed authentication attempt for username: %s", req.GetUsername()) // Log failed login
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}

	// Success, generate token
	token := "mock-token-for-" + req.GetUsername()
	log.Printf("Authentication successful for username: %s, token: %s", req.GetUsername(), token) // Log success

	// Return the response with the generated token
	return &api.AuthResponse{
		Token: token,
	}, nil
}

func main() {
	// Set up the listener on port 50051
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the AuthService with the server
	api.RegisterAuthServiceServer(grpcServer, &server{})

	// Start the server and block until it’s done
	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}