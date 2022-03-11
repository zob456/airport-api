package data

import (
	"AirportApi/rest-api/models"
	"database/sql"
)

func SelectAuthedUser(db *sql.DB, email string, password string) (*models.AuthedUser, error) {
	/*language=PostgreSQL*/
	const query = `SELECT
       "UserID",
       "UserName"
	FROM "UserData"."vw_User" WHERE "UserName" = $1 AND "Password" = $2 AND "Active" = true LIMIT 1;`

	authedUser := &models.AuthedUser{}
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(email, password)
	err = row.Scan(
		&authedUser.UserID,
		&authedUser.UserName,
	)
	if err != nil {
		return nil, err
	}
	return authedUser, nil
}

func AddJWT(db *sql.DB, value string, userID string) error {
	/*language=PostgreSQL*/
	const query = `UPDATE "UserData"."user" SET "JWT" = $1 where "UserID" = $2`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(value, userID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJWT(db *sql.DB, userID string) error {
	/*language=PostgreSQL*/
	const query = `UPDATE "UserData"."user" SET "JWT" = $1 where "UserID" = $2`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(nil, userID)
	if err != nil {
		return err
	}
	return nil
}
