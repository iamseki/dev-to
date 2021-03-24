<h1 align='center'>Dev.to community posts</h1>

<p align='center'>This repository contains all code base that I used to write my posts in <a href="https://dev.to/chseki">dev.to</a> community</p>

## Usage :books:

- I create the branches with the name similar to the **title** of the post

## Testing :scroll:

`go test ./usecases/...`

We can also see the coverage results in a pretty interface with built in go tool:

- `go test ./... -coverprofile cover.out`
- `go tool cover -html cover.out`

## Build and Run :whale:

- `docker build -t dev-to .`
- `docker run --name go-web -p 8080:8080 -d dev-to`

---

The web server expose only one route and we can make a `HTTP GET` request with _curl_ as follow:

```sh
curl localhost:8080/events
```

