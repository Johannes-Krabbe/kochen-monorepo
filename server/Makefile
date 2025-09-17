gen_templ:
	templ generate ./internal/ui/

gen_sqlc:
	sqlc generate -f ./sqlc.yaml

gen_tailwind:
	npx @tailwindcss/cli -c ./internal/ui/tailwind.config.js -i ./internal/ui/main.css -o ./internal/ui/static/main.css

codegen:
	make gen_templ
	make gen_sqlc
	make gen_tailwind
	
migrate_up:
	go run cmd/migrator/main.go

build_server: gen_templ gen_tailwind
	rm -rf ./bin/server
	go build -o bin/server/server cmd/server/main.go
	mkdir -p ./bin/server/internal/ui
	cp -r ./internal/ui/static ./bin/server/internal/ui/static

build_migrator: gen_templ gen_tailwind
	rm -rf ./bin/migrator
	go build -o bin/migrator/migrator cmd/migrator/main.go
	mkdir -p ./bin/migrator/internal/database
	cp -r ./internal/database/schema ./bin/migrator/internal/database


run_server: codegen
	go run cmd/server/main.go

dev_db_start:
	docker-compose --file ./dev/compose-db.yml up -d

dev: dev_db_start
	air
