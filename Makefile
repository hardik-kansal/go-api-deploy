DB_URL=postgresql://root:123@localhost:5432/simplebank?sslmode=disable
rm:
	docker stop pg16
	docker rm pg16
postgres:
	docker run --name pg16 -p 5432:5432 -e POSTGRES_PASSWORD=123 -e POSTGRES_USER=root -d postgres:16-alpine
createdb:
	docker exec -it pg16 createdb --username=root --owner=root simplebank
dropdb:
	docker exec -it pg16 dropdb simplebank
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down



	

