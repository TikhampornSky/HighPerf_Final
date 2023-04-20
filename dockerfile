FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy && go mod verify

RUN go build -o myapp

ENTRYPOINT ["./myapp"]