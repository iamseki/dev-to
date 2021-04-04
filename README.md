<h1 align='center'>Dev.to community posts</h1>

<p align='center'>This repository contains all code base that I used to write my posts in <a href="https://dev.to/chseki">dev.to</a> community</p>

## Usage :books:

- I create the branches with the name similar to the **title** of the post

## Running apps :scroll:

- `docker-compose up -d`

## Endpoints :clipboard:

| App     | Endpoints              | Port    |
| ------- |:----------------------:| -------:|
| Golang  | `/hc`   `/fibonacci/n` | **8080**|
| Nodejs  | `/hc`   `/fibonacci/n` | **9090**|

## Running tests :muscle:

- Golang: `k6 run ./k6-test.js -e API_BASE=GOLANG`

- Nodejs: `k6 run ./k6-test.js`