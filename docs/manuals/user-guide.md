# Manual de usuario

## Inicio

1. Entrar al sitio del cliente.
2. Iniciar sesion con correo y contraseña.
3. Seleccionar empresa desde el selector lateral.
4. Confirmar que las tiendas disponibles correspondan al usuario.

## Personalizacion

Ruta: `Mi empresa > Personalizacion`

Permite configurar:

- nombre de la aplicacion,
- logo,
- favicon,
- colores,
- correo de soporte,
- dominio/subdominio,
- textos,
- plantillas de correo,
- archivos de marca,
- importaciones iniciales.

## Importacion inicial

1. Elegir tipo: productos, proveedores o inventario.
2. Subir CSV/XLSX.
3. Revisar preview.
4. Mapear columnas requeridas.
5. Ejecutar importacion.
6. Revisar procesadas y errores.

Campos minimos:

- productos: nombre y SKU.
- proveedores: nombre, ID fiscal y email.
- inventario: ID producto tienda, ID bodega y cantidad.

## Compras

1. Crear solicitud o crear orden de compra directa.
2. Seleccionar proveedor y productos de tienda.
3. Confirmar precios y cantidades.
4. El sistema genera token y envia correo al proveedor.
5. El proveedor revisa y aprueba/rechaza desde enlace externo.

## Recepcion

1. Abrir orden de compra.
2. Crear recepcion.
3. Adjuntar documentos.
4. Confirmar cantidades recibidas.
5. Completar recepcion.
6. El stock se mueve desde bodega de transicion a bodega destino.

## Inventario

Permite revisar:

- stock por tienda,
- stock por bodega,
- movimientos,
- alertas de minimo/maximo,
- conteos de inventario.

## Usuarios y permisos

Los usuarios ven empresas y tiendas segun perfiles asignados. Para multi-cliente, revisar siempre que el usuario tenga solo:

- `company:<id>` de su empresa,
- `store:<id>` de sus tiendas,
- permisos funcionales necesarios.
