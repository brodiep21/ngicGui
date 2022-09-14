package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Labor struct {
	Zip           string    `json:"ZIP"`
	County        string    `json:"County"`
	BodyRate      string    `json:"BodyRate"`
	PaintRate     string    `json:"PaintRate"`
	MechRate      string    `json:"MechRate"`
	FrameRate     string    `json:"FrameRate"`
	PaintSupplies string    `json:"PaintSupplies"`
	Createdat     time.Time `json:"Createdat"`
}

func (l *Labor) getRateByZip(db *sql.DB) error {
	return db.QueryRow("SELECT ZIP, County, BodyRate, PaintRate, MechRate, FrameRate, PaintSupplies, Createdat FROM labor WHERE ZIP=$1", l.Zip).Scan(&l.Zip, &l.County, &l.BodyRate, &l.PaintRate, &l.MechRate, &l.FrameRate, &l.PaintSupplies, &l.Createdat)
}

func (l *Labor) getRateByCounty(db *sql.DB, start, counter int) ([]Labor, error) {

	rows, err := db.Query("SELECT ZIP, County, BodyRate, PaintRate, MechRate, FrameRate, PaintSupplies, Createdat FROM labor WHERE County=$1", counter, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	labors := []Labor{}

	for rows.Next() {
		var l Labor
		if err := rows.Scan(&l.Zip, &l.County, &l.BodyRate, &l.PaintRate, &l.MechRate, &l.FrameRate, &l.PaintSupplies, &l.Createdat); err != nil {
			return nil, err
		}
		labors = append(labors, l)
	}
	return labors, nil
}

func (l *Labor) updateRate(db *sql.DB) (string, error) {
	_, err := db.Exec("UPDATE labor SET BodyRate=$1, PaintRate=$2, MechRate=$3, FrameRate=$4, PaintSupplies=$5, Createdat=$6 WHERE ZIP=$7", &l.BodyRate, &l.PaintRate, &l.MechRate, &l.FrameRate, &l.PaintSupplies, &l.Createdat, &l.Zip)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfully updated %s rates", l.Zip), nil
}

func (l *Labor) createLabor(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO labor (ZIP, County, BodyRate, PaintRate, MechRate, FrameRate, PaintSupplies, Createdat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", &l.Zip, &l.County, &l.BodyRate, &l.PaintRate, &l.MechRate, &l.FrameRate, &l.PaintSupplies, &l.Createdat)
	if err != nil {
		return err
	}
	return err
}

func (l *Labor) deleteLabor(db *sql.DB) (Labor, error) {

	lab := Labor{}
	response := db.QueryRow("ZIP, County, BodyRate, PaintRate, MechRate, FrameRate, PaintSupplies, Createdat WHERE ZIP=$1", l.Zip)

	if err := response.Scan(&l.Zip, &l.County, &l.BodyRate, &l.PaintRate, &l.MechRate, &l.FrameRate, &l.PaintSupplies, &l.Createdat); err != nil {
		return Labor{}, nil
	}
	_, err := db.Exec("DELETE from labor WHERE ZIP=$1", l.Zip)
	if err != nil {
		return Labor{}, err
	}

	return lab, nil
}

// CREATE TABLE Labor (
// 	ZIP VARCHAR(5) UNIQUE NOT NULL,
// 	County VARCHAR(50) NOT NULL,
// 	BodyRate VARCHAR(3) NOT NULL,
// 	PaintRate VARCHAR(3) NOT NULL,
// 	MechRate VARCHAR(3) ,
// 	FrameRate VARCHAR(3) ,
// 	Paintsupplies VARCHAR(3) NOT NULL,
// 	Createdat VARCHAR(50) NOT NULL,
// );
