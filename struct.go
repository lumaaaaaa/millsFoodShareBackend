package main

import "database/sql"

type Donations struct {
	db *sql.DB
}

type Donation struct {
	Name    string
	Needed  int
	Current int
}

type ItemsResponse struct {
	Items []Donation
}
