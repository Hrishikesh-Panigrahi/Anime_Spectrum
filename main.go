package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/Anime_Spectrum/middleware"
	api "github.com/Hrishikesh-Panigrahi/Anime_Spectrum/proto"
	"github.com/Hrishikesh-Panigrahi/Anime_Spectrum/server"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	api.RegisterAnimeServiceServer(grpcServer, &server.AnimeServer{})

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		// Allow all origins.
		return true
	}))

	router := chi.NewRouter()

	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		middleware.NewGrpcWebMiddleware(wrappedGrpc).Handler,
	)

	// router.Get("/article-proxy", proxy.Article)

	log.Println("Serving API on http://127.0.0.1:8080")

	
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed starting http2 server: %v", err)
	}

	fmt.Println("Hello World")
}
