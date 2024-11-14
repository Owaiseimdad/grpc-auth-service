// cmd/client.go
package main

import (
	"context"
	"fmt"
	"grpc-auth-service/api" // Import the generated Go files
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new AuthService client
	client := api.NewAuthServiceClient(conn)

	// Call the Authenticate method
	req := &api.AuthRequest{
		Username: "user1",
		Password: "password123",
	}

	resp, err := client.Authenticate(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			// Handle different status codes
			switch st.Code() {
			case codes.InvalidArgument:
				log.Printf("Invalid input: %s", st.Message())
			case codes.Unauthenticated:
				log.Printf("Authentication failed: %s", st.Message())
			default:
				log.Printf("Error occurred: %s", st.Message())
			}
		} else {
			log.Fatalf("could not authenticate: %v", err)
		}
		return
	}

	// Print the token received from the server
	fmt.Printf("Received token: %s\n", resp.GetToken())
}
