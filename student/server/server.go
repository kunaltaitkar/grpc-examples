package main

import (
	"context"
	"grpc-examples/student/studentpb"
	"log"
	"net"
	"strconv"
	"time"

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

//unary request
func (*server) RegisterStudent(ctx context.Context, req *studentpb.RegisterStudentRequest) (*studentpb.RegisterStudentResponse, error) {

	firstName := req.GetStudent().GetFirstName()
	lastName := req.GetStudent().GetLastName()

	return &studentpb.RegisterStudentResponse{Result: "hello " + firstName + " " + lastName}, nil
}

//server streaming
func (*server) GetStudents(req *studentpb.GetStudentListRequest, stream studentpb.StudentService_GetStudentsServer) error {

	collegeName := req.GetCollegeName()

	for i := 0; i < 10; i++ {
		response := &studentpb.GetStudentListResponse{
			CollegeName: collegeName,
			StudentName: "Student " + strconv.Itoa(i+1),
		}
		stream.Send(response)
		time.Sleep(1 * time.Second)
	}

	return nil
}
