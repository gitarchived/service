all:
	@make api
	@make lister
	@make updater
	@make deleter

run:
	@sh ./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))

api:
	@go build -o ./bin/api ./cmd/api/main.go

lister:
	@go build -o ./bin/lister ./cmd/lister/main.go

updater:
	@go build -o ./bin/updater ./cmd/updater/main.go

deleter:
	@go build -o ./bin/deleter ./cmd/deleter/main.go
