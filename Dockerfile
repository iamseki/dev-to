FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM busybox

WORKDIR /app

COPY --from=builder /app/dev-to /usr/bin/

ENTRYPOINT ["dev-to"]