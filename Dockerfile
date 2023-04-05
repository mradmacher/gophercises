FROM golang:1.20

RUN go install golang.org/x/tools/cmd/godoc@latest
