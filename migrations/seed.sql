INSERT INTO public.country OVERRIDING SYSTEM VALUE VALUES (1, 'CL', 'Chile');
INSERT INTO public.currency OVERRIDING SYSTEM VALUE VALUES (1, 'CLP', 152, 'Peso Chileno', '$', 0, 1, true);
INSERT INTO public.currency OVERRIDING SYSTEM VALUE VALUES (2, 'USD', 840, 'Dólar Estadounidense', 'US$', 2, 970.87, true);
INSERT INTO public.currency OVERRIDING SYSTEM VALUE VALUES (3, 'EUR', 978, 'Euro', '€', 2, 1125.20, true);
INSERT INTO public.currency OVERRIDING SYSTEM VALUE VALUES (4, 'PEN', 604, 'Sol Peruano', 'S/', 2, 273.21, true);
INSERT INTO public.currency OVERRIDING SYSTEM VALUE VALUES (5, 'COP', 170, 'Peso Colombiano', '$', 0, 0.24, true);
INSERT INTO public.currency OVERRIDING SYSTEM VALUE VALUES (6, 'ARS', 32, 'Peso Argentino', '$', 2, 0.73, true);
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (1, 'GTIN/EAN-8', 'Código de barras de 8 dígitos usado para identificar productos pequeños en el comercio minorista');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (2, 'GTIN/UPC-A', 'Código de barras de 12 dígitos ampliamente utilizado en Norteamérica para identificar productos');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (3, 'GTIN/EAN-13', 'Código de barras de 13 dígitos usado internacionalmente para identificar productos');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (4, 'GTIN-14', 'Código de 14 dígitos que identifica agrupaciones o embalajes logísticos de productos');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (5, 'PLU', 'Código numérico de 4 o 5 dígitos usado para identificar productos agrícolas vendidos a granel');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (6, 'SSCC', 'Código de 18 dígitos que identifica de forma única unidades logísticas como pallets o contenedores');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (7, 'ISBN', 'Código único para identificar libros, revistas y publicaciones monográficas');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (8, 'GS1-128', 'Formato de código de barras que permite incluir múltiples datos estructurados según estándares GS1');
INSERT INTO public.code_kind OVERRIDING SYSTEM VALUE VALUES (9, 'GIAI', 'Identificador único para activos individuales de una empresa, como equipos o maquinaria');
INSERT INTO public.economic_activity_class OVERRIDING SYSTEM VALUE VALUES (1, 1, 'Clasificación Industrial Internacional Uniforme', 'CIIU-REV4-CL-01');
INSERT INTO public.economic_activity_class OVERRIDING SYSTEM VALUE VALUES (2, 1, 'Clasificación Industrial Internacional Uniforme', 'CIIU-REV4-CL-02');
INSERT INTO public.economic_activity_class OVERRIDING SYSTEM VALUE VALUES (3, 1, 'Clasificación Industrial Internacional Uniforme', 'CIIU-REV4-CL-03');
INSERT INTO public.economic_activity OVERRIDING SYSTEM VALUE VALUES (1, 1, 'Comercio al por menor de productos alimenticios', '4711', 'Venta de alimentos y bebidas en locales especializados');
INSERT INTO public.economic_activity OVERRIDING SYSTEM VALUE VALUES (2, 1, 'Comercio al por menor de bebidas alcohólicas', '4725', 'Venta de cervezas, vinos y licores');
INSERT INTO public.economic_activity OVERRIDING SYSTEM VALUE VALUES (3, 2, 'Desarrollo de software a medida', '6201', 'Servicios de desarrollo, mantenimiento y soporte de software personalizado');
INSERT INTO public.economic_activity OVERRIDING SYSTEM VALUE VALUES (4, 2, 'Consultoría en tecnologías de la información', '6202', 'Asesoría y gestión de proyectos tecnológicos');
INSERT INTO public.economic_activity OVERRIDING SYSTEM VALUE VALUES (5, 3, 'Transporte de carga por carretera', '4941', 'Transporte de mercancías dentro y fuera de la ciudad');
INSERT INTO public.economic_activity OVERRIDING SYSTEM VALUE VALUES (6, 3, 'Almacenamiento y distribución', '5210', 'Servicios de bodegaje y logística de distribución');
INSERT INTO public.fiscal_data VALUES ('cdb6f2eb-596f-41e7-acc5-52c5b567f8cf', 'CL-76.123.456-7', '76123456-7', 'Distribuidora Los Andes Ltda.', 'Av. Libertador Bernardo O’Higgins 1234', 'Región Metropolitana', 'Santiago', 'contacto@losandes.cl');
INSERT INTO public.fiscal_data VALUES ('d0c56423-7127-49a3-a2b0-2d7023141cf0', 'CL-77.987.654-3', '77987654-3', 'Soluciones Digitales SpA', 'Calle Nueva Providencia 456', 'Región Metropolitana', 'Santiago', 'ventas@solucionesdigitales.cl');
INSERT INTO public.fiscal_data VALUES ('e1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c', 'CL-78.654.321-2', '78654321-2', 'Proveedor Ejemplo S.A.', 'Av. Ejemplo 123', 'Región Metropolitana', 'Santiago', 'contacto@proveedorejemplo.cl');
INSERT INTO public.fiscal_data VALUES ('f1a1b1c1-d1e1-4f1a-8b1c-0d1e2f3a4b1c', 'CL-76.111.222-3', '76111222-3', 'Frutas del Valle Ltda.', 'Camino Valle 123', 'Región de O’Higgins', 'Rancagua', 'contacto@frutasdelvalle.cl');
INSERT INTO public.fiscal_data VALUES ('f2a2b2c2-d2e2-4f2a-8b2c-0d2e2f3a4b2c', 'CL-77.222.333-4', '77222333-4', 'Panadería El Trigal SpA', 'Av. Trigal 456', 'Región Metropolitana', 'Santiago', 'ventas@eltrigal.cl');
INSERT INTO public.fiscal_data VALUES ('f3a3b3c3-d3e3-4f3a-8b3c-0d3e2f3a4b3c', 'CL-78.333.444-5', '78333444-5', 'Carnes Premium S.A.', 'Calle Carnes 789', 'Región del Maule', 'Talca', 'info@carnespremium.cl');
INSERT INTO public.fiscal_data VALUES ('f4a4b4c4-d4e4-4f4a-8b4c-0d4e2f3a4b4c', 'CL-79.444.555-6', '79444555-6', 'Verduras Frescas Ltda.', 'Camino Verde 321', 'Región de Valparaíso', 'Valparaíso', 'contacto@verdurasfrescas.cl');
INSERT INTO public.fiscal_data VALUES ('f5a5b5c5-d5e5-4f5a-8b5c-0d5e2f3a4b5c', 'CL-80.555.666-7', '80555666-7', 'Bebidas y Más SpA', 'Av. Bebidas 654', 'Región de Los Lagos', 'Puerto Montt', 'ventas@bebidasymas.cl');
INSERT INTO public.fiscal_data VALUES ('a6b6c6d6-e6f6-4a6b-8c6d-0e6f2a3b4c6d', 'CL-81.111.777-8', '81111777-8', 'Aceites del Pacífico S.A.', 'Av. Pacífico 101', 'Región de Tarapacá', 'Iquique', 'contacto@aceitespacifico.cl');
INSERT INTO public.fiscal_data VALUES ('a7b7c7d7-e7f7-4a7b-8c7d-0e7f2a3b4c7d', 'CL-82.222.888-9', '82222888-9', 'Lácteos del Sur Ltda.', 'Calle Sur 202', 'Región de Los Ríos', 'Valdivia', 'info@lacteosdelsur.cl');
INSERT INTO public.fiscal_data VALUES ('a8b8c8d8-e8f8-4a8b-8c8d-0e8f2a3b4c8d', 'CL-83.333.999-0', '83333999-0', 'Pastas Italia SpA', 'Av. Italia 303', 'Región Metropolitana', 'Santiago', 'ventas@pastasitalia.cl');
INSERT INTO public.fiscal_data VALUES ('b5c5d5e5-f5a5-4b5c-8d5e-0f5a2b3c4d5e', 'CL-90.000.666-7', '90000666-7', 'Cereales Premium SpA', 'Calle Premium 1010', 'Región de Los Lagos', 'Osorno', 'info@cerealespremium.cl');
INSERT INTO public.fiscal_data VALUES ('b4c4d4e4-f4a4-4b4c-8d4e-0f4a2b3c4d4e', 'CL-89.999.555-6', '89999555-6', 'Snacks Express Ltda.', 'Av. Express 909', 'Región de Valparaíso', 'Viña del Mar', 'contacto@snacksexpress.cl');
INSERT INTO public.fiscal_data VALUES ('b3c3d3e3-f3a3-4b3c-8d3e-0f3a2b3c4d3e', 'CL-88.888.444-5', '88888444-5', 'Bebidas del Centro S.A.', 'Calle Centro 808', 'Región Metropolitana', 'Santiago', 'ventas@bebidasdelcentro.cl');
INSERT INTO public.fiscal_data VALUES ('b2c2d2e2-f2a2-4b2c-8d2e-0f2a2b3c4d2e', 'CL-87.777.333-4', '87777333-4', 'Legumbres Chile Ltda.', 'Av. Legumbres 707', 'Región del Maule', 'Talca', 'info@legumbreschile.cl');

INSERT INTO public.fiscal_per_activity OVERRIDING SYSTEM VALUE VALUES (1, 'cdb6f2eb-596f-41e7-acc5-52c5b567f8cf', 1);
INSERT INTO public.fiscal_per_activity OVERRIDING SYSTEM VALUE VALUES (2, 'cdb6f2eb-596f-41e7-acc5-52c5b567f8cf', 2);
INSERT INTO public.fiscal_per_activity OVERRIDING SYSTEM VALUE VALUES (3, 'd0c56423-7127-49a3-a2b0-2d7023141cf0', 3);
INSERT INTO public.fiscal_per_activity OVERRIDING SYSTEM VALUE VALUES (4, 'd0c56423-7127-49a3-a2b0-2d7023141cf0', 4);

INSERT INTO public.company VALUES ('43f96ad5-6c50-4ce8-b7cf-203210741a18', 1, 'cdb6f2eb-596f-41e7-acc5-52c5b567f8cf', 'Distribuidora Los Andes', 'Mayorista y minorista de alimentos y bebidas.', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmI_HZ9XAtXhZDFVRA8o5rpI_JPUptsCK93Q&s', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.company VALUES ('5d3984dd-58ef-42df-84ce-deb13dae434c', 1, 'd0c56423-7127-49a3-a2b0-2d7023141cf0', 'Soluciones Digitales', 'Empresa de software y servicios digitales.', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSbTPUd8HcXh3XeK9K0Lep-NcqJKixK0OJaSQ&s', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');

INSERT INTO public.power_account_categories VALUES ('e10af3a0-67af-42a8-88c4-0d10d332f6a9', 'Usuarios', 'Permisos relacionados con gestión de usuarios', false);
INSERT INTO public.power_account_categories VALUES ('605bbd6c-1fb2-4266-9881-d20b231ef44e', 'Perfiles', 'Permisos relacionados con gestión de perfiles', false);
INSERT INTO public.power_account_categories VALUES ('0e496576-580c-4a9d-8009-1855b28d8f96', 'Propiedad_Tienda', 'Permisos relacionados con gestión de perfiles', true);
INSERT INTO public.power_account_categories VALUES ('1fdbaca4-e12b-4378-a670-fb15bb509e93', 'Propiedad_Empresa', 'Permisos relacionados con gestión de perfiles', true);
INSERT INTO public.power_account_categories VALUES ('6465e10a-dbb9-4000-919b-9156a3762af3', 'Compras_Tienda', 'Permisos relacionados con gestión de compras en tienda', false);
INSERT INTO public.power_account_categories VALUES ('db358d6d-2cb8-476d-8a95-41fffb7d0a8c', 'Solicitudes_Tienda', 'Permisos relacionados con gestión de solicitudes en tienda', false);
INSERT INTO public.power_account_categories VALUES ('89c53c8d-5c75-4c95-9810-46e6903b1bcc', 'Empresas', 'Permisos relacionados con gestión de empresas', false);
INSERT INTO public.power_account_categories VALUES ('a1584804-7c4e-4ebc-b762-96b5bf1cc901', 'Tiendas', 'Permisos relacionados con gestion de tiendas',false);
INSERT INTO public.power_account_categories VALUES ('b1e5287d-b503-43d3-b555-ebfb8af5dbaa', 'Proveedores', 'Permisos relacionados con gestión de proveedores', false);
INSERT INTO public.power_account_categories VALUES ('f2a8c4e1-9b3d-4f6a-8c2e-1d5f7a9b3e4c', 'Bodegas', 'Permisos relacionados con gestión de bodegas', false);
INSERT INTO public.power_account_categories VALUES ('a3b9d5e2-8c4f-4a7b-9d3e-2f6a8c1e4b5d', 'Productos', 'Permisos relacionados con gestión de productos', false);
INSERT INTO public.power_account_categories VALUES ('c4d0e6f3-7b5a-4c8d-ae4f-3a7b9d2f5c6e', 'Productos_Empresa', 'Permisos relacionados con gestión de productos por empresa', false);
INSERT INTO public.power_account_categories VALUES ('d5e1f7a4-6c8b-4d9e-bf5a-4b8c0e3a6d7f', 'Conteo_Inventario', 'Permisos relacionados con conteo de inventario', false);
INSERT INTO public.power_account_categories VALUES ('e6f2a8b5-5d9c-4e0f-ca6b-5c9d1f4b7e8a', 'Movimiento_Productos', 'Permisos relacionados con movimiento de productos', false);
INSERT INTO public.power_account_categories VALUES ('c1f4d3e2-5f4e-4d3a-9f4e-8c6f4e2b5a0b', 'Plantillas de productos', 'Permisos relaciones con la gestion de plantillas de productos', false);

INSERT INTO public.power_account_categories VALUES ('c2f4e8a9-4d6b-4c7e-9f8a-1b3d5e7f9a2c', 'Unidades Medida', 'Permisos relacionados con gestión de unidades de medida', false);
INSERT INTO public.power_account_categories VALUES ('d3a5b9c7-8e2f-4a1d-bf3c-6e9d4f7a2b8e', 'Productos_Tienda', 'Permisos relacionados con gestión de productos por tienda', false);
INSERT INTO public.power_account_categories VALUES ('f4b6c9d8-7e3f-4a2b-8c5d-9e0f1a2b3c4d', 'Notas_Entrega_Compra', 'Permisos relacionados con gestión de notas de entrega de compra', false);
INSERT INTO public.power_accounts VALUES ('1a866441-6c83-43ee-9061-8b37ecdba341', 'user:create', 'Crear usuario', 'Permite la creación de nuevos usuarios', 'e10af3a0-67af-42a8-88c4-0d10d332f6a9');
INSERT INTO public.power_accounts VALUES ('cad12d0b-1dbe-426d-9164-abacbca58e3e', 'user:update', 'Editar usuario', 'Permite editar la información de un usuario', 'e10af3a0-67af-42a8-88c4-0d10d332f6a9');
INSERT INTO public.power_accounts VALUES ('9d70c941-d97d-4eea-adfe-23bd9430dd5f', 'user:enable-disable', 'Activar/Desactivar usuario', 'Permite activar o desactivar un usuario de la plataforma', 'e10af3a0-67af-42a8-88c4-0d10d332f6a9');
INSERT INTO public.power_accounts VALUES ('1587fa74-c5c4-4221-9a03-c8e1a47b6851', 'profile:create', 'Crear perfil', 'Permite la creación de nuevos perfiles', '605bbd6c-1fb2-4266-9881-d20b231ef44e');
INSERT INTO public.power_accounts VALUES ('de092967-af75-4142-8390-6478b23b51be', 'profile:update', 'Editar perfil', 'Permite editar la información de un perfil existente', '605bbd6c-1fb2-4266-9881-d20b231ef44e');
INSERT INTO public.power_accounts VALUES ('3880287a-f927-462c-ad6e-bce4c3a1b372', 'profile:delete', 'Eliminar un perfil', 'Permite eliminar un perfil', '605bbd6c-1fb2-4266-9881-d20b231ef44e');
INSERT INTO public.power_accounts VALUES ('0ac945fd-183d-4593-8d19-9deaa312183b', 'request:create', 'Crear Solicitud', 'Permite la creación de solicitudes en la tienda', 'db358d6d-2cb8-476d-8a95-41fffb7d0a8c');
INSERT INTO public.power_accounts VALUES ('330aefab-a77d-4b58-8c9f-8f36be9d8bfe', 'request:update', 'Editar Solicitud', 'Permite editar la información de una solicitud en la tienda', 'db358d6d-2cb8-476d-8a95-41fffb7d0a8c');
INSERT INTO public.power_accounts VALUES ('a1f2e3d4-c5b6-47a8-9b0c-1d2e3f4a5b6c', 'request:createForStore', 'Crear Solicitud para Tienda', 'Permite crear solicitudes de inventario para tiendas', 'db358d6d-2cb8-476d-8a95-41fffb7d0a8c');
INSERT INTO public.power_accounts VALUES ('b2f3e4d5-c6b7-48a9-0c1d-2e3f4a5b6c7d', 'request:createForCompany', 'Crear Solicitud para Empresa', 'Permite crear solicitudes de inventario para empresas', 'db358d6d-2cb8-476d-8a95-41fffb7d0a8c');
INSERT INTO public.power_accounts VALUES ('c3f4e5d6-c7b8-49a0-1d2e-3f4a5b6c7d8e', 'request:createForSupplier', 'Crear Solicitud para Proveedor', 'Permite crear solicitudes de inventario para proveedores', 'db358d6d-2cb8-476d-8a95-41fffb7d0a8c');
INSERT INTO public.power_accounts VALUES ('57eec64a-9e5a-4b2c-8ba2-7e18b845f7bb', 'purchase:create', 'Crear compra', 'Permite la creación de compras en la tienda', '6465e10a-dbb9-4000-919b-9156a3762af3');
INSERT INTO public.power_accounts VALUES ('03e72386-429c-4dfa-b33b-b6ca48b5154f', 'purchase:update', 'Editar compra', 'Permite editar la información de una compra en la tienda', '6465e10a-dbb9-4000-919b-9156a3762af3');
INSERT INTO public.power_accounts VALUES ('f8a1b2c3-d4e5-4f6a-7b8c-9d0e1f2a3b4c', 'purchase:approve', 'Aprobar compra', 'Permite aprobar una solicitud de compra en la tienda', '6465e10a-dbb9-4000-919b-9156a3762af3');
INSERT INTO public.power_accounts VALUES ('6ca241a0-d1e4-40e6-bc6b-cb2981dc8caf', 'request:resolve', 'Resolver compra', 'Permite resolver el estado de una compra en la tienda', '6465e10a-dbb9-4000-919b-9156a3762af3');
INSERT INTO public.power_accounts VALUES ('685b363e-170e-4e95-9d2c-22ca1682376e', 'store:ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'Propiedad Sucursal Santiago', 'Este permiso permite el acceso a la tienda Sucursal Santiago', '0e496576-580c-4a9d-8009-1855b28d8f96');
INSERT INTO public.power_accounts VALUES ('4cb22fae-bbce-4fc0-954d-fdf2b3622617', 'store:9f666bab-35e3-46e6-b147-75354262dc84', 'Propiedad Oficina Central', 'Este permiso permite el acceso a la tienda Oficina Central', '0e496576-580c-4a9d-8009-1855b28d8f96');
INSERT INTO public.power_accounts VALUES ('4ead438b-364f-4ef3-9b76-0145ecab4021', 'store:86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'Propiedad Bodega Central', 'Este permiso permite el acceso a la tienda Bodega Central', '0e496576-580c-4a9d-8009-1855b28d8f96');
INSERT INTO public.power_accounts VALUES ('3f44db76-f194-46b7-a306-e071488bc0ca', 'store:90775541-4b56-4327-84b7-0927f46122d7', 'Propiedad Centro de Desarrollo', 'Este permiso permite el acceso a la tienda Centro de Desarrollo', '0e496576-580c-4a9d-8009-1855b28d8f96');
INSERT INTO public.power_accounts VALUES ('3980c25a-f147-4f69-aa44-5af480675109', 'company:create', 'Crear empresa', 'Permite la creación de nuevas empresas', '89c53c8d-5c75-4c95-9810-46e6903b1bcc');
INSERT INTO public.power_accounts VALUES ('7b1f3f4e-3e2e-4d3a-9f4e-8c6f4e2b5a1d', 'company:update', 'Editar empresa', 'Permite editar la información de una empresa', '89c53c8d-5c75-4c95-9810-46e6903b1bcc');
INSERT INTO public.power_accounts VALUES ('fd2b8b08-0183-482b-b982-6d61cd9d1bc5', 'company:43f96ad5-6c50-4ce8-b7cf-203210741a18', 'Propiedad Empresa Distribuidora Los Andes', 'Este permiso permite el acceso a la empresa Distribuidora Los Andes', '1fdbaca4-e12b-4378-a670-fb15bb509e93');
INSERT INTO public.power_accounts VALUES ('dc8c704e-9bcb-4dc4-afa0-fc66b41f1f40', 'company:5d3984dd-58ef-42df-84ce-deb13dae434c', 'Propiedad Empresa Soluciones Digitales', 'Este permiso permite el acceso a la empresa Soluciones Digitales', '1fdbaca4-e12b-4378-a670-fb15bb509e93');
INSERT INTO public.power_accounts VALUES ('6568de4f-f4cb-4483-af91-6cde266be10e', 'store:create', 'Crear Tienda', 'Permite crear nuevas tiendas', 'a1584804-7c4e-4ebc-b762-96b5bf1cc901');
INSERT INTO public.power_accounts VALUES ('eb72f588-7ba1-4160-a706-b38a7dc04cef', 'store:update', 'Actualizar Tienda', 'Permite Actualizar tiendas', 'a1584804-7c4e-4ebc-b762-96b5bf1cc901');
INSERT INTO public.power_accounts VALUES ('d7dd6155-f6df-442f-9267-d9fb858dab56', 'supplier:create', 'Crear proveedor', 'Permite la creación de nuevos proveedores', 'b1e5287d-b503-43d3-b555-ebfb8af5dbaa');
INSERT INTO public.power_accounts VALUES ('385ffa1d-371d-4790-b83b-a60ff2718f4f', 'supplier:update', 'Editar proveedor', 'Permite la edición de proveedores', 'b1e5287d-b503-43d3-b555-ebfb8af5dbaa');

-- Warehouse powers
INSERT INTO public.power_accounts VALUES ('a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'warehouse:create', 'Crear bodega', 'Permite la creación de nuevas bodegas', 'f2a8c4e1-9b3d-4f6a-8c2e-1d5f7a9b3e4c');
INSERT INTO public.power_accounts VALUES ('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'warehouse:update', 'Editar bodega', 'Permite la edición de bodegas', 'f2a8c4e1-9b3d-4f6a-8c2e-1d5f7a9b3e4c');

-- Product powers
INSERT INTO public.power_accounts VALUES ('c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f', 'product:create', 'Crear producto', 'Permite la creación de nuevos productos', 'a3b9d5e2-8c4f-4a7b-9d3e-2f6a8c1e4b5d');
INSERT INTO public.power_accounts VALUES ('d4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a', 'product:update', 'Editar producto', 'Permite la edición de productos', 'a3b9d5e2-8c4f-4a7b-9d3e-2f6a8c1e4b5d');

-- Product Company powers
INSERT INTO public.power_accounts VALUES ('e5f6a7b8-c9d0-4e1f-2a3b-4c5d6e7f8a9b', 'product_company:create', 'Crear producto empresa', 'Permite la creación de productos por empresa', 'c4d0e6f3-7b5a-4c8d-ae4f-3a7b9d2f5c6e');
INSERT INTO public.power_accounts VALUES ('f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c', 'product_company:update', 'Editar producto empresa', 'Permite la edición de productos por empresa', 'c4d0e6f3-7b5a-4c8d-ae4f-3a7b9d2f5c6e');

-- Inventory Count powers
INSERT INTO public.power_accounts VALUES ('a7b8c9d0-e1f2-4a3b-4c5d-6e7f8a9b0c1d', 'inventory_count:create', 'Crear conteo inventario', 'Permite la creación de conteos de inventario', 'd5e1f7a4-6c8b-4d9e-bf5a-4b8c0e3a6d7f');
INSERT INTO public.power_accounts VALUES ('b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e', 'inventory_count:update', 'Editar conteo inventario', 'Permite la edición de conteos de inventario', 'd5e1f7a4-6c8b-4d9e-bf5a-4b8c0e3a6d7f');

-- Product Movement powers
INSERT INTO public.power_accounts VALUES ('c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f', 'product_movement:create', 'Crear movimiento producto', 'Permite la creación de movimientos de productos', 'e6f2a8b5-5d9c-4e0f-ca6b-5c9d1f4b7e8a');
INSERT INTO public.power_accounts VALUES ('d0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a', 'product_movement:update', 'Editar movimiento producto', 'Permite la edición de movimientos de productos', 'e6f2a8b5-5d9c-4e0f-ca6b-5c9d1f4b7e8a');
INSERT INTO public.power_accounts VALUES ('f4e5d2c2-1f4e-4d3a-9f4e-8c6f4e2b5a2e', 'template_product:create', 'Crear Plantilla de producto', 'Permite la creación de plantillas de productos', 'c1f4d3e2-5f4e-4d3a-9f4e-8c6f4e2b5a0b');
INSERT INTO public.power_accounts VALUES ('a3b2c1d4-e5f6-4789-abcd-ef0123456789', 'template_product:update', 'Editar Plantilla de producto', 'Permite la edición de plantillas de productos', 'c1f4d3e2-5f4e-4d3a-9f4e-8c6f4e2b5a0b');
INSERT INTO public.power_accounts VALUES ('f8e9d7c6-5b4a-3c2d-1e0f-9a8b7c6d5e4f', 'measurement:create', 'Crear unidad de medida', 'Permite la creación de nuevas unidades de medida', 'c2f4e8a9-4d6b-4c7e-9f8a-1b3d5e7f9a2c');

-- Store Product powers
INSERT INTO public.power_accounts VALUES ('e1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c', 'store_product:create', 'Crear producto tienda', 'Permite la creación de productos en tiendas', 'd3a5b9c7-8e2f-4a1d-bf3c-6e9d4f7a2b8e');
INSERT INTO public.power_accounts VALUES ('f2b3c4d5-e6f7-4a8b-9c0d-1e2f3a4b5c6d', 'store_product:update', 'Editar producto tienda', 'Permite la edición de productos en tiendas', 'd3a5b9c7-8e2f-4a1d-bf3c-6e9d4f7a2b8e');

-- Delivery Purchase Note powers
INSERT INTO public.power_accounts VALUES ('a5b6c7d8-e9f0-4a1b-2c3d-4e5f6a7b8c9d', 'delivery_purchase_note:create', 'Crear nota de entrega de compra', 'Permite la creación de notas de entrega de compra', 'f4b6c9d8-7e3f-4a2b-8c5d-9e0f1a2b3c4d');
INSERT INTO public.power_accounts VALUES ('b6c7d8e9-f0a1-4b2c-3d4e-5f6a7b8c9d0e', 'delivery_purchase_note:update', 'Editar nota de entrega de compra', 'Permite la edición de notas de entrega de compra', 'f4b6c9d8-7e3f-4a2b-8c5d-9e0f1a2b3c4d');


INSERT INTO tag (tag_name, description) VALUES ('Materia Prima', 'Productos que son materias primas o ingredientes base'),('Receta', 'Productos que son resultado de una receta o preparación');

INSERT INTO measurement_unit (abbreviation, unit_name, description, basic_unit) VALUES
-- PESO
('kg', 'kilogramo', 'Equivale a 1000 gramos', TRUE),
('g', 'gramo', 'Unidad de masa del Sistema Internacional', TRUE),

-- VOLUMEN
('l', 'litro', 'Equivale a 1000 mililitros', TRUE),
('ml', 'mililitro', 'Unidad de volumen equivalente a un cm³', TRUE),

-- CANTIDAD
('u', 'unidad simple', 'Una unidad simple', TRUE);

INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor) VALUES
-- PESO
(1, 2, 1000),
(2, 1, 0.001),

-- VOLUMEN
(3, 4, 1000),
(4, 3, 0.001);

-- Unidades base adicionales
INSERT INTO measurement_unit (abbreviation, unit_name, description, basic_unit) VALUES
('lb', 'libra', 'Unidad de masa del sistema imperial. 1 lb = 16 oz, 1 lb = 0.45359237 kg', TRUE),
('oz', 'onza', 'Unidad de masa del sistema imperial. 1 oz = 28.349523125 g', TRUE),
('gal', 'galón', 'Unidad de volumen del sistema imperial. 1 gal = 3.785411784 l', TRUE);

-- Unidades personalizadas (no base)
INSERT INTO measurement_unit (abbreviation, unit_name, description, basic_unit) VALUES
('saco2kg', 'saco de 2 kg', 'Unidad personalizada equivalente a 2 kilogramos', FALSE),
('saco5kg', 'saco de 5 kg', 'Unidad personalizada equivalente a 5 kilogramos', FALSE),
('bot1_5l', 'botella 1.5 l', 'Unidad personalizada equivalente a 1.5 litros', FALSE),
('pack4u', 'pack 4 u', 'Paquete de 4 unidades', FALSE),
('pack6u', 'pack 6 u', 'Paquete de 6 unidades', FALSE),
('pack12u', 'pack 12 u', 'Paquete de 12 unidades', FALSE),
('caja12u', 'caja 12 u', 'Caja de 12 unidades', FALSE),
('caja24u', 'caja 24 u', 'Caja de 24 unidades', FALSE);

-- Conversiones entre unidades base (permitidas en ambos sentidos)
-- kg ↔ lb
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'kg'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'lb'),
	   2.20462262185
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'lb'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'kg'),
	   0.45359237;

-- g ↔ oz
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'g'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'oz'),
	   0.03527396195
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'oz'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'g'),
	   28.349523125;

-- lb ↔ oz
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'lb'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'oz'),
	   16
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'oz'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'lb'),
	   0.0625;

-- l ↔ gal
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'l'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'gal'),
	   0.26417205236
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'gal'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'l'),
	   3.785411784;

-- Conversiones de unidades no base hacia sus respectivas unidades base (solo en un sentido)
-- Saco 2 kg → kg ; Saco 5 kg → kg
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'saco2kg'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'kg'),
	   2
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'saco5kg'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'kg'),
	   5;

-- Botella 1.5 l → l
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'bot1_5l'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'l'),
	   1.5;

-- Packs y Cajas → u
INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'pack4u'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'u'),
	   4
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'pack6u'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'u'),
	   6
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'pack12u'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'u'),
	   12
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'caja12u'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'u'),
	   12
UNION ALL
SELECT (SELECT id FROM measurement_unit WHERE abbreviation = 'caja24u'),
	   (SELECT id FROM measurement_unit WHERE abbreviation = 'u'),
	   24;

INSERT INTO product_category (category_name, description, available)
VALUES 
  ('Carnes y pescados', 'Cortes de carne, pollo, cerdo y productos del mar para cocina', true),
  ('Verduras y frutas frescas', 'Productos frescos utilizados para ensaladas, guarniciones y preparaciones', true),
  ('Lácteos y huevos', 'Leche, quesos, crema, mantequilla, yogurt y huevos para producción', true),
  ('Congelados', 'Productos congelados como papas pre fritas, verduras, mariscos y postres', true),
  ('Secos y despensa', 'Arroz, pastas, legumbres, harinas, aceites, condimentos y conservas', true);

INSERT INTO public.product VALUES ('c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt natural', 'SKU-7802820002019', 'Yogurt natural cremoso y saludable, perfecto para un desayuno equilibrado.', 'https://cc-media.dotsolutions.cl/smart-order/development/6c924e58b77547c894514232fdb966fb_20260102115457.png', 2500.00, '2024-03-12 09:11:30+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche entera', 'SKU-7802820001012', 'Leche entera fresca y nutritiva, ideal para el consumo diario y la preparación de recetas.', 'https://cc-media.dotsolutions.cl/smart-order/development/2a54e9fae22a463a8ba32caa8b75a029_20260102115946.png', 2500.00, '2024-01-15 10:23:45+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso mantecoso', 'SKU-7802820003016', 'Queso mantecoso de textura suave y sabor tradicional, ideal para acompañar tus platos favoritos.', 'https://cc-media.dotsolutions.cl/smart-order/development/2f49882308e44b40ba7118a3648ada4a_20260102115248.png', 4500.00, '2024-05-20 14:45:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla sin sal', 'SKU-7802820004013', 'Mantequilla sin sal elaborada con ingredientes de alta calidad para un sabor puro.', 'https://cc-media.dotsolutions.cl/smart-order/development/55b90657f79441758d66aa36eea22b0d_20260102114756.png', 3800.00, '2024-07-01 16:10:15+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de leche', 'SKU-7802820005010', 'Crema de leche fresca para realzar tus recetas y postres con un toque cremoso.', 'https://cc-media.dotsolutions.cl/smart-order/development/abfcc0eb0980499bbade177fbf0fa277_20260102115518.png', 3500.00, '2024-08-23 11:35:50+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'SKU-7802820006017', 'Bebida gaseosa Coca-Cola en presentación familiar, refrescante y con sabor único.', 'https://cc-media.dotsolutions.cl/smart-order/development/3b2ae93c421d4a2e96fd466c39f86eb9_20260102113441.png', 2500.00, '2024-02-18 12:00:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua mineral 500ml', 'SKU-7802820007014', 'Agua mineral natural embotellada para mantenerte hidratado en cualquier momento.', 'https://cc-media.dotsolutions.cl/smart-order/development/5840b3bdfa2d4156aa491647346056a0_20260102112652.png', 1500.00, '2024-04-25 15:22:10+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de naranja', 'SKU-7802820008011', 'Jugo de naranja 100% natural, fuente de vitamina C y sabor refrescante.', 'https://cc-media.dotsolutions.cl/smart-order/development/7b006ad15c9b4e73ac29a521b6bf0807_20260102113921.png', 3200.00, '2024-06-17 13:18:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té helado', 'SKU-7802820009018', 'Té helado listo para tomar con sabor suave y refrescante, perfecto para el verano.', 'https://cc-media.dotsolutions.cl/smart-order/development/ddf008f6517b477c9403e1f1c2c2cf30_20260102112858.png', 2800.00, '2024-09-03 09:05:30+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida energética', 'SKU-7802820010014', 'Bebida energética para revitalizar tu día con un impulso extra de energía.', 'https://cc-media.dotsolutions.cl/smart-order/development/b3053d3ae87f476bbd2547e471291990_20260102115124.png', 3500.00, '2024-10-10 08:45:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz grano largo', 'SKU-7802820011011', 'Arroz grano largo de alta calidad, ideal para acompañar cualquier comida.', 'https://cc-media.dotsolutions.cl/smart-order/development/861d50ac779248b7b28cee5c36679390_20260102115219.png', 3500.00, '2024-01-20 08:00:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de trigo', 'SKU-7802820012018', 'Harina de trigo fina y versátil, perfecta para repostería y panadería casera.', 'https://cc-media.dotsolutions.cl/smart-order/development/36e9050a69e34e9bb6e177aa53a4c379_20260102120058.png', 2000.00, '2024-03-14 07:50:30+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite vegetal', 'SKU-7802820013015', 'Aceite vegetal puro para cocinar y freír con un sabor neutro y saludable.', 'https://cc-media.dotsolutions.cl/smart-order/development/4b1f2824876f481ba6b4d09837d7ab5b_20260102112932.png', 4000.00, '2024-05-09 16:45:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('912c9f07-5810-438f-ad4c-11790d181910', 'Fideos espagueti', 'SKU-7802820014012', 'Fideos espagueti de textura firme, ideales para tus recetas italianas favoritas.', 'https://cc-media.dotsolutions.cl/smart-order/development/40a5758e5f664588b7117a498773fd89_20260102115044.png', 1500.00, '2024-07-30 10:15:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar refinada', 'SKU-7802820015019', 'Azúcar refinada blanca, perfecta para endulzar y preparar todo tipo de postres.', 'https://cc-media.dotsolutions.cl/smart-order/development/09ca4e4bacd94402953705b8848b50ac_20260102115332.png', 2800.00, '2024-08-18 14:20:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 'Papas fritas clásicas', 'SKU-7802820016016', 'Papas fritas clásicas crujientes y deliciosas, el snack perfecto para cualquier ocasión.', 'https://cc-media.dotsolutions.cl/smart-order/development/ea65afacd9ba44e8b943738c16ca6ded_20260102113413.png', 2000.00, '2024-02-02 11:00:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 'Galletas de chocolate', 'SKU-7802820017013', 'Galletas de chocolate con trozos generosos y un sabor irresistible.', 'https://cc-media.dotsolutions.cl/smart-order/development/e170d80ac3804191a10ade2e7440db02_20260102115400.png', 1800.00, '2024-04-16 09:35:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 'Maní salado', 'SKU-7802820018010', 'Maní salado tostado, ideal para acompañar tus bebidas o como snack saludable.', 'https://cc-media.dotsolutions.cl/smart-order/development/561e6fef9b50496db1885357b5abccaf_20260102112802.png', 1200.00, '2024-06-25 13:40:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('92599d39-eddb-4e9d-9012-b209675ff9bd', 'Chocolate en barra', 'SKU-7802820019017', 'Chocolate en barra con sabor intenso, ideal para disfrutar o para repostería.', 'https://cc-media.dotsolutions.cl/smart-order/development/81dcaa5349f54e68b11546b71aaf2795_20260102115104.png', 2200.00, '2024-09-12 15:55:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 'Barra de cereal', 'SKU-7802820020013', 'Barra de cereal nutritiva y práctica, perfecta para llevar y disfrutar en cualquier momento.', 'https://cc-media.dotsolutions.cl/smart-order/development/b5d3bd85c4f64fd695b7bd1253c767e7_20260102115612.png', 2800.00, '2024-11-05 10:05:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('db682432-362e-4e7b-af54-0d87f91832da', 'Pechuga de pollo', 'SKU-7802820021010', 'Pechuga de pollo fresca y jugosa, lista para preparar platos saludables y deliciosos.', 'https://cc-media.dotsolutions.cl/smart-order/development/cc9b6ac3330544dab4830fe4abeab0f9_20260102115543.png', 8500.00, '2024-01-05 07:30:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 'Carne molida', 'SKU-7802820022017', 'Carne molida de alta calidad, ideal para hamburguesas, albóndigas y más.', 'https://cc-media.dotsolutions.cl/smart-order/development/1494e3e30dce4e18a725f1af0aee97c7_20260102113823.png', 9500.00, '2024-03-21 12:10:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('95523bb0-c72f-4a4c-affe-f14c9a32870e', 'Chuleta de cerdo', 'SKU-7802820023014', 'Chuleta de cerdo tierna y sabrosa, perfecta para asar o guisar.', 'https://cc-media.dotsolutions.cl/smart-order/development/99ce86117ec9485dba0a7575ed37e47d_20260102115151.png', 7500.00, '2024-06-10 09:45:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('e2432d39-68ca-4a02-81d2-e200e47ea620', 'Filete de vacuno', 'SKU-7802820024011', 'Filete de vacuno selecto, con textura suave y sabor excepcional.', 'https://cc-media.dotsolutions.cl/smart-order/development/ada204ef767f494da9d993e9e8ea6c9c_20260102115823.png', 12000.00, '2024-08-14 14:30:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 'Longaniza', 'SKU-7802820025018', 'Longaniza tradicional, con el sabor auténtico para tus comidas.', 'https://cc-media.dotsolutions.cl/smart-order/development/853d62c470344f4bb4caf5ec9b5d1997_20260102113745.png', 3500.00, '2024-10-22 11:20:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate fresco', 'SKU-7802820026015', 'Tomate fresco y jugoso, ideal para ensaladas y salsas.', 'https://cc-media.dotsolutions.cl/smart-order/development/2898d0103952402d9aa784a9e6946818_20260102114714.png', 2000.00, '2024-02-28 10:10:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 'Lechuga', 'SKU-7802820027012', 'Lechuga fresca y crujiente para ensaladas saludables.', 'https://cc-media.dotsolutions.cl/smart-order/development/74175c8e631942e0a000099f274657f8_20260102115307.png', 1500.00, '2024-05-18 13:25:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('f0dddfbb-a31c-4247-ba0b-d37727672f51', 'Zanahoria', 'SKU-7802820028019', 'Zanahoria dulce y crujiente, excelente para cocinar o comer fresca.', 'https://cc-media.dotsolutions.cl/smart-order/development/8362f4ce99c34b07b14f77b6f7911583_20260102120033.png', 1200.00, '2024-07-07 09:00:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('25fbe34d-f25f-43c9-a192-f0b24bac0a66', 'Pimentón rojo', 'SKU-7802820029016', 'Pimentón rojo dulce y colorido, ideal para agregar sabor y color a tus platos.', 'https://cc-media.dotsolutions.cl/smart-order/development/6baf7093b607443eab825e9ee571d869_20260102113510.png', 0, '2024-09-28 12:40:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 'Papa amarilla', 'SKU-7802820030012', 'Papa amarilla de textura suave, perfecta para purés y guisos.', 'https://cc-media.dotsolutions.cl/smart-order/development/8abf06f1f7e84e90bbe05d2d0226789f_20260102113346.png', 1800.00, '2024-11-15 08:55:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana roja', 'SKU-7802820031019', 'Manzana roja crujiente y jugosa, perfecta para snacks y postres.', 'https://cc-media.dotsolutions.cl/smart-order/development/71dc03c9bc884160bfe940d0c9617164_20260102120014.png', 3000.00, '2024-01-12 10:50:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('01337057-887f-4ba3-be7f-7e3296d680e8', 'Banana', 'SKU-7802820032016', 'Banana madura y dulce, fuente natural de energía.', 'https://cc-media.dotsolutions.cl/smart-order/development/5fd10fa68a074092b06aaf6197cf3687_20260102113130.png', 0, '2024-03-30 11:15:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('fec88f6a-af7f-49b2-bf11-f73c6236b580', 'Naranja', 'SKU-7802820033013', 'Naranja fresca con alto contenido de vitamina C y sabor refrescante.', 'https://cc-media.dotsolutions.cl/smart-order/development/6f68d4ca7fd2427b8adc2ed6db73200d_20260102120204.png', 0, '2024-06-22 14:00:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 'Uva morada', 'SKU-7802820034010', 'Uva morada dulce y jugosa, ideal para consumir fresca o en postres.', 'https://cc-media.dotsolutions.cl/smart-order/development/b74ca583fea64fcea12feaf28e58bc13_20260102113849.png', 0, '2024-08-05 09:35:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('69f51fca-89eb-49db-a3c0-84fb149023a9', 'Sandía', 'SKU-7802820035017', 'Sandía fresca y dulce, perfecta para los días calurosos.', 'https://cc-media.dotsolutions.cl/smart-order/development/cd22d2dfb72c4d20aa64503f5b0b47f1_20260102114000.png', 8000.00, '2024-10-01 15:45:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('faffb338-771d-4bd3-871d-4f3e1c37268d', 'Pan marraqueta', 'SKU-7802820036014', 'Pan marraqueta recién horneado, crujiente por fuera y suave por dentro.', 'https://cc-media.dotsolutions.cl/smart-order/development/7067a487ccd0496493050e91c39195d6_20260102120123.png', 2000.00, '2024-02-10 07:50:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('397d445c-3b56-429a-9f52-7835fb340bb6', 'Pan de molde', 'SKU-7802820037011', 'Pan de molde suave y esponjoso, ideal para sándwiches y tostadas.', 'https://cc-media.dotsolutions.cl/smart-order/development/833436438d404631969c2a7c874b94c5_20260102113552.png', 3500.00, '2024-04-27 09:20:00+00', '2025-09-05 21:54:06.935859+00');
INSERT INTO public.product VALUES ('e30b3445-fe86-4914-a2d4-97f10e798922', 'Croissant', 'SKU-7802820038018', 'Croissant hojaldrado y mantecoso, perfecto para el desayuno.', 'https://cc-media.dotsolutions.cl/smart-order/development/6e953f94158a481597abbf6ecb92db4a_20260102115920.png', 5000.00, '2024-06-15 11:30:00+00', '2025-09-05 21:54:06.935859+00');


INSERT INTO public.product_per_category (product_id, category_id) VALUES ('e8aac862-1028-4353-a3fd-928ba671ddb9', 3);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('e8aac862-1028-4353-a3fd-928ba671ddb9', 3, '7802820001012');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('e8aac862-1028-4353-a3fd-928ba671ddb9', 4, '17802820001019');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('c2d0f299-f53d-4e45-99cb-4ff3c946417f', 3);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('c2d0f299-f53d-4e45-99cb-4ff3c946417f', 3, '7802820002019');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('c2d0f299-f53d-4e45-99cb-4ff3c946417f', 2, '078282000201');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('a350d0f1-7720-48fe-8cf8-7cba62879f8e', 3);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('a350d0f1-7720-48fe-8cf8-7cba62879f8e', 3, '7802820003016');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('a350d0f1-7720-48fe-8cf8-7cba62879f8e', 4, '17802820003013');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('800596d0-5e8f-49aa-876a-f45f9418f74d', 3);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('800596d0-5e8f-49aa-876a-f45f9418f74d', 3, '7802820004013');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('c4d42258-52e4-4735-bb02-e163145577ae', 3);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('c4d42258-52e4-4735-bb02-e163145577ae', 3, '7802820005010');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('2099754d-54c7-4b5d-93d8-3990c687981e', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('2099754d-54c7-4b5d-93d8-3990c687981e', 3, '7802820006017');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('2099754d-54c7-4b5d-93d8-3990c687981e', 4, '17802820006014');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 3, '7802820007014');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 1, '78028207');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('67da0a90-ffe5-42bc-88c5-058fca657658', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('67da0a90-ffe5-42bc-88c5-058fca657658', 3, '7802820008011');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('11dc275a-bb9c-4591-98e6-f3e1fc026a61', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('11dc275a-bb9c-4591-98e6-f3e1fc026a61', 3, '7802820009018');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('946d986b-33c5-43af-b6d7-564f461a5b0b', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('946d986b-33c5-43af-b6d7-564f461a5b0b', 3, '7802820010014');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('99de8b79-a9ff-41ec-9791-61117d71900a', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('99de8b79-a9ff-41ec-9791-61117d71900a', 3, '7802820011011');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('f7500114-702b-45bd-bef7-64e6d6d9740c', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('f7500114-702b-45bd-bef7-64e6d6d9740c', 3, '7802820012018');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 3, '7802820013015');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('912c9f07-5810-438f-ad4c-11790d181910', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('912c9f07-5810-438f-ad4c-11790d181910', 3, '7802820014012');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('a55fb2a9-8ba6-476d-b814-8392572ae99d', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('a55fb2a9-8ba6-476d-b814-8392572ae99d', 3, '7802820015019');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 3, '7802820016016');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 3, '7802820017013');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 3, '7802820018010');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('92599d39-eddb-4e9d-9012-b209675ff9bd', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('92599d39-eddb-4e9d-9012-b209675ff9bd', 3, '7802820019017');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 3, '7802820020013');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('db682432-362e-4e7b-af54-0d87f91832da', 1);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('db682432-362e-4e7b-af54-0d87f91832da', 3, '7802820021010');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('db682432-362e-4e7b-af54-0d87f91832da', 5, '4011');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 1);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 3, '7802820022017');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('95523bb0-c72f-4a4c-affe-f14c9a32870e', 1);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('95523bb0-c72f-4a4c-affe-f14c9a32870e', 3, '7802820023014');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('e2432d39-68ca-4a02-81d2-e200e47ea620', 1);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('e2432d39-68ca-4a02-81d2-e200e47ea620', 3, '7802820024011');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 1);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 3, '7802820025018');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('7ddda93a-fe87-4385-9299-232ec0f5f7ee', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('7ddda93a-fe87-4385-9299-232ec0f5f7ee', 3, '7802820026015');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('7ddda93a-fe87-4385-9299-232ec0f5f7ee', 5, '4664');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 3, '7802820027012');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 5, '4061');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('f0dddfbb-a31c-4247-ba0b-d37727672f51', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('f0dddfbb-a31c-4247-ba0b-d37727672f51', 3, '7802820028019');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('f0dddfbb-a31c-4247-ba0b-d37727672f51', 5, '4562');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('25fbe34d-f25f-43c9-a192-f0b24bac0a66', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('25fbe34d-f25f-43c9-a192-f0b24bac0a66', 3, '7802820029016');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 3, '7802820030012');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 5, '4725');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 3, '7802820031019');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 5, '4016');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('01337057-887f-4ba3-be7f-7e3296d680e8', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('01337057-887f-4ba3-be7f-7e3296d680e8', 3, '7802820032016');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('01337057-887f-4ba3-be7f-7e3296d680e8', 5, '94011');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('fec88f6a-af7f-49b2-bf11-f73c6236b580', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('fec88f6a-af7f-49b2-bf11-f73c6236b580', 3, '7802820033013');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('fec88f6a-af7f-49b2-bf11-f73c6236b580', 5, '3107');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 3, '7802820034010');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('69f51fca-89eb-49db-a3c0-84fb149023a9', 2);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('69f51fca-89eb-49db-a3c0-84fb149023a9', 3, '7802820035017');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('69f51fca-89eb-49db-a3c0-84fb149023a9', 5, '4032');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('faffb338-771d-4bd3-871d-4f3e1c37268d', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('faffb338-771d-4bd3-871d-4f3e1c37268d', 3, '7802820036014');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('397d445c-3b56-429a-9f52-7835fb340bb6', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('397d445c-3b56-429a-9f52-7835fb340bb6', 3, '7802820037011');
INSERT INTO public.product_per_category (product_id, category_id) VALUES ('e30b3445-fe86-4914-a2d4-97f10e798922', 5);
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('e30b3445-fe86-4914-a2d4-97f10e798922', 3, '7802820038018');
INSERT INTO public.product_code (product_id, kind_id, code_value) VALUES ('e30b3445-fe86-4914-a2d4-97f10e798922', 1, '78028388');
-- INSERT INTO public.product VALUES ('e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche entera', 'Leche entera fresca y nutritiva, ideal para el consumo diario y la preparación de recetas.', NULL, '2024-01-15 10:23:45+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt natural', 'Yogurt natural cremoso y saludable, perfecto para un desayuno equilibrado.', NULL, '2024-03-12 09:11:30+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso mantecoso', 'Queso mantecoso de textura suave y sabor tradicional, ideal para acompañar tus platos favoritos.', NULL, '2024-05-20 14:45:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla sin sal', 'Mantequilla sin sal elaborada con ingredientes de alta calidad para un sabor puro.', NULL, '2024-07-01 16:10:15+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de leche', 'Crema de leche fresca para realzar tus recetas y postres con un toque cremoso.', NULL, '2024-08-23 11:35:50+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'Bebida gaseosa Coca-Cola en presentación familiar, refrescante y con sabor único.', NULL, '2024-02-18 12:00:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua mineral 500ml', 'Agua mineral natural embotellada para mantenerte hidratado en cualquier momento.', NULL, '2024-04-25 15:22:10+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de naranja', 'Jugo de naranja 100% natural, fuente de vitamina C y sabor refrescante.', NULL, '2024-06-17 13:18:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té helado', 'Té helado listo para tomar con sabor suave y refrescante, perfecto para el verano.', NULL, '2024-09-03 09:05:30+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida energética', 'Bebida energética para revitalizar tu día con un impulso extra de energía.', NULL, '2024-10-10 08:45:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz grano largo', 'Arroz grano largo de alta calidad, ideal para acompañar cualquier comida.', NULL, '2024-01-20 08:00:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de trigo', 'Harina de trigo fina y versátil, perfecta para repostería y panadería casera.', NULL, '2024-03-14 07:50:30+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite vegetal', 'Aceite vegetal puro para cocinar y freír con un sabor neutro y saludable.', NULL, '2024-05-09 16:45:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('912c9f07-5810-438f-ad4c-11790d181910', 'Fideos espagueti', 'Fideos espagueti de textura firme, ideales para tus recetas italianas favoritas.', NULL, '2024-07-30 10:15:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar refinada', 'Azúcar refinada blanca, perfecta para endulzar y preparar todo tipo de postres.', NULL, '2024-08-18 14:20:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 'Papas fritas clásicas', 'Papas fritas clásicas crujientes y deliciosas, el snack perfecto para cualquier ocasión.', NULL, '2024-02-02 11:00:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 'Galletas de chocolate', 'Galletas de chocolate con trozos generosos y un sabor irresistible.', NULL, '2024-04-16 09:35:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 'Maní salado', 'Maní salado tostado, ideal para acompañar tus bebidas o como snack saludable.', NULL, '2024-06-25 13:40:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('92599d39-eddb-4e9d-9012-b209675ff9bd', 'Chocolate en barra', 'Chocolate en barra con sabor intenso, ideal para disfrutar o para repostería.', NULL, '2024-09-12 15:55:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 'Barra de cereal', 'Barra de cereal nutritiva y práctica, perfecta para llevar y disfrutar en cualquier momento.', NULL, '2024-11-05 10:05:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('db682432-362e-4e7b-af54-0d87f91832da', 'Pechuga de pollo', 'Pechuga de pollo fresca y jugosa, lista para preparar platos saludables y deliciosos.', NULL, '2024-01-05 07:30:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 'Carne molida', 'Carne molida de alta calidad, ideal para hamburguesas, albóndigas y más.', NULL, '2024-03-21 12:10:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('95523bb0-c72f-4a4c-affe-f14c9a32870e', 'Chuleta de cerdo', 'Chuleta de cerdo tierna y sabrosa, perfecta para asar o guisar.', NULL, '2024-06-10 09:45:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('e2432d39-68ca-4a02-81d2-e200e47ea620', 'Filete de vacuno', 'Filete de vacuno selecto, con textura suave y sabor excepcional.', NULL, '2024-08-14 14:30:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 'Longaniza', 'Longaniza tradicional, con el sabor auténtico para tus comidas.', NULL, '2024-10-22 11:20:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate fresco', 'Tomate fresco y jugoso, ideal para ensaladas y salsas.', NULL, '2024-02-28 10:10:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 'Lechuga', 'Lechuga fresca y crujiente para ensaladas saludables.', NULL, '2024-05-18 13:25:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('f0dddfbb-a31c-4247-ba0b-d37727672f51', 'Zanahoria', 'Zanahoria dulce y crujiente, excelente para cocinar o comer fresca.', NULL, '2024-07-07 09:00:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('25fbe34d-f25f-43c9-a192-f0b24bac0a66', 'Pimentón rojo', 'Pimentón rojo dulce y colorido, ideal para agregar sabor y color a tus platos.', NULL, '2024-09-28 12:40:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 'Papa amarilla', 'Papa amarilla de textura suave, perfecta para purés y guisos.', NULL, '2024-11-15 08:55:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana roja', 'Manzana roja crujiente y jugosa, perfecta para snacks y postres.', NULL, '2024-01-12 10:50:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('01337057-887f-4ba3-be7f-7e3296d680e8', 'Banana', 'Banana madura y dulce, fuente natural de energía.', NULL, '2024-03-30 11:15:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('fec88f6a-af7f-49b2-bf11-f73c6236b580', 'Naranja', 'Naranja fresca con alto contenido de vitamina C y sabor refrescante.', NULL, '2024-06-22 14:00:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 'Uva morada', 'Uva morada dulce y jugosa, ideal para consumir fresca o en postres.', NULL, '2024-08-05 09:35:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('69f51fca-89eb-49db-a3c0-84fb149023a9', 'Sandía', 'Sandía fresca y dulce, perfecta para los días calurosos.', NULL, '2024-10-01 15:45:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('faffb338-771d-4bd3-871d-4f3e1c37268d', 'Pan marraqueta', 'Pan marraqueta recién horneado, crujiente por fuera y suave por dentro.', NULL, '2024-02-10 07:50:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('397d445c-3b56-429a-9f52-7835fb340bb6', 'Pan de molde', 'Pan de molde suave y esponjoso, ideal para sándwiches y tostadas.', NULL, '2024-04-27 09:20:00+00', '2025-09-05 21:54:06.935859+00');
-- INSERT INTO public.product VALUES ('e30b3445-fe86-4914-a2d4-97f10e798922', 'Croissant', 'Croissant hojaldrado y mantecoso, perfecto para el desayuno.', NULL, '2024-06-15 11:30:00+00', '2025-09-05 21:54:06.935859+00');

/* DEPRECATED: product_company table has been deprecated. Data now lives in product_per_store.
INSERT INTO public.product_company VALUES ('26d261de-5b9d-44fa-bdf6-c5fbf0f15ad3', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 1, 'SKU-LECHE-001', 'Leche entera', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 900, 'Leche entera fresca y nutritiva, ideal para el consumo diario y la preparación de recetas.', 900, 900,10,100,1);
INSERT INTO public.product_company VALUES ('7e3fafcd-0197-4f62-84b4-ed5c753921ac', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 1, 'SKU-LECHE-002', 'Leche entera', true, true, true, false, false, 5, 5, '{}'::jsonb, 950, 'Leche entera fresca y nutritiva, ideal para el consumo diario y la preparación de recetas.', 950, 950,10,100,1);
INSERT INTO public.product_company VALUES ('45a120db-2844-4fe0-8541-49906750509c', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 1, 'SKU-DISTRIBUIDORALOSANDES-YOGURTNATURAL', 'Yogurt natural', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2455, 'Yogurt natural cremoso y saludable, perfecto para un desayuno equilibrado.', 1242, 2198,10,100,1);
INSERT INTO public.product_company VALUES ('81865ab1-c78b-4db4-9f92-2755eb102b9b', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 1, 'SKU-SOLUCIONESDIGITALES-YOGURTNATURAL', 'Yogurt natural', true, true, true, false, false, 5, 5, '{}'::jsonb, 3641, 'Yogurt natural cremoso y saludable, perfecto para un desayuno equilibrado.', 5141, 1371,10,100,1);
INSERT INTO public.product_company VALUES ('35bf2c4c-af8a-469a-bf71-1b1b340b2322', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 1, 'SKU-DISTRIBUIDORALOSANDES-QUESOMANTECOSO', 'Queso mantecoso', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 3415, 'Queso mantecoso de textura suave y sabor tradicional, ideal para acompañar tus platos favoritos.', 4896, 5064,10,100,1);
INSERT INTO public.product_company VALUES ('723d45e6-58c2-434e-8a96-c2540fec5c4d', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 1, 'SKU-SOLUCIONESDIGITALES-QUESOMANTECOSO', 'Queso mantecoso', true, true, true, false, false, 5, 5, '{}'::jsonb, 5849, 'Queso mantecoso de textura suave y sabor tradicional, ideal para acompañar tus platos favoritos.', 5107, 1755,10,100,1);
INSERT INTO public.product_company VALUES ('52bbed11-7f3d-42c1-82ca-604b3d74b4b4', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '800596d0-5e8f-49aa-876a-f45f9418f74d', 1, 'SKU-DISTRIBUIDORALOSANDES-MANTEQUILLASINSAL', 'Mantequilla sin sal', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 3359, 'Mantequilla sin sal elaborada con ingredientes de alta calidad para un sabor puro.', 3229, 3583,10,100,1);
INSERT INTO public.product_company VALUES ('602c2189-c685-47f1-a0ef-cb02c0f7958f', '5d3984dd-58ef-42df-84ce-deb13dae434c', '800596d0-5e8f-49aa-876a-f45f9418f74d', 1, 'SKU-SOLUCIONESDIGITALES-MANTEQUILLASINSAL', 'Mantequilla sin sal', true, true, true, false, false, 5, 5, '{}'::jsonb, 3365, 'Mantequilla sin sal elaborada con ingredientes de alta calidad para un sabor puro.', 3901, 5216,10,100,1);
INSERT INTO public.product_company VALUES ('17f8cc95-ea0b-4313-a370-5a55b8e7b76b', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'c4d42258-52e4-4735-bb02-e163145577ae', 1, 'SKU-DISTRIBUIDORALOSANDES-CREMADELECHE', 'Crema de leche', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5105, 'Crema de leche fresca para realzar tus recetas y postres con un toque cremoso.', 3426, 5807,10,100,1);
INSERT INTO public.product_company VALUES ('f2266d2e-a22a-40b6-bdad-b178d4203ee6', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'c4d42258-52e4-4735-bb02-e163145577ae', 1, 'SKU-SOLUCIONESDIGITALES-CREMADELECHE', 'Crema de leche', true, true, true, false, false, 5, 5, '{}'::jsonb, 2425, 'Crema de leche fresca para realzar tus recetas y postres con un toque cremoso.', 3451, 3510,10,100,1);
INSERT INTO public.product_company VALUES ('b37d59ec-a709-4435-b854-0e7e496320ee', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '2099754d-54c7-4b5d-93d8-3990c687981e', 1, 'SKU-DISTRIBUIDORALOSANDES-COCA-COLA1.5L', 'Coca-Cola 1.5L', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2912, 'Bebida gaseosa Coca-Cola en presentación familiar, refrescante y con sabor único.', 2847, 3034,10,100,1);
INSERT INTO public.product_company VALUES ('63750a09-7986-44cf-b467-bf5a901e5155', '5d3984dd-58ef-42df-84ce-deb13dae434c', '2099754d-54c7-4b5d-93d8-3990c687981e', 1, 'SKU-SOLUCIONESDIGITALES-COCA-COLA1.5L', 'Coca-Cola 1.5L', true, true, true, false, false, 5, 5, '{}'::jsonb, 3236, 'Bebida gaseosa Coca-Cola en presentación familiar, refrescante y con sabor único.', 5752, 2706,10,100,1);
INSERT INTO public.product_company VALUES ('d96e872a-8cad-420a-a229-480ed233ad46', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 1, 'SKU-DISTRIBUIDORALOSANDES-AGUAMINERAL500ML', 'Agua mineral 500ml', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4441, 'Agua mineral natural embotellada para mantenerte hidratado en cualquier momento.', 5314, 5159,10,100,1);
INSERT INTO public.product_company VALUES ('366c98f4-c33e-41a8-a435-cee3c150ffda', '5d3984dd-58ef-42df-84ce-deb13dae434c', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 1, 'SKU-SOLUCIONESDIGITALES-AGUAMINERAL500ML', 'Agua mineral 500ml', true, true, true, false, false, 5, 5, '{}'::jsonb, 2538, 'Agua mineral natural embotellada para mantenerte hidratado en cualquier momento.', 1390, 4102,10,100,1);
INSERT INTO public.product_company VALUES ('6c88f769-511f-4a89-bc8c-79affdcdbaa5', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '67da0a90-ffe5-42bc-88c5-058fca657658', 1, 'SKU-DISTRIBUIDORALOSANDES-JUGODENARANJA', 'Jugo de naranja', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2124, 'Jugo de naranja 100% natural, fuente de vitamina C y sabor refrescante.', 2511, 1911,10,100,1);
INSERT INTO public.product_company VALUES ('3ee8ed9b-e3a8-4a67-b5eb-c5f85e74cd51', '5d3984dd-58ef-42df-84ce-deb13dae434c', '67da0a90-ffe5-42bc-88c5-058fca657658', 1, 'SKU-SOLUCIONESDIGITALES-JUGODENARANJA', 'Jugo de naranja', true, true, true, false, false, 5, 5, '{}'::jsonb, 2077, 'Jugo de naranja 100% natural, fuente de vitamina C y sabor refrescante.', 1515, 2149,10,100,1);
INSERT INTO public.product_company VALUES ('8780d84e-693d-4645-8286-17d782ba6f24', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 1, 'SKU-DISTRIBUIDORALOSANDES-TÉHELADO', 'Té helado', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2196, 'Té helado listo para tomar con sabor suave y refrescante, perfecto para el verano.', 1575, 4124,10,100,1);
INSERT INTO public.product_company VALUES ('78168930-9e29-49bd-a1a9-336484b499f0', '5d3984dd-58ef-42df-84ce-deb13dae434c', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 1, 'SKU-SOLUCIONESDIGITALES-TÉHELADO', 'Té helado', true, true, true, false, false, 5, 5, '{}'::jsonb, 3412, 'Té helado listo para tomar con sabor suave y refrescante, perfecto para el verano.', 5928, 5480,10,100,1);
INSERT INTO public.product_company VALUES ('d0220269-aad5-4764-a253-c372d2dced23', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '946d986b-33c5-43af-b6d7-564f461a5b0b', 1, 'SKU-DISTRIBUIDORALOSANDES-BEBIDAENERGÉTICA', 'Bebida energética', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5939, 'Bebida energética para revitalizar tu día con un impulso extra de energía.', 2287, 4624,10,100,1);
INSERT INTO public.product_company VALUES ('6fe150f1-dad8-4c62-9add-0e5ec09d2710', '5d3984dd-58ef-42df-84ce-deb13dae434c', '946d986b-33c5-43af-b6d7-564f461a5b0b', 1, 'SKU-SOLUCIONESDIGITALES-BEBIDAENERGÉTICA', 'Bebida energética', true, true, true, false, false, 5, 5, '{}'::jsonb, 1100, 'Bebida energética para revitalizar tu día con un impulso extra de energía.', 2940, 5328,10,100,1);
INSERT INTO public.product_company VALUES ('2498ad2f-d703-4aef-88ef-e13b2ccae704', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '99de8b79-a9ff-41ec-9791-61117d71900a', 1, 'SKU-DISTRIBUIDORALOSANDES-ARROZGRANOLARGO', 'Arroz grano largo', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4477, 'Arroz grano largo de alta calidad, ideal para acompañar cualquier comida.', 3996, 5922,10,100,1);
INSERT INTO public.product_company VALUES ('6e1435b9-18cf-40b4-a74c-42ae64317058', '5d3984dd-58ef-42df-84ce-deb13dae434c', '99de8b79-a9ff-41ec-9791-61117d71900a', 1, 'SKU-SOLUCIONESDIGITALES-ARROZGRANOLARGO', 'Arroz grano largo', true, true, true, false, false, 5, 5, '{}'::jsonb, 1473, 'Arroz grano largo de alta calidad, ideal para acompañar cualquier comida.', 1675, 2714,10,100,1);
INSERT INTO public.product_company VALUES ('aadba3e3-b6c2-4df6-903d-9c26e6fcfe84', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 1, 'SKU-DISTRIBUIDORALOSANDES-HARINADETRIGO', 'Harina de trigo', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 1050, 'Harina de trigo fina y versátil, perfecta para repostería y panadería casera.', 3782, 1966,10,100,1);
INSERT INTO public.product_company VALUES ('09527901-79e8-40fe-99f3-ab766d33cb48', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 1, 'SKU-SOLUCIONESDIGITALES-HARINADETRIGO', 'Harina de trigo', true, true, true, false, false, 5, 5, '{}'::jsonb, 5284, 'Harina de trigo fina y versátil, perfecta para repostería y panadería casera.', 3210, 5569,10,100,1);
INSERT INTO public.product_company VALUES ('6b71a9ad-1b51-4c3a-9093-aff78ab0f052', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 1, 'SKU-DISTRIBUIDORALOSANDES-ACEITEVEGETAL', 'Aceite vegetal', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5251, 'Aceite vegetal puro para cocinar y freír con un sabor neutro y saludable.', 2812, 4694,10,100,1);
INSERT INTO public.product_company VALUES ('8468d6b3-f6f4-4e27-9767-d727b22046da', '5d3984dd-58ef-42df-84ce-deb13dae434c', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 1, 'SKU-SOLUCIONESDIGITALES-ACEITEVEGETAL', 'Aceite vegetal', true, true, true, false, false, 5, 5, '{}'::jsonb, 3899, 'Aceite vegetal puro para cocinar y freír con un sabor neutro y saludable.', 4974, 4666,10,100,1);
INSERT INTO public.product_company VALUES ('35555aa9-0bcd-4c11-901b-6712a0f29253', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '912c9f07-5810-438f-ad4c-11790d181910', 1, 'SKU-DISTRIBUIDORALOSANDES-FIDEOSESPAGUETI', 'Fideos espagueti', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4759, 'Fideos espagueti de textura firme, ideales para tus recetas italianas favoritas.', 3112, 3516,10,100,1);
INSERT INTO public.product_company VALUES ('6108caaa-7d24-4320-8913-10cacca37397', '5d3984dd-58ef-42df-84ce-deb13dae434c', '912c9f07-5810-438f-ad4c-11790d181910', 1, 'SKU-SOLUCIONESDIGITALES-FIDEOSESPAGUETI', 'Fideos espagueti', true, true, true, false, false, 5, 5, '{}'::jsonb, 1665, 'Fideos espagueti de textura firme, ideales para tus recetas italianas favoritas.', 5899, 5413,10,100,1);
INSERT INTO public.product_company VALUES ('ae5f5d66-1c9a-4f5b-b810-833fd1a99bd8', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 1, 'SKU-DISTRIBUIDORALOSANDES-AZÚCARREFINADA', 'Azúcar refinada', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 3463, 'Azúcar refinada blanca, perfecta para endulzar y preparar todo tipo de postres.', 4707, 1970,10,100,1);
INSERT INTO public.product_company VALUES ('15832b25-4891-4c74-abab-2fafccb3b919', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 1, 'SKU-SOLUCIONESDIGITALES-AZÚCARREFINADA', 'Azúcar refinada', true, true, true, false, false, 5, 5, '{}'::jsonb, 4742, 'Azúcar refinada blanca, perfecta para endulzar y preparar todo tipo de postres.', 3529, 3383,10,100,1);
INSERT INTO public.product_company VALUES ('090ca8d1-3e7a-456a-8628-614c3210c20a', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 1, 'SKU-DISTRIBUIDORALOSANDES-PAPASFRITASCLÁSICAS', 'Papas fritas clásicas', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2097, 'Papas fritas clásicas crujientes y deliciosas, el snack perfecto para cualquier ocasión.', 2687, 3831,10,100,1);
INSERT INTO public.product_company VALUES ('cbb0584f-3622-4229-9c65-5fc02073f71d', '5d3984dd-58ef-42df-84ce-deb13dae434c', '1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 1, 'SKU-SOLUCIONESDIGITALES-PAPASFRITASCLÁSICAS', 'Papas fritas clásicas', true, true, true, false, false, 5, 5, '{}'::jsonb, 5209, 'Papas fritas clásicas crujientes y deliciosas, el snack perfecto para cualquier ocasión.', 2661, 4079,10,100,1);
INSERT INTO public.product_company VALUES ('cc2837ed-994a-4d44-b1b3-1c4e8419ee24', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 1, 'SKU-DISTRIBUIDORALOSANDES-GALLETASDECHOCOLATE', 'Galletas de chocolate', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5744, 'Galletas de chocolate con trozos generosos y un sabor irresistible.', 1854, 3770,10,100,1);
INSERT INTO public.product_company VALUES ('e4e3b668-5f4c-4261-a0ed-3e88508b199e', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 1, 'SKU-SOLUCIONESDIGITALES-GALLETASDECHOCOLATE', 'Galletas de chocolate', true, true, true, false, false, 5, 5, '{}'::jsonb, 3626, 'Galletas de chocolate con trozos generosos y un sabor irresistible.', 5502, 3626,10,100,1);
INSERT INTO public.product_company VALUES ('e01e8805-18fd-4b28-bc27-04d4f10e00db', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 1, 'SKU-DISTRIBUIDORALOSANDES-MANÍSALADO', 'Maní salado', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4895, 'Maní salado tostado, ideal para acompañar tus bebidas o como snack saludable.', 4258, 5120,10,100,1);
INSERT INTO public.product_company VALUES ('4af1f64a-8aed-4269-bd14-c00a6b6a8a11', '5d3984dd-58ef-42df-84ce-deb13dae434c', '0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 1, 'SKU-SOLUCIONESDIGITALES-MANÍSALADO', 'Maní salado', true, true, true, false, false, 5, 5, '{}'::jsonb, 2006, 'Maní salado tostado, ideal para acompañar tus bebidas o como snack saludable.', 3760, 1438,10,100,1);
INSERT INTO public.product_company VALUES ('2f93d86e-3f81-46e7-9607-b3bf047c9d1e', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '92599d39-eddb-4e9d-9012-b209675ff9bd', 1, 'SKU-DISTRIBUIDORALOSANDES-CHOCOLATEENBARRA', 'Chocolate en barra', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4996, 'Chocolate en barra con sabor intenso, ideal para disfrutar o para repostería.', 3040, 4742,10,100,1);
INSERT INTO public.product_company VALUES ('24e3ac1e-7849-4c26-b9b5-21568d41ea95', '5d3984dd-58ef-42df-84ce-deb13dae434c', '92599d39-eddb-4e9d-9012-b209675ff9bd', 1, 'SKU-SOLUCIONESDIGITALES-CHOCOLATEENBARRA', 'Chocolate en barra', true, true, true, false, false, 5, 5, '{}'::jsonb, 3823, 'Chocolate en barra con sabor intenso, ideal para disfrutar o para repostería.', 1289, 5429,10,100,1);
INSERT INTO public.product_company VALUES ('d8f1b637-785a-48ec-8b97-e9bc0591c832', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 1, 'SKU-DISTRIBUIDORALOSANDES-BARRADECEREAL', 'Barra de cereal', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5819, 'Barra de cereal nutritiva y práctica, perfecta para llevar y disfrutar en cualquier momento.', 3349, 5768,10,100,1);
INSERT INTO public.product_company VALUES ('f3c07f48-d141-404e-9c74-65ff91d3a715', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 1, 'SKU-SOLUCIONESDIGITALES-BARRADECEREAL', 'Barra de cereal', true, true, true, false, false, 5, 5, '{}'::jsonb, 3360, 'Barra de cereal nutritiva y práctica, perfecta para llevar y disfrutar en cualquier momento.', 1280, 2830,10,100,1);
INSERT INTO public.product_company VALUES ('752ca137-fe2c-46d1-bb44-62b935625dea', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'db682432-362e-4e7b-af54-0d87f91832da', 1, 'SKU-DISTRIBUIDORALOSANDES-PECHUGADEPOLLO', 'Pechuga de pollo', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2395, 'Pechuga de pollo fresca y jugosa, lista para preparar platos saludables y deliciosos.', 2962, 4766,10,100,1);
INSERT INTO public.product_company VALUES ('ee34f25d-564d-404c-84d3-6a6cabe2e01a', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'db682432-362e-4e7b-af54-0d87f91832da', 1, 'SKU-SOLUCIONESDIGITALES-PECHUGADEPOLLO', 'Pechuga de pollo', true, true, true, false, false, 5, 5, '{}'::jsonb, 5785, 'Pechuga de pollo fresca y jugosa, lista para preparar platos saludables y deliciosos.', 3443, 1299,10,100,1);
INSERT INTO public.product_company VALUES ('38d0ec4e-3251-4967-9377-78b53279013d', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 1, 'SKU-DISTRIBUIDORALOSANDES-CARNEMOLIDA', 'Carne molida', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 1670, 'Carne molida de alta calidad, ideal para hamburguesas, albóndigas y más.', 4832, 4010,10,100,1);
INSERT INTO public.product_company VALUES ('54ef3630-482c-4e5d-9433-effa3dccafde', '5d3984dd-58ef-42df-84ce-deb13dae434c', '60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 1, 'SKU-SOLUCIONESDIGITALES-CARNEMOLIDA', 'Carne molida', true, true, true, false, false, 5, 5, '{}'::jsonb, 4410, 'Carne molida de alta calidad, ideal para hamburguesas, albóndigas y más.', 1118, 1004,10,100,1);
INSERT INTO public.product_company VALUES ('3ab99192-d518-4b2b-b010-ff5d35a59e28', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '95523bb0-c72f-4a4c-affe-f14c9a32870e', 1, 'SKU-DISTRIBUIDORALOSANDES-CHULETADECERDO', 'Chuleta de cerdo', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4473, 'Chuleta de cerdo tierna y sabrosa, perfecta para asar o guisar.', 1076, 1929,10,100,1);
INSERT INTO public.product_company VALUES ('8d70a09d-3af8-484b-acd5-9ff86bc86971', '5d3984dd-58ef-42df-84ce-deb13dae434c', '95523bb0-c72f-4a4c-affe-f14c9a32870e', 1, 'SKU-SOLUCIONESDIGITALES-CHULETADECERDO', 'Chuleta de cerdo', true, true, true, false, false, 5, 5, '{}'::jsonb, 2868, 'Chuleta de cerdo tierna y sabrosa, perfecta para asar o guisar.', 5779, 1902,10,100,1);
INSERT INTO public.product_company VALUES ('ab29c1cf-7029-4f32-8b76-45f88bd69705', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'e2432d39-68ca-4a02-81d2-e200e47ea620', 1, 'SKU-DISTRIBUIDORALOSANDES-FILETEDEVACUNO', 'Filete de vacuno', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4393, 'Filete de vacuno selecto, con textura suave y sabor excepcional.', 5873, 1639,10,100,1);
INSERT INTO public.product_company VALUES ('a7d21404-d4f5-42b0-a197-77281d37881e', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'e2432d39-68ca-4a02-81d2-e200e47ea620', 1, 'SKU-SOLUCIONESDIGITALES-FILETEDEVACUNO', 'Filete de vacuno', true, true, true, false, false, 5, 5, '{}'::jsonb, 5465, 'Filete de vacuno selecto, con textura suave y sabor excepcional.', 5104, 2712,10,100,1);
INSERT INTO public.product_company VALUES ('27b62db1-6324-4ba5-8d5d-dd395a9019f9', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 1, 'SKU-DISTRIBUIDORALOSANDES-LONGANIZA', 'Longaniza', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4250, 'Longaniza tradicional, con el sabor auténtico para tus comidas.', 3783, 1368,10,100,1);
INSERT INTO public.product_company VALUES ('d1bccbdc-e72d-43c2-9817-5c2698d1a02a', '5d3984dd-58ef-42df-84ce-deb13dae434c', '3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 1, 'SKU-SOLUCIONESDIGITALES-LONGANIZA', 'Longaniza', true, true, true, false, false, 5, 5, '{}'::jsonb, 4259, 'Longaniza tradicional, con el sabor auténtico para tus comidas.', 2283, 1003,10,100,1);
INSERT INTO public.product_company VALUES ('0c0e11a0-6a54-433f-92e3-620486e9de0d', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 1, 'SKU-DISTRIBUIDORALOSANDES-TOMATEFRESCO', 'Tomate fresco', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2911, 'Tomate fresco y jugoso, ideal para ensaladas y salsas.', 4410, 5700,10,100,1);
INSERT INTO public.product_company VALUES ('dc5c1c96-e7eb-4e88-97d2-7b34abccd8c2', '5d3984dd-58ef-42df-84ce-deb13dae434c', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 1, 'SKU-SOLUCIONESDIGITALES-TOMATEFRESCO', 'Tomate fresco', true, true, true, false, false, 5, 5, '{}'::jsonb, 5100, 'Tomate fresco y jugoso, ideal para ensaladas y salsas.', 3462, 5535,10,100,1);
INSERT INTO public.product_company VALUES ('ed07d17b-fa3b-4891-bb40-49c3346ec58d', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 1, 'SKU-DISTRIBUIDORALOSANDES-LECHUGA', 'Lechuga', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4617, 'Lechuga fresca y crujiente para ensaladas saludables.', 1707, 4551,10,100,1);
INSERT INTO public.product_company VALUES ('288ac5c3-4c0c-4b2c-bb9e-19685e72c456', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 1, 'SKU-SOLUCIONESDIGITALES-LECHUGA', 'Lechuga', true, true, true, false, false, 5, 5, '{}'::jsonb, 3574, 'Lechuga fresca y crujiente para ensaladas saludables.', 2626, 3910,10,100,1);
INSERT INTO public.product_company VALUES ('579ecda9-71c2-4613-b40c-18b442efe431', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'f0dddfbb-a31c-4247-ba0b-d37727672f51', 1, 'SKU-DISTRIBUIDORALOSANDES-ZANAHORIA', 'Zanahoria', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 1062, 'Zanahoria dulce y crujiente, excelente para cocinar o comer fresca.', 2807, 3992,10,100,1);
INSERT INTO public.product_company VALUES ('a49cf3b8-03f9-4795-a773-f123aca1dfb9', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'f0dddfbb-a31c-4247-ba0b-d37727672f51', 1, 'SKU-SOLUCIONESDIGITALES-ZANAHORIA', 'Zanahoria', true, true, true, false, false, 5, 5, '{}'::jsonb, 2605, 'Zanahoria dulce y crujiente, excelente para cocinar o comer fresca.', 2330, 1987,10,100,1);
INSERT INTO public.product_company VALUES ('90fcf214-abca-44f8-bd17-21e20a939891', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '25fbe34d-f25f-43c9-a192-f0b24bac0a66', 1, 'SKU-DISTRIBUIDORALOSANDES-PIMENTÓNROJO', 'Pimentón rojo', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 1954, 'Pimentón rojo dulce y colorido, ideal para agregar sabor y color a tus platos.', 4665, 2042,10,100,1);
INSERT INTO public.product_company VALUES ('eeeba08f-acb5-476b-ab98-03558c89d095', '5d3984dd-58ef-42df-84ce-deb13dae434c', '25fbe34d-f25f-43c9-a192-f0b24bac0a66', 1, 'SKU-SOLUCIONESDIGITALES-PIMENTÓNROJO', 'Pimentón rojo', true, true, true, false, false, 5, 5, '{}'::jsonb, 5773, 'Pimentón rojo dulce y colorido, ideal para agregar sabor y color a tus platos.', 4509, 2750,10,100,1);
INSERT INTO public.product_company VALUES ('7573ad47-efda-48e3-af8b-1daf4870cb55', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 1, 'SKU-DISTRIBUIDORALOSANDES-PAPAAMARILLA', 'Papa amarilla', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 2913, 'Papa amarilla de textura suave, perfecta para purés y guisos.', 2815, 1989,10,100,1);
INSERT INTO public.product_company VALUES ('356a90aa-abb7-41f5-9d87-6970e53df229', '5d3984dd-58ef-42df-84ce-deb13dae434c', '19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 1, 'SKU-SOLUCIONESDIGITALES-PAPAAMARILLA', 'Papa amarilla', true, true, true, false, false, 5, 5, '{}'::jsonb, 5856, 'Papa amarilla de textura suave, perfecta para purés y guisos.', 1970, 4568,10,100,1);
INSERT INTO public.product_company VALUES ('497d1bd5-e037-4105-abc7-fec911094e47', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 1, 'SKU-DISTRIBUIDORALOSANDES-MANZANAROJA', 'Manzana roja', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 1962, 'Manzana roja crujiente y jugosa, perfecta para snacks y postres.', 2165, 5537,10,100,1);
INSERT INTO public.product_company VALUES ('4e5fa6b3-8254-4d54-92ab-4992399fb78a', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 1, 'SKU-SOLUCIONESDIGITALES-MANZANAROJA', 'Manzana roja', true, true, true, false, false, 5, 5, '{}'::jsonb, 4145, 'Manzana roja crujiente y jugosa, perfecta para snacks y postres.', 3550, 3963,10,100,1);
INSERT INTO public.product_company VALUES ('cd5f8712-488e-485e-ba9f-e2618c5ecbac', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '01337057-887f-4ba3-be7f-7e3296d680e8', 1, 'SKU-DISTRIBUIDORALOSANDES-BANANA', 'Banana', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4470, 'Banana madura y dulce, fuente natural de energía.', 2581, 1695,10,100,1);
INSERT INTO public.product_company VALUES ('a07e85fa-3350-4e2c-a047-bf7470bc4ace', '5d3984dd-58ef-42df-84ce-deb13dae434c', '01337057-887f-4ba3-be7f-7e3296d680e8', 1, 'SKU-SOLUCIONESDIGITALES-BANANA', 'Banana', true, true, true, false, false, 5, 5, '{}'::jsonb, 3334, 'Banana madura y dulce, fuente natural de energía.', 4266, 5609,10,100,1);
INSERT INTO public.product_company VALUES ('7c97c897-ee57-438d-9504-09dbc0f3a8a9', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'fec88f6a-af7f-49b2-bf11-f73c6236b580', 1, 'SKU-DISTRIBUIDORALOSANDES-NARANJA', 'Naranja', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 3743, 'Naranja fresca con alto contenido de vitamina C y sabor refrescante.', 4516, 2010,10,100,1);
INSERT INTO public.product_company VALUES ('7ba236e5-3cc1-4d05-ba3b-55e27bb86a3a', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'fec88f6a-af7f-49b2-bf11-f73c6236b580', 1, 'SKU-SOLUCIONESDIGITALES-NARANJA', 'Naranja', true, true, true, false, false, 5, 5, '{}'::jsonb, 2944, 'Naranja fresca con alto contenido de vitamina C y sabor refrescante.', 4567, 5441,10,100,1);
INSERT INTO public.product_company VALUES ('aaea6f21-3d2a-4fdd-93c0-f920c6877032', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 1, 'SKU-DISTRIBUIDORALOSANDES-UVAMORADA', 'Uva morada', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5423, 'Uva morada dulce y jugosa, ideal para consumir fresca o en postres.', 1567, 3368,10,100,1);
INSERT INTO public.product_company VALUES ('7587328d-0474-40f3-a988-a8d90b8a98a1', '5d3984dd-58ef-42df-84ce-deb13dae434c', '66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 1, 'SKU-SOLUCIONESDIGITALES-UVAMORADA', 'Uva morada', true, true, true, false, false, 5, 5, '{}'::jsonb, 3714, 'Uva morada dulce y jugosa, ideal para consumir fresca o en postres.', 4928, 5369,10,100,1);
INSERT INTO public.product_company VALUES ('4d869257-21c4-4ce8-9ae4-1b850607df7b', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '69f51fca-89eb-49db-a3c0-84fb149023a9', 1, 'SKU-DISTRIBUIDORALOSANDES-SANDÍA', 'Sandía', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 5927, 'Sandía fresca y dulce, perfecta para los días calurosos.', 3550, 3657,10,100,1);
INSERT INTO public.product_company VALUES ('fde77882-0c73-4f7e-b0fd-f55a68e60291', '5d3984dd-58ef-42df-84ce-deb13dae434c', '69f51fca-89eb-49db-a3c0-84fb149023a9', 1, 'SKU-SOLUCIONESDIGITALES-SANDÍA', 'Sandía', true, true, true, false, false, 5, 5, '{}'::jsonb, 5718, 'Sandía fresca y dulce, perfecta para los días calurosos.', 5968, 5634,10,100,1);
INSERT INTO public.product_company VALUES ('3e10b54f-6a54-4b41-a6e7-2d2bf5098ef1', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'faffb338-771d-4bd3-871d-4f3e1c37268d', 1, 'SKU-DISTRIBUIDORALOSANDES-PANMARRAQUETA', 'Pan marraqueta', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4886, 'Pan marraqueta recién horneado, crujiente por fuera y suave por dentro.', 2038, 1152,10,100,1);
INSERT INTO public.product_company VALUES ('a50cabfa-8800-470b-a55a-ce8d8b8a06bd', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'faffb338-771d-4bd3-871d-4f3e1c37268d', 1, 'SKU-SOLUCIONESDIGITALES-PANMARRAQUETA', 'Pan marraqueta', true, true, true, false, false, 5, 5, '{}'::jsonb, 5315, 'Pan marraqueta recién horneado, crujiente por fuera y suave por dentro.', 5724, 1009,10,100,1);
INSERT INTO public.product_company VALUES ('6ccaff61-a317-46a0-b264-dcf00ee93f09', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '397d445c-3b56-429a-9f52-7835fb340bb6', 1, 'SKU-DISTRIBUIDORALOSANDES-PANDEMOLDE', 'Pan de molde', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 3172, 'Pan de molde suave y esponjoso, ideal para sándwiches y tostadas.', 5528, 3027,10,100,1);
INSERT INTO public.product_company VALUES ('b42bd3bb-d418-45df-a4a8-c8aea9fca749', '5d3984dd-58ef-42df-84ce-deb13dae434c', '397d445c-3b56-429a-9f52-7835fb340bb6', 1, 'SKU-SOLUCIONESDIGITALES-PANDEMOLDE', 'Pan de molde', true, true, true, false, false, 5, 5, '{}'::jsonb, 3509, 'Pan de molde suave y esponjoso, ideal para sándwiches y tostadas.', 2065, 4603,10,100,1);
INSERT INTO public.product_company VALUES ('fc3ec16a-7743-4050-b055-78527d6b3b73', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'e30b3445-fe86-4914-a2d4-97f10e798922', 1, 'SKU-DISTRIBUIDORALOSANDES-CROISSANT', 'Croissant', true, true, true, false, false, 1, 1, '{"2": {"name": "gramos","abbreviation": "g","description": "unidad de gramos","factor": 0.001}}'::jsonb, 4946, 'Croissant hojaldrado y mantecoso, perfecto para el desayuno.', 5420, 2038,10,100,1);
INSERT INTO public.product_company VALUES ('60a2bcd7-d49c-4bf6-bef3-dcb394a9eb81', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'e30b3445-fe86-4914-a2d4-97f10e798922', 1, 'SKU-SOLUCIONESDIGITALES-CROISSANT', 'Croissant', true, true, true, false, false, 5, 5, '{}'::jsonb, 4091, 'Croissant hojaldrado y mantecoso, perfecto para el desayuno.', 3060, 2974,10,100,1);
END DEPRECATED */

INSERT INTO public.profile_accounts VALUES ('8f20d7a3-5160-4124-a743-a354b39e6507', 'tienda_sucursal_santiago', 'Perfil asociado a la tienda Sucursal Santiago');
INSERT INTO public.profile_accounts VALUES ('90cfcfa1-090e-426e-aefc-0353dc481489', 'tienda_oficina_central', 'Perfil asociado a la tienda Oficina Central');
INSERT INTO public.profile_accounts VALUES ('50cc64e1-62bc-4f58-b626-6c217d74d490', 'tienda_bodega_central', 'Perfil asociado a la tienda Bodega Central');
INSERT INTO public.profile_accounts VALUES ('7d385e85-7032-4fee-8dcc-b9d49dc41c70', 'tienda_centro_de_desarrollo', 'Perfil asociado a la tienda Centro de Desarrollo');

INSERT INTO public.profile_accounts VALUES ('e3a7b420-0c17-470e-8769-6d99ef0b2a3b', 'empresa_distribuidora_los_andes', 'Perfil asociado a la empresa Distribuidora Los Andes');
INSERT INTO public.profile_accounts VALUES ('c5045b0e-b49b-43fd-a733-3de160fabe48', 'empresa_soluciones_digitales', 'Perfil asociado a la empresa Soluciones Digitales');

INSERT INTO public.profile_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'admin', 'Administrador con todos los permisos');
INSERT INTO public.profile_accounts VALUES ('9fc2d9e2-c47a-412a-a574-506d8c03709f', 'gerente', 'Gerente general');
INSERT INTO public.profile_accounts VALUES ('584ab451-114d-45f6-8fac-119b89c9f607', 'proveedor', 'Proveedor externo');
INSERT INTO public.profile_accounts VALUES ('24ce1667-f111-4cfd-a7e5-42fe622c3bc2', 'bodeguero', 'Encargado de bodega');

INSERT INTO public.profile_account_per_power_accounts VALUES ('e3a7b420-0c17-470e-8769-6d99ef0b2a3b', 'fd2b8b08-0183-482b-b982-6d61cd9d1bc5');
INSERT INTO public.profile_account_per_power_accounts VALUES ('c5045b0e-b49b-43fd-a733-3de160fabe48', 'dc8c704e-9bcb-4dc4-afa0-fc66b41f1f40');

INSERT INTO public.profile_account_per_power_accounts VALUES ('50cc64e1-62bc-4f58-b626-6c217d74d490', '4ead438b-364f-4ef3-9b76-0145ecab4021');
INSERT INTO public.profile_account_per_power_accounts VALUES ('7d385e85-7032-4fee-8dcc-b9d49dc41c70', '3f44db76-f194-46b7-a306-e071488bc0ca');
INSERT INTO public.profile_account_per_power_accounts VALUES ('90cfcfa1-090e-426e-aefc-0353dc481489', '4cb22fae-bbce-4fc0-954d-fdf2b3622617');
INSERT INTO public.profile_account_per_power_accounts VALUES ('8f20d7a3-5160-4124-a743-a354b39e6507', '685b363e-170e-4e95-9d2c-22ca1682376e');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '1a866441-6c83-43ee-9061-8b37ecdba341');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'cad12d0b-1dbe-426d-9164-abacbca58e3e');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '1587fa74-c5c4-4221-9a03-c8e1a47b6851');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'de092967-af75-4142-8390-6478b23b51be');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '3880287a-f927-462c-ad6e-bce4c3a1b372');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '57eec64a-9e5a-4b2c-8ba2-7e18b845f7bb');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '03e72386-429c-4dfa-b33b-b6ca48b5154f');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'f8a1b2c3-d4e5-4f6a-7b8c-9d0e1f2a3b4c');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '6ca241a0-d1e4-40e6-bc6b-cb2981dc8caf');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '0ac945fd-183d-4593-8d19-9deaa312183b');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '330aefab-a77d-4b58-8c9f-8f36be9d8bfe');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'a1f2e3d4-c5b6-47a8-9b0c-1d2e3f4a5b6c'); -- request:createForStore
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'b2f3e4d5-c6b7-48a9-0c1d-2e3f4a5b6c7d'); -- request:createForCompany
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'c3f4e5d6-c7b8-49a0-1d2e-3f4a5b6c7d8e'); -- request:createForSupplier
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '3980c25a-f147-4f69-aa44-5af480675109');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '7b1f3f4e-3e2e-4d3a-9f4e-8c6f4e2b5a1d');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'd7dd6155-f6df-442f-9267-d9fb858dab56');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '385ffa1d-371d-4790-b83b-a60ff2718f4f');
INSERT INTO public.profile_account_per_power_accounts VALUES ('9fc2d9e2-c47a-412a-a574-506d8c03709f', '57eec64a-9e5a-4b2c-8ba2-7e18b845f7bb');
INSERT INTO public.profile_account_per_power_accounts VALUES ('9fc2d9e2-c47a-412a-a574-506d8c03709f', '03e72386-429c-4dfa-b33b-b6ca48b5154f');
INSERT INTO public.profile_account_per_power_accounts VALUES ('9fc2d9e2-c47a-412a-a574-506d8c03709f', '6ca241a0-d1e4-40e6-bc6b-cb2981dc8caf');
INSERT INTO public.profile_account_per_power_accounts VALUES ('9fc2d9e2-c47a-412a-a574-506d8c03709f', '0ac945fd-183d-4593-8d19-9deaa312183b');
INSERT INTO public.profile_account_per_power_accounts VALUES ('9fc2d9e2-c47a-412a-a574-506d8c03709f', '330aefab-a77d-4b58-8c9f-8f36be9d8bfe');
INSERT INTO public.profile_account_per_power_accounts VALUES ('24ce1667-f111-4cfd-a7e5-42fe622c3bc2', '0ac945fd-183d-4593-8d19-9deaa312183b');
INSERT INTO public.profile_account_per_power_accounts VALUES ('24ce1667-f111-4cfd-a7e5-42fe622c3bc2', '330aefab-a77d-4b58-8c9f-8f36be9d8bfe');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '6568de4f-f4cb-4483-af91-6cde266be10e');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'eb72f588-7ba1-4160-a706-b38a7dc04cef');

-- Admin: Warehouse powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e');

-- Admin: Product powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'd4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a');

-- Admin: Product Company powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'e5f6a7b8-c9d0-4e1f-2a3b-4c5d6e7f8a9b');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c');

-- Admin: Inventory Count powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'a7b8c9d0-e1f2-4a3b-4c5d-6e7f8a9b0c1d');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e');

-- Admin: Product Movement powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'd0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a');

-- Se agregan a admin los permisos de plantillas de productos
--f4e5d2c2-1f4e-4d3a-9f4e-8c6f4e2b5a2e
--a3b2c1d4-e5f6-4789-abcd-ef0123456789

INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'f4e5d2c2-1f4e-4d3a-9f4e-8c6f4e2b5a2e');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'a3b2c1d4-e5f6-4789-abcd-ef0123456789');

-- Admin: Store Product powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'e1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'f2b3c4d5-e6f7-4a8b-9c0d-1e2f3a4b5c6d');

-- Admin: Delivery Purchase Note powers
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'a5b6c7d8-e9f0-4a1b-2c3d-4e5f6a7b8c9d');
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'b6c7d8e9-f0a1-4b2c-3d4e-5f6a7b8c9d0e');

-- Admin: Ownership powers for Companies and Stores (seed fixed IDs)
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'fd2b8b08-0183-482b-b982-6d61cd9d1bc5'); -- company:43f96ad5-6c50-4ce8-b7cf-203210741a18
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', 'dc8c704e-9bcb-4dc4-afa0-fc66b41f1f40'); -- company:5d3984dd-58ef-42df-84ce-deb13dae434c
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '685b363e-170e-4e95-9d2c-22ca1682376e'); -- store:ce79763f-175b-44ee-8cdd-47fc95d4a8ce
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '4cb22fae-bbce-4fc0-954d-fdf2b3622617'); -- store:9f666bab-35e3-46e6-b147-75354262dc84
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '4ead438b-364f-4ef3-9b76-0145ecab4021'); -- store:86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6
INSERT INTO public.profile_account_per_power_accounts VALUES ('e481d36c-ab9e-4dac-8636-a8a278b5cd0e', '3f44db76-f194-46b7-a306-e071488bc0ca'); -- store:90775541-4b56-4327-84b7-0927f46122d7
------------------------------------


INSERT INTO public.store VALUES ('90775541-4b56-4327-84b7-0927f46122d7', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'Centro de Desarrollo', 'Calle Innovación 101', 'Centro de desarrollo y soporte técnico', 'CC-004', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.store VALUES ('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'Sucursal Santiago', 'Av. Libertador Bernardo O’Higgins 1234', 'Sucursal principal en Santiago', 'CC-001', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.store VALUES ('9f666bab-35e3-46e6-b147-75354262dc84', '5d3984dd-58ef-42df-84ce-deb13dae434c', 'Oficina Central', 'Calle Nueva Providencia 456', 'Oficina administrativa y de desarrollo', 'CC-002', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.store VALUES ('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', '43f96ad5-6c50-4ce8-b7cf-203210741a18', 'Bodega Central', 'Calle Comercio 789', 'Bodega principal para almacenamiento de productos', 'CC-003', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');

/* DEPRECATED: request_restriction table has been deprecated. Data now lives in product_per_store.max_quantity.
INSERT INTO public.request_restriction VALUES ('fd924e89-d705-414d-9899-c7bbf53ee9d6', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '26d261de-5b9d-44fa-bdf6-c5fbf0f15ad3', 10);
INSERT INTO public.request_restriction VALUES ('dc582cb2-50e3-427f-b464-869cc6853597', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '7e3fafcd-0197-4f62-84b4-ed5c753921ac', 10);
END DEPRECATED */

INSERT INTO public.supplier VALUES ('fc594468-422b-4dbc-8e67-76b97eea3c21', 'cdb6f2eb-596f-41e7-acc5-52c5b567f8cf', 1, 'Proveedor de Alimentos S.A.', 'Suministra productos alimenticios variados.', true, '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.supplier VALUES ('924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'd0c56423-7127-49a3-a2b0-2d7023141cf0', 1, 'Tech Supplies Ltda.', 'Provee hardware y software para empresas.', true, '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.supplier VALUES ('a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'e1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c', 1, 'Proveedor Ejemplo S.A.', 'Proveedor de insumos y servicios.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', 'f1a1b1c1-d1e1-4f1a-8b1c-0d1e2f3a4b1c', 1, 'Frutas del Valle Ltda.', 'Proveedor de frutas frescas y procesadas.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('a2b2c2d2-e2f2-4a2b-8c2d-0e2f2a3b4c2d', 'f2a2b2c2-d2e2-4f2a-8b2c-0d2e2f3a4b2c', 1, 'Panadería El Trigal SpA', 'Proveedor de pan y productos de pastelería.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', 'f3a3b3c3-d3e3-4f3a-8b3c-0d3e2f3a4b3c', 1, 'Carnes Premium S.A.', 'Proveedor de carnes selectas y embutidos.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', 'f4a4b4c4-d4e4-4f4a-8b4c-0d4e2f3a4b4c', 1, 'Verduras Frescas Ltda.', 'Proveedor de verduras y hortalizas frescas.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', 'f5a5b5c5-d5e5-4f5a-8b5c-0d5e2f3a4b5c', 1, 'Bebidas y Más SpA', 'Proveedor de bebidas y jugos.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', 'a6b6c6d6-e6f6-4a6b-8c6d-0e6f2a3b4c6d', 1, 'Aceites del Pacífico S.A.', 'Proveedor de aceites y grasas vegetales.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', 'a7b7c7d7-e7f7-4a7b-8c7d-0e7f2a3b4c7d', 1, 'Lácteos del Sur Ltda.', 'Proveedor de leche, quesos y yogurt.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e', 'a8b8c8d8-e8f8-4a8b-8c8d-0e8f2a3b4c8d', 1, 'Pastas Italia SpA', 'Proveedor de pastas y salsas italianas.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'b5c5d5e5-f5a5-4b5c-8d5e-0f5a2b3c4d5e', 1, 'Cereales Premium SpA', 'Proveedor de cereales y barras energéticas.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', 'b4c4d4e4-f4a4-4b4c-8d4e-0f4a2b3c4d4e', 1, 'Snacks Express Ltda.', 'Proveedor de snacks y productos para cafetería.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', 'b3c3d3e3-f3a3-4b3c-8d3e-0f3a2b3c4d3e', 1, 'Bebidas del Centro S.A.', 'Proveedor de bebidas y jugos.', true, NOW(), NOW());
INSERT INTO public.supplier VALUES ('c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f', 'b2c2d2e2-f2a2-4b2c-8d2e-0f2a2b3c4d2e', 1, 'Legumbres Chile Ltda.', 'Proveedor de legumbres y granos.', true, NOW(), NOW());

INSERT INTO public.supplier_contact VALUES ('b0940623-aaea-4ed8-b8cd-6239544c3e58', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'Juan Pérez', 'Gerente de ventas', 'juan.perez@proveedordealimentos.cl', '+56912345678', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.supplier_contact VALUES ('14f19e21-539b-4a0e-b9b8-011bd7c812d1', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'María López', 'Ejecutiva de ventas', 'maria.lopez@techsupplies.cl', '+56987654321', '2025-09-05 21:54:06.914962', '2025-09-05 21:54:06.914962');
INSERT INTO public.supplier_contact VALUES ('c1d2e3f4-a5b6-4c7d-8e9f-0a1b2c3d4e5f', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'Ana Ejemplo', 'Ejecutiva Comercial', 'ana.ejemplo@proveedorejemplo.cl', '+56912345679', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('c1d1e1f1-a1b1-4c1d-8e1f-0a1b2c3d4e1f', 'a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', 'Sofía Valle', 'Ejecutiva Comercial', 'sofia.valle@frutasdelvalle.cl', '+56911112223', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f', 'a2b2c2d2-e2f2-4a2b-8c2d-0e2f2a3b4c2d', 'Luis Trigal', 'Jefe de Ventas', 'luis.trigal@eltrigal.cl', '+56922223334', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', 'a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', 'Marcela Premium', 'Ejecutiva Comercial', 'marcela.premium@carnespremium.cl', '+56933334445', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', 'a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', 'Pedro Verde', 'Jefe de Compras', 'pedro.verde@verdurasfrescas.cl', '+56944445556', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', 'Andrea Más', 'Ejecutiva Comercial', 'andrea.mas@bebidasymas.cl', '+56955556667', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('d6e6f6a6-b6c6-4d6e-8f6a-0b6c2d3e4f6a', 'b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', 'Patricia Pacífico', 'Jefa Comercial', 'patricia.pacifico@aceitespacifico.cl', '+56966667778', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('d7e7f7a7-b7c7-4d7e-8f7a-0b7c2d3e4f7a', 'b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', 'Ricardo Sur', 'Ejecutivo de Ventas', 'ricardo.sur@lacteosdelsur.cl', '+56977778889', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('d8e8f8a8-b8c8-4d8e-8f8a-0b8c2d3e4f8a', 'b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e', 'Giovanni Italia', 'Gerente Comercial', 'giovanni.italia@pastasitalia.cl', '+56988889990', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('e5f5a5b5-c5d5-4e5f-8a5b-0c5d2e3f4a5b', 'c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'Ignacio Premium', 'Jefe de Ventas', 'ignacio.premium@cerealespremium.cl', '+56955556667', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('e4f4a4b4-c4d4-4e4f-8a4b-0c4d2e3f4a4b', 'c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', 'Valeria Express', 'Ejecutiva Comercial', 'valeria.express@snacksexpress.cl', '+56944445556', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('e3f3a3b3-c3d3-4e3f-8a3b-0c3d2e3f4a3b', 'c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', 'Martín Centro', 'Jefe de Ventas', 'martin.centro@bebidasdelcentro.cl', '+56933334445', NOW(), NOW());
INSERT INTO public.supplier_contact VALUES ('e2f2a2b2-c2d2-4e2f-8a2b-0c2d2e3f4a2b', 'c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f', 'Carolina Chile', 'Ejecutiva Comercial', 'carolina.chile@legumbreschile.cl', '+56922223334', NOW(), NOW());

INSERT INTO public.supplier_per_company (id, supplier_id, company_id, created_at) VALUES ('a46d3ca1-060e-4acf-9c52-fbc44200a204', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '43f96ad5-6c50-4ce8-b7cf-203210741a18', '2025-09-05 21:54:06.914962');
INSERT INTO public.supplier_per_company (id, supplier_id, company_id, created_at) VALUES ('1f622dd8-d5a0-41cb-9343-121f5b4aaba7', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '5d3984dd-58ef-42df-84ce-deb13dae434c', '2025-09-05 21:54:06.914962');

INSERT INTO public.supplier_product (id, supplier_id, product_id, product_name, description, sku, unit_price, purchase_unit_id, available, created_at, updated_at) VALUES 
('913b50df-b1f9-42d8-912a-bdfdb22f2575', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche entera', 'Leche entera fresca y nutritiva', 'SUP-LECHE-001', 850, 1, true, '2025-09-05 21:54:06.938377', '2025-09-05 21:54:06.938377'),
('91ec47cd-4097-4a01-9e9f-b202d9f085e1', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 'Carne molida', 'Carne molida de alta calidad', 'SUP-CARNE-001', 5500, 1, true, '2025-09-05 21:54:06.942818', '2025-09-05 21:54:06.942818'),
('b5c3aaa8-b3b8-4d3b-b950-d299817dc708', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana roja', 'Manzana roja crujiente y jugosa', 'SUP-MANZ-001', 1800, 1, true, '2025-09-05 21:54:06.943469', '2025-09-05 21:54:06.943469'),
('4cfb1798-9712-4429-b712-4aeb55587276', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt Natural', 'Yogurt natural cremoso y saludable', 'SUP-YOGURT-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('47b29ea7-8d34-4198-a866-cda0729accd1', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso Mantecoso', 'Queso mantecoso de textura suave y sabor tradicional', 'SUP-QUESO-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('f34e3cb6-268c-4dbf-9792-8a0a733a47e1', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla Sin Sal', 'Mantequilla sin sal elaborada con ingredientes de alta calidad', 'SUP-MANTEQUILLA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('4b18a78e-2b07-4ea3-b3d7-eda2d2eeb7e1', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de Leche', 'Crema de leche fresca para realzar tus recetas', 'SUP-CREMA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('65969010-803b-4f7c-9e0a-bc5f9feae1a5', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'Bebida gaseosa Coca-Cola en presentación familiar', 'SUP-COCA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('5b39fada-3aed-42be-8725-078977dc6419', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua Mineral 500ml', 'Agua mineral natural embotellada', 'SUP-AGUA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('d7e663d6-20a3-4bb1-8040-451c1c9abe7a', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de Naranja', 'Jugo de naranja 100% natural y refrescante', 'SUP-JUGO-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('ba7f4155-744b-4de9-baca-76a4bff3e092', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té Helado', 'Té helado listo para tomar con sabor suave', 'SUP-TE-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('9d697908-a6ba-44eb-925a-01412b7e4a73', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida Energética', 'Bebida energética para revitalizar tu día', 'SUP-ENERGETICA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('ed53ff8c-7e88-4147-9da5-e8bbde34f085', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz Grano Largo', 'Arroz grano largo de alta calidad', 'SUP-ARROZ-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('c9ebeca8-840c-4fd8-81c2-ab689bb7a004', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de Trigo', 'Harina de trigo fina y versátil', 'SUP-HARINA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('d6ad3811-8c92-4c07-8ce1-e0d83cb9e094', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite Vegetal', 'Aceite vegetal puro para cocinar', 'SUP-ACEITE-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('23cd2701-6e91-4359-846d-14a43f35fbe7', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '912c9f07-5810-438f-ad4c-11790d181910', 'Fideos Espagueti', 'Fideos espagueti de textura firme', 'SUP-FIDEOS-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('64d2fdd0-c41f-4497-91b0-a9a763d86434', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar Refinada', 'Azúcar refinada blanca y versátil', 'SUP-AZUCAR-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('fce49a3e-05b1-472d-9777-45b0c9c80c36', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 'Papas Fritas Clásicas', 'Papas fritas crujientes y deliciosas', 'SUP-PAPAS-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('d7d0ee70-520b-478f-95a8-eb3d84d5cc6d', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 'Galletas de Chocolate', 'Galletas de chocolate con trozos generosos', 'SUP-GALLETAS-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('0575ca88-a233-421a-96c0-0635707f6ee8', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 'Maní Salado', 'Maní salado tostado, ideal como snack', 'SUP-MANI-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('93f65612-51a4-4a8c-96e0-caefc69c215b', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '92599d39-eddb-4e9d-9012-b209675ff9bd', 'Chocolate en Barra', 'Chocolate en barra con sabor intenso', 'SUP-CHOCOLATE-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('5e59ff66-8a0b-40d0-9afa-6f2d1a1ec026', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 'Barra de Cereal', 'Barra de cereal nutritiva y práctica', 'SUP-BARRA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('ccc81d7e-c523-45db-a00b-dd6ede9f2ea1', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'db682432-362e-4e7b-af54-0d87f91832da', 'Pechuga de Pollo', 'Pechuga de pollo fresca y jugosa', 'SUP-PECHUGA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('20a73156-e282-41d6-b379-a777965a81a4', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '95523bb0-c72f-4a4c-affe-f14c9a32870e', 'Chuleta de Cerdo', 'Chuleta de cerdo tierna y sabrosa', 'SUP-CHULETA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('ef96851b-dbfc-4f36-a15d-fb3f13ad6479', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'e2432d39-68ca-4a02-81d2-e200e47ea620', 'Filete de Vacuno', 'Filete de vacuno selecto y de alta calidad', 'SUP-FILETE-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('dcea4308-ea2b-4850-b0f3-c5bdeba7621f', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 'Longaniza', 'Longaniza tradicional con un sabor auténtico', 'SUP-LONGANIZA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('3c67c2df-a9f1-4431-990d-20ca89c1cb0c', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate Fresco', 'Tomate fresco y jugoso, ideal para ensaladas', 'SUP-TOMATE-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('010a55c4-3750-44d4-8013-2a25623b6aa6', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 'Lechuga', 'Lechuga fresca y crujiente para ensaladas', 'SUP-LECHUGA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('69892e6d-8dba-4e91-8445-8d6e9140be63', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'f0dddfbb-a31c-4247-ba0b-d37727672f51', 'Zanahoria', 'Zanahoria dulce y crujiente, ideal para cocinar', 'SUP-ZANAHORIA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('24411dc0-8d55-4ea6-ae84-9b3ccafd5c89', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '25fbe34d-f25f-43c9-a192-f0b24bac0a66', 'Pimentón Rojo', 'Pimentón rojo dulce y colorido', 'SUP-PIMENTON-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('c1612f43-0e64-4682-a6e1-86901d61438b', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 'Papa Amarilla', 'Papa amarilla de textura suave', 'SUP-PAPA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('67df6da6-b89d-4894-be98-4ab87c675234', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '01337057-887f-4ba3-be7f-7e3296d680e8', 'Banana', 'Banana madura y dulce, fuente de energía', 'SUP-BANANA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('2e2b04d6-4783-4d07-a352-5cebb97d3b10', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'fec88f6a-af7f-49b2-bf11-f73c6236b580', 'Naranja', 'Naranja fresca con alto contenido de vitamina C', 'SUP-NARANJA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('883656c2-d59a-4c53-bdd9-57e22bc311dd', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 'Uva Morada', 'Uva morada dulce y jugosa', 'SUP-UVA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('7304efbe-22a3-4cb4-bb1d-5f8bb829a426', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '69f51fca-89eb-49db-a3c0-84fb149023a9', 'Sandía', 'Sandía fresca y dulce, ideal para el verano', 'SUP-SANDIA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('90981b11-75a0-418f-aa9b-81727959a81e', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'faffb338-771d-4bd3-871d-4f3e1c37268d', 'Pan Marraqueta', 'Pan marraqueta recién horneado', 'SUP-MARRAQUETA-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('f0b32070-ac99-4ed8-83fd-66517f1c0ad3', 'fc594468-422b-4dbc-8e67-76b97eea3c21', '397d445c-3b56-429a-9f52-7835fb340bb6', 'Pan de Molde', 'Pan de molde suave y esponjoso', 'SUP-MOLDE-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('af694a1b-3ff8-472e-9783-82c3f3474c2e', 'fc594468-422b-4dbc-8e67-76b97eea3c21', 'e30b3445-fe86-4914-a2d4-97f10e798922', 'Croissant', 'Croissant hojaldrado y mantecoso', 'SUP-CROISSANT-001', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('4c7ff00c-23e0-43ce-8fcb-002daa293d98', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt Natural', 'Yogurt natural cremoso y saludable', 'SUP-YOGURT-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('2617ebfa-5215-409d-a294-97d0fca14f28', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso Mantecoso', 'Queso mantecoso de textura suave y sabor tradicional', 'SUP-QUESO-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('d3a7f9dc-1735-4476-a13b-31232842695f', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla Sin Sal', 'Mantequilla sin sal elaborada con ingredientes de alta calidad', 'SUP-MANTEQUILLA-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('181db0c8-6b0f-4c79-9409-8727f1319a0f', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de Leche', 'Crema de leche fresca para realzar tus recetas', 'SUP-CREMA-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('844c61a2-27aa-48df-a676-a1283bd92340', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'Bebida gaseosa Coca-Cola en presentación familiar', 'SUP-COCA-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('a1143ac7-69c9-4431-aa58-6aded06e8703', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua Mineral 500ml', 'Agua mineral natural embotellada', 'SUP-AGUA-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('87299350-6cf0-432a-912a-5380de9b49c3', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de Naranja', 'Jugo de naranja 100% natural y refrescante', 'SUP-JUGO-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('2a753ddd-4135-4cc2-bcc9-1f2ea361863d', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té Helado', 'Té helado listo para tomar con sabor suave', 'SUP-TE-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('dc12d07a-6cb9-420c-b8dd-5d57c3932b33', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida Energética', 'Bebida energética para revitalizar tu día', 'SUP-ENERGETICA-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('1d9e6e97-e6c5-4cee-8b42-3ee87adfbbf7', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz Grano Largo', 'Arroz grano largo de alta calidad', 'SUP-ARROZ-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('faf42b6b-5d60-4016-bd5c-a26d2e55960a', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de Trigo', 'Harina de trigo fina y versátil', 'SUP-HARINA-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('3494dd6c-8158-4163-9c01-dc80299d8840', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite Vegetal', 'Aceite vegetal puro para cocinar', 'SUP-ACEITE-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('738b1e49-10bc-4cab-9a24-fd5b8ed4e516', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', '912c9f07-5810-438f-ad4c-11790d181910', 'Fideos Espagueti', 'Fideos espagueti de textura firme', 'SUP-FIDEOS-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271'),
('c4a1a9c4-fdec-42bd-b4a2-846ffc44dfd8', '924d14d0-5d46-4af0-bef0-878c8d1ff10c', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar Refinada', 'Azúcar refinada blanca y versátil', 'SUP-AZUCAR-002', 1000, 1, true, '2025-09-08 15:30:52.661271', '2025-09-08 15:30:52.661271');

INSERT INTO public.supplier_product (id, supplier_id, product_id, product_name, description, sku, unit_price, purchase_unit_id, available, created_at, updated_at) VALUES
-- Proveedor: Frutas del Valle Ltda. (a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d)
('f1a1b1c1-0001-4001-a001-000000000001', 'a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana Roja', 'Manzana roja fresca del valle', 'FV-MANZANA-001', 2500, 1, true, NOW(), NOW()),
('f1a1b1c1-0001-4001-a001-000000000002', 'a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', '01337057-887f-4ba3-be7f-7e3296d680e8', 'Banana', 'Banana madura ecuatoriana', 'FV-BANANA-001', 1800, 1, true, NOW(), NOW()),
('f1a1b1c1-0001-4001-a001-000000000003', 'a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', 'fec88f6a-af7f-49b2-bf11-f73c6236b580', 'Naranja', 'Naranja fresca con vitamina C', 'FV-NARANJA-001', 2200, 1, true, NOW(), NOW()),
('f1a1b1c1-0001-4001-a001-000000000004', 'a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', '66d0b490-5f5b-42aa-8f31-598f94cb9eaf', 'Uva Morada', 'Uva morada dulce sin pepas', 'FV-UVA-001', 3200, 1, true, NOW(), NOW()),
('f1a1b1c1-0001-4001-a001-000000000005', 'a1b1c1d1-e1f1-4a1b-8c1d-0e1f2a3b4c1d', '69f51fca-89eb-49db-a3c0-84fb149023a9', 'Sandía', 'Sandía fresca y jugosa', 'FV-SANDIA-001', 4500, 5, true, NOW(), NOW()),

-- Proveedor: Proveedor Ejemplo S.A. (a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d)
('b1c1d1e1-0012-4012-a012-000000000001', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche Entera', 'Leche entera 1 litro', 'PE-LECHE-001', 1300, 3, true, NOW(), NOW()),
('b1c1d1e1-0012-4012-a012-000000000002', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'db682432-362e-4e7b-af54-0d87f91832da', 'Pechuga de Pollo', 'Pechuga de pollo fresca 1kg', 'PE-POLLO-001', 7800, 1, true, NOW(), NOW()),
('b1c1d1e1-0012-4012-a012-000000000003', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate Fresco', 'Tomate fresco 1kg', 'PE-TOMATE-001', 1700, 1, true, NOW(), NOW()),
('b1c1d1e1-0012-4012-a012-000000000004', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', '2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'Bebida Coca-Cola 1.5 litros', 'PE-COCA-001', 2300, 5, true, NOW(), NOW()),
('b1c1d1e1-0012-4012-a012-000000000005', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', '99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz Grano Largo', 'Arroz 1 kilogramo', 'PE-ARROZ-001', 3300, 1, true, NOW(), NOW()),
('b1c1d1e1-0012-4012-a012-000000000006', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', '1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 'Papas Fritas', 'Papas fritas clásicas 170g', 'PE-PAPAS-001', 2000, 5, true, NOW(), NOW()),

-- Proveedor: Panadería El Trigal SpA (a2b2c2d2-e2f2-4a2b-8c2d-0e2f2a3b4c2d)
('f2a2b2c2-0002-4002-a002-000000000001', 'a2b2c2d2-e2f2-4a2b-8c2d-0e2f2a3b4c2d', 'faffb338-771d-4bd3-871d-4f3e1c37268d', 'Pan Marraqueta', 'Pan marraqueta recién horneado', 'PT-PAN-001', 1200, 5, true, NOW(), NOW()),
('f2a2b2c2-0002-4002-a002-000000000002', 'a2b2c2d2-e2f2-4a2b-8c2d-0e2f2a3b4c2d', '397d445c-3b56-429a-9f52-7835fb340bb6', 'Pan de Molde', 'Pan de molde integral', 'PT-MOLDE-001', 2800, 5, true, NOW(), NOW()),
('f2a2b2c2-0002-4002-a002-000000000003', 'a2b2c2d2-e2f2-4a2b-8c2d-0e2f2a3b4c2d', 'e30b3445-fe86-4914-a2d4-97f10e798922', 'Croissant', 'Croissant francés mantequilla', 'PT-CROISSANT-001', 3500, 5, true, NOW(), NOW()),

-- Proveedor: Carnes Premium S.A. (a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d)
('f3a3b3c3-0003-4003-a003-000000000001', 'a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', 'db682432-362e-4e7b-af54-0d87f91832da', 'Pechuga de Pollo', 'Pechuga de pollo premium', 'CP-POLLO-001', 7500, 1, true, NOW(), NOW()),
('f3a3b3c3-0003-4003-a003-000000000002', 'a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', '60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 'Carne Molida', 'Carne molida premium', 'CP-MOLIDA-001', 8500, 1, true, NOW(), NOW()),
('f3a3b3c3-0003-4003-a003-000000000003', 'a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', 'e2432d39-68ca-4a02-81d2-e200e47ea620', 'Filete de Vacuno', 'Filete de vacuno selecto', 'CP-FILETE-001', 15000, 1, true, NOW(), NOW()),
('f3a3b3c3-0003-4003-a003-000000000004', 'a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', '95523bb0-c72f-4a4c-affe-f14c9a32870e', 'Chuleta de Cerdo', 'Chuleta de cerdo tierna', 'CP-CHULETA-001', 6800, 1, true, NOW(), NOW()),
('f3a3b3c3-0003-4003-a003-000000000005', 'a3b3c3d3-e3f3-4a3b-8c3d-0e3f2a3b4c3d', '3fc0c0e8-f6c5-48f3-bc18-5854232bdbbb', 'Longaniza', 'Longaniza tradicional', 'CP-LONGANIZA-001', 5500, 1, true, NOW(), NOW()),

-- Proveedor: Verduras Frescas Ltda. (a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d)
('f4a4b4c4-0004-4004-a004-000000000001', 'a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate Fresco', 'Tomate tipo pera fresco', 'VF-TOMATE-001', 1800, 1, true, NOW(), NOW()),
('f4a4b4c4-0004-4004-a004-000000000002', 'a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', 'a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 'Lechuga', 'Lechuga fresca hidropónica', 'VF-LECHUGA-001', 1500, 5, true, NOW(), NOW()),
('f4a4b4c4-0004-4004-a004-000000000003', 'a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', 'f0dddfbb-a31c-4247-ba0b-d37727672f51', 'Zanahoria', 'Zanahoria dulce y crujiente', 'VF-ZANAHORIA-001', 1200, 1, true, NOW(), NOW()),
('f4a4b4c4-0004-4004-a004-000000000004', 'a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', '25fbe34d-f25f-43c9-a192-f0b24bac0a66', 'Pimentón Rojo', 'Pimentón rojo dulce', 'VF-PIMENTON-001', 2500, 1, true, NOW(), NOW()),
('f4a4b4c4-0004-4004-a004-000000000005', 'a4b4c4d4-e4f4-4a4b-8c4d-0e4f2a3b4c4d', '19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 'Papa Amarilla', 'Papa amarilla tiquichoca', 'VF-PAPA-001', 1600, 1, true, NOW(), NOW()),

-- Proveedor: Bebidas y Más SpA (a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d)
('f5a5b5c5-0005-4005-a005-000000000001', 'a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', '2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'Bebida Coca-Cola 1.5 litros', 'BM-COCA-001', 2200, 5, true, NOW(), NOW()),
('f5a5b5c5-0005-4005-a005-000000000002', 'a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua Mineral 500ml', 'Agua mineral 500ml pack 24', 'BM-AGUA-001', 1400, 16, true, NOW(), NOW()),
('f5a5b5c5-0005-4005-a005-000000000003', 'a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', '67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de Naranja', 'Jugo 100% natural 1 litro', 'BM-JUGO-001', 3000, 5, true, NOW(), NOW()),
('f5a5b5c5-0005-4005-a005-000000000004', 'a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té Helado', 'Té helado 1.5 litros', 'BM-TE-001', 2600, 5, true, NOW(), NOW()),
('f5a5b5c5-0005-4005-a005-000000000005', 'a5b5c5d5-e5f5-4a5b-8c5d-0e5f2a3b4c5d', '946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida Energética', 'Bebida energética 250ml pack 6', 'BM-ENERGETICA-001', 3300, 13, true, NOW(), NOW()),

-- Proveedor: Aceites del Pacífico S.A. (b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e)
('a6b6c6d6-0006-4006-a006-000000000001', 'b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite Vegetal', 'Aceite vegetal 1 litro', 'AP-ACEITE-001', 3800, 3, true, NOW(), NOW()),
('a6b6c6d6-0006-4006-a006-000000000002', 'b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', '800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla Sin Sal', 'Mantequilla sin sal 250g', 'AP-MANTEQUILLA-001', 4200, 5, true, NOW(), NOW()),
('a6b6c6d6-0006-4006-a006-000000000003', 'b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', 'c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de Leche', 'Crema de leche 250ml', 'AP-CREMA-001', 3900, 4, true, NOW(), NOW()),
('a6b6c6d6-0006-4006-a006-000000000004', 'b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso Mantecoso', 'Queso mantecoso premium 300g', 'AP-QUESO-001', 4800, 5, true, NOW(), NOW()),
('a6b6c6d6-0006-4006-a006-000000000005', 'b6c6d6e6-f6a6-4b6c-8d6e-0f6a2b3c4d6e', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche Entera', 'Leche entera descremada 1L', 'AP-LECHE-001', 1500, 3, true, NOW(), NOW()),

-- Proveedor: Lácteos del Sur Ltda. (b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e)
('a7b7c7d7-0007-4007-a007-000000000001', 'b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche Entera', 'Leche entera 1 litro', 'LS-LECHE-001', 1200, 3, true, NOW(), NOW()),
('a7b7c7d7-0007-4007-a007-000000000002', 'b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt Natural', 'Yogurt natural sin azúcar 1L', 'LS-YOGURT-001', 2800, 3, true, NOW(), NOW()),
('a7b7c7d7-0007-4007-a007-000000000003', 'b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso Mantecoso', 'Queso mantecoso 500g', 'LS-QUESO-001', 4200, 1, true, NOW(), NOW()),
('a7b7c7d7-0007-4007-a007-000000000004', 'b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', '800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla Sin Sal', 'Mantequilla 200g', 'LS-MANTEQUILLA-001', 3600, 5, true, NOW(), NOW()),
('a7b7c7d7-0007-4007-a007-000000000005', 'b7c7d7e7-f7a7-4b7c-8d7e-0f7a2b3c4d7e', 'c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de Leche', 'Crema de leche 200ml', 'LS-CREMA-001', 3200, 4, true, NOW(), NOW()),

-- Proveedor: Pastas Italia SpA (b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e)
('a8b8c8d8-0008-4008-a008-000000000001', 'b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e', '912c9f07-5810-438f-ad4c-11790d181910', 'Fideos Espagueti', 'Fideos espagueti 500g', 'PI-FIDEOS-001', 1600, 5, true, NOW(), NOW()),
('a8b8c8d8-0008-4008-a008-000000000002', 'b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de Trigo', 'Harina de trigo para pastas 1kg', 'PI-HARINA-001', 1800, 1, true, NOW(), NOW()),
('a8b8c8d8-0008-4008-a008-000000000003', 'b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate Fresco', 'Tomate para salsa italiana 1kg', 'PI-TOMATE-001', 1900, 1, true, NOW(), NOW()),
('a8b8c8d8-0008-4008-a008-000000000004', 'b8c8d8e8-f8a8-4b8c-8d8e-0f8a2b3c4d8e', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite de Oliva', 'Aceite de oliva virgen extra 750ml', 'PI-ACEITE-001', 5200, 3, true, NOW(), NOW()),

-- Proveedor: Legumbres Chile Ltda. (c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f)
('b2c2d2e2-0011-4011-a011-000000000001', 'c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f', '99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz Grano Largo', 'Arroz 1 kilogramo', 'LC-ARROZ-001', 3200, 1, true, NOW(), NOW()),
('b2c2d2e2-0011-4011-a011-000000000002', 'c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de Trigo', 'Harina de trigo 1kg', 'LC-HARINA-001', 1900, 1, true, NOW(), NOW()),
('b2c2d2e2-0011-4011-a011-000000000003', 'c2d2e2f2-a2b2-4c2d-8e2f-0a2b2c3d4e2f', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar Refinada', 'Azúcar blanca 1kg', 'LC-AZUCAR-001', 2600, 1, true, NOW(), NOW()),

-- Proveedor: Bebidas del Centro S.A. (c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f)
('b3c3d3e3-0012-4012-a012-000000000001', 'c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', '2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', 'Bebida Coca-Cola 1.5 litros', 'BC-COCA-001', 2400, 5, true, NOW(), NOW()),
('b3c3d3e3-0012-4012-a012-000000000002', 'c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua Mineral 500ml', 'Agua mineral 500ml pack 24', 'BC-AGUA-001', 1600, 16, true, NOW(), NOW()),
('b3c3d3e3-0012-4012-a012-000000000003', 'c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', '67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de Naranja', 'Jugo natural 100% 1 litro', 'BC-JUGO-001', 3100, 5, true, NOW(), NOW()),
('b3c3d3e3-0012-4012-a012-000000000004', 'c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té Helado', 'Té helado natural 1.5 litros', 'BC-TE-001', 2700, 5, true, NOW(), NOW()),
('b3c3d3e3-0012-4012-a012-000000000005', 'c3d3e3f3-a3b3-4c3d-8e3f-0a3b2c3d4e3f', '946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida Energética', 'Bebida energética 250ml pack 6', 'BC-ENERGETICA-001', 3400, 13, true, NOW(), NOW()),

-- Proveedor: Snacks Express Ltda. (c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f)
('b4c4d4e4-0010-4010-a010-000000000001', 'c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', '1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 'Papas Fritas Clásicas', 'Papas fritas 170g', 'SE-PAPAS-001', 1900, 5, true, NOW(), NOW()),
('b4c4d4e4-0010-4010-a010-000000000002', 'c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', 'a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 'Galletas de Chocolate', 'Galletas chocolate 200g', 'SE-GALLETAS-001', 1700, 5, true, NOW(), NOW()),
('b4c4d4e4-0010-4010-a010-000000000003', 'c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', '0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 'Maní Salado', 'Maní salado 100g', 'SE-MANI-001', 1100, 5, true, NOW(), NOW()),
('b4c4d4e4-0010-4010-a010-000000000004', 'c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', '92599d39-eddb-4e9d-9012-b209675ff9bd', 'Chocolate en Barra', 'Chocolate 100g', 'SE-CHOCOLATE-001', 2100, 5, true, NOW(), NOW()),
('b4c4d4e4-0010-4010-a010-000000000005', 'c4d4e4f4-a4b4-4c4d-8e4f-0a4b2c3d4e4f', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana Verde', 'Manzana verde fresca 1kg', 'SE-MANZANA-001', 2400, 1, true, NOW(), NOW()),

-- Proveedor: Cereales Premium SpA (c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f)
('b5c5d5e5-0009-4009-a009-000000000001', 'c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 'Barra de Cereal', 'Barra de cereal 30g pack 5', 'CP-BARRA-001', 2500, 5, true, NOW(), NOW()),
('b5c5d5e5-0009-4009-a009-000000000002', 'c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche Entera', 'Leche entera fortificada 1L', 'CP-LECHE-001', 1400, 3, true, NOW(), NOW()),
('b5c5d5e5-0009-4009-a009-000000000003', 'c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt Natural', 'Yogurt natural con cereales 1L', 'CP-YOGURT-001', 2900, 3, true, NOW(), NOW()),
('b5c5d5e5-0009-4009-a009-000000000004', 'c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar Refinada', 'Azúcar refinada 1kg', 'CP-AZUCAR-001', 2700, 1, true, NOW(), NOW()),
('b5c5d5e5-0009-4009-a009-000000000005', 'c5d5e5f5-a5b5-4c5d-8e5f-0a5b2c3d4e5f', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana Roja', 'Manzana roja fresca 1kg', 'CP-MANZANA-001', 2600, 1, true, NOW(), NOW());


/*
INSERT INTO public.supplier_product_per_store VALUES ('3defed48-74f8-49ab-8e99-7f3e53aeb3b1', '913b50df-b1f9-42d8-912a-bdfdb22f2575', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 4, true, 'General', 5616, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('adda2715-c36a-48f0-ba5b-6903d554a366', 'a57ae7fc-a39b-4398-beb1-a466b19ade22', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 5, true, 'General', 1158, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('0faa4977-346a-4256-9712-ac2c2d9a71f4', '91ec47cd-4097-4a01-9e9f-b202d9f085e1', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 10, true, 'General', 5990, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('6389b73f-9db5-4897-9f1c-a8203c71b58d', 'a9d3fa5d-7501-491c-b41e-edbafa14de5f', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 9, true, 'General', 1252, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('aac79688-59f1-4817-9cd1-d57f163ec223', 'b5c3aaa8-b3b8-4d3b-b950-d299817dc708', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 5366, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('263d3b6d-3d46-4571-beca-44bd468d92eb', 'b92493b0-3c8e-491d-b832-81a8ebaf638c', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 9, true, 'General', 3039, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('4b638d6a-ad35-48b9-87fc-e36143967857', '4cfb1798-9712-4429-b712-4aeb55587276', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 4, true, 'General', 2779, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('d631fbf7-d56a-40c0-a040-9364d2baa7ac', '47b29ea7-8d34-4198-a866-cda0729accd1', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 6, true, 'General', 2513, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('3bfcf81f-d71e-4c9d-9854-a85539461008', 'f34e3cb6-268c-4dbf-9792-8a0a733a47e1', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 10, true, 'General', 3578, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('347f4919-1533-4dc3-a69b-6d4afe10dbe4', '4b18a78e-2b07-4ea3-b3d7-eda2d2eeb7e1', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 5128, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('674a3f0d-4c03-47ac-9301-66c6de94defa', '65969010-803b-4f7c-9e0a-bc5f9feae1a5', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 3, true, 'General', 5757, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('4a4a2055-fa91-470a-8f73-a76f6c1710a9', '5b39fada-3aed-42be-8725-078977dc6419', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 2, true, 'General', 5363, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('d5ca206c-87df-4dad-8d56-a91643592441', 'd7e663d6-20a3-4bb1-8040-451c1c9abe7a', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 1753, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('512190be-06c6-4194-8553-96ff5eb81747', 'ba7f4155-744b-4de9-baca-76a4bff3e092', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 4498, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('979074f1-b60a-40be-a15a-8c1a582dd04c', '9d697908-a6ba-44eb-925a-01412b7e4a73', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 6, true, 'General', 1857, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('749a3bf1-a6d2-40fe-a85d-8b9caaf0a57d', 'ed53ff8c-7e88-4147-9da5-e8bbde34f085', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 1, true, 'General', 2883, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('f64fdcc5-8d82-4f57-bfca-e331dd843b0f', 'c9ebeca8-840c-4fd8-81c2-ab689bb7a004', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 1392, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('2607a8d2-dbd8-415b-b17e-1734080ccc21', 'd6ad3811-8c92-4c07-8ce1-e0d83cb9e094', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 5433, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('548092cc-d6c6-4a6f-a494-f9663209c9f4', '23cd2701-6e91-4359-846d-14a43f35fbe7', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 2, true, 'General', 4785, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('84a2b858-54a3-432d-b41f-f3e82de15da3', '64d2fdd0-c41f-4497-91b0-a9a763d86434', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 6, true, 'General', 2309, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('e3cb732d-7926-4766-bf51-8bb756411751', 'fce49a3e-05b1-472d-9777-45b0c9c80c36', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 10, true, 'General', 3073, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('000ddc01-57c6-460b-9b75-515b75be9539', 'd7d0ee70-520b-478f-95a8-eb3d84d5cc6d', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 3150, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('a5256909-6da4-4c77-afd3-1eeb5cc08cee', '0575ca88-a233-421a-96c0-0635707f6ee8', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 2, true, 'General', 2273, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('a6971f5a-9c5b-423e-a77e-7660e0bd13e6', '93f65612-51a4-4a8c-96e0-caefc69c215b', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 3, true, 'General', 5730, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('f44a52c7-9c28-43e9-97c8-370eab83c5c7', '5e59ff66-8a0b-40d0-9afa-6f2d1a1ec026', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 1, true, 'General', 2102, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('7e1d689f-95cc-492d-ba39-a1d339ab70d1', 'ccc81d7e-c523-45db-a00b-dd6ede9f2ea1', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 10, true, 'General', 3502, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('9f00fd4e-5f60-4bc5-8099-6bf7eaa84926', '20a73156-e282-41d6-b379-a777965a81a4', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 2, true, 'General', 1885, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('b031e634-e0fa-4502-ad13-3fa68f240c5f', 'ef96851b-dbfc-4f36-a15d-fb3f13ad6479', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 4415, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('8ca2c453-d39f-4058-9c0c-93449ee836a4', 'dcea4308-ea2b-4850-b0f3-c5bdeba7621f', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 3151, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('aa66d0eb-80ad-42ee-a110-a815756cbb5b', '3c67c2df-a9f1-4431-990d-20ca89c1cb0c', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 10, true, 'General', 3138, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('2ae4e7e7-a70d-41d1-a25d-56f5d8172666', '010a55c4-3750-44d4-8013-2a25623b6aa6', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 5, true, 'General', 4385, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('3b042199-4564-4c17-9c9a-9839b394c7c4', '69892e6d-8dba-4e91-8445-8d6e9140be63', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 1, true, 'General', 3835, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('8b177328-d44c-4d8d-89e1-52a3998c27de', '24411dc0-8d55-4ea6-ae84-9b3ccafd5c89', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 4207, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('2c4fbeb6-3a7f-41fb-b1c2-73900076e0ee', 'c1612f43-0e64-4682-a6e1-86901d61438b', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 10, true, 'General', 1107, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('6154512b-0ff7-4cbd-85c2-92a3ba4eeb2b', '67df6da6-b89d-4894-be98-4ab87c675234', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 9, true, 'General', 1669, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('47a04fbd-5c39-4362-8997-1d7770d6fae2', '2e2b04d6-4783-4d07-a352-5cebb97d3b10', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 9, true, 'General', 1219, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('f8daed53-be96-4daa-bef3-284a91862d76', '883656c2-d59a-4c53-bdd9-57e22bc311dd', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 9, true, 'General', 4459, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('50140fa9-df4b-43b8-8c8e-d907104b0040', '7304efbe-22a3-4cb4-bb1d-5f8bb829a426', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 3, true, 'General', 2072, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('f256d948-8343-41ea-b9af-add7318cbb06', '90981b11-75a0-418f-aa9b-81727959a81e', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 6, true, 'General', 3105, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('8f0ce229-14ab-4138-ba9a-283930209824', 'f0b32070-ac99-4ed8-83fd-66517f1c0ad3', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 2, true, 'General', 2327, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('0647a24f-4bba-42de-8cc8-667e99c68c61', 'af694a1b-3ff8-472e-9783-82c3f3474c2e', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 1449, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('558757ce-e430-44f3-b1b3-d5d38f6ff015', '4c7ff00c-23e0-43ce-8fcb-002daa293d98', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 2, true, 'General', 3170, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('1905726f-b4fd-4f99-9536-a1516ef82640', '2617ebfa-5215-409d-a294-97d0fca14f28', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 4, true, 'General', 4943, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('c81e4f26-8fb0-4689-b547-a3472531d899', 'd3a7f9dc-1735-4476-a13b-31232842695f', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 5, true, 'General', 4187, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('5c8db4aa-9954-4b6f-891e-122d5a6d1238', '181db0c8-6b0f-4c79-9409-8727f1319a0f', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 1, true, 'General', 4935, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('ed1b8a61-54fb-4ad7-879c-d30f1b74ed87', '844c61a2-27aa-48df-a676-a1283bd92340', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 1015, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('68fa5e91-26ea-49ca-8cb6-9be354a03399', 'a1143ac7-69c9-4431-aa58-6aded06e8703', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 4, true, 'General', 2130, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('71d48bce-e3b0-4e8a-ac6f-a24c3158a95d', '87299350-6cf0-432a-912a-5380de9b49c3', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 4, true, 'General', 5761, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('90749df2-9e96-4f3d-a77b-a2c48296fc33', '2a753ddd-4135-4cc2-bcc9-1f2ea361863d', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 5, true, 'General', 3430, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('b40b4813-38a0-4e00-8442-20caca3c8c4e', 'dc12d07a-6cb9-420c-b8dd-5d57c3932b33', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 2516, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('491299f8-b1b0-459e-b147-74cf965887ea', '1d9e6e97-e6c5-4cee-8b42-3ee87adfbbf7', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 8, true, 'General', 5480, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('2f1ba770-a4eb-4df8-8ca4-236255f12b49', 'faf42b6b-5d60-4016-bd5c-a26d2e55960a', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 3808, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('90ef6a96-f39b-4510-a96b-d27c2cb97bf9', '3494dd6c-8158-4163-9c01-dc80299d8840', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 5, true, 'General', 4310, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('6791ee67-23ee-4b7e-834d-c14588c09d1c', '738b1e49-10bc-4cab-9a24-fd5b8ed4e516', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 2993, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
INSERT INTO public.supplier_product_per_store VALUES ('86ce1720-d6b1-41c8-857b-2d2d10d9f9a8', 'c4a1a9c4-fdec-42bd-b4a2-846ffc44dfd8', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 7, true, 'General', 4628, 1, '30 días', 5, true, '2025-09-08 15:32:21.860113', '2025-09-08 15:32:21.860113');
*/
INSERT INTO public.user_accounts VALUES ('5b877be1-890c-4aed-a33c-38e9607d4a83', 'gerente', 'gerente@dotsolutions.cl', 'Gerente del sistema', 'gerente_password', true, true, NULL, '2025-09-05 21:54:06.930692', '2025-09-05 21:54:06.930692');
INSERT INTO public.user_accounts VALUES ('d3b5030e-8789-4e7a-8fc7-de5936a91dfd', 'bodeguero', 'bodeguero@dotsolutions.cl', 'Bodeguero del sistema', 'bodeguero_password', true, true, NULL, '2025-09-05 21:54:06.931133', '2025-09-05 21:54:06.931133');
INSERT INTO public.user_accounts VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', 'admin', 'admin@dotsolutions.cl', 'Administrador del sistema', 'a015ed36bb30b44489c18731ca4b4b5852e1b4573265e8fab92c30c935a2952a', true, false, NULL, '2025-09-05 21:54:06.929357', '2025-09-08 14:57:02.066274');

INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', '8f20d7a3-5160-4124-a743-a354b39e6507');
INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', '90cfcfa1-090e-426e-aefc-0353dc481489');
INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', '50cc64e1-62bc-4f58-b626-6c217d74d490');
INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', '7d385e85-7032-4fee-8dcc-b9d49dc41c70');
INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', 'e481d36c-ab9e-4dac-8636-a8a278b5cd0e');
INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', 'e3a7b420-0c17-470e-8769-6d99ef0b2a3b');
INSERT INTO public.user_account_per_profiles VALUES ('6e161905-a38a-46a3-b7a7-8429c0069b82', 'c5045b0e-b49b-43fd-a733-3de160fabe48');

INSERT INTO public.user_account_per_profiles VALUES ('5b877be1-890c-4aed-a33c-38e9607d4a83', '50cc64e1-62bc-4f58-b626-6c217d74d490');
INSERT INTO public.user_account_per_profiles VALUES ('5b877be1-890c-4aed-a33c-38e9607d4a83', '9fc2d9e2-c47a-412a-a574-506d8c03709f');
INSERT INTO public.user_account_per_profiles VALUES ('5b877be1-890c-4aed-a33c-38e9607d4a83', 'e3a7b420-0c17-470e-8769-6d99ef0b2a3b');

INSERT INTO public.user_account_per_profiles VALUES ('d3b5030e-8789-4e7a-8fc7-de5936a91dfd', '50cc64e1-62bc-4f58-b626-6c217d74d490');
INSERT INTO public.user_account_per_profiles VALUES ('d3b5030e-8789-4e7a-8fc7-de5936a91dfd', '24ce1667-f111-4cfd-a7e5-42fe622c3bc2');
INSERT INTO public.user_account_per_profiles VALUES ('d3b5030e-8789-4e7a-8fc7-de5936a91dfd', 'e3a7b420-0c17-470e-8769-6d99ef0b2a3b');

-- Ejemplo: Sucursal Santiago
INSERT INTO public.warehouse (
	id, store_id, description, warehouse_name, warehouse_address, warehouse_phone, created_at
) VALUES (
	'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'Bodega principal de la sucursal Santiago', 'Bodega Central Santiago', 'Av. Libertador Bernardo O’Higgins 1234', '+56 2 2345 6789','2025-09-05 21:54:06.914962'
);
-- Bodega de traspaso: Sucursal Santiago
INSERT INTO public.warehouse (
	id, store_id, warehouse_name, is_momevent_warehouse, created_at
) VALUES (
	gen_random_uuid(), 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'Bodega de traspaso - Sucursal Santiago', true, '2025-09-05 21:54:06.914962'
);
-- Ejemplo: Oficina Central
INSERT INTO public.warehouse (
	id, store_id, description, warehouse_name, warehouse_address, warehouse_phone, created_at
) VALUES (
	'b1e2c3d4-5678-4abc-9def-1234567890ab', '9f666bab-35e3-46e6-b147-75354262dc84', 'Bodega de insumos de oficina central', 'Bodega Oficina Central', 'Calle Nueva Providencia 456', '+56 2 8765 4321', '2025-09-05 21:54:06.914962'
);

-- Bodega de traspaso: Oficina Central
INSERT INTO public.warehouse (
	id, store_id, warehouse_name, is_momevent_warehouse, created_at
) VALUES (
	gen_random_uuid(), '9f666bab-35e3-46e6-b147-75354262dc84', 'Bodega de traspaso - Oficina Central', true, '2025-09-05 21:54:06.914962'
);

-- Ejemplo: Bodega Central
INSERT INTO public.warehouse (
	id, store_id, description, warehouse_name, warehouse_address, warehouse_phone, created_at
) VALUES (
	'c2e3f4a5-6789-4bcd-8efa-2345678901bc', '86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'Bodega central de distribución', 'Bodega Central', 'Av. Central 789', '+56 2 3456 7890', '2025-09-05 21:54:06.914962'
);

-- Bodega de traspaso: Bodega Central
INSERT INTO public.warehouse (
	id, store_id, warehouse_name, is_momevent_warehouse, created_at
) VALUES (
	gen_random_uuid(), '86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'Bodega de traspaso - Bodega Central', true, '2025-09-05 21:54:06.914962'
);

-- Ejemplo: Centro de Desarrollo
INSERT INTO public.warehouse (
	id, store_id, description, warehouse_name, warehouse_address, warehouse_phone, created_at
) VALUES (
	'd3f4a5b6-7890-4cde-9fab-3456789012cd', '90775541-4b56-4327-84b7-0927f46122d7', 'Bodega de desarrollo y pruebas', 'Bodega Centro de Desarrollo', 'Av. Innovación 321', '+56 2 4567 8901', '2025-09-05 21:54:06.914962'
);

-- Bodega de traspaso: Centro de Desarrollo
INSERT INTO public.warehouse (
	id, store_id, warehouse_name, is_momevent_warehouse, created_at
) VALUES (
	gen_random_uuid(), '90775541-4b56-4327-84b7-0927f46122d7', 'Bodega de traspaso - Centro de Desarrollo', true, '2025-09-05 21:54:06.914962'
);

-- Bodega 2 asociada a Sucursal Santiago
INSERT INTO public.warehouse (
    id, store_id, description, warehouse_name, warehouse_address, warehouse_phone, created_at
) VALUES (
    '22222222-2222-2222-2222-222222222222', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'Bodega secundaria de Sucursal Santiago', 'Bodega 2', 'Av. Libertador Bernardo O’Higgins 1234', '+56 2 2345 6790','2025-09-05 21:54:06.914962'
);

-- Bodega 3 asociada a Sucursal Santiago
INSERT INTO public.warehouse (
    id, store_id, description, warehouse_name, warehouse_address, warehouse_phone, created_at
) VALUES (
    '33333333-3333-3333-3333-333333333333', 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'Bodega secundaria de Sucursal Santiago', 'Bodega 3', 'Av. Libertador Bernardo O’Higgins 1234', '+56 2 8765 4322', '2025-09-05 21:54:06.914962'
);


-- ===== SUCURSAL SANTIAGO (Distribuidora Los Andes) =====
-- Productos enfocados en consumo general: lácteos, carnes, verduras y despensa
INSERT INTO public.product_per_store (store_id, product_id, product_name, item_sale, use_recipe, unit_inventory_id, description, minimal_stock, maximal_stock, max_quantity, created_at, updated_at) VALUES
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'e8aac862-1028-4353-a3fd-928ba671ddb9', 'Leche entera 1L', true, false, 3, 'Leche entera fresca 1 litro', 15, 50, 20, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'c2d0f299-f53d-4e45-99cb-4ff3c946417f', 'Yogurt natural 500ml', true, false, 3, 'Yogurt natural sin azúcar 500ml', 12, 40, 15, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'db682432-362e-4e7b-af54-0d87f91832da', 'Pechuga de pollo 1kg', true, false, 1, 'Pechuga de pollo fresca 1 kilogramo', 5, 25, 10, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '60da2f24-c4d9-407a-99e6-cd10d2e5cbb5', 'Carne molida 1kg', true, false, 1, 'Carne molida de vacuno 1 kilogramo', 4, 20, 8, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '7ddda93a-fe87-4385-9299-232ec0f5f7ee', 'Tomate fresco 1kg', true, false, 1, 'Tomate fresco tipo pera 1 kilogramo', 10, 35, 20, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'a4037a6c-fd92-49a4-aa32-80fdd0b1c2c3', 'Lechuga fresca', true, false, 5, 'Lechuga fresca verde variedad mantequilla', 8, 30, 15, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '99de8b79-a9ff-41ec-9791-61117d71900a', 'Arroz grano largo 1kg', true, false, 1, 'Arroz grano largo integral 1 kilogramo', 10, 45, 25, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', '15bf5684-3363-499d-ab0c-fdfbc5d84c6b', 'Aceite vegetal 1L', true, false, 3, 'Aceite vegetal refinado 1 litro', 8, 30, 12, NOW(), NOW()),
('ce79763f-175b-44ee-8cdd-47fc95d4a8ce', 'faffb338-771d-4bd3-871d-4f3e1c37268d', 'Pan marraqueta', true, false, 5, 'Pan marraqueta fresco recién horneado', 20, 80, 40, NOW(), NOW());

-- ===== BODEGA CENTRAL (Distribuidora Los Andes) =====
-- Productos enfocados en distribución mayorista: carnes, verduras, frutas, despensa y panadería
INSERT INTO public.product_per_store (store_id, product_id, product_name, item_sale, use_recipe, unit_inventory_id, description, minimal_stock, maximal_stock, max_quantity, created_at, updated_at) VALUES
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'e2432d39-68ca-4a02-81d2-e200e47ea620', 'Filete de vacuno 1kg', true, false, 1, 'Filete de vacuno premium 1 kilogramo', 3, 20, 8, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', '95523bb0-c72f-4a4c-affe-f14c9a32870e', 'Chuleta de cerdo 1kg', true, false, 1, 'Chuleta de cerdo fresca 1 kilogramo', 4, 22, 10, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'f0dddfbb-a31c-4247-ba0b-d37727672f51', 'Zanahoria 1kg', true, false, 1, 'Zanahoria fresca y limpia 1 kilogramo', 15, 50, 30, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', '19d5a6e2-2992-47e9-80bd-fc13e7a6083c', 'Papa amarilla 1kg', true, false, 1, 'Papa amarilla variedad tiquichoca 1 kg', 20, 60, 40, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'f0993ce1-93f6-4fe8-bf4b-f6265b65260c', 'Manzana roja 1kg', true, false, 1, 'Manzana roja fuerte y dulce 1 kilogramo', 12, 45, 25, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'f7500114-702b-45bd-bef7-64e6d6d9740c', 'Harina de trigo 1kg', true, false, 1, 'Harina de trigo sin preparar 1 kilogramo', 15, 50, 30, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'a55fb2a9-8ba6-476d-b814-8392572ae99d', 'Azúcar refinada 1kg', true, false, 1, 'Azúcar blanca refinada 1 kilogramo', 12, 40, 20, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', '397d445c-3b56-429a-9f52-7835fb340bb6', 'Pan de molde 570g', true, false, 5, 'Pan de molde integral fresco 570 gramos', 25, 100, 50, NOW(), NOW()),
('86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6', 'e30b3445-fe86-4914-a2d4-97f10e798922', 'Croissant 6 unidades', true, false, 5, 'Croissant francés mantequilla caja x6', 15, 60, 30, NOW(), NOW());

-- ===== OFICINA CENTRAL (Soluciones Digitales) =====
-- Productos enfocados en oficina y break room: bebidas, snacks, lácteos y despensa
INSERT INTO public.product_per_store (store_id, product_id, product_name, item_sale, use_recipe, unit_inventory_id, description, minimal_stock, maximal_stock, max_quantity, created_at, updated_at) VALUES
('9f666bab-35e3-46e6-b147-75354262dc84', '2099754d-54c7-4b5d-93d8-3990c687981e', 'Coca-Cola 1.5L', true, false, 3, 'Bebida gaseosa Coca-Cola 1.5 litros', 25, 100, 40, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', '087eae2b-ec75-4ab5-ade1-dcba24f0c5cf', 'Agua mineral 500ml', true, false, 3, 'Agua mineral pura 500 mililitros pack 24', 30, 120, 50, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', '11dc275a-bb9c-4591-98e6-f3e1fc026a61', 'Té helado 1.5L', true, false, 3, 'Bebida de té helado natural 1.5 litros', 20, 80, 30, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', '1dfbc4d7-c67c-42b6-9674-7b2bb7554fbe', 'Papas fritas clásicas 170g', true, false, 5, 'Papas fritas clásicas saladas bolsa 170g', 20, 80, 40, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', 'a8b94ec1-ce2b-4a8b-9ece-9dde391251a1', 'Galletas de chocolate 200g', true, false, 5, 'Galletas de chocolate premium 200 gramos', 15, 60, 30, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', 'c4d42258-52e4-4735-bb02-e163145577ae', 'Crema de leche 200ml', true, false, 3, 'Crema de leche fresca 200 mililitros', 8, 30, 12, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', '912c9f07-5810-438f-ad4c-11790d181910', 'Fideos espagueti 500g', true, false, 1, 'Fideos espagueti tipo 00 marca premium 500g', 10, 40, 20, NOW(), NOW()),
('9f666bab-35e3-46e6-b147-75354262dc84', '92599d39-eddb-4e9d-9012-b209675ff9bd', 'Chocolate en barra 100g', true, false, 5, 'Chocolate amargo 70% cacao 100 gramos', 12, 50, 25, NOW(), NOW());

-- ===== CENTRO DE DESARROLLO (Soluciones Digitales) =====
-- Productos enfocados en equipo desarrollo: bebidas, snacks, café/cereales, panadería
INSERT INTO public.product_per_store (store_id, product_id, product_name, item_sale, use_recipe, unit_inventory_id, description, minimal_stock, maximal_stock, max_quantity, created_at, updated_at) VALUES
('90775541-4b56-4327-84b7-0927f46122d7', '67da0a90-ffe5-42bc-88c5-058fca657658', 'Jugo de naranja 1L', true, false, 3, 'Jugo de naranja natural 100% 1 litro', 20, 80, 35, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', '946d986b-33c5-43af-b6d7-564f461a5b0b', 'Bebida energética 250ml', true, false, 3, 'Bebida energética premium lata 250ml pack 6', 15, 60, 25, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', '0f5636d1-bef9-48cb-beb8-5c7ccedbffc6', 'Maní salado 100g', true, false, 5, 'Maní tostado y salado 100 gramos', 18, 70, 35, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', 'e0ec1d66-c844-469c-9bc6-49dfa31ae13f', 'Barra de cereal 30g', true, false, 5, 'Barra de cereal y frutos secos 30g pack 5', 20, 80, 40, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', 'a350d0f1-7720-48fe-8cf8-7cba62879f8e', 'Queso mantecoso 400g', true, false, 1, 'Queso mantecoso fresco y cremoso 400g', 6, 25, 10, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', '800596d0-5e8f-49aa-876a-f45f9418f74d', 'Mantequilla sin sal 200g', true, false, 5, 'Mantequilla sin sal premium 200 gramos', 8, 30, 15, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', 'faffb338-771d-4bd3-871d-4f3e1c37268d', 'Pan marraqueta fresco', true, false, 5, 'Pan marraqueta recién horneado cada mañana', 25, 100, 50, NOW(), NOW()),
('90775541-4b56-4327-84b7-0927f46122d7', '69f51fca-89eb-49db-a3c0-84fb149023a9', 'Sandía 4-5kg', true, false, 5, 'Sandía fresca roja jugosa 4-5 kilogramos', 3, 15, 8, NOW(), NOW());






/*
INSERT INTO public.warehouse_per_product VALUES ('1b811743-ecd7-4d9f-a3e0-28f016a94d0d', '26d261de-5b9d-44fa-bdf6-c5fbf0f15ad3', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 0, 0, 20);
INSERT INTO public.warehouse_per_product VALUES ('42e377d3-95de-4607-aefa-b09aaf82b78e', '7e3fafcd-0197-4f62-84b4-ed5c753921ac', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 1, 0, 20);
INSERT INTO public.warehouse_per_product VALUES ('8b0dc02b-3c09-4f2c-9e50-57abf3f7a92f', '45a120db-2844-4fe0-8541-49906750509c', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 3, 0, 20);
INSERT INTO public.warehouse_per_product VALUES ('c1a0ad9a-6f69-4c30-bd1e-9b7b4f8b9f7a', '81865ab1-c78b-4db4-9f92-2755eb102b9b', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 5, 0, 20);
INSERT INTO public.warehouse_per_product VALUES ('7a287f0d-8a6c-4e08-b87d-5e32cc9b460f', '35bf2c4c-af8a-469a-bf71-1b1b340b2322', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 8, 0, 20);

INSERT INTO public.warehouse_per_product VALUES
('8d775f02-a03d-4eab-b17f-5043e990e969', '52bbed11-7f3d-42c1-82ca-604b3d74b4b4', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 13, 0, 20),
('79924c68-2ccf-4024-9111-32935b9d8080', '17f8cc95-ea0b-4313-a370-5a55b8e7b76b', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 21, 0, 20),
('ab3b4d32-60a8-49e6-9106-f9310dab5d0f', 'b37d59ec-a709-4435-b854-0e7e496320ee', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 34, 0, 20),
('8ae15787-20c0-4034-8062-7941c53929e2', 'd96e872a-8cad-420a-a229-480ed233ad46', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 55, 0, 20),
('8229ce51-0e00-43f6-9974-0ae25581df8d', '6c88f769-511f-4a89-bc8c-79affdcdbaa5', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 34, 0, 20),
('59c6b364-19e4-4b0d-a30b-54c21e7a4daf', '8780d84e-693d-4645-8286-17d782ba6f24', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 21, 0, 20),
('3b4e6ece-dd8a-43e1-bbdf-4e6107235171', 'd0220269-aad5-4764-a253-c372d2dced23', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 13, 0, 20),
('8b15aaf1-a2c6-4bbb-80a8-9462e33378f9', '2498ad2f-d703-4aef-88ef-e13b2ccae704', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 8, 0, 20),
('01be1eab-440e-41c7-821c-26cd3582163a', 'aadba3e3-b6c2-4df6-903d-9c26e6fcfe84', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 5, 0, 20),
('28c4cc7d-68b6-466c-a1bf-7c3eba41697d', '6b71a9ad-1b51-4c3a-9093-aff78ab0f052', 'aac08583-b8f7-4a09-b710-7503a7d6f1e3', 3, 0, 20);
*/
-- =======================
-- Poblar warehouse_per_product (stock inicial)
-- =======================
-- Nota: Solo bodegas principales (is_momevent_warehouse = false)
--       No se generan filas para bodegas de traspaso

-- (Nota) No ajustamos cost_avg del producto; se mantiene por compatibilidad.

-- ===== SUCURSAL SANTIAGO (Distribuidora Los Andes) =====
-- Relaciona productos de la tienda con su bodega principal y genera stock aleatorio
INSERT INTO public.warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
SELECT 
	pps.id AS store_product_id,
	w.id AS warehouse_id,
	GREATEST(0, ROUND(pps.minimal_stock + random() * (pps.maximal_stock - pps.minimal_stock))) AS in_stock,
	ROUND((p.cost_estimated * (0.95 + random() * 0.10))::numeric, 2) AS cost_avg
FROM public.product_per_store pps
JOIN public.product p ON p.id = pps.product_id
JOIN public.warehouse w ON w.store_id = pps.store_id AND w.is_momevent_warehouse = false
LEFT JOIN public.warehouse_per_product wpp ON wpp.store_product_id = pps.id AND wpp.warehouse_id = w.id
WHERE pps.store_id = 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce'
	AND wpp.store_product_id IS NULL;

-- ===== BODEGA CENTRAL (Distribuidora Los Andes) =====
INSERT INTO public.warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
SELECT 
	pps.id AS store_product_id,
	w.id AS warehouse_id,
	GREATEST(0, ROUND(pps.minimal_stock + random() * (pps.maximal_stock - pps.minimal_stock))) AS in_stock,
	ROUND((p.cost_estimated * (0.95 + random() * 0.10))::numeric, 2) AS cost_avg
FROM public.product_per_store pps
JOIN public.product p ON p.id = pps.product_id
JOIN public.warehouse w ON w.store_id = pps.store_id AND w.is_momevent_warehouse = false
LEFT JOIN public.warehouse_per_product wpp ON wpp.store_product_id = pps.id AND wpp.warehouse_id = w.id
WHERE pps.store_id = '86a0f5ae-94c5-4fb3-80d9-b7da3a2a0fe6'
	AND wpp.store_product_id IS NULL;

-- ===== OFICINA CENTRAL (Soluciones Digitales) =====
INSERT INTO public.warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
SELECT 
	pps.id AS store_product_id,
	w.id AS warehouse_id,
	GREATEST(0, ROUND(pps.minimal_stock + random() * (pps.maximal_stock - pps.minimal_stock))) AS in_stock,
	ROUND((p.cost_estimated * (0.95 + random() * 0.10))::numeric, 2) AS cost_avg
FROM public.product_per_store pps
JOIN public.product p ON p.id = pps.product_id
JOIN public.warehouse w ON w.store_id = pps.store_id AND w.is_momevent_warehouse = false
LEFT JOIN public.warehouse_per_product wpp ON wpp.store_product_id = pps.id AND wpp.warehouse_id = w.id
WHERE pps.store_id = '9f666bab-35e3-46e6-b147-75354262dc84'
	AND wpp.store_product_id IS NULL;

-- ===== CENTRO DE DESARROLLO (Soluciones Digitales) =====
INSERT INTO public.warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
SELECT 
	pps.id AS store_product_id,
	w.id AS warehouse_id,
	GREATEST(0, ROUND(pps.minimal_stock + random() * (pps.maximal_stock - pps.minimal_stock))) AS in_stock,
	ROUND((p.cost_estimated * (0.95 + random() * 0.10))::numeric, 2) AS cost_avg
FROM public.product_per_store pps
JOIN public.product p ON p.id = pps.product_id
JOIN public.warehouse w ON w.store_id = pps.store_id AND w.is_momevent_warehouse = false
LEFT JOIN public.warehouse_per_product wpp ON wpp.store_product_id = pps.id AND wpp.warehouse_id = w.id
WHERE pps.store_id = '90775541-4b56-4327-84b7-0927f46122d7'
	AND wpp.store_product_id IS NULL;


-- ===== BODEGA 2 (Sucursal Santiago) =====
INSERT INTO public.warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
SELECT 
    pps.id AS store_product_id,
    '22222222-2222-2222-2222-222222222222' AS warehouse_id,
    GREATEST(0, ROUND(pps.minimal_stock + random() * (pps.maximal_stock - pps.minimal_stock))) AS in_stock,
    ROUND((p.cost_estimated * (0.95 + random() * 0.10))::numeric, 2) AS cost_avg
FROM public.product_per_store pps
JOIN public.product p ON p.id = pps.product_id
WHERE pps.store_id = 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce'
  AND p.sku LIKE 'BOD2-%';

-- ===== BODEGA 3 (Sucursal Santiago) =====
INSERT INTO public.warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
SELECT 
    pps.id AS store_product_id,
    '33333333-3333-3333-3333-333333333333' AS warehouse_id,
    GREATEST(0, ROUND(pps.minimal_stock + random() * (pps.maximal_stock - pps.minimal_stock))) AS in_stock,
    ROUND((p.cost_estimated * (0.95 + random() * 0.10))::numeric, 2) AS cost_avg
FROM public.product_per_store pps
JOIN public.product p ON p.id = pps.product_id
WHERE pps.store_id = 'ce79763f-175b-44ee-8cdd-47fc95d4a8ce'
  AND p.sku LIKE 'BOD3-%';

-- ===== ESTADOS DE SOLICITUDES (request_status) =====
INSERT INTO public.request_status (name) VALUES 
    ('Creada'),
    ('Con conflicto'),
    ('Aprobada'),
    ('Cancelada'),
    ('Finalizada');