# Procedimiento de soporte

## Severidades

- `S1`: sistema caido, login imposible, perdida de datos o compras/inventario bloqueados para todo el cliente.
- `S2`: flujo principal bloqueado para una sucursal, errores 5xx repetidos o correo/archivos sin funcionar.
- `S3`: problema funcional con alternativa manual.
- `S4`: consulta, ajuste menor o mejora.

## Datos minimos para abrir caso

- Cliente y empresa afectada.
- Usuario afectado.
- Fecha y hora aproximada.
- URL o pantalla.
- Pasos para reproducir.
- Captura o mensaje de error.
- ID de compra, recepcion, producto, tienda o bodega si aplica.

## Diagnostico inicial

1. Confirmar `/api/v1/health`.
2. Revisar metricas y alertas.
3. Revisar logs por ventana horaria.
4. Reproducir con usuario admin de soporte si existe.
5. Confirmar permisos `company:<id>` y `store:<id>` del usuario.
6. Confirmar estado de Postgres, Redis, Meilisearch, correo y bucket.

## Escalamiento

- `S1`: respuesta inmediata, workaround o rollback, informe al cliente cada 30 minutos.
- `S2`: respuesta el mismo dia, informe cada 2 horas.
- `S3`: planificar correccion en siguiente release menor.
- `S4`: backlog comercial/producto.

## Cierre

Todo caso cerrado debe incluir:

- Causa probable o confirmada.
- Accion correctiva.
- Prueba realizada.
- Si requiere cambio preventivo, issue/tarea asociada.
