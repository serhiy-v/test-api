.PHONY: run
run:
	go run cmd/main.go

.PHONY: start-db
start-db: ## Run local postgres
	if [ ! "$$(docker ps | grep postgres13)" ]; then \
    		docker run --rm --name postgres13 \
    			--env POSTGRES_PASSWORD=secret \
    			--env POSTGRES_USER=postgres \
    			--env POSTGRES_DB=test-api \
    			--publish 5432:5432 \
    			-d postgres:13-alpine ;\
    	fi

.PHONY: stop-db
stop-db: ## Stop local postgres
	docker stop postgres13

.PHONY: migrate
migrate: ## Run local migrations
	migrate -path pkg/db/migrations -database 'postgres://postgres:secret@localhost:5432/test-api?sslmode=disable' up