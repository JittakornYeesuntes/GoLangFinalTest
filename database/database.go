package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sql.DB, error) {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	err = CreateTable(conn)
	if err!= nil {
		return nil, err
	}
	return conn, err
}

func CreateTable(conn *sql.DB) error{
	createTb := `
	CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`

	_, err := conn.Exec(createTb)
	if err != nil {
		return err
	}
	return nil
}

func InsertCustomer(name, email, status string)(*sql.Row, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id, name, email, status")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(name, email, status)
	return row, nil
}

func SelectByID(id string)(*sql.Row, error){
	conn, err:= Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT id, name, email, status FROM customers WHERE id = $1")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)
	return row, nil
}

func SelectAll() (*sql.Rows, error) {
	conn, err:= Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return rows, err
	}

	return rows, err
}


func UpdateByID(id, name, email, status string)(*sql.Row, error){
	conn, err:= Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("UPDATE customers SET name=$2, email=$3, status=$4  WHERE id=$1  RETURNING  id, name, email, status")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id, name, email, status)
	return row, nil
}

func DeleteByID(id string) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}