-- name: CreateMedicalRecord :exec
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
    
);

-- name: GetMedicalRecord :many
SELECT *
FROM MEDICAL_RECORD;
