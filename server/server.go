package server

import (
	"context"
	"fmt"
	"log"
	"net"

	api "github.com/Hrishikesh-Panigrahi/Anime_Spectrum/proto"
	"google.golang.org/grpc"
)

const(
	port = ":8080"
)

type AnimeServer struct {
	api.UnimplementedAnimeServiceServer
}

func (s *AnimeServer) GetAnimeSuggestions(ctx context.Context, req *api.AnimeRequest) (*api.AnimeResponse, error) {

	animeList := []string{"Naruto", "One Piece", "Attack on Titan"}

	if req.Genre == "action" {
		animeList = append(animeList, "My Hero Academia")
	} else if req.Genre == "romance" {
		animeList = append(animeList, "Your Lie in April")
	}

	res := &api.AnimeResponse{
		AnimeTitles: animeList[:req.Limit],
	}

	return res, nil
}

// RunServer starts the gRPC server.
func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterAnimeServiceServer(grpcServer, &AnimeServer{})

	fmt.Println("gRPC server is running on port 8080...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
