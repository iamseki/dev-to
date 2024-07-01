CREATE TABLE shifts (
    id SERIAL PRIMARY KEY,
    doctor_name TEXT NOT NULL,
    shift_id INTEGER NOT NULL,
    on_call BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO shifts (shift_id, doctor_name, on_call) VALUES (1, 'Bob', TRUE);
INSERT INTO shifts (shift_id, doctor_name, on_call) VALUES (1, 'Alice', TRUE);

INSERT INTO shifts (shift_id, doctor_name, on_call) VALUES (2, 'Jack', TRUE);
INSERT INTO shifts (shift_id, doctor_name, on_call) VALUES (2, 'John', TRUE);

INSERT INTO shifts (shift_id, doctor_name, on_call) VALUES (3, 'Thamires', TRUE);
INSERT INTO shifts (shift_id, doctor_name, on_call) VALUES (3, 'Rafaella', TRUE);

-- Function to Manage On Call Status with Advisory Locks
CREATE OR REPLACE FUNCTION update_on_call_status_with_advisory_lock(shift_id INT, doctor_name TEXT, on_call BOOLEAN)
RETURNS VOID AS $$
DECLARE
    lock_id BIGINT;
    on_call_count INT;
BEGIN
    lock_id := shift_id; -- Use shift_id as the lock ID for the shift-level lock
    PERFORM pg_advisory_lock(lock_id);

    -- Check the current number of doctors on call for this shift
    SELECT COUNT(*) INTO on_call_count FROM shifts WHERE shift_id = shift_id AND on_call = TRUE;

    IF on_call = FALSE AND on_call_count = 1 THEN
        RAISE EXCEPTION 'Cannot set on_call to FALSE. At least one doctor must be on call for this shift.';
    ELSE
        UPDATE shifts
        SET on_call = on_call
        WHERE shift_id = shift_id AND doctor_name = doctor_name;
    END IF;

    PERFORM pg_advisory_unlock(lock_id);
END;
$$ LANGUAGE plpgsql;

-- Function to Manage On Call Status with serializable snapshot isolation
CREATE OR REPLACE FUNCTION update_on_call_status_with_serializable_isolation(shift_id INT, doctor_name TEXT, on_call BOOLEAN)
RETURNS VOID AS $$
DECLARE
    on_call_count INT;
BEGIN
    -- Check the current number of doctors on call for this shift
    SELECT COUNT(*) INTO on_call_count FROM shifts WHERE shift_id = shift_id AND on_call = TRUE;

    IF on_call = FALSE AND on_call_count = 1 THEN
        RAISE EXCEPTION 'Cannot set on_call to FALSE. At least one doctor must be on call for this shift.';
    ELSE
        UPDATE shifts
        SET on_call = on_call
        WHERE shift_id = shift_id AND doctor_name = doctor_name;
    END IF;

END;
$$ LANGUAGE plpgsql;