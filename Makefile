all:
	@make api

run:
	@sh ./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))

api:
	@go build -o ./bin/api ./cmd/api/main.go

clean:
	@rm -rf ./bin
