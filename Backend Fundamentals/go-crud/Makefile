# Using goose to migrate database as well as maintain a rollback in case I want to clear db!
# Up migrations in goose
migrationUp:
	cd sql/schema && goose postgres "postgres://cheems:cheems@localhost:5432/crud?sslmode=disable" up && cd .. && cd ..

# Down migrations in goose
migrationDown:
	cd sql/schema && goose postgres "postgres://cheems:cheems@localhost:5432/crud?sslmode=disable" down && cd .. && cd ..