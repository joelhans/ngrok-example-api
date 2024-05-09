# API Gateway gallery

This repo contains two assets:

- A simple Go app that creates a simple API with a single route and connects it to an existing ngrok Edge.
- A `policies/` folder of YAML-based policies for the ngrok API gateway.

To run the Go-based API:

```
mkdir legends-api
cd legends-api
go mod init legends-api
go get golang.ngrok.com/ngrok github.com/gorilla/mux
```

Paste your ngrok authtoken and the label for your Edge into the command below:

```
NGROK_AUTHTOKEN=<YOUR-NGROK-AUTHTOKEN> NGROK_LABEL=<YOUR-EDGE-LABEL> go run main.go
```
