package models

import (
	"github.com/OmarDardery/home-passwords/db"
	"github.com/gosimple/slug"
)

type Precious struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"`
	Key  string `json:"key" binding:"required"`
}

func (p *Precious) Save() {
	p.Name = slug.Make(p.Name)
	r, err := db.Db.Exec("INSERT INTO precious (name, key) VALUES (?, ?)", p.Name, p.Key)
	if err != nil {
		panic(err)
	}
	p.ID, err = r.LastInsertId()
	if err != nil {
		panic(err)
	}
}

func GetAllPrecious() []Precious {
	rows, err := db.Db.Query("SELECT * FROM precious")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var preciousList []Precious
	for rows.Next() {
		var p Precious
		if err := rows.Scan(&p.ID, &p.Name, &p.Key); err != nil {
			panic(err)
		}
		preciousList = append(preciousList, p)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return preciousList
}

func GetPreciousByName(n string) (Precious, error) {
	row := db.Db.QueryRow("SELECT * FROM precious WHERE name = ?", n)
	var myPrecious Precious
	err := row.Scan(&myPrecious.ID, &myPrecious.Name, &myPrecious.Key)
	if err != nil {
		return Precious{}, err
	}
	return myPrecious, nil
}
