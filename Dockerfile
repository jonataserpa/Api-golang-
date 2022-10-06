FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o server

ENTRYPOINT ["/app/server"]