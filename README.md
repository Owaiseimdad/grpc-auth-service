# Go gRPC Authentication & Authorization Service

## Description

This project is a microservice built using **gRPC** and **Go** for **Authentication** and **Authorization**. It provides a simple, fast, and scalable way to authenticate users by checking their credentials (username/password) and returning a token upon successful authentication.

The service can be extended to integrate with various authentication mechanisms, such as OAuth, JWT, or traditional username/password systems. The current implementation is a simple mock, where valid users get a hardcoded token.

This service is designed to be easy to integrate into your existing microservice architecture, making it ideal for securing APIs and ensuring user access control.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
   - [Running the Server](#running-the-server)
   - [Running the Client](#running-the-client)
3. [gRPC API](#grpc-api)
   - [Authentication Service](#authentication-service)
4. [Contributing](#contributing)
5. [License](#license)
6. [Roadmap](#roadmap)

## Installation

To get started with this project, follow these steps to install the dependencies and set up the project:

### Prerequisites

Make sure you have the following installed:

- **Go** (v1.18+): The programming language used for this service.
- **Protocol Buffers (protoc)**: A tool to compile `.proto` files into Go code.
- **gRPC Go Plugin**: A plugin for generating Go code from the protobuf definitions.

### Step 1: Clone the repository

```bash
git clone https://github.com/yourusername/authentication-service.git
cd authentication-service
```

## Usage

### Running the Server

Once you have set up the project and installed dependencies, you can run the **gRPC server** to start the authentication service.

1. **Start the server** by running the following command from the project root:

   ```bash
   go run cmd/server.go
   ```

2. The server will listen on port `50051`. You should see output similar to this indicating the server is running:

   ```bash
   Server is running on port 50051...
   ```

3. The server is now ready to receive authentication requests via gRPC.

### Running the Client

To test the service, you can run the **gRPC client**, which sends a mock authentication request to the server.

1. **Run the client** by using the following command:

   ```bash
   go run cmd/client.go
   ```

2. The client will send a predefined username and password to the server. If the credentials are valid, it will receive a mock authentication token and print it out:

   Example output:

   ```bash
   Received token: mock-token-for-user1
   ```

   If the credentials are incorrect, the client will receive an error message based on the server's response.

### Example Client Request

The client is pre-configured to send the following `AuthRequest`:

```go
req := &api.AuthRequest{
    Username: "user1",  // You can modify the username for testing
    Password: "password123",  // Modify the password for testing
}
```

## gRPC API

### Authentication Service

The `AuthService` is a gRPC service that provides user authentication via the `Authenticate` method.

#### Method: `Authenticate`

- **Input**:

  - `AuthRequest`:
    - `username` (string): The username of the user trying to authenticate.
    - `password` (string): The password of the user.

- **Output**:
  - `AuthResponse`:
    - `token` (string): The authentication token returned if the credentials are valid.

#### Example gRPC Call:

Here’s how to make an RPC call to the `Authenticate` method using a gRPC client:

```go
req := &api.AuthRequest{
    Username: "user1", // Enter your username
    Password: "password123", // Enter your password
}

resp, err := client.Authenticate(context.Background(), req)
if err != nil {
    log.Fatalf("could not authenticate: %v", err)
}
fmt.Printf("Received token: %s\n", resp.GetToken())
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
