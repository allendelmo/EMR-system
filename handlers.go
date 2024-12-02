package main

import (
	"net/http"
	"time"
)

type MedicalRecord struct {
	PATIENT_ID int
	FIRST_NAME string
	LAST_NAME  string
	DOB        time.Time
	GENDER     string
	CONTACT    int32
	EMAIL      string
}

func (cfg *config) InsertMedicalRecord(w http.ResponseWriter, r *http.Request) {

}
