// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: medical_record.sql

package database

import (
	"context"
	"database/sql"
)

const createMedicalRecord = `-- name: CreateMedicalRecord :exec
INSERT INTO MEDICAL_RECORD(
  FIRST_NAME,
  LAST_NAME,
  DOB,
  GENDER,
  CONTACT,
  EMAIL
)
VALUES
(
 
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
    
)
`

type CreateMedicalRecordParams struct {
	FirstName sql.NullString
	LastName  sql.NullString
	Dob       sql.NullTime
	Gender    sql.NullString
	Contact   sql.NullInt64
	Email     sql.NullString
}

func (q *Queries) CreateMedicalRecord(ctx context.Context, arg CreateMedicalRecordParams) error {
	_, err := q.db.ExecContext(ctx, createMedicalRecord,
		arg.FirstName,
		arg.LastName,
		arg.Dob,
		arg.Gender,
		arg.Contact,
		arg.Email,
	)
	return err
}

const getMedicalRecord = `-- name: GetMedicalRecord :many
SELECT patient_id, first_name, last_name, dob, gender, contact, email
FROM MEDICAL_RECORD
`

func (q *Queries) GetMedicalRecord(ctx context.Context) ([]MEDICALRECORD, error) {
	rows, err := q.db.QueryContext(ctx, getMedicalRecord)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MEDICALRECORD
	for rows.Next() {
		var i MEDICALRECORD
		if err := rows.Scan(
			&i.PatientID,
			&i.FirstName,
			&i.LastName,
			&i.Dob,
			&i.Gender,
			&i.Contact,
			&i.Email,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}