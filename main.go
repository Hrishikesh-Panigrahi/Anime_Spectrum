package main

import (
	"fmt"

	"github.com/Hrishikesh-Panigrahi/Anime_Spectrum/server"
)

func main() {
	/*
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

		log.Println("Serving API on http://127.0.0.1:8080")
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatalf("failed starting http2 server: %v", err)
		}
	*/

	server.RunServer()

	fmt.Println("Hello World")
}
