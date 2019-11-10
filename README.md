# Simple-Chat-Server

## Introduction
This is a simple demo of Golang chat application. Client can send and receive message from other client go through a central server.

## How to use
### Start a server
```bash
cd server
go run server.go
```

### Start multiple clients
*You can open clients as many as you can.
```bash
cd client
go run client.go
```

### Send message
To send a message to other clients.
You need to enter the correct command from client CUI to specify the opponent id and message. See the information from the cmd.
eg:
-t id -m message
