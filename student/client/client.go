package main

import (
	"fmt"
	"grpc-examples/student/studentpb"
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

	fmt.Println("%v", studentClient)

}
