We use golang-migrate to handle DB migrations

- migrate create -ext sql -dir db/migration -seq init_schema
