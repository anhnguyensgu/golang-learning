FROM golang:alpine

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./app ./main.go

# COPY /api/app.out .

ENTRYPOINT ["./app"]
