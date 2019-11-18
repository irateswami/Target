FROM golang:1.13.4

LABEL maintainer="Bryan English <bryanenglish@protonmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]