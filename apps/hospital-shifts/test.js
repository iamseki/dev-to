import http from 'k6/http';
import { check } from 'k6';
import { Counter } from 'k6/metrics';

const BASE_URL = 'http://localhost:9092'; // Replace with your actual base URL

let errors = new Counter('errors');
let iterationCount = new Counter('iterations');

let writeSkewShiftId1 = new Counter('write_skew_shift_id_1');
let writeSkewShiftId2 = new Counter('write_skew_shift_id_2');
let writeSkewShiftId3 = new Counter('write_skew_shift_id_3');

export const options = {
  vus: 1, // Enforce only one user
  duration: '5s', // Duration of the test
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

const getRequestsForReadCommittedIsolation = () => {
  const payloadsForSerialazible = [
    JSON.stringify({
      shiftId: 3,
      doctorName: "Thamires",
      onCall: false
    }),
    JSON.stringify({
      shiftId: 3,
      doctorName: "Rafaella",
      onCall: false
    }),
  ]

  return payloadsForSerialazible.map(payload => {
    const params = {
      headers: { 'Content-Type': 'application/json', },
    };

    return {
      method: 'POST',
      url: `${BASE_URL}/update`,
      body: payload,
      params
    }
  });
}

export default function () {
  console.log(`\n--\n======== Iteration ID: ${__ITER} START ========`);
  iterationCount.add(1); // Increment the iteration counter

  const requestsForAdvisoryLock = getRequestsForAdvisoryLock();
  const requestsForSerialazibleIsolation = getRequestsForSerialazibleIsolation();
  const requestForReadCommittedIsolation = getRequestsForReadCommittedIsolation();

  const responses = http.batch([...requestsForAdvisoryLock, ...requestsForSerialazibleIsolation, ...requestForReadCommittedIsolation]);

  responses.forEach((res, idx) => {
    if (res.status !== 200) {
      errors.add(1);
      console.error(`Request failed: ${res.status} - ${res.body}`);
    }
  });

  const shiftsRes = http.get(`${BASE_URL}/shift`);
  const shifts = shiftsRes.json().sort((a, b) => a.shiftId - b.shiftId);

  // increment write skew with 1 for each shifts when write skew occurs
  if (shifts.filter(s => s.shiftId === 1 && s.onCall === false).length > 1) {
    writeSkewShiftId1.add(1);
  }

  if (shifts.filter(s => s.shiftId === 2 && s.onCall === false).length > 1) {
    writeSkewShiftId2.add(1);
  }

  if (shifts.filter(s => s.shiftId === 3 && s.onCall === false).length > 1) {
    writeSkewShiftId3.add(1);
  }

  console.log(`<------ Shifts Table ------>`);
  for (const shift of shifts) {
    console.log(`doctor: ${shift.doctorName}, shiftId: ${shift.shiftId}, onCall => ${shift.onCall}`);
  }
  console.log(`<------      -      ------>`);

  check(shifts, {
    'at least one doctor on call for shiftId=1': (shifts) => shifts.filter(s => s.shiftId === 1 && s.onCall === true).length > 0,
    'at least one doctor on call for shiftId=2': (shifts) => shifts.filter(s => s.shiftId === 2 && s.onCall === true).length > 0, 
    'at least one doctor on call for shiftId=3': (shifts) => shifts.filter(s => s.shiftId === 3 && s.onCall === true).length > 0, 
  });

  http.post(`${BASE_URL}/reset/shift`);

  console.log(`======== Iteration ID: ${__ITER} END ========`);
}
