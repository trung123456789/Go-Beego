package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Car_20200205_180956 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Car_20200205_180956{}
	m.Created = "20200205_180956"

	migration.Register("Car_20200205_180956", m)
}

// Run the migrations
func (m *Car_20200205_180956) Up() {
	query := `
			CREATE TABLE car (
				id INT PRIMARY KEY NOT NULL,
				name VARCHAR(255), 
				type VARCHAR(255), 
				color VARCHAR(255)
			);
		`
	m.SQL(query)
}

// Reverse the migrations
func (m *Car_20200205_180956) Down() {
	m.SQL("DROP TABLE car")

}
