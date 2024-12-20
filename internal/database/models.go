// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type EMERGENCYCONTACT struct {
	PatientID interface{}
	EName     sql.NullString
	Contact   sql.NullInt64
}

type MEDICALHISTORY struct {
	PatientID      interface{}
	MedicalHistory sql.NullString
	Allergies      sql.NullString
}

type MEDICALRECORD struct {
	PatientID int64
	FirstName sql.NullString
	LastName  sql.NullString
	Dob       sql.NullTime
	Gender    sql.NullString
	Contact   sql.NullInt64
	Email     sql.NullString
}
