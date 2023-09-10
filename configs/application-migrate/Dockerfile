FROM postgres:latest

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app

COPY . .
RUN chmod +x /app/migrate.sh

CMD ["/app/migrate.sh"]