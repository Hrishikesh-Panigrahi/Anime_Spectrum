package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sort"

	api "github.com/Hrishikesh-Panigrahi/Anime_Spectrum/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// AnimeServer implements the AnimeService gRPC server
type AnimeServer struct {
	api.UnimplementedAnimeServiceServer
}

type Anime struct {
	Title       string
	Description string
	Genre       string
	Rating      float32
	ReleaseDate string
	Reviews     []string
}

// GetAnimeSuggestions is the unary RPC to get anime suggestions based on the request
func (s *AnimeServer) GetAnimeSuggestions(ctx context.Context, req *api.AnimeRequest) (*api.AnimeResponse, error) {

	// Sample anime data
	animeDB := []Anime{
		{Title: "Naruto", Genre: "action", Rating: 8.3, ReleaseDate: "2002", Reviews: []string{"Epic", "Great Story"}},
		{Title: "One Piece", Genre: "action", Rating: 9.0, ReleaseDate: "1999", Reviews: []string{"Adventure packed", "Exciting"}},
		{Title: "Attack on Titan", Genre: "action", Rating: 9.1, ReleaseDate: "2013", Reviews: []string{"Mind-blowing", "Intense"}},
		{Title: "My Hero Academia", Genre: "action", Rating: 7.9, ReleaseDate: "2016", Reviews: []string{"Fun", "Inspiring"}},
		{Title: "Your Lie in April", Genre: "romance", Rating: 8.7, ReleaseDate: "2014", Reviews: []string{"Heartbreaking", "Beautiful"}},
		{Title: "Steins;Gate", Genre: "sci-fi", Rating: 9.1, ReleaseDate: "2011", Reviews: []string{"Mind-bending", "Sci-fi masterpiece"}},
		{Title: "Death Note", Genre: "thriller", Rating: 9.0, ReleaseDate: "2006", Reviews: []string{"Thrilling", "Psychological battle"}},
		{Title: "Fullmetal Alchemist: Brotherhood", Genre: "action", Rating: 9.2, ReleaseDate: "2009", Reviews: []string{"Masterpiece", "Excellent characters"}},
		{Title: "Clannad", Genre: "romance", Rating: 8.5, ReleaseDate: "2007", Reviews: []string{"Emotional", "Heartwarming"}},
		{Title: "Cowboy Bebop", Genre: "sci-fi", Rating: 8.9, ReleaseDate: "1998", Reviews: []string{"Cool and stylish", "Timeless"}},
	}

	var filteredAnime []Anime
	for _, anime := range animeDB {
		if req.Genre == "" || anime.Genre == req.Genre {
			filteredAnime = append(filteredAnime, anime)
		}
	}

	switch req.SortBy {
	case "rating":
		sort.Slice(filteredAnime, func(i, j int) bool {
			return filteredAnime[i].Rating > filteredAnime[j].Rating
		})
	case "release_date":
		sort.Slice(filteredAnime, func(i, j int) bool {
			return filteredAnime[i].ReleaseDate > filteredAnime[j].ReleaseDate
		})
	}

	if req.Limit > int32(len(filteredAnime)) {
		req.Limit = int32(len(filteredAnime))
	}
	filteredAnime = filteredAnime[:req.Limit]

	var animeDetails []*api.AnimeDetail
	for _, anime := range filteredAnime {
		animeDetail := &api.AnimeDetail{
			Title:       anime.Title,
			Description: anime.Description,
			Genre:       anime.Genre,
			ReleaseDate: anime.ReleaseDate,
		}

		if req.IncludeRatings {
			animeDetail.Rating = anime.Rating
		}
		if req.IncludeReviews {
			animeDetail.Reviews = anime.Reviews
		}

		animeDetails = append(animeDetails, animeDetail)
	}

	res := &api.AnimeResponse{
		AnimeDetails: animeDetails,
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
