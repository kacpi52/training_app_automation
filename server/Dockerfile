FROM golang:1.23-alpine3.20 AS base

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o main .

# Skrypt entrypoint
COPY migrateAndTest.sh /app/migrateAndTest.sh
RUN chmod +x /app/migrateAndTest.sh
# @TODO check that migrations in main are needed 
FROM base AS test
CMD ["/app/migrateAndTest.sh"]

FROM base AS production
CMD ["/app/main"]