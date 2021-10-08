.PHONY: clean test security build run

APP_NAME = bluebird
BUILD_DIR = $(PWD)/build
MIGRATIONS_DIR = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost/postgres?sslmode=disable

clean:
	rm -rf ./build

# security:
# 	gosec -quiet ./..

# test: security
# 	go test -v timeout 30s -coverprofile=cover.out -cover ./..

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

docker.run: docker.network docker.postgres docker.nginx swag docker.fiber migrate.up

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network


docker.fiber.build: docker build -t fiber

docker.fiber: docker.fiber.build
	docker run --rm -d \
	--name dev-fiber \
	--network dev-network \
	-p 5000:5000 \
	fiber

docker.postgres:
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.stop: docker.stop.fiber docker.stop.postgres

docker.stop.fiber:
	docker stop dev-fiber

docker.stop.postgres:
	docker stop dev-postgres

swag:
	swag init
