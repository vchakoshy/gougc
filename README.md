# Gougc - Golang user generated content platform

![Coverage](https://img.shields.io/badge/Coverage-89.2%25-brightgreen)

## Run database

```bash
docker run --name go-postgres --rm -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres
```

## run app

```bash
make api
```

## Regenerate Swagger api

```bash
make swagger
```
