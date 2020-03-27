package main

import (
	"context"
	"fmt"
	"grpc-examples/student/studentpb"
	"io"
	"log"
	"time"

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

	clientStreaming(studentClient)

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

func clientStreaming(client studentpb.StudentServiceClient) {
	payload := []studentpb.SendStudentDataRequest{
		studentpb.SendStudentDataRequest{
			Student: &studentpb.Student{
				FirstName: "Kunal",
				LastName:  "Taitkar",
			},
		},
		studentpb.SendStudentDataRequest{
			Student: &studentpb.Student{
				FirstName: "Aditya",
				LastName:  "Taitkar",
			},
		},
		studentpb.SendStudentDataRequest{
			Student: &studentpb.Student{
				FirstName: "Saurabh",
				LastName:  "Nimbarte",
			},
		},
		studentpb.SendStudentDataRequest{
			Student: &studentpb.Student{
				FirstName: "Shubham",
				LastName:  "Londase",
			},
		},
		studentpb.SendStudentDataRequest{
			Student: &studentpb.Student{
				FirstName: "Ishan",
				LastName:  "Jamjare",
			},
		},
	}

	stream, err := client.SendStudentData(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(payload); i++ {

		fmt.Printf("\nSending %v", payload[i])

		err = stream.Send(&payload[i])
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second)

	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCLIENT STEAMING RESPONSE:%v", response)
}
