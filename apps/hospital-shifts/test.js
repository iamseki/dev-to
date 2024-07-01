import http from 'k6/http';
import { check } from 'k6';
import { Counter } from 'k6/metrics';

const BASE_URL = 'http://localhost:9092'; // Replace with your actual base URL

let errors = new Counter('errors');
let iterationCount = new Counter('iterations');

export const options = {
  vus: 1, // Enforce only one user
  duration: '30s', // Duration of the test
};

const getRequestsForAdvisoryLock = () => {
  const payloadsForAdvisoryLock = [
    JSON.stringify({
      shiftId: 1,
      doctorName: "Bob",
      onCall: false
    }),
    JSON.stringify({
      shiftId: 1,
      doctorName: "Alice",
      onCall: false
    }),
  ]

  return payloadsForAdvisoryLock.map(payload => {
    const params = {
      headers: { 'Content-Type': 'application/json', },
    };

    return {
      method: 'POST',
      url: `${BASE_URL}/update-with-advisory`,
      body: payload,
      params
    }
  });
}

const getRequestsForSerialazibleIsolation = () => {
  const payloadsForSerialazible = [
    JSON.stringify({
      shiftId: 2,
      doctorName: "Jack",
      onCall: false
    }),
    JSON.stringify({
      shiftId: 2,
      doctorName: "John",
      onCall: false
    }),
  ]

  return payloadsForSerialazible.map(payload => {
    const params = {
      headers: { 'Content-Type': 'application/json', },
    };

    return {
      method: 'POST',
      url: `${BASE_URL}/update-with-serializable`,
      body: payload,
      params
    }
  });
}

export default function () {
  iterationCount.add(1); // Increment the iteration counter

  const requestsForAdvisoryLock = getRequestsForAdvisoryLock();
  const requestsForSerialazibleIsolation = getRequestsForSerialazibleIsolation();

  const responses = http.batch([...requestsForAdvisoryLock, ...requestsForSerialazibleIsolation]);

  responses.forEach((res, idx) => {
    const success = check(res, {
      'status is 200': (r) => r.status === 200,
      'response contains success': (r) => r.json().status === 'success',
    });

    if (!success) {
      errors.add(1);
      console.error(`Request failed: ${res.status} - ${res.body}`);
    }
  });

  const resetResponse = http.post(`${BASE_URL}/reset/shift`);

  check(resetResponse, {
    'reset status is 200': (r) => r.status === 200,
    'reset response contains success': (r) => r.json().status === 'success',
  });

  if (resetResponse.status !== 200) {
    errors.add(1);
    console.error(`Reset request failed: ${resetResponse.status} - ${resetResponse.body}`);
  }
}
