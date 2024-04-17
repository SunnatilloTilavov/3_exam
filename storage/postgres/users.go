package postgres

import (
	"clone/3_exam/api/models"
	// "clone/3_exam/pkg"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) UserRepo {
	return UserRepo{
		db: db,
	}
}


func (c *UserRepo) Create(ctx context.Context, User models.CreateUser) (string, error) {
	id := uuid.New()
	query := ` INSERT INTO users (
		id,       
		first_name,
		last_name ,
		mail,    
		phone,
		password,
		sex,
		active    )
		VALUES($1,$2,$3,$4,$5,$6,$7,$8) `

	_, err := c.db.Exec(ctx, query,
		id.String(),
		User.First_name, User.Last_name,
		User.Mail,User.Phone,
		User.Password,User.Sex,User.Active)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}




func (c *UserRepo) Update(ctx context.Context, User models.UpdateUser) (string, error) {

	query := ` UPDATE Users set      
			first_name=$1,
			last_name=$2,
			mail=$3,    
			phone=$4,
			sex=$5,
			active=$6,
			updated_at=CURRENT_TIMESTAMP
			WHERE id = $7
	`

	_, err := c.db.Exec(ctx, query,
		User.First_name, User.Last_name,
		User.Mail, User.Phone, User.Sex,User.Active,User.Id)

	if err != nil {
		return "", err
	}

	return User.Id, nil
}


func (c *UserRepo) GetAllUsers(ctx context.Context, req models.GetAllUsersRequest) (models.GetAllUsersResponse, error) {
	var (
		resp   = models.GetAllUsersResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` where first_name ILIKE  '%%%v%%' `, req.Search)
	}

	fmt.Println("req.Search",req.Search)

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	
	rows, err := c.db.Query(context.Background(),`select 
				count(id) OVER(),
				"id",
				"first_name",
				"last_name",
				"mail",
				"phone",
				"sex",
				"active",
				"created_at",
				"updated_at"
	  FROM users ` + filter + ``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			id,sex,first_name string
			createdAt sql.NullTime
			active bool
			last_name sql.NullString
			mail sql.NullString
			phone sql.NullString
			updateAt sql.NullTime
		)

		if err := rows.Scan(
			&resp.Count,
			&id,
			&first_name,
			&last_name,
			&mail,
			&phone,
			&sex,
			&active,
			&createdAt,
			&updateAt,
			); err != nil {
			return resp, err
		}

		resp.User = append(resp.User, models.GetAllUser{
			Id:id        ,
			Mail:mail.String       ,
			First_name:first_name ,
			Last_name:last_name.String  ,
			Phone:phone.String     ,
			Sex:sex   ,
			Active: active  ,
			CreatedAt:createdAt.Time.String()  ,
			UpdatedAt:updateAt.Time.String()  ,
		})
	}
	return resp, nil
}

func (c *UserRepo) GetByIDUser(ctx context.Context, id string) (models.GetAllUser, error) {
    User := models.GetAllUser{}
    var (
        createdAt sql.NullTime    
        updatedAt sql.NullTime
    )
    if err := c.db.QueryRow(ctx, `
        SELECT
            first_name,
            last_name,
            mail,
            phone,
            sex,
            active,
            created_at,
            updated_at
        FROM Users WHERE id = $1`, id).Scan(
            &User.First_name,
            &User.Last_name,
            &User.Mail,
            &User.Phone,
            &User.Sex,
            &User.Active,
            &createdAt,
            &updatedAt,
        ); err != nil {
        return models.GetAllUser{}, err
    }
    if createdAt.Valid {
        User.CreatedAt = createdAt.Time.String()
    }
    if updatedAt.Valid {
        User.UpdatedAt = updatedAt.Time.String()
    }
    return User, nil
}



func (c *UserRepo) Delete(ctx context.Context, id string) error {

	query := ` DELETE FROM users WHERE  id = $1;
	`

	_, err := c.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}



func (c *UserRepo) UpdatePassword(ctx context.Context, User models.PasswordUser) (string, error) {

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(User.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing new password")
	}

	var currentPassword string
	err = c.db.QueryRow(ctx, `SELECT password FROM Users WHERE mail = $1`, User.Mail).Scan(&currentPassword)
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(User.OldPassword))
	if err != nil {
		return "", errors.New("invalid old password")
	}

	_, err = c.db.Exec(ctx, `UPDATE Users SET password = $1 WHERE mail = $2`, hashedNewPassword, User.Mail)
	if err != nil {
		return "", errors.New("error updating password")
	}

	return "Password updated successfully", nil
}



func (c *UserRepo) GetPassword (ctx context.Context, phone string) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM Users
	WHERE mail = $1 AND deleted_at = 0`

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


func (c *UserRepo) GetByLogin(ctx context.Context, mail string) (models.GetIdPassword, error) {
	User := models.GetIdPassword{}
    var (
        createdAt sql.NullTime    
        updatedAt sql.NullTime
    )
    if err := c.db.QueryRow(ctx, `
        SELECT
            first_name,
            last_name,
            mail,
            phone,
            sex,
            active,
			password,
            created_at,
            updated_at
        FROM Users WHERE mail = $1`, mail).Scan(
            &User.First_name,
            &User.Last_name,
            &User.Mail,
            &User.Phone,
            &User.Sex,
            &User.Active,
			&User.Password,
            &createdAt,
            &updatedAt,
        ); err != nil {
        return models.GetIdPassword{}, err
    }
    if createdAt.Valid {
        User.CreatedAt = createdAt.Time.String()
    }
    if updatedAt.Valid {
        User.UpdatedAt = updatedAt.Time.String()
    }
    return User, nil
}



func (c *UserRepo) UpdatePasswordForget(ctx context.Context, User models.Forgetpassword2) (string, error) {

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing new password")
	}

	_, err = c.db.Exec(ctx, `UPDATE Users SET password = $1 WHERE mail = $2`, hashedNewPassword, User.Mail)
	if err != nil {
		return "", errors.New("error updating password")
	}

	return "Password updated successfully", nil
}



func (c *UserRepo) UpdateStatus(ctx context.Context, User models.UpdateStatus) (string, error) {

	query := ` UPDATE Users set      
			active=$1,
			updated_at=CURRENT_TIMESTAMP
			WHERE id = $2
	`

	_, err := c.db.Exec(ctx, query,User.Active,User.Id)

	if err != nil {
		return "", err
	}

	return User.Id, nil
}