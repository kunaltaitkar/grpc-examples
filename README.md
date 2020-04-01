# **grpc-examples**
This repo contains examples of all types of grpc api.

- **Unary**
  - The client sends a single request and gets back a single response.
- **Server Streaming**
  - The server returns a stream of messages in response to a client’s request
- **Client Streaming**
  - The client sends a stream of messages to the server instead of a single message
- **BiDirectional Streaming**
  -  The two streams are independent, the client and server can read and write messages in any order.


## **How to run?**

1. >go run ./student/server/server.go

2. >go run ./student/client/client.go


## **Output**

![](output.gif)