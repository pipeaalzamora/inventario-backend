# Checklist go-live

## Tecnico

- [ ] `DEBUG=false`.
- [ ] `SECRET` y `JWT_SECRET` reales.
- [ ] Dominio frontend configurado.
- [ ] Dominio API configurado.
- [ ] CORS limitado a `FRONT_URL`.
- [ ] Postgres productivo configurado.
- [ ] Redis productivo configurado.
- [ ] Meilisearch productivo configurado.
- [ ] Correo real configurado.
- [ ] Bucket GCP real configurado.
- [ ] Migraciones aplicadas.
- [ ] Backup inicial tomado.
- [ ] Restore probado en ambiente separado.

## Seguridad

- [ ] Usuario admin temporal cambiado o eliminado.
- [ ] Contraseñas con bcrypt.
- [ ] Usuarios con solo empresas/tiendas necesarias.
- [ ] Proveedor externo validado por token temporal.
- [ ] Endpoints `/health` y `/metrics` monitoreados.

## Funcional

- [ ] Empresa creada.
- [ ] Tiendas creadas.
- [ ] Bodegas creadas.
- [ ] Usuarios y perfiles creados.
- [ ] Branding configurado.
- [ ] Plantillas de correo revisadas.
- [ ] Productos importados.
- [ ] Proveedores importados/asignados.
- [ ] Inventario inicial cargado.
- [ ] Compra de prueba creada.
- [ ] Correo de proveedor recibido.
- [ ] Recepcion de prueba completada.
- [ ] Stock final validado.

## Operacion

- [ ] Prometheus activo.
- [ ] Alertas activas.
- [ ] Responsable de soporte asignado.
- [ ] Manual de usuario entregado.
- [ ] Procedimiento de soporte entregado.
