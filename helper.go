package main

import (
	"database/sql"
	"log"
)

const file string = "hackathon.db"
const create string = `
  CREATE TABLE IF NOT EXISTS donations (
      item VARCHAR(25),
      needed INTEGER,
      current INTEGER
  );`

func NewDonations() (*Donations, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &Donations{
		db: db,
	}, nil
}

func (c *Donations) Insert(donation Donation) (int, error) {
	res, err := c.db.Exec("INSERT INTO donations VALUES(?,?,?);", donation.Name, donation.Needed, donation.Current)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (c *Donations) Retrieve(rowNum int) (Donation, error) {
	log.Printf("Getting row %d", rowNum)

	// Query DB row based on ID
	row := c.db.QueryRow("SELECT item, needed, current FROM donations LIMIT 1 OFFSET ?", rowNum)

	// Parse row into Activity struct
	donation := Donation{}
	var err error
	if err = row.Scan(&donation.Name, &donation.Needed, &donation.Current); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return Donation{}, sql.ErrNoRows
	}
	return donation, err
}
