# Hospital Shifts Concurrency Write Skew Example

TODO

- something is wrong with the endpoint, need to check

## Testing

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