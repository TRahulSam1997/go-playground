package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/TRahulSam1997/go-playground/grpc/proto"
)

const (
	address = "localhost:50051"
)

func main() {
 conn, err := grpc.NewClient(address)
 if err != nil {
  log.Fatalf("Did not connect: %v", err)
 }
 defer conn.Close()

 client := pb.NewTaskServiceClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()

 createResp, err := client.CreateTask(ctx, &pb.CreateTaskRequest{
  Title:       "Complete gRPC tutorial",
  Description: "Finish writing the gRPC tutorial with Go and MongoDB",
  DueDate:     timestamppb.New(time.Now().Add(24 * time.Hour)),
 })
 if err != nil {
  log.Fatalf("Could not create task: %v", err)
 }
 fmt.Printf("Created task: %v\n", createResp)

 getResp, err := client.GetTask(ctx, &pb.GetTaskRequest{Id: createResp.Id})
 if err != nil {
  log.Fatalf("Could not get task: %v", err)
 }
 fmt.Printf("Got task: %v\n", getResp)

 updateResp, err := client.UpdateTask(ctx, &pb.UpdateTaskRequest{
  Id:          createResp.Id,
  Title:       "Complete gRPC tutorial (updated)",
  Description: "Finish writing and reviewing the gRPC tutorial with Go and MongoDB",
  Completed:   true,
  DueDate:     timestamppb.New(time.Now().Add(48 * time.Hour)),
 })
 if err != nil {
  log.Fatalf("Could not update task: %v", err)
 }
 fmt.Printf("Updated task: %v\n", updateResp)

 listResp, err := client.ListTasks(ctx, &pb.ListTasksRequest{
  Page:     1,
  PageSize: 10,
 })
 if err != nil {
  log.Fatalf("Could not list tasks: %v", err)
 }
 fmt.Printf("Listed %d tasks, total count: %d\n", len(listResp.Tasks), listResp.TotalCount)
 for _, task := range listResp.Tasks {
  fmt.Printf("- %s: %s\n", task.Id, task.Title)
 }

 deleteResp, err := client.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: createResp.Id})
 if err != nil {
  log.Fatalf("Could not delete task: %v", err)
 }
 fmt.Printf("Deleted task, success: %v\n", deleteResp.Success)
}