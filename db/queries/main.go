package queries

import (
	"fmt"
)

func CreateUser(name string, mail string) string {

	insertDataQuery := fmt.Sprintf(`
	INSERT INTO users (name, mail) VALUES ('%s', '%s')
`, name, mail)

	return insertDataQuery
}

func FetchOneUser(name string, mail string) string {
	fetchUserQuery := fmt.Sprintf(`
	SELECT * 
	FROM users
	WHERE name='%s' AND mail='%s'
	LIMIT 1;`, name, mail)

	return fetchUserQuery
}
