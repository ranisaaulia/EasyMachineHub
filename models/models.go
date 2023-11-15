// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package models

// import (
// 	"database/sql"
// )

type Employee struct {
	EmployeeID int32         `db:"employee_id" json:"employeeId"`
	UserID     int32 `db:"user_id" json:"userId"`
	Name       string        `db:"name" json:"name"`
	Position   string        `db:"position" json:"position"`
}

type User struct {
	UserID   int32  `db:"user_id" json:"userId"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}
