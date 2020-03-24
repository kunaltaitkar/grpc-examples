package main

import (
	"context"
	"fmt"
	"grpc-examples/student/studentpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	studentClient := studentpb.NewStudentServiceClient(conn)

	unaryRequest(studentClient)

	serverStreamingRequest(studentClient)

}

func unaryRequest(client studentpb.StudentServiceClient) {

	payload := &studentpb.RegisterStudentRequest{
		Student: &studentpb.Student{
			FirstName: "Kunal",
			LastName:  "Taitkar",
		},
	}

	resp, err := client.RegisterStudent(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UNARY RESPONSE:")
	fmt.Println(resp)

}

func serverStreamingRequest(client studentpb.StudentServiceClient) {

	payload := &studentpb.GetStudentListRequest{
		CollegeName: "GHRCE",
	}

	resp, err := client.GetStudents(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("STEAMING RESPONSE")

	//read stream
	for {

		response, err := resp.Recv()

		if err == io.EOF {
			//reached end of stream
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(response)
	}

}
