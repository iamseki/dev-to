package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type DoctorShift struct {
	ID         int    `json:"id" db:"id"`
	ShiftID    int    `json:"shiftId" db:"shift_id"`
	DoctorName string `json:"doctorName" db:"doctor_name"`
	OnCall     bool   `json:"onCall" db:"on_call"`
}

func main() {
	e := echo.New()

	db, err := sqlx.Connect("postgres", "user=local password=local dbname=hospital_shifts sslmode=disable application_name=shift_app")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Logger())

	// Routes
	e.POST("/update-with-advisory", func(c echo.Context) error { return updateWithAdvisoryLock(c, db) })
	e.POST("/update-with-serializable", func(c echo.Context) error { return updateWithSerializableIsolation(c, db) })
	e.POST("/update", func(c echo.Context) error { return updateWithDefaultIsolation(c, db) })
	e.POST("/reset/shift", func(c echo.Context) error { return resetShifts(c, db) })
	e.GET("/shift", func(c echo.Context) error { return listShifts(c, db) })

	// Start server
	e.Logger.Fatal(e.Start(":9092"))
}
