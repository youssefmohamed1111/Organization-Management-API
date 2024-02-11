# use official Golang image
FROM golang:alpine

WORKDIR /app

# Copy everything from the source code
COPY . .

RUN go mod download

RUN go build -o my-app

EXPOSE 8000

CMD ["./api"]
