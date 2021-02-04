module server

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/labstack/echo/v4 v4.1.17
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	service v0.0.0
)

replace (
	model => ./model
	service => ./service
	conndb => ./conndb
)
