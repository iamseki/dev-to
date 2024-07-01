package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type DoctorShift struct {
	ShiftID    int    `json:"shiftId"`
	DoctorName string `json:"doctorName"`
	OnCall     bool   `json:"onCall"`
}

func main() {
	e := echo.New()

	db, err := sqlx.Connect("postgres", "user=local password=local dbname=hospital_shifts sslmode=disable")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Logger())

	// Routes
	e.POST("/update-with-advisory", func(c echo.Context) error { return updateWithAdvisoryLock(c, db) })
	e.POST("/update-with-serializable", func(c echo.Context) error { return updateWithSerializableIsolation(c, db) })
	e.POST("/reset/shift", func(c echo.Context) error { return resetShifts(c, db) })

	// Start server
	e.Logger.Fatal(e.Start(":9092"))
}
