<h1 align='center'>Dev.to community posts</h1>

<p align='center'>This repository contains all code base that I used to write my posts in <a href="https://dev.to/chseki">dev.to</a> community</p>

## Usage :books:

- I create the branches with the name similar to the **title** of the post

## Running apps :scroll:

- `docker-compose up -d`
- `k6 run k6-test.js -e API_URL='NGINX'`
- `k6 run k6-test.js`

## Endpoints :clipboard:

| App     | Endpoints              | Port    |
| ------- |:----------------------:| -------:|
| Api     |   `/fibonacci/n`       | **8080**|
| Nginx   |   `/fibonacci/n`       | **3000**|