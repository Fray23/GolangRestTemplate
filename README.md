RUN migrations
goose postgres "postgres://postgres:postgres@database:5432/gorm3?sslmode=disable" up

