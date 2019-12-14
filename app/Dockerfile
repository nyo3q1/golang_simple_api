FROM golang:1.13

WORKDIR /app
COPY . .
RUN go build -o simple_api

EXPOSE 8080
CMD ["./simple_api"]