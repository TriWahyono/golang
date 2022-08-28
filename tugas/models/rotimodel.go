package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/TUGAS/config"
	"github.com/jeypc/TUGAS/entities"
)

type RotiModel struct {
	conn *sql.DB
}

func NewRotiModel() *RotiModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &RotiModel{
		conn: conn,
	}
}

func (r *RotiModel) FindAll() ([]entities.Roti, error) {
	rows, err := r.conn.Query("select * from roti")
	if err != nil {
		return []entities.Roti{}, err
	}
	defer rows.Close()

	var dataRoti []entities.Roti
	for rows.Next() {
		var roti entities.Roti
		rows.Scan(&roti.Id,
			&roti.Nama_roti,
			&roti.Tittle,
			&roti.Description,
			&roti.Rating,
			&roti.Image)

		dataRoti = append(dataRoti, roti)
	}

	return dataRoti, nil
}

func (r *RotiModel) Create(roti entities.Roti) bool {
	result, err := r.conn.Exec("insert into roti (nama_roti, tittle, description, rating, image) values (?,?,?,?,?)", roti.Nama_roti, roti.Tittle, roti.Description, roti.Rating, roti.Image)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
