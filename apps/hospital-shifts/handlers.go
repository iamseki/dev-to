package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func updateWithDefaultIsolation(c echo.Context, db *sqlx.DB) error {
	ds := &DoctorShift{}
	if err := c.Bind(&ds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	tx, err := db.Beginx()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer tx.Rollback()

	// Check the current number of doctors on call for this shift
	onCallCount := 0
	err = tx.Get(&onCallCount, "SELECT COUNT(*) FROM shifts WHERE shift_id = $1 AND on_call = TRUE", ds.ShiftID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Logic to ensure at least one doctor is always on call
	if ds.OnCall == false && onCallCount == 1 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf(`[ReadCommittedIsolation] Cannot set on_call to FALSE. At least one doctor must be on call for this shiftId: %v.`, ds.ShiftID)})
	}

	_, err = tx.Exec(`UPDATE shifts SET on_call = $1 WHERE shift_id = $2 AND doctor_name = $3`, ds.OnCall, ds.ShiftID, ds.DoctorName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func updateWithAdvisoryLock(c echo.Context, db *sqlx.DB) error {
	ds := &DoctorShift{}
	if err := c.Bind(&ds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	c.Logger().Infof("handle update with advisory lock for %v \n", ds)

	_, err := db.Exec(`SELECT update_on_call_status_with_advisory_lock($1, $2, $3)`, ds.ShiftID, ds.DoctorName, ds.OnCall)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func updateWithSerializableIsolation(c echo.Context, db *sqlx.DB) error {
	ds := &DoctorShift{}
	if err := c.Bind(&ds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	c.Logger().Infof("handle update with serializable isolation for %v \n", ds)

	// Init transaction with serializable isolation level
	tx, err := db.BeginTxx(c.Request().Context(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	defer tx.Rollback()

	_, err = tx.Exec(`SELECT update_on_call_status_with_serializable_isolation($1, $2, $3)`, ds.ShiftID, ds.DoctorName, ds.OnCall)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func resetShifts(c echo.Context, db *sqlx.DB) error {
	c.Logger().Info("handle reset shifts")

	_, err := db.Exec(`UPDATE shifts SET on_call = true`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func listShifts(c echo.Context, db *sqlx.DB) error {
	doctors := &[]DoctorShift{}
	err := db.Select(doctors, "SELECT * from shifts")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, doctors)
}
