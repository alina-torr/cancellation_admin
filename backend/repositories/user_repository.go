package repositories

import (
	ent "booking/entities"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
}

func NewUserRepository(dbpool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		dbpool: dbpool,
	}
}

func (ur UserRepository) CreateManager(user ent.ManagerData, hotelId int64) (int64, error) {
	tx, err := ur.dbpool.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())

	var userId int64
	err = tx.QueryRow(context.Background(),
		`INSERT INTO userscheme.manager (hotel_id, login, password)
			VALUES ($1, $2, $3) RETURNING id`, hotelId, user.Login, user.Password).Scan(&userId)

	if err != nil {
		return -1, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (ur UserRepository) CreateClient(user ent.ClientRegister) (int64, error) {
	tx, err := ur.dbpool.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())

	var userId int64
	err = tx.QueryRow(context.Background(),
		`INSERT INTO userscheme.client (phone_number, email, first_name, last_name, middle_name, country)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, user.Phone_number, user.Email, user.First_name, user.Last_name, user.Middle_name, user.Country).Scan(&userId)

	if err != nil {
		return -1, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (ur UserRepository) GetClientByPhone(phone_number string) (user ent.Client, err error) {
	rows, err := ur.dbpool.Query(context.Background(),
		`SELECT * FROM userscheme.client u WHERE phone_number = $1`, phone_number)
	if err != nil {
		return
	}
	user, err = pgx.CollectOneRow[ent.Client](rows, pgx.RowToStructByName[ent.Client])
	return
}

func (ur UserRepository) GetClientByEmail(email string) (user ent.Client, err error) {
	rows, err := ur.dbpool.Query(context.Background(),
		`SELECT * FROM userscheme.client u WHERE email = $1`, email)
	if err != nil {
		return
	}
	user, err = pgx.CollectOneRow[ent.Client](rows, pgx.RowToStructByName[ent.Client])
	return
}

func (ur UserRepository) GetManagerByLogin(login string) (user ent.Manager, err error) {
	rows, err := ur.dbpool.Query(context.Background(),
		`SELECT * FROM userscheme.manager u WHERE login = $1`, login)
	if err != nil {
		return
	}
	user, err = pgx.CollectOneRow[ent.Manager](rows, pgx.RowToStructByName[ent.Manager])
	return
}

func (ur UserRepository) GetManagerById(id int64) (user ent.Manager, err error) {
	rows, err := ur.dbpool.Query(context.Background(),
		`SELECT * FROM userscheme.manager u WHERE id = $1`, id)
	if err != nil {
		return
	}
	user, err = pgx.CollectOneRow[ent.Manager](rows, pgx.RowToStructByName[ent.Manager])
	return
}

// func (ur UserRepository) GetUserById(id int) (user ent.User, err error) {
// 	rows, err := ur.dbpool.Query(context.Background(),
// 		`SELECT * FROM userscheme.user u WHERE id = $1`, id)
// 	if err != nil {
// 		return
// 	}
// 	user, err = pgx.CollectOneRow[ent.User](rows, pgx.RowToStructByName[ent.User])
// 	return
// }

// func (ur UserRepository) SelectUserByNickname(nickname string)
