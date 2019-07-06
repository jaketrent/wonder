package users

import (
	"database/sql"
)

func findAll(db *sql.DB) ([]*User, error) {
	rows, err := db.Query(`
		select id
		, name
		from users
  `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil

}
