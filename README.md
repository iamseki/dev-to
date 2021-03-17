# Dev-to blog posts

<p align='center'>This repository contains all code base that I used to write my posts in <a href="https://dev.to/chseki">dev.to</a> community</p>

## Usage :books:

- I create the branchs with the name likely the **title** of the post

## Running :scroll:

`go test ./usecases/... -coverprofile cover.out`

We can also see the coverage results in an pretty interface with built in go tool:

- `go test ./usecases/... -coverprofile cover.out`
- `go tool cover -html cover.out`
