package repositorie

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"AntonyChR/golang-hexagonal-arch-template/internal/core/domain"
)

type MySqlTodoRepository struct {
	DB    *sql.DB
	Table string
}

func (repo *MySqlTodoRepository) GetById(id string) (domain.Todo, error) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE id='%s'", repo.Table, id)
	res, err := repo.DB.Query(q)
	var t domain.Todo
	if err != nil {
		return domain.Todo{}, err
	}
	for res.Next() {
		err = res.Scan(&t.ID, &t.Title, &t.Description, &t.Date, &t.Done)
		fmt.Println(t)
	}
	return t, err
}

func (repo *MySqlTodoRepository) GetAll() ([]domain.Todo, error) {
	todos := []domain.Todo{}
	res, err := repo.DB.Query("SELECT * FROM " + repo.Table)
	if err != nil {
		return []domain.Todo{}, err
	}

	for res.Next() {
		var temp domain.Todo
		err = res.Scan(&temp.ID, &temp.Title, &temp.Description, &temp.Date, &temp.Done)
		if err != nil {
			continue
		}
		todos = append(todos, temp)
	}
	return todos, err

}
func (repo *MySqlTodoRepository) Create(data domain.Todo) (string, error) {
	id := uuid.New().String()
	done := "0"
	if data.Done {
		done = "1"
	}
	values := fmt.Sprintf("VALUES('%s','%s','%s','%s', %s)", id, data.Title, data.Description, data.Date, done)
	q := fmt.Sprintf("INSERT INTO %s (id, title, description, date, done) %s", repo.Table, values)
	_, err := repo.DB.Query(q)
	if err != nil {
		return "", err
	}
	return id, err
}
func (repo *MySqlTodoRepository) Update(id, field string, value string) error {
	q := fmt.Sprintf("UPDATE %s SET %s='%s' WHERE id='%s'", repo.Table, field, value, id)
	_, err := repo.DB.Query(q)
	return err
}

func (repo *MySqlTodoRepository) DeleteById(id string) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE id='%s'", repo.Table, id)
	_, err := repo.DB.Query(q)
	return err
}

type MysqlConfig struct {
	Database string
	Table    string
	Host     string
	Port     string
	User     string
	pass     string
}

func InitializeMysqlRepository(config MysqlConfig) *MySqlTodoRepository {
	configString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.pass, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", configString)
	if err != nil {
		log.Fatal(err)
	}
	return &MySqlTodoRepository{
		DB:    db,
		Table: config.Table,
	}
}
