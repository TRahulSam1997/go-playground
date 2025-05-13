package main

import (
	"context"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pb "github.com/TRahulSam1997/go-playground/grpc/proto"
	"github.com/TRahulSam1997/go-playground/grpc/proto/server/taskservice"
)

const (
	port = ":50051"
	mongoURI = "mongodb://localhost:27017"
	dbName = "taskdb"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect to MongoDB: %v", err)
		}
	}()

	err = client.Ping(ctx, nil) 
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	taskCollection := client.Database(dbName).Collection("tasks")
	taskService := taskservice.NewTaskService(taskCollection)

	s := grpc.NewServer()

	pb.RegisterTaskServiceServer(s, taskService)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listeninmg on %s", port)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}