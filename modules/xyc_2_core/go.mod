module github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core

go 1.18

replace github.com/jictyvoo/multi_client_rest_api/services/apicontracts => ../../services/apicontracts

require (
	github.com/golang/mock v1.6.0
	github.com/jictyvoo/multi_client_rest_api/services/apicontracts v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.6
	github.com/nyaruka/phonenumbers v1.0.75
	github.com/wrapped-owls/goremy-di/remy v1.2.1
)

require (
	github.com/golang/protobuf v1.3.2 // indirect
	golang.org/x/text v0.3.7 // indirect
)
