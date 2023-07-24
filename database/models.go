package database

import (
	"database/sql"
	"errors"
)

type DataCenterInterface interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type DataCenter struct {
	DBconn DataCenterInterface
}

type Post struct {
	id    int
	Title string
	Body  string
}

func (dbExec *DataCenter) CreatePost(title, body string) error {
	sql := `INSERT INTO blog(title, body) VALUES(?, ?)`

	_, err := dbExec.DBconn.Exec(sql, title, body)
	if err != nil {
		return err
	}

	return nil
}

func (dbExec *DataCenter) ReadPost(postID int64) (*Post, error) {
	p := Post{}

	sql := `SELECT title, body FROM blog WHERE id=(?)`

	err := dbExec.DBconn.QueryRow(sql, postID).Scan(&p.Title, &p.Body)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (dbExec *DataCenter) DeletePost(postID int64) error {
	sql := `DELETE FROM blog WHERE id=(?)`

	res, err := dbExec.DBconn.Exec(sql, postID)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if ra == 0 {
		return errors.New("enr")
	}

	return nil
}
