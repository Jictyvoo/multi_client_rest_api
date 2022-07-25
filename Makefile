# Install prerequisites
dev-deps:
	go install github.com/golang/mock/mockgen@v1.6.0

gen-mock:
	cd ./modules && mockgen -source=./internal/domain/interfaces/repositories.go -destination=./internal/domain/mocks/repositories_mock.go -package=mocks

# Build the binary
build:
	cd ./server && go mod download
	cd ./server && go build -o ../multiservice-server .
