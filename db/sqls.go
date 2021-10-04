package sqls

import (
	"database/sql"
	"fmt"
	structures "pos-tech/models"

	_ "github.com/mattn/go-sqlite3"
)

type SqlDb struct {
	DB *sql.DB
}

func (m *SqlDb) Get(id int) (*structures.Task, error) {

	task := &structures.Task{}
	stmt := "SELECT * FROM demo WHERE id = ?"
	err := m.DB.QueryRow(stmt, id).Scan(&task.Id, &task.Name, &task.Task, &task.Author)
	if err != nil {
		return nil, err
	}

	return task, nil
}
func (m *SqlDb) WriteToSql(data *structures.Task) error {
	stmt := "INSERT INTO demo(name, task, author) VALUES (?, ?, ?)"
	statement, err := m.DB.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = statement.Exec(data.Name, data.Task, data.Author)
	if err != nil {
		return err
	}
	return nil
}
func (m *SqlDb) DeleteFromSql(id int) error {
	_, err := m.DB.Exec("DELETE FROM demo WHERE id = ?", id)
	return err
}
func (m *SqlDb) UpdateSql(data *structures.Task) error {
	_, err := m.DB.Exec("update demo set name = ?, task = ?, author = ? where id = ?", data.Name, data.Task, data.Author, data.Id)
	fmt.Println(err)
	return err
}
