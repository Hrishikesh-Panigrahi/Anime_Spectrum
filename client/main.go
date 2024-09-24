package main

import (
	"context"
	"fmt"
	"log"
	"strings"
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

	fmt.Println("======== Anime Suggestions ========")
	for _, anime := range res.AnimeDetails {
		fmt.Printf("Title: %s\n", anime.Title)
		fmt.Printf("Genre: %s\n", anime.Genre)
		fmt.Printf("Rating: %.1f\n", anime.Rating)
		fmt.Printf("Release Date: %s\n", anime.ReleaseDate)
		fmt.Printf("Description: %s\n", anime.Description)

		if len(anime.Reviews) > 0 {
			fmt.Print("Reviews:")
			fmt.Printf("%s\n", strings.Join(anime.Reviews, "\n \t"))
		}
		fmt.Println("-----------------------------------")
	}
	fmt.Println("===================================")

}
