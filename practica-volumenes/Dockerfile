FROM postgres:16-alpine
ENV POSTGRES_USER=estudiante
ENV POSTGRES_PASSWORD=secreto123
COPY init.sql /docker-entrypoint-initdb.d/
VOLUME ["/var/lib/postgresql/data"]
EXPOSE 5432