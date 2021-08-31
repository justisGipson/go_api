.PHONY: clean test security build run

APP_NAME = bluebird
BUILD_DIR = $(PWD)/build
MIGRATIONS_DIR = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost/postgres?sslmode=disable

clean:
	rm -rf ./build

security:
	gosec -quiet ./..

test: security
	go test -v timeout 30s -coverprofile=cover.out -cover ./..

build: clean test
	CGO_ENABLED=0 go build -ldflags='-w -s' -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" force $(version)

docker.run: docker.network docker.postgres swag docker.fiber migrate.up

# TODO: Finish docker commands + single swag command
