# Anime_Spectrum

**Anime_Spectrum** is an anime recommendation system built using **gRPC** and **Go**. It allows clients to query for anime suggestions based on genres, limits, and other filters.

## Features

- **gRPC API** for requesting anime suggestions.
- Support for filters like genre, rating, and reviews.
- Built with middleware to handle **gRPC-Web** requests for frontend compatibility.
- Easily extensible to use databases or external APIs for dynamic anime data.

## Project Structure
```
Anime_Spectrum/
├── client/
├── middleware/
├── proto/
├── server/
├── main.go
└── go.mod
```

## How It Works

The **Anime_Spectrum** project consists of:

1. **Client**: Sends requests to the server for anime recommendations. You can specify filters such as genre, limit, and sorting options.
2. **Middleware**: Provides support for **gRPC-Web**, making it possible for the **React frontend** to communicate with the **gRPC backend** over HTTP.
3. **Proto Definitions**: Contains the **.proto** file defining the gRPC service and message structure for requesting and receiving anime suggestions.
4. **Server**: Implements the gRPC server logic that processes client requests and returns anime suggestions. Currently, the server uses a hardcoded anime list, but it can be extended to use a database or external API.

## Installation and Setup

### Prerequisites

- **Go** (1.23+)
- **gRPC** and **gRPC-Web** libraries
- **Protobuf Compiler (protoc)**
- **MySQL** (optional if using a database)

### Steps

1. **Clone the repository:**
```bash
   git clone https://github.com/Hrishikesh-Panigrahi/Anime_Spectrum.git
   cd Anime_Spectrum
```

2. **Install Go dependencies:**
```bash
   go mod tidy
```

3. **Generate gRPC code from proto file: Make sure you have the protoc compiler installed, then run:**
```bash
   protoc --go_out=. --go-grpc_out=. proto/anime.proto
```

4. **Run the gRPC Server: Navigate to the root directory and run:**
```bash
   go run main.go
```
The gRPC server will start on localhost:8080.

5. **Run the Client: In a separate terminal, navigate to the client/ directory and run:**
```bash
   go run main.go
```
This will make a gRPC request to the server and print anime suggestions.

