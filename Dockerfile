# Compiles into >15mb!!!!
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser
WORKDIR ~/Projects/Target
COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/hello /go/bin/hello
USER appuser
ENTRYPOINT ["/go/bin/hello"]