# Logs, monitoreo y alertas

## Logs

El backend escribe logs HTTP estructurados con `slog` por cada request:

- `method`
- `path`
- `status`
- `latency_ms`
- `client_ip`

En produccion ejecutar con `DEBUG=false` y recolectar stdout/stderr desde el orquestador usado por el cliente.

## Healthcheck

Endpoint:

```bash
curl https://api.example.com/api/v1/health
```

Respuesta esperada:

```json
{"status":"ok"}
```

## Metricas

Endpoint Prometheus:

```bash
curl https://api.example.com/api/v1/metrics
```

Metricas expuestas:

- `sofia_http_requests_total`
- `sofia_http_errors_total`
- `sofia_http_in_flight_requests`
- `sofia_http_request_duration_seconds_avg`
- `sofia_http_responses_total{class="2xx|3xx|4xx|5xx"}`

## Prometheus local

```bash
docker compose -f docker-compose.monitoring.yaml up -d
```

Abrir `http://localhost:9090`.

## Alertas iniciales

Las reglas viven en `ops/alerts.yml`:

- API caida por mas de 2 minutos.
- Mas de 5 respuestas 5xx en 5 minutos.
- Latencia promedio superior a 1 segundo durante 10 minutos.

## Runbook rapido

1. Revisar `/api/v1/health`.
2. Revisar `/api/v1/metrics`.
3. Revisar logs del contenedor API.
4. Revisar Postgres, Redis y Meilisearch.
5. Si el problema empezo despues de deploy, restaurar imagen anterior y revisar migraciones.
6. Si hay perdida o corrupcion de datos, detener escritura, tomar backup inmediato y restaurar desde ultimo backup verificado.
