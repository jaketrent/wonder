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

func insertWonder(db *sql.DB, userWonders []*UserWonder) ([]*UserWonder, error) {
	var tx *sql.Tx
	var err error
	tx, err = db.Begin()

	for _, userWonder := range userWonders {
		const query = `
insert into user_wonders
( user_id
, description
) values
( $1
, $2
) returning id
, created
`

		err = tx.QueryRow(query, userWonder.UserID, userWonder.Description).Scan(&userWonder.ID, &userWonder.Created)
		if err != nil {
			err = tx.Rollback()
			break
		}
	}

	err = tx.Commit()

	return userWonders, err
}
