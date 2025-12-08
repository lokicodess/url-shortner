package model

import (
	"database/sql"
	"errors"
	"time"
)

type UrlModel struct {
	DB *sql.DB
}

type Url struct {
	ID         int       `json:"id"`
	Short_Code string    `json:"short_code"`
	Actual_Url string    `json:"actual_url"`
	Created_At time.Time `json:"created_at"`
	Expires_At time.Time `json:"expires_at"`
}

func (um UrlModel) Get(short_code string) (string, error) {
	stmt := `
	SELECT actual_url from url_table 
	WHERE short_code = ?;
	`
	var actual_url string
	err := um.DB.QueryRow(stmt, short_code).Scan(&actual_url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrRowNotFound
		}
		return "", nil
	}
	return actual_url, nil
}

func (um UrlModel) Post(short_code, actual_url string, expires_at int) error {
	stmt := `
 INSERT INTO url_table(short_code,actual_url,created_at,expires_at)
 VALUES(?,?,UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
`
	_, err := um.DB.Exec(stmt, short_code, actual_url, expires_at)
	if err != nil {
		return err
	}
	return nil
}
