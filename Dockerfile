FROM golang:alpine
EXPOSE 8000

WORKDIR /build
ADD go.mod .
COPY . .

RUN go build -o main src/server/main.go

CMD ["./main"]