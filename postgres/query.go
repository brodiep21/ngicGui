package postgres

import (
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Labor struct {
	Zip           string `json:"ZIP"`
	County        string `json:"County"`
	BodyRate      string `json:"BodyRate"`
	PaintRate     string `json:"PaintRate"`
	MechRate      string `json:"MechRate"`
	FrameRate     string `json:"FrameRate"`
	PaintSupplies string `json:"PaintSupplies"`
	Createdat time.Time `json:"Createdat"`
}

func (l *Labor) GetRateByZip(db *sql.DB) error {
	return db.QueryRow("SELECT ZIP, County, BodyRate, PaintRate, MechRate, FrameRate, PaintSupplies, Createdat FROM labor WHERE ZIP=$1", l.ZIP).Scan(&l.Zip, &l.County, &l.BodyRate, &l.PaintRate,&l.MechRate,&l.FrameRate,&l.PaintSupplies,&l.Createdat)
}

func (l *Labor) GetRateByCounty(db *sql.DB) error {
	return db.QueryRow("SELECT ZIP, County, BodyRate, PaintRate, MechRate, FrameRate, PaintSupplies, Createdat FROM labor WHERE ZIP=$1", l.County).Scan(&l.Zip, &l.County, &l.BodyRate, &l.PaintRate,&l.MechRate,&l.FrameRate,&l.PaintSupplies,&l.Createdat)
}

func (l *Labor) UpdateRate) error {
	result, err := db.Exec("UPDATE labor SET BodyRate=$1, PaintRate=$2, MechRate=$3, FrameRate=$4, PaintSupplies=$5, Createdat=$6 WHERE ZIP=$7", &l.BodyRate, &l.PaintRate,&l.MechRate,&l.FrameRate,&l.PaintSupplies,&l.Createdat, &l.Zip)
}


CREATE TABLE Labor (
	ZIP VARCHAR(5) UNIQUE NOT NULL,
	County VARCHAR(50) NOT NULL,
	BodyRate VARCHAR(3) NOT NULL,
	PaintRate VARCHAR(3) NOT NULL,
	MechRate VARCHAR(3) ,
	FrameRate VARCHAR(3) ,
	Paintsupplies VARCHAR(3) NOT NULL,
	Createdat VARCHAR(50) NOT NULL,
);