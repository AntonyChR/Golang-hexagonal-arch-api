package repositoryImplementations

import (
	"chat/domain/entities"
	"database/sql"
	"fmt"
)

func NewMySqlUserRepository(table string, db *sql.DB) *MySqlUserRepository {
	return &MySqlUserRepository{Table: table, DB: db} 
}

type MySqlUserRepository struct {
	DB    *sql.DB
	Table string
}

func (r *MySqlUserRepository) GetAll() ([]*entities.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", r.Table)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*entities.User{}
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = ""
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MySqlUserRepository) GetById(id string) (*entities.User, error) {
	return r.GetBy("id", id)
}

func (r *MySqlUserRepository) GetBy(field, value string) (*entities.User, error) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s=?", r.Table, field)
	resp, err := r.DB.Query(q, value)

	if err != nil {
		return nil, err
	}

	defer resp.Close()

	if !resp.Next() {
		return nil, fmt.Errorf("user with %s %s not found", field, value)
	}

	var user entities.User
	err = resp.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MySqlUserRepository) Save(user *entities.User) error {
	_, err := r.GetBy("email", user.Email)
	if err == nil {
		return fmt.Errorf("the email is already in use")
	}
	q := fmt.Sprintf("INSERT INTO %s (id, name, email, password) VALUES (?, ?, ?, ?)", r.Table)
	stmt, err := r.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *MySqlUserRepository) DeleteById(id string) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=?", r.Table)
	_, err := r.DB.Exec(q, id)
	if err != nil {
		return err
	}
	return nil

}
