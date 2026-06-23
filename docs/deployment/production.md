# Deploy productivo

## Variables

1. Copiar `.env.production.example` a `.env.production`.
2. Reemplazar todos los secretos y URLs por valores reales.
3. Usar `DEBUG=false`.
4. Mantener `SECRET` y `JWT_SECRET` largos, aleatorios y distintos.

## Migraciones

Antes de levantar una version nueva:

```bash
set -a
source .env.production
set +a
scripts/backup-postgres.sh
scripts/apply-migrations.sh
```

`scripts/apply-migrations.sh` aplica solo `migrations/20*.sql`. `schema.sql` y `seed.sql` son para inicializar entornos nuevos, no para actualizar produccion.

## Backup y restore

Los backups quedan en `BACKUP_DIR` o en `./backups` si no se define. Para restaurar:

```bash
pg_restore -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -d "$POSTGRES_DB" --clean --if-exists backups/archivo.dump
```

## Compose base

```bash
SOFIA_BACKEND_IMAGE=registry.example.com/sofia-backend:tag docker compose -f docker-compose.prod.yaml up -d
```

PostgreSQL queda fuera del compose productivo por defecto para poder usar una base administrada, backups separados y politicas de retencion propias.
