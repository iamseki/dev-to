import http from 'k6/http';
import { sleep } from 'k6';

const SLEEP_DURATION = 0.1;

export let options = {
    stages: [
        { duration: "30s", target: 100 },
        { duration: "30s", target: 0 },
    ],
    thresholds: {
        http_req_duration: ['p(99)<100'] // 99% request must complete below 100ms
    }
}

const NGINX_LOCAL_URL = "http://localhost:3000"
const API_LOCAL_URL = "http://localhost:8080"

const BASE_URL = __ENV.API_URL === "NGINX" ? NGINX_LOCAL_URL : API_LOCAL_URL 
const HEADERS = { "Content-Type": "application/json" }

const randomInteger = (min, max) => Math.floor(Math.random() * (max - min + 1)) + min;

export default () => {
    http.get(`${BASE_URL}/fibonacci/${randomInteger(99999991,99999999)}`);
    sleep(SLEEP_DURATION);
}