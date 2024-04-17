package postgres


import (
	"time"
	"context"
	"clone/rent_car_us/api/models"
	// "clone/rent_car_us/pkg"
	"database/sql"
	 "fmt"
	 "github.com/google/uuid"
	 "github.com/jackc/pgx/v5/pgxpool"
)

type carRepo struct {
	db *pgxpool.Pool
}

func NewCar(db *pgxpool.Pool) carRepo {
	return carRepo{
		db: db,
	}
}



func (c *carRepo) Create(ctx context.Context,car models.Car) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		model,
		hourse_power,
		colour,
		engine_cap)
		VALUES($1,$2,$3,$4,$5,$6,$7) 
	`

	_, err := c.db.Exec(ctx,query,
		id.String(),
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *carRepo) Update(ctx context.Context,car models.Car) (string, error) {

	query := ` UPDATE cars set
			name=$1,
			brand=$2,
			model=$3,
			hourse_power=$4,
			colour=$5,
			engine_cap=$6,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $7 AND deleted_at=0
	`

	_, err := c.db.Exec(ctx,query,
		car.Name, car.Brand,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap, car.Id)

	if err != nil {
		return "", err
	}

	return car.Id, nil
}

func (c carRepo) GetAllCars(ctx context.Context,req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	
	rows, err := c.db.Query(context.Background(),`select 
				count(id) OVER(),
				id, 
				name,
				brand,
				model,
				hourse_power,
				colour,
				engine_cap,
				created_at::date,
				updated_at
	  FROM cars WHERE deleted_at = 0 ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			updateAt sql.NullString
			id sql.NullString
			name sql.NullString
			brand sql.NullString
			model sql.NullString
			hourse_power sql.NullInt64
			colour sql.NullString
			engine_cap sql.NullFloat64
			createdAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&brand,
			&model,
			&hourse_power,
			&colour,
			&engine_cap,
			&createdAt,
			&updateAt,
			); err != nil {
			return resp, err
		}

		resp.Cars = append(resp.Cars, models.Car{
			Id:id.String ,
			Name: name.String,
			Brand: brand.String,
			Model: model.String,
			HoursePower: int(hourse_power.Int64),
			Colour: colour.String,
			EngineCap: float32(engine_cap.Float64),
			CreatedAt: createdAt.String,
			UpdatedAt: updateAt.String,
		})
	}
	return resp, nil
}

func (c *carRepo) GetByID(ctx context.Context,id string) (models.Car, error) {
	car := models.Car{}
	if err := c.db.QueryRow(ctx,`select id,name,
	brand,model,hourse_power,
	colour,engine_cap from cars where id = $1`, id).Scan(
		&car.Id,
		&car.Name,
		&car.Brand,
		&car.Model,
		&car.HoursePower,
		&car.Colour,
		&car.EngineCap,	
	); err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c *carRepo) Delete(ctx context.Context,id string) error {

	query := ` UPDATE cars set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.db.Exec(ctx,query, id)

	if err != nil {
		return err
	}

	return nil
}


func (c carRepo) GetAllCarsFree(ctx context.Context,req models.GetAllCarsRequest) (models.GetAllCarsResponse, error) {
	var (
		resp   = models.GetAllCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}
	now := time.Now()
	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)
	rows, err := c.db.Query(context.Background(), `
    SELECT 
        COUNT(cars.id) OVER(),
        cars.id, 
        cars.name,
        cars.brand,
        cars.model,
        cars.hourse_power,
        cars.colour,
        cars.engine_cap,
        cars.created_at::date,
        cars.updated_at
    FROM 
        cars
    LEFT JOIN 
        orders ON cars.id = orders.car_id
    WHERE 
        deleted_at = 0 AND $1 > orders.to_date `+filter,
    now)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (

			updateAt sql.NullString
			id sql.NullString
			name sql.NullString
			brand sql.NullString
			model sql.NullString
			hourse_power sql.NullInt64
			colour sql.NullString
			engine_cap sql.NullFloat64
			createdAt sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&brand,
			&model,
			&hourse_power,
			&colour,
			&engine_cap,
			&createdAt,
			&updateAt,
			); err != nil {
			return resp, err
		}

		resp.Cars = append(resp.Cars, models.Car{
			Id:id.String ,
			Name: name.String,
			Brand: brand.String,
			Model: model.String,
			HoursePower: int(hourse_power.Int64),
			Colour: colour.String,
			EngineCap: float32(engine_cap.Float64),
			CreatedAt: createdAt.String,
			UpdatedAt: updateAt.String,
		})
	}
	return resp, nil
}



