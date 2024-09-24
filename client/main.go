package main

import (
	"context"
	"fmt"
	"log"
	"time"

	api "github.com/Hrishikesh-Panigrahi/Anime_Spectrum/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewAnimeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.AnimeRequest{
		Genre:          "action",
		Limit:          3,
		SortBy:         "rating",
		IncludeRatings: true,
		IncludeReviews: true,
	}

	res, err := client.GetAnimeSuggestions(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling GetAnimeSuggestions: %v", err)
	}

	for _, anime := range res.AnimeDetails {
		fmt.Printf("Title: %s\n", anime.Title)
		fmt.Printf("Description: %s\n", anime.Description)
		fmt.Printf("Genre: %s\n", anime.Genre)
		fmt.Printf("Rating: %.1f\n", anime.Rating)
		fmt.Printf("Release Date: %s\n", anime.ReleaseDate)

		if len(anime.Reviews) > 0 {
			fmt.Println("Reviews:")
			for _, review := range anime.Reviews {
				fmt.Printf("- %s\n", review)
			}
		}
		fmt.Println("-----")
	}

}
