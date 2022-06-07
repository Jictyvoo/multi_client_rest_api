# Install prerequisites
#
i-dev-deps:
	go install github.com/golang/mock/mockgen@v1.6.0

gen-mock:
	cd ./modules/abz_1_core && mockgen -source=./internal/domain/interfaces/repositories.go -destination=./internal/domain/mocks/repositories_mock.go -package=mocks
	cd ./modules/xyc_2_core && mockgen -source=./internal/domain/interfaces/repositories.go -destination=./internal/domain/mocks/repositories_mock.go -package=mocks
