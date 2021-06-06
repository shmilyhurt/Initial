module Initial

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	gopkg.in/go-playground/validator.v9 v9.29.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace github.com/gin-contrib/sse v0.1.0 => github.com/e421083458/sse v0.1.1
