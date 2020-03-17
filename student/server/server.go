package main

import (
	"grpc-examples/student/studentpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// server is used to implement studentpb.UnimplementedStudentServiceServer.
type server struct {
	studentpb.UnimplementedStudentServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
