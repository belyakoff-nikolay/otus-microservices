package pgstorage

import "github.com/belyakoff-nikolay/otus-microservices/hw06/internal/app/model"

type UserRepo struct {
	storage *Storage
}

func (r *UserRepo) GetByID(ID int64) (*model.User, error) {
	user := model.User{}
	err := r.storage.db.Get(&user, "select ID, FirstName, LastName, Email from users where ID=$1", ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	var userID int64
	err := r.storage.db.QueryRow(
		"insert into users(FirstName, LastName, Email) values ($1, $2, $3) RETURNING ID",
		u.FirstName, u.LastName, u.Email).Scan(&userID)
	if err != nil {
		return nil, err
	}

	newUser := *u
	newUser.ID = userID
	return &newUser, nil
}

func (r *UserRepo) Update(u *model.User) error {
	_, err := r.storage.db.Exec(
		"update users set FirstName=$1, LastName=$2, Email=$3 where ID=$4",
		u.FirstName, u.LastName, u.Email, u.ID)
	return err
}

func (r *UserRepo) Drop(userID int64) error {
	_, err := r.storage.db.Exec("delete from users where ID=$1", userID)
	return err
}
