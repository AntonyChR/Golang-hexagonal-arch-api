package repositoryImplementations

import (
	"chat/domain/entities"
	"database/sql"
	"fmt"
)

func NewMySqlMessageRepository(table string, db *sql.DB) *MySqlMessageRepository {
	return &MySqlMessageRepository{Table: table, DB: db} 
}

type MySqlMessageRepository struct {
	Table string
	DB    *sql.DB
}

func (r *MySqlMessageRepository) GetChat(from, to string) ([]*entities.Message, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE (`from` = ? AND `to` = ?) OR (`from` = ? AND `to` = ?) ORDER BY `date`", r.Table)
	rows, err := r.DB.Query(query, from, to, to, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []*entities.Message{}
	for rows.Next() {
		message := &entities.Message{}
		err := rows.Scan(&message.ID, &message.Date, &message.Content, &message.From, &message.To)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MySqlMessageRepository) Save(msg *entities.Message) error {
	q := fmt.Sprintf("INSERT INTO %s (`id`, `date`, `content`, `from`, `to`) VALUES (?, ?, ?, ?, ?)", r.Table)
	stmt, err := r.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(msg.ID, msg.Date, msg.Content, msg.From, msg.To)
	if err != nil {
		return err
	}

	return nil


}

func (r *MySqlMessageRepository) DeleteChat(from, to string) error {
	return nil
}
func (r *MySqlMessageRepository) DeleteMessage(id string) error {
	return nil
}

