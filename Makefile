start:
	touch config/keto/sqlite/db.sqlite
	docker-compose up -d
	go run main.go