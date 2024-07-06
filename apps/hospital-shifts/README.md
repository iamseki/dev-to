# Hospital Shifts Race Condition Example

This project demonstrates a race condition scenario that can happen with concurrent writes and how to fix it using PostgreSQL's [serialization isolation level](https://www.postgresql.org/docs/current/transaction-iso.html#XACT-SERIALIZABLE) or [advisory locks](https://www.postgresql.org/docs/current/explicit-locking.html#ADVISORY-LOCKS).

**Inspiration**: Designing Data-Intensive Applications, Chapter 7 - Transactions "Weak Isolation Levels"

## Project Requirement

We have a table that tracks which doctor is on call for a specific shift:
```sql
CREATE TABLE shifts (
    id SERIAL PRIMARY KEY,
    doctor_name TEXT NOT NULL,
    shift_id INTEGER NOT NULL,
    on_call BOOLEAN NOT NULL DEFAULT FALSE
);
```

The only requirement is:

- Each shift must *always* have **at least one** doctor on call.

## Race Condition Scenario

The race condition occurs when two doctors from the same shift try to go off call at nearly the same time. To reproduce:

1. Start the server: `yarn nx serve hospital-shifts`
2. Run the k6 test:  `yarn nx test hospital-shifts`

## Understanding the Test

The test simulates the "nearly the same time" scenario by sending concurrent requests to different endpoints. Since the issue is non-deterministic, the test runs many iterations to increase the chance of reproducing it. Endpoints Tested:

1. `/update-with-advisory`: This endpoint uses _advisory locks_ to prevent the race condition.
2. `update-with-serializable` This endpoint uses a _serializable transaction_ for isolation.
3. `/update` This is the naive implementation without any concurrency control which shows the problem.

Each iteration displays the current on-call status and any errors related to preventing the concurrency issue. 

***Pay attention to*** `shiftId 3`: This shift is updated by the endpoint without any concurrency protection. Sometimes, it violates the requirement of having one doctor on call, leaving the shift uncovered.

```sh
======== Iteration ID: 337 START ======== 
ERRO[0005] Request failed: 500 - {"error":"pq: [AdvisoryLock] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: 1"}
ERRO[0005] Request failed: 500 - {"error":"pq: [SerializableIsolation] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: 2"}
INFO[0005] <------ Shifts Table ------>                  
INFO[0005] doctor: Alice, shiftId: 1, onCall => true     
INFO[0005] doctor: Bob, shiftId: 1, onCall => false      
INFO[0005] doctor: John, shiftId: 2, onCall => true      
INFO[0005] doctor: Jack, shiftId: 2, onCall => false     
INFO[0005] doctor: Rafaella, shiftId: 3, onCall => false  
INFO[0005] doctor: Thamires, shiftId: 3, onCall => false  
INFO[0005] <------      -      ------>                   
INFO[0005] ======== Iteration ID: 337 END ========       
INFO[0005] 
--
======== Iteration ID: 338 START ========  
ERRO[0005] Request failed: 500 - {"error":"pq: [AdvisoryLock] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: 1"}  
ERRO[0005] Request failed: 500 - {"error":"pq: could not serialize access due to read/write dependencies among transactions"}  
ERRO[0005] Request failed: 500 - {"error":"[ReadCommittedIsolation] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: 3."}  
INFO[0005] <------ Shifts Table ------>                  
INFO[0005] doctor: Bob, shiftId: 1, onCall => true       
INFO[0005] doctor: Alice, shiftId: 1, onCall => false    
INFO[0005] doctor: Jack, shiftId: 2, onCall => true      
INFO[0005] doctor: John, shiftId: 2, onCall => false     
INFO[0005] doctor: Rafaella, shiftId: 3, onCall => true  
INFO[0005] doctor: Thamires, shiftId: 3, onCall => false  
INFO[0005] <------      -      ------>                   
INFO[0005] ======== Iteration ID: 338 END ========       

``` 


---

## Debugging Tips

Here are some curl commands to interact with the API:

```sh
curl -X POST http://localhost:9092/update-with-advisory \
     -H "Content-Type: application/json" \
     -d '{
           "shiftId": 1,
           "doctorName": "Bob",
           "onCall": false
         }'
```


```sh
curl -X POST http://localhost:9092/reset/shift
```