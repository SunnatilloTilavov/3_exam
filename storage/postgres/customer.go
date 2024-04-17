package postgres

import (
	"clone/rent_car_us/api/models"
	// "clone/rent_car_us/pkg"
	"context"
	"database/sql"
	"fmt"
"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomer(db *pgxpool.Pool) customerRepo {
	return customerRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *customerRepo) Create(ctx context.Context, customer models.CreateCustomer) (string, error) {

	id := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return "password hashing error", err
	}

	query := ` INSERT INTO customers (
		id ,       
		first_name,
		last_name ,
		gmail,    
		phone,
		password     )
		VALUES($1,$2,$3,$4,$5,$6) `

	_, err = c.db.Exec(ctx, query,
		id.String(),
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone, hashedPassword)
	if err != nil {
		return "", err
	}

	

	return id.String(), nil
}









func (c *customerRepo) Update(ctx context.Context, customer models.GetCustomer) (string, error) {

	query := ` UPDATE customers set
	        first_name=$1,
	        last_name=$2,
	        gmail=$3,
	        phone=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $5 
	`

	_, err := c.db.Exec(ctx, query,
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone, customer.Id)

	if err != nil {
		return "", err
	}

	return customer.Id, nil
}

func (c *customerRepo) GetAllCustomers(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}
		// filter = ""
	)
// 	offset := (req.Page - 1) * req.Limit
// 	if req.Search != "" {
// 		filter += fmt.Sprintf(` and first_name ILIKE '%%%v%%'`, req.Search)
// 	}

// 	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
// 	fmt.Println("filter:", filter)

// 	query := `SELECT
//     cu.id AS customer_id,
//     cu.first_name AS customer_first_name,
//     cu.last_name AS customer_last_name,
//     cu.gmail AS customer_gmail,
//     cu.phone AS customer_phone,
//     COUNT(o.id) AS order_count,
//     COUNT(ca.id) AS car_count,
//     cu.created_at,
//     cu.updated_at,
//     o.id AS order_id,
//     o.from_date,
//     o.to_date,
//     o.status,
//     o.paid,
//     o.amount
// FROM
//     customers cu
// JOIN
//     orders o ON cu.id = o.customer_id
// JOIN
//     cars ca ON ca.id = o.car_id
// GROUP BY
//     cu.id,
//     cu.first_name,
//     cu.last_name,
//     cu.gmail,
//     cu.phone,
//     cu.created_at,
//     cu.updated_at,
//     o.id,
//     o.from_date,
//     o.to_date,
//     o.status,
//     o.paid,
//     o.amount`

// 	rows, err := c.db.Query(ctx, query)
// 	if err != nil {
// 		return resp, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var (
// 			customer = models.GetAllCustomer{
// 				Order: models.GetOrder{},
// 			}
// 			updateAt sql.NullString
// 		)
// 		if err := rows.Scan(
// 			&customer.Id,
// 			&customer.FirstName,
// 			&customer.LastName,
// 			&customer.Gmail,
// 			&customer.Phone,
// 			&customer.OrderCount,
// 			&customer.CarsCount,
// 			&customer.CreatedAt,
// 			&updateAt,
// 			&customer.Order.Id,
// 			&customer.Order.FromDate,
// 			&customer.Order.ToDate,
// 			&customer.Order.Status,
// 			&customer.Order.Paid,
// 			&customer.Order.Amount); err != nil {
// 			return resp, err
// 		}
// 		customer.UpdatedAt = pkg.NullStringToString(updateAt)
// 		resp.Customer = append(resp.Customer, customer)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return resp, err
// 	}
// 	countQuery := `Select count(*) from customers`
// 	err = c.db.QueryRow(ctx, countQuery).Scan(&resp.Count)
// 	if err != nil {
// 		return resp, err
// 	}
	return resp, nil
}

func (c *customerRepo) GetByID(ctx context.Context, id string) (models.GetCustomer, error) {
	customer := models.GetCustomer{}

	if err := c.db.QueryRow(ctx, `select id,
	first_name,
	last_name,gmail,
	phone
	from customers where id = $1`, id).Scan(
		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Gmail,
		&customer.Phone); err != nil {
		return models.GetCustomer{}, err
	}
	return customer, nil
}

func (c *customerRepo) Delete(ctx context.Context, id string) error {

	query := ` UPDATE customers set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 
	`

	_, err := c.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepo) GetAllCustomerCars(ctx context.Context, req models.GetAllCustomerCarsRequest) (models.GetAllCustomerCarsResponse, error) {
	var (
		resp   = models.GetAllCustomerCarsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(` and ca.name ILIKE '%%%v%%'`, req.Search)
	}

	if req.Id != "" {
		filter += fmt.Sprintf(` Where cu.id= '%%%v%%'`, req.Id)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query := `Select 
	cu.id as customer_id,
	ca.name as car_name,
	o.created_ad as creatad_at,
	o.amount as price
	From customers cu JOIN orders o ON  cu.id = o.customer_id Join cars ca  ON ca.id=o.car_id 
	`

	rows, err := c.db.Query(ctx, query+filter+``)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			customer = models.GetAllCustomerCars{}
		)
		if err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.CreatedAt,
			&customer.Amount); err != nil {
			return resp, err
		}
		resp.Customer = append(resp.Customer, customer)
	}
	if err = rows.Err(); err != nil {
		return resp, err
	}
	countQuery := `Select count(*) from customers`
	err = c.db.QueryRow(context.Background(), countQuery).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}


func (c *customerRepo) UpdatePassword(ctx context.Context, customer models.PasswordCustomer) (string, error) {

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(customer.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing new password")
	}

	var currentPassword string
	err = c.db.QueryRow(ctx, `SELECT password FROM customers WHERE phone = $1`, customer.Phone).Scan(&currentPassword)
	if err != nil {
		return "", errors.New("customer not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(customer.OldPassword))
	if err != nil {
		return "", errors.New("invalid old password")
	}

	_, err = c.db.Exec(ctx, `UPDATE customers SET password = $1 WHERE phone = $2`, hashedNewPassword, customer.Phone)
	if err != nil {
		return "", errors.New("error updating password")
	}

	return "Password updated successfully", nil
}



func (c *customerRepo) GetPassword (ctx context.Context, phone string) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM customers
	WHERE phone = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query, phone).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect phone")
		} else {
			return "", err
		}
	}

	return hashedPass, nil
}


func (c *customerRepo) GetByLogin(ctx context.Context, login string) (models.GetAllCustomer, error) {
	var (
		firstname sql.NullString
		lastname  sql.NullString
		phone     sql.NullString
		email     sql.NullString
		createdat sql.NullString
		updatedat sql.NullString
	)

	query := `SELECT 
		id, 
		first_name, 
		last_name, 
		phone,
		gmail,
		created_at, 
		updated_at,
		password
		FROM customers WHERE phone = $1 AND deleted_at = 0`

	row := c.db.QueryRow(ctx, query, login)

	customer := models.GetAllCustomer{
	}

	err := row.Scan(
		&customer.Id,
		&firstname,
		&lastname,
		&phone,
		&email,
		&createdat,
		&updatedat,
		&customer.Password,
	)

	if err != nil {
		// c.logger.Error("failed to scan customer by LOGIN from database", logger.Error(err))
		return models.GetAllCustomer{}, err
	}

	customer.FirstName = firstname.String
	customer.LastName = lastname.String
	customer.Phone = phone.String
	customer.Gmail = email.String
	customer.CreatedAt = createdat.String
	customer.UpdatedAt = updatedat.String

	fmt.Println(customer)

	return customer, nil


}




