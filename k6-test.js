import http from 'k6/http';
import { sleep } from 'k6';

const SLEEP_DURATION = 0.1;

export let options = {
    stages: [
        { duration: "1m", target: 100 },
        { duration: "2m", target: 100 },
        { duration: "1m", target: 0 }
    ],
    thresholds: {
        http_req_duration: ['p(95)<1000'] // 99% request must complete below 1s
    }
}

const BASE_URL = __ENV.API_BASE === "GOLANG" ? "http://localhost:8080" : "http://localhost:9090" 
const HEADERS = { "Content-Type": "application/json" }

export default () => {
    http.get(`${BASE_URL}/fibonacci/9999999`);
    sleep(SLEEP_DURATION);
}