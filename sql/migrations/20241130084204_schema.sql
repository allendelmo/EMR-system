-- +goose Up
CREATE TABLE MEDICAL_RECORD(
    PATIENT_ID INTEGER NOT NULL PRIMARY KEY,
    FIRST_NAME TEXT,
    LAST_NAME TEXT,
    DOB DATE,
    GENDER TEXT,
    CONTACT INTEGER,
    EMAIL TEXT
);

CREATE TABLE MEDICAL_HISTORY(
    PATIENT_ID NOT NULL,
    MEDICAL_HISTORY TEXT,
    ALLERGIES TEXT,
    FOREIGN KEY (PATIENT_ID) REFERENCES
    MEDICAL_RECORD(PATIENT_ID)
);

CREATE TABLE EMERGENCY_CONTACT(
    PATIENT_ID NOT NULL,
    E_NAME TEXT,
    CONTACT INTEGER,
    FOREIGN KEY (PATIENT_ID) REFERENCES
    MEDICAL_RECORD(PATIENT_ID)
);




-- +goose Down
DROP TABLE IF EXISTS MEDICAL_RECORD; 
DROP TABLE IF EXISTS MEDICAL_HISTORY; 
DROP TABLE IF EXISTS EMERGENCY_CONTACT;
