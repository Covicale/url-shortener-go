FROM postgres:16.0

COPY internal/db/schema/schema.sql /docker-entrypoint-initdb.d/