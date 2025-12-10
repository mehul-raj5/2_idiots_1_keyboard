FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod init relay
RUN go mod tidy
RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
