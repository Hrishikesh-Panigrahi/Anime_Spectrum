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

	req := &api.AnimeRequest{Genre: "action", Limit: 3}
	res, err := client.GetAnimeSuggestions(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling GetAnimeSuggestions: %v", err)
	}

	fmt.Println("Anime suggestions:", res.AnimeTitles)
}
