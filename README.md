<h1 align='center'>Dev.to community posts</h1>

<p align='center'>This repository contains all code base that I used to write my posts in <a href="https://dev.to/chseki">dev.to</a> community</p>

## Usage :books:

- I create the branches with the name similar to the **title** of the post

## Running :scroll:

`go test ./usecases/...`

We can also see the coverage results in a pretty interface with built in go tool:

- `go test ./usecases/... -coverprofile cover.out`
- `go tool cover -html cover.out`
