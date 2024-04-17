package postgres
import (
	"clone/3_exam/api/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"errors"
)
func (c *UserRepo)UserRegisterCreateConfirm(ctx context.Context, User models.LoginUser) (string, error) {

	id := uuid.New()
	fmt.Println("PASSWORD-----------",User.Password)
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

func (c *UserRepo) GetGmail (ctx context.Context, gmail string) (string, error) {
	var id string

	query := `SELECT id
	FROM Users
	WHERE mail = $1 `

	err := c.db.QueryRow(ctx, query, gmail).Scan(&id)

	if err != nil {
		return id,nil
	}

	return id, errors.New(" get gmail address is already registered")
}

