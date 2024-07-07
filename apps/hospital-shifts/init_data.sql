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
CREATE OR REPLACE FUNCTION update_on_call_status_with_advisory_lock(shift_id_to_update INT, doctor_name_to_update TEXT, on_call_to_update BOOLEAN)
RETURNS VOID AS $$
DECLARE
    on_call_count INT;
BEGIN
    -- Attempt to acquire advisory lock and handle failure with NOTICE
    IF NOT pg_try_advisory_lock(shift_id_to_update) THEN
        RAISE NOTICE 'Could not acquire advisory lock for shift_id: %', shift_id_to_update;
        RETURN;
    END IF;

    -- Check the current number of doctors on call for this shift
    SELECT COUNT(*) INTO on_call_count FROM shifts s WHERE s.shift_id = shift_id_to_update AND s.on_call = TRUE;

    IF on_call_to_update = FALSE AND on_call_count = 1 THEN
        RAISE EXCEPTION '[AdvisoryLock] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: %', shift_id_to_update;
    ELSE
        UPDATE shifts s
        SET on_call = on_call_to_update
        WHERE s.shift_id = shift_id_to_update AND s.doctor_name = doctor_name_to_update;
    END IF;

    PERFORM pg_advisory_unlock(shift_id_to_update);
END;
$$ LANGUAGE plpgsql;

-- Function to Manage On Call Status with serializable snapshot isolation
CREATE OR REPLACE FUNCTION update_on_call_status_with_serializable_isolation(shift_id_to_update INT, doctor_name_to_update TEXT, on_call_to_update BOOLEAN)
RETURNS VOID AS $$
DECLARE
    on_call_count INT;
BEGIN
    -- Check the current number of doctors on call for this shift
    SELECT COUNT(*) INTO on_call_count FROM shifts s WHERE s.shift_id = shift_id_to_update AND s.on_call = TRUE;

    IF on_call_to_update = FALSE AND on_call_count = 1 THEN
        RAISE EXCEPTION '[SerializableIsolation] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: %', shift_id_to_update;
    ELSE
        UPDATE shifts s
        SET on_call = on_call_to_update
        WHERE s.shift_id = shift_id_to_update AND s.doctor_name = doctor_name_to_update;
    END IF;

END;
$$ LANGUAGE plpgsql;