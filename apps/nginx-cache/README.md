# Nginx Cache

- **[Blog post](https://dev.to/chseki/caching-with-nginx-2ob2)**

## Running the application :scroll:

- `yarn nx serve nginx-cache`

## Load Testing Comparison

Tests will fail if the 99th percentile latency exceeds 100ms.

- `yarn nx load-test nginx-cache`
- `yarn nx load-test-with-nginx nginx-cache`

## Endpoints :clipboard:

| App     | Endpoints              | Port    |
| ------- |:----------------------:| -------:|
| Api     |   `/fibonacci/n`       | **8080**|
| Nginx   |   `/fibonacci/n`       | **3000**|