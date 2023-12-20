FROM golang:1.21

WORKDIR /app
COPY go.sum go.mod ./ 
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /hms-api

EXPOSE 8080

CMD ["/hms-api"]