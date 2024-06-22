# Load test with javascript

This project demonstrates how to use K6 for load testing simple endpoints that calculate the Fibonacci sequence. No optimizations were made, and Golang outperformed the competition since it's a CPU-bound task and the Node.js server doesn't spawn threads. Checkout this blog post for more details: **[Blog post on Load Testing with javascript](https://dev.to/chseki/load-test-with-javascript-1f64)**.

## Running the applications

- To run both apps execute `yarn nx serve javascript-load-testing`

## Load Testing Comparison

- `yarn nx load-test-golang javascript-load-testing` 
- `yarn nx load-test-nodejs javascript-load-testing`

## Endpoints

| App     | Endpoints              | Port    |
| ------- |:----------------------:| -------:|
| Golang  | `/hc`   `/fibonacci/n` | **8080**|
| Nodejs  | `/hc`   `/fibonacci/n` | **9090**|

---

### Technologies Used

- [K6](https://k6.io/open-source/)
- [Golang](https://go.dev/)
- [Nodejs](https://nodejs.org)
