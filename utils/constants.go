package utils

const (
	// DbType is the type of database to use
	DbType = "DBTYPE"

	// Sqlite is the name of the sqlite database
	Sqlite = "Sqlite"

	// Postgres is the name of the postgres database
	Postgres = "POSTGRES"

	// Mysql is the name of the mysql database
	Mysql = "MYSQL"

	// BasepathV1 is the base path for v1 of the api
	BasepathV1 = "/api/v1"

	// BasepathV2 is the base path for v2 of the api
	BasepathV2 = "/api/v2"
)

//Message is a simple message
type Message struct {
	Message string `json:"message"`
}
