package repository

const (
	createUserQuery = `INSERT INTO users(name, email, password) VALUES (:name, :email, :password) RETURNING id`
)
