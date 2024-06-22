# Nginx Cache

This project demonstrates using Nginx for caching purposes, check out this blog post: **[Blog post on caching with Nginx](https://dev.to/chseki/caching-with-nginx-2ob2)**.

## Running the application

- `yarn nx serve nginx-cache`

## Load Testing Comparison

Tests will fail if the 99th percentile latency exceeds 100ms.

- `yarn nx load-test nginx-cache`
- `yarn nx load-test-with-nginx nginx-cache`

## Endpoints

| App     | Endpoints              | Port    |
| ------- |:----------------------:| -------:|
| Api     |   `/fibonacci/n`       | **8080**|
| Nginx   |   `/fibonacci/n`       | **3000**|

---

### Technologies Used

- [Docker](https://www.docker.com/)
- [K6](https://k6.io/open-source/)
- [Golang](https://go.dev/)
- [Nginx](https://www.f5.com/company/blog/nginx/nginx-caching-guide)
