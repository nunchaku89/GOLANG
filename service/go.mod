module service

go 1.15

require (
	conndb v0.0.0
	github.com/labstack/echo/v4 v4.1.17
	gopkg.in/guregu/null.v3 v3.5.0
	model v0.0.0
)

replace (
	conndb => ../conndb
	model => ../model
)
