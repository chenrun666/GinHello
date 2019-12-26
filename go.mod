module GinHello

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/magiconair/properties v1.8.1
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	GinHello/handler => ./handler
	GinHello/initDB => ./initDB
	GinHello/initRoute => ./initRoute
	GinHello/middleware => ./middleware
	GinHello/model => ./model
)
