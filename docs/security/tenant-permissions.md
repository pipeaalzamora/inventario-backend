# Auditoria de permisos multi-cliente

## Regla operativa

Todo endpoint autenticado que recibe `companyId` debe validar `company:<companyId>`.
Todo endpoint autenticado que recibe `storeId`, `warehouseId`, `storeProductId`, compra, recepcion o inventario debe terminar validando el `store:<storeId>` propietario y, cuando exista `companyId`, tambien `company:<companyId>`.

## Flujos revisados y cerrados

- Personalizacion de empresa: marca, archivos, plantillas e importaciones validan `company:<id>`.
- Importacion de inventario: antes de escribir stock valida que `product_per_store`, `warehouse` y `store.company_id` pertenecen a la empresa del job.
- Tiendas por empresa: `GetStoresByCompanyID` ya no devuelve tiendas no incluidas en los poderes `store:<id>`.
- Recepciones por tienda: `GetAllDeliveryPurchaseNotes` valida `store:<id>` antes de consultar.
- Solicitudes de inventario: crear solicitudes usa la tienda enviada y valida que pertenece a la empresa indicada.
- Productos por tienda, compras, movimientos e inventarios operativos validan poderes de tienda en servicio.

## Puntos que deben seguir revisandose antes de cada venta

- Los endpoints externos de proveedor con token no usan usuario autenticado; su frontera de seguridad es el token temporal de OC.
- Proveedores base son catalogo global, pero la asignacion visible/usable por cliente debe pasar por `supplier_per_company`.
- Cualquier nuevo reporte debe filtrar por companias/tiendas derivadas del contexto, no por parametros libres del cliente.

## Comando de apoyo

Ejecutar:

```bash
scripts/audit-tenant-permissions.sh
```

El script lista servicios/facades con parametros sensibles y ayuda a detectar metodos nuevos sin `EveryPower`, `SomePower` o validacion de pertenencia.
