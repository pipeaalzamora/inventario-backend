# Manual de instalacion

## Requisitos

- Docker y Docker Compose.
- PostgreSQL 14 o superior.
- Redis.
- Meilisearch.
- Servicio de correo compatible con la configuracion del backend.
- Bucket GCP para archivos e imagenes.
- Dominio para frontend y API.

## Variables

1. Crear `.env.production` desde `.env.production.example`.
2. Configurar `DEBUG=false`.
3. Configurar `BASE_URL` y `FRONT_URL` con dominios reales.
4. Configurar secretos fuertes para `SECRET` y `JWT_SECRET`.
5. Configurar Postgres, Redis, Meilisearch, correo y bucket.

## Base de datos nueva

Para una instalacion vacia:

```bash
psql -h "$POSTGRES_HOST" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -v ON_ERROR_STOP=1 -f migrations/schema.sql
psql -h "$POSTGRES_HOST" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -v ON_ERROR_STOP=1 -f migrations/seed.sql
scripts/apply-migrations.sh
```

## Actualizacion de version

```bash
set -a
source .env.production
set +a
scripts/backup-postgres.sh
scripts/apply-migrations.sh
SOFIA_BACKEND_IMAGE=registry.example.com/sofia-backend:tag docker compose -f docker-compose.prod.yaml up -d
```

## Verificacion

```bash
curl "$BASE_URL/api/v1/health"
curl "$BASE_URL/api/v1/metrics"
```

Luego iniciar sesion desde el frontend y validar:

- empresas visibles,
- tiendas visibles,
- dashboard,
- personalizacion,
- importacion,
- compras,
- recepcion,
- stock.
