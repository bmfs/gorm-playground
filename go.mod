module gorm.io/playground

go 1.16

require (
	github.com/denisenkom/go-mssqldb v0.12.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.13.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/microsoft/go-mssqldb v0.19.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	gorm.io/driver/mysql v1.4.4
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlite v1.4.3
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.2
)

replace gorm.io/gorm => ./gorm
