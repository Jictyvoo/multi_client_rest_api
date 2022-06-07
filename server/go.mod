module github.com/jictyvoo/multi_client_rest_api/server

go 1.18

replace (
	github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core => ../modules/abz_1_core
	github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core => ../modules/xyc_2_core
	github.com/jictyvoo/multi_client_rest_api/services/apicontracts => ../services/apicontracts
)

require (
	github.com/gofiber/fiber/v2 v2.34.0
	github.com/gofiber/jwt/v3 v3.2.12
	github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core v0.0.0-00010101000000-000000000000
	github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core v0.0.0-00010101000000-000000000000
	github.com/jictyvoo/multi_client_rest_api/services/apicontracts v0.0.0
	github.com/pelletier/go-toml/v2 v2.0.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.12.0
	github.com/wrapped-owls/goremy-di/remy v1.1.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nyaruka/phonenumbers v1.0.75 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.37.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)
