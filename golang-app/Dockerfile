FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o http-iris .

FROM scratch

WORKDIR /app

COPY --from=builder /app/http-iris /usr/bin/

ENTRYPOINT ["http-iris"]
