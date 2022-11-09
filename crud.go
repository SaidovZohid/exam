package main

import (
	"database/sql"
	"time"
)

type DBManager struct {
	db *sql.DB
}

type Car struct {
	Id         int
	Name       string
	Model      string
	Year       time.Time
	Color      string
	HorsePower int
	Km         float64
	Images     []*Images
	Image      string
}

type Images struct {
	Id              int
	Car_id          int
	Sequence_number int
	Image_url       string
}

type GetAllCar struct {
	AllCars []*Car
}

type GetAllParams struct {
	Limit  int
	Page   int
	Search string
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{
		db: db,
	}
}

// to create new car info and images of this car in database
func (d *DBManager) CreateAvtomobile(c *Car) (int, error) {
	tx, err := d.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	query := `
		INSERT INTO avtomobile(
			name,
			model,
			year,
			color,
			horse_power,
			km
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	row := tx.QueryRow(
		query,
		c.Name,
		c.Model,
		c.Year,
		c.Color,
		c.HorsePower,
		c.Km,
	)
	var car_id int
	err = row.Scan(
		&car_id,
	)
	if err != nil { 
		tx.Rollback()
		return 0, err
	}
	queryInsertImage := `
		INSERT INTO images(
			car_id, 
			sequence_number,
			image_url
		) VALUES ($1, $2, $3)
	`
	for _, v := range c.Images {
		_, err := tx.Exec(
			queryInsertImage,
			car_id,
			v.Sequence_number,
			v.Image_url,
		)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return car_id, nil
}

func (d *DBManager) GetCar(car_id int) (*Car, error) {
	query := `
		SELECT 
			a.id, 
			a.name, 
			a.model, 
			a.year, 
			a.color, 
			a.horse_power, 
			a.km 
		FROM avtomobile a WHERE a.id = $1
	`
	var car Car
	row := d.db.QueryRow(
		query,
		car_id,
	)
	err := row.Scan(
		&car.Id,
		&car.Name,
		&car.Model,
		&car.Year,
		&car.Color,
		&car.HorsePower,
		&car.Km,
	)
	if err != nil {
		return nil, err
	}
	queryTakeImages := `
		SELECT 
			i.id,
			i.sequence_number,
			i.image_url
		FROM images i WHERE car_id = $1
	`
	rows, err := d.db.Query(queryTakeImages, car_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i Images
		err := rows.Scan(
			&i.Id,
			&i.Sequence_number,
			&i.Image_url,
		)
		if err != nil {
			return nil, err
		}
		car.Images = append(car.Images, &i)
	}
	return &car, nil
}

func (d *DBManager) UpdateCar(c *Car) error {
	tx, err := d.db.Begin()

	if err != nil {
		tx.Rollback()
		return err
	}

	queryUpdateCar := `
		UPDATE avtomobile SET 
			name = $1,
			model = $2,
			year = $3,
			color = $4,
			horse_power = $5,
			km = $6
		WHERE id = $7 
	`

	row, err := tx.Exec(
		queryUpdateCar,
		c.Name,
		c.Model,
		c.Year,
		c.Color,
		c.HorsePower,
		c.Km,
		c.Id,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	resCount, err := row.RowsAffected()

	if err != nil {
		tx.Rollback()
		return err
	}

	if resCount == 0 {
		tx.Rollback()
		return sql.ErrNoRows
	}

	queryDeleteImages := `
		DELETE FROM images WHERE car_id = $1
	`

	_, err = tx.Exec(queryDeleteImages, c.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryInsertImage := `
		INSERT INTO images(
			car_id,
			sequence_number,
			image_url
		) VALUES ($1, $2, $3)
	`

	for _, v := range c.Images {
		_, err = tx.Exec(
			queryInsertImage,
			c.Id,
			v.Sequence_number,
			v.Image_url,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *DBManager) DeleteCar(car_id int) error {
	queryDeleteImages := `
		DELETE FROM images where car_id = $1
	`
	_, err := d.db.Exec(queryDeleteImages, car_id)
	if err != nil {
		return err
	}
	queryDeleteCar := `
		DELETE FROM avtomobile where id = $1
	`
	_, err = d.db.Exec(queryDeleteCar, car_id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBManager) GetAllCarsInfo(p *GetAllParams) (*GetAllCar, error) {
	offset := (p.Page - 1) * p.Limit
	filter := ""
	if p.Search != "" {
		filter = " WHERE a.name ilike '%s'" + "%" + p.Search + "%"
	}
	query := `
		SELECT * FROM car_view a` + filter + ` ORDER BY a.id ASC LIMIT $1 OFFSET $2
	`
	rows, err := d.db.Query(query, p.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars GetAllCar
	for rows.Next() {
		var c Car
		err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Model,
			&c.Year,
			&c.Color,
			&c.HorsePower,
			&c.Km,
			&c.Image,
		)
		if err != nil {
			return nil, err
		}
		cars.AllCars = append(cars.AllCars, &c)
	}
	return &cars, nil
}
