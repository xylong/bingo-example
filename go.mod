module bingo-example

go 1.18

replace github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220816024939-bc8df83d7b9d => github.com/antlr/antlr4 v0.0.0-20200712162734-eb1adaa8a7a6

require (
	github.com/KenmyZhang/aliyun-communicate v0.0.0-20180308134849-7997edc57454
	github.com/gin-gonic/gin v1.8.2
	github.com/go-playground/validator/v10 v10.11.1
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gorilla/websocket v1.5.0
	github.com/graphql-go/graphql v0.8.0
	github.com/graphql-go/handler v0.2.3
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/mojocn/base64Captcha v1.3.5
	github.com/olivere/elastic/v7 v7.0.32
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.0.2
	github.com/spf13/cast v1.5.0
	github.com/spf13/viper v1.15.0
	github.com/swaggo/files v1.0.0
	github.com/swaggo/gin-swagger v1.5.3
	github.com/swaggo/swag v1.8.10
	github.com/thedevsaddam/govalidator v1.9.10
	github.com/ulule/limiter/v3 v3.11.1
	github.com/xylong/bingo v0.2.6
	go.mongodb.org/mongo-driver v1.10.1
	go.uber.org/zap v1.23.0
	golang.org/x/crypto v0.7.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.3.4
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.0
)

replace github.com/xylong/bingo => ../bingo

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/antlr/antlr4 v0.0.0-20200712162734-eb1adaa8a7a6 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.8 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/shenyisyn/goft-expr v0.3.0 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
