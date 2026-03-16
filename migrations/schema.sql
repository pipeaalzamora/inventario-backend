CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE user_accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_name character varying NOT NULL,
    user_email character varying NOT NULL UNIQUE,
    description character varying,
    user_password character varying NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    is_new_account BOOLEAN NOT NULL DEFAULT TRUE,
    deleted_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_user UUID,
    to_user UUID NOT NULL,
    read_at TIMESTAMPTZ NULL,
    send_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    payload JSONB NOT NULL,
    notification_type character varying NOT NULL,
    FOREIGN KEY (from_user) REFERENCES user_accounts(id) ON DELETE SET NULL,
    FOREIGN KEY (to_user) REFERENCES user_accounts(id) ON DELETE CASCADE
);

CREATE TABLE profile_accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    profile_name character varying NOT NULL,
    description character varying
);

CREATE TABLE user_account_per_profiles (
    user_account_id UUID NOT NULL,
    profile_account_id UUID NOT NULL,
    PRIMARY KEY (user_account_id, profile_account_id),
    FOREIGN KEY (user_account_id) REFERENCES user_accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (profile_account_id) REFERENCES profile_accounts(id) ON DELETE CASCADE
);

CREATE TABLE power_account_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_name character varying NOT NULL UNIQUE,
    description character varying,
    ownable BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE power_accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    power_name character varying NOT NULL UNIQUE,
    power_display character varying NOT NULL,
    description character varying,
    power_account_category_id UUID NOT NULL,
    FOREIGN KEY (power_account_category_id) REFERENCES power_account_categories(id) ON DELETE CASCADE
);

CREATE TABLE profile_account_per_power_accounts (
    profile_account_id UUID NOT NULL,
    power_account_id UUID NOT NULL,
    PRIMARY KEY (profile_account_id, power_account_id),
    FOREIGN KEY (profile_account_id) REFERENCES profile_accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (power_account_id) REFERENCES power_accounts(id) ON DELETE CASCADE
);

CREATE TABLE country (
  id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  iso_code character varying NOT NULL UNIQUE,
  country_name character varying NOT NULL
);

CREATE TABLE fiscal_data (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  id_fiscal character varying NOT NULL UNIQUE,
  raw_fiscal_id character varying NOT NULL,
  fiscal_name character varying NOT NULL,
  fiscal_address character varying NOT NULL,
  fiscal_state character varying NOT NULL,
  fiscal_city character varying NOT NULL,
  email character varying NOT NULL
);

CREATE TABLE economic_activity_class (
  id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  country_id integer NOT NULL,
  system_name character varying NOT NULL,
  system_code character varying NOT NULL UNIQUE,
  FOREIGN KEY (country_id) REFERENCES country(id) ON DELETE CASCADE
);

CREATE TABLE economic_activity (
  id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  economic_activity_class_id integer NOT NULL,
  activity_name character varying NOT NULL,
  activity_code character varying NOT NULL UNIQUE,
  description character varying NOT NULL,
  FOREIGN KEY (economic_activity_class_id) REFERENCES economic_activity_class(id) ON DELETE CASCADE
);

CREATE TABLE fiscal_per_activity (
  id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  fiscal_data_id UUID NOT NULL,
  economic_activity_id integer NOT NULL,
  FOREIGN KEY (fiscal_data_id) REFERENCES fiscal_data(id) ON DELETE CASCADE,
  FOREIGN KEY (economic_activity_id) REFERENCES economic_activity(id) ON DELETE CASCADE
);

CREATE TABLE company (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    country_id integer NOT NULL,
    fiscal_data_id UUID NOT NULL,
    company_name character varying NOT NULL,
    description character varying,
    image_logo character varying,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (country_id) REFERENCES country(id) ON DELETE RESTRICT, -- Block deletion
    FOREIGN KEY (fiscal_data_id) REFERENCES fiscal_data(id) ON DELETE RESTRICT -- Block deletion
);

CREATE TABLE store (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID NOT NULL,
    store_name character varying NOT NULL,
    store_address character varying NOT NULL,
    description character varying NOT NULL,
    id_cost_center character varying,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE RESTRICT -- Block deletion
);

CREATE TABLE supplier (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fiscal_data_id UUID NOT NULL,
    country_id integer NOT NULL,
    supplier_name character varying NOT NULL,
    description character varying,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (fiscal_data_id) REFERENCES fiscal_data(id) ON DELETE RESTRICT,
    FOREIGN KEY (country_id) REFERENCES country(id) ON DELETE RESTRICT
);

CREATE TABLE supplier_per_company (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    supplier_id UUID NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE,
    UNIQUE (supplier_id, company_id)
);

-- Migración de datos legacy: si existe supplier_per_store se traspasan asignaciones a nivel compañía
DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = 'public'
          AND table_name = 'supplier_per_store'
    ) THEN
        INSERT INTO supplier_per_company (supplier_id, company_id)
        SELECT DISTINCT sps.supplier_id, st.company_id
        FROM supplier_per_store sps
        INNER JOIN store st ON st.id = sps.store_id
        ON CONFLICT DO NOTHING;

        DROP TABLE supplier_per_store;
    END IF;
END $$;

CREATE TABLE supplier_contact (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    supplier_id UUID NOT NULL,
    contact_name character varying NOT NULL,
    description character varying NOT NULL,
    email character varying NOT NULL,
    phone character varying NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON DELETE CASCADE
);

------------- Unidades de medida --------------
CREATE TABLE measurement_unit
(
    id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    abbreviation character varying NOT NULL,
    unit_name character varying NOT NULL,
    description character varying,
    basic_unit boolean NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id),
    UNIQUE (unit_name),
    UNIQUE (abbreviation)
);

CREATE TABLE tag
(
    id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    tag_name character varying NOT NULL,
    description character varying,
    PRIMARY KEY (id),
    UNIQUE (tag_name)
);

CREATE TABLE unit_conversion
(
    id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    from_unit_id INTEGER NOT NULL,
    to_unit_id INTEGER NOT NULL,
    conversion_factor NUMERIC NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (from_unit_id) REFERENCES measurement_unit(id) ON DELETE CASCADE,
    FOREIGN KEY (to_unit_id) REFERENCES measurement_unit(id) ON DELETE CASCADE,
    UNIQUE (from_unit_id, to_unit_id)
);

------------- Productos Base --------------
CREATE TABLE product
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    product_name character varying NOT NULL,
    sku character varying NOT NULL,
    description character varying,
    image character varying,
    cost_estimated float NOT NULL DEFAULT 0,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    UNIQUE (sku)
);

CREATE TABLE product_category
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    category_name character varying NOT NULL,
    description character varying,
    available boolean NOT NULL DEFAULT true,
    PRIMARY KEY (id),
    UNIQUE (category_name)
);

CREATE TABLE product_per_category
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
    product_id uuid NOT NULL,
    category_id integer NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (product_id)
        REFERENCES product (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    FOREIGN KEY (category_id)
        REFERENCES product_category (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);

CREATE TABLE code_kind
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    code_name character varying NOT NULL,
    description character varying,
    PRIMARY KEY (id),
    UNIQUE (code_name)
);

CREATE TABLE product_code
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    product_id uuid NOT NULL,
    kind_id integer NOT NULL,
    code_value character varying NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (product_id)
        REFERENCES product (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    FOREIGN KEY (kind_id)
        REFERENCES code_kind (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);

CREATE INDEX idx_product_code_kind_value ON product_code(kind_id, code_value);

------------- Producto por Tienda (NUEVA ESTRUCTURA) --------------
-- Reemplaza product_company y request_restriction
-- Cada tienda tiene su propia configuración de productos con costos independientes
CREATE TABLE IF NOT EXISTS product_per_store
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    store_id uuid NOT NULL,
    product_id uuid NOT NULL,
    tag_id integer NOT NULL DEFAULT 1,
    product_name character varying NOT NULL,
    -- DEPRECATED: item_purchase boolean NOT NULL,
    item_sale boolean NOT NULL, 
    -- DEPRECATED: item_inventory boolean NOT NULL,
    -- DEPRECATED: is_frozen boolean NOT NULL,
    use_recipe boolean NOT NULL,
    -- DEPRECATED: unit_purchase_id INTEGER NOT NULL,
    unit_inventory_id INTEGER NOT NULL,
    -- unit_matrix JSONB DEFAULT '[]'::jsonb,
    -- DEPRECATED: cost_last float NOT NULL DEFAULT 0,
    description character varying NOT NULL DEFAULT '',
    -- DEPRECATED: cost_estimated float NOT NULL DEFAULT 0,
    cost_avg float NOT NULL DEFAULT 0,
    minimal_stock float NOT NULL DEFAULT 0,
    maximal_stock float NOT NULL DEFAULT 0,
    -- DEPRECATED: minimal_order float NOT NULL DEFAULT 0,
    max_quantity float, -- Campo migrado de request_restriction (nullable)
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE RESTRICT,
    -- DEPRECATED: FOREIGN KEY (unit_purchase_id) REFERENCES measurement_unit(id),
    FOREIGN KEY (unit_inventory_id) REFERENCES measurement_unit(id),
    UNIQUE (store_id, product_id)
);

CREATE TABLE IF NOT EXISTS product_per_store_per_measurement_unit
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    store_product_id uuid NOT NULL,
    measurement_unit_id INTEGER NOT NULL,
    conversion_factor NUMERIC NOT NULL,
    FOREIGN KEY (store_product_id) REFERENCES product_per_store(id) ON DELETE CASCADE,
    FOREIGN KEY (measurement_unit_id) REFERENCES measurement_unit(id) ON DELETE CASCADE,
    UNIQUE (store_product_id, measurement_unit_id)
);

CREATE TABLE IF NOT EXISTS supplier_per_store_product
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    store_product_id uuid NOT NULL,
    supplier_id uuid NOT NULL,
    priority INTEGER NOT NULL CHECK (priority >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (store_product_id) REFERENCES product_per_store(id) ON DELETE CASCADE,
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON DELETE CASCADE,
    UNIQUE (store_product_id, supplier_id)
);

------------- DEPRECATED: product_company (soft delete - comentado) --------------
-- Esta tabla fue reemplazada por product_per_store
-- Se mantiene comentada para referencia durante la migración
/*
CREATE TABLE IF NOT EXISTS product_company
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id uuid NOT NULL,
    product_id uuid NOT NULL,
    tag_id integer NOT NULL,
    sku character varying NOT NULL,
    product_name character varying NOT NULL,
    item_purchase boolean NOT NULL,
    item_sale boolean NOT NULL, 
    item_inventory boolean NOT NULL,
    is_frozen boolean NOT NULL,
    use_recipe boolean NOT NULL,
    unit_purchase_id INTEGER NOT NULL,
    unit_inventory_id INTEGER NOT NULL,
    unit_matrix JSONB DEFAULT '{}'::jsonb,
    cost_last float NOT NULL,
    description character varying NOT NULL,
    cost_estimated float NOT NULL,
    cost_avg float NOT NULL,
    minimal_stock float NOT NULL,
    maximal_stock float NOT NULL,
    minimal_order float NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    FOREIGN KEY (unit_purchase_id) REFERENCES measurement_unit(id),
    FOREIGN KEY (unit_inventory_id) REFERENCES measurement_unit(id),
    UNIQUE (company_id, product_id)
);
*/


CREATE TABLE currency
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    iso_code character varying NOT NULL,
    numeric_code integer NOT NULL,
    currency_name character varying NOT NULL,
    currency_symbol character varying NOT NULL,
    decimal_places integer NOT NULL,
    rate numeric NOT NULL,
    available boolean NOT NULL DEFAULT true,
    PRIMARY KEY (id),
    UNIQUE (iso_code),
    UNIQUE (numeric_code)
);


-------------- BODEGAS --------------------
CREATE TABLE warehouse
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    store_id uuid NOT NULL,
    description character varying NULL,
    warehouse_name character varying NULL,
    warehouse_address character varying NULL,
    warehouse_phone character varying NULL, 
    is_momevent_warehouse boolean NOT NULL DEFAULT FALSE,
    --delivery_instructions character varying,
    --working_hours jsonb NULL DEFAULT ,
    --working_timezone character varying NOT NULL,
    created_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE warehouse_per_product
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    store_product_id uuid NOT NULL,
    warehouse_id uuid NOT NULL,
    warehouse_id_reference uuid NULL,
    direction character varying NULL,
    in_stock FLOAT NOT NULL,
    cost_avg FLOAT NOT NULL,
    --in_transit FLOAT NOT NULL,
    --ordered FLOAT NOT NULL,
    FOREIGN KEY (warehouse_id)
        REFERENCES warehouse (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    FOREIGN KEY (store_product_id)
        REFERENCES product_per_store (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE supplier_product (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    supplier_id UUID NOT NULL,
    product_id UUID NOT NULL,
    product_name CHARACTER VARYING NOT NULL,
    description CHARACTER VARYING,
    sku CHARACTER VARYING NOT NULL,
    unit_price NUMERIC NOT NULL,
    purchase_unit_id INTEGER NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    UNIQUE (supplier_id, product_id)
);

------------- DEPRECATED: request_restriction (soft delete - comentado) --------------
-- Esta tabla fue unificada con product_per_store (campo max_quantity)
/*
CREATE TABLE request_restriction
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    store_id uuid NOT NULL,
    product_company_id uuid NOT NULL,
    max_quantity numeric NOT NULL,
    FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
    FOREIGN KEY (product_company_id) REFERENCES product_company(id) ON DELETE CASCADE,
    UNIQUE (store_id, product_company_id)
);
*/


-- CREATE TYPE request_status AS ENUM
-- (
--     'pending',
--     'approved',
--     'rejected',
--     'conflicted',
--     'completed',
--     'cancelled'
-- );

-- CREATE TYPE request_type AS ENUM
-- (
--     'supplier_request', -- Solicitud a proveedor
--     'purchase_request', -- Otra empresa
--     'internal_request' -- Transferencia entre bodegas
-- );

-- CREATE TABLE inventory_request
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     display_id character varying(20) GENERATED ALWAYS AS ('SOL-' || upper(right(id::text, 5))) STORED,
--     store_id uuid NOT NULL,
--     warehouse_id uuid NOT NULL,
--     status request_status NOT NULL DEFAULT 'pending',
--     request_type request_type NOT NULL DEFAULT 'supplier_request',
--     requester_id uuid NOT NULL,
--     created_at timestamp with time zone NOT NULL,
--     updated_at timestamp with time zone NOT NULL,
--     PRIMARY KEY (id),
--     FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
--     FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE,
--     FOREIGN KEY (requester_id) REFERENCES user_accounts(id) ON DELETE SET NULL
-- );

-- CREATE TABLE inventory_request_item
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     inventory_request_id uuid NOT NULL,
--     store_product_id uuid NOT NULL,
--     quantity numeric NOT NULL,
--     PRIMARY KEY (id),
--     FOREIGN KEY (inventory_request_id) REFERENCES inventory_request(id) ON DELETE CASCADE,
--     FOREIGN KEY (store_product_id) REFERENCES product_per_store(id) ON DELETE CASCADE
-- );

-- CREATE TABLE inventory_request_history
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     inventory_request_id uuid NOT NULL,
--     new_status request_status NOT NULL,
--     observation text,
--     changed_by uuid,
--     changed_at timestamp NOT NULL DEFAULT now(),
--     PRIMARY KEY (id),
--     FOREIGN KEY (inventory_request_id) REFERENCES inventory_request(id) ON DELETE CASCADE,
--     FOREIGN KEY (changed_by) REFERENCES user_accounts(id) ON DELETE SET NULL
-- );

/*
CREATE TABLE supplier_product_per_store (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    supplier_product_id UUID NOT NULL,
    store_id UUID NOT NULL,

    priority INTEGER NOT NULL DEFAULT 9999, -- menor = más prioritario
    preferred BOOLEAN NOT NULL DEFAULT FALSE,
    service_zone character varying,
    price NUMERIC NOT NULL,
    min_order_quantity NUMERIC DEFAULT 1,
    payment_terms character varying,
    lead_time_days INTEGER NOT NULL,

    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    FOREIGN KEY (supplier_product_id) REFERENCES supplier_product(id) ON DELETE CASCADE,
    FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
    UNIQUE (supplier_product_id, store_id)
);
*/

-- CREATE TYPE purchase_status AS ENUM
-- (
--     'pending',     -- recién creada
--     'rejected',    -- rechazada por el proveedor
--     'completed',   -- entregada y cerrada
--     'sunk',        -- no concluida
--     'on_delivery', -- en camino
--     'arrived',     -- llegada a bodega pero no cerrada
--     'cancelled',     -- anulada
--     'edited'    -- corregida
-- );
-- CREATE TABLE purchase
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     display_id character varying(20) GENERATED ALWAYS AS ('ORD-' || upper(right(id::text, 5))) STORED,
--     supplier_id uuid NOT NULL,
--     store_id uuid NOT NULL,
--     inventory_request_id uuid NOT NULL,
--     delivery_purchase_note_id uuid,
--     status purchase_status NOT NULL DEFAULT 'pending',
--     created_at timestamp with time zone NOT NULL DEFAULT now(),
--     updated_at timestamp with time zone NOT NULL DEFAULT now(),
--     PRIMARY KEY (id),
--     FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON DELETE CASCADE,
--     FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
--     FOREIGN KEY (inventory_request_id) REFERENCES inventory_request(id) ON DELETE CASCADE
-- );

-- CREATE TABLE purchase_hierarchy
-- (
--     parent_purchase_id uuid NOT NULL,
--     child_purchase_id uuid NOT NULL,
--     PRIMARY KEY (parent_purchase_id, child_purchase_id),
--     FOREIGN KEY (parent_purchase_id) REFERENCES purchase(id) ON DELETE CASCADE,
--     FOREIGN KEY (child_purchase_id) REFERENCES purchase(id) ON DELETE CASCADE
-- );

-- CREATE TABLE purchase_history
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     purchase_id uuid NOT NULL,
--     new_status purchase_status NOT NULL,
--     observation text,
--     changed_at timestamp with time zone NOT NULL DEFAULT now(),
--     PRIMARY KEY (id),
--     FOREIGN KEY (purchase_id) REFERENCES purchase(id) ON DELETE CASCADE
-- );

-- CREATE TYPE purchase_item_status AS ENUM
-- (
--     'pending',   -- enviada al proveedor
--     'approved',  -- proveedor confirma
--     'rejected',  -- proveedor rechaza
--     'no_supplier', -- no hay proveedor
--     'retried' -- reintentado
-- );

-- CREATE TABLE purchase_item
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     purchase_id uuid NOT NULL,
--     store_product_id uuid NOT NULL,
--     quantity numeric NOT NULL,
--     unit_price numeric NOT NULL,
--     subtotal numeric GENERATED ALWAYS AS (quantity * unit_price) STORED,
--     available_suppliers jsonb NOT NULL DEFAULT '[]'::jsonb,
--     status purchase_item_status NOT NULL DEFAULT 'pending',
--     PRIMARY KEY (id),
--     FOREIGN KEY (purchase_id) REFERENCES purchase(id) ON DELETE CASCADE,
--     FOREIGN KEY (store_product_id) REFERENCES product_per_store(id) ON DELETE CASCADE
-- );

-- CREATE TABLE supplier_token
-- (
--     id bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
--     purchase_id character varying NOT NULL,
--     token_hash character varying NOT NULL,
--     exp_time TIMESTAMP NOT NULL,
--     used boolean NOT NULL,
--     PRIMARY KEY (id)
-- );

-- CREATE TYPE delivery_note_status AS ENUM (
--     'pending',   -- recién creada
--     'completed',  -- completada
--     'disputed',
--     'cancelled'   -- anulada
-- );

-- CREATE TABLE delivery_purchase_note
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     display_id character varying(20) GENERATED ALWAYS AS ('DPN-' || upper(right(id::text, 5))) STORED,
--     supplier_id uuid NOT NULL,
--     folio_invoice character varying NOT NULL DEFAULT '',
--     folio_guide character varying NOT NULL DEFAULT '',
--     store_id uuid NOT NULL,
--     warehouse_id uuid NOT NULL,
--     purchase_id uuid NOT NULL,
--     due_date date NOT NULL,
--     comment character varying,
--     note_status delivery_note_status NOT NULL DEFAULT 'pending',
--     total numeric NOT NULL,
--     user_id uuid NOT NULL,
--     created_at timestamp with time zone NOT NULL DEFAULT now(),
--     updated_at timestamp with time zone NOT NULL DEFAULT now(),
--     PRIMARY KEY (id),
--     UNIQUE (purchase_id),
--     FOREIGN KEY (supplier_id) REFERENCES supplier(id) ON DELETE CASCADE,
--     FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
--     FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE,
--     FOREIGN KEY (user_id) REFERENCES user_accounts(id) ON DELETE SET NULL
-- );


-- -- create foreign key to purchase
-- ALTER TABLE purchase ADD CONSTRAINT fk_delivery_purchase_note FOREIGN KEY (delivery_purchase_note_id) REFERENCES delivery_purchase_note(id) ON DELETE SET NULL;

-- CREATE TYPE delivery_note_status_item AS ENUM
-- (
--     'accepted',   -- recién creada
--     'substock',
--     'rejected',
--     'suprastock'
-- );

-- CREATE TABLE delivery_purchase_note_item
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     delivery_purchase_note_id uuid NOT NULL,
--     store_product_id uuid NOT NULL,
--     item_status delivery_note_status_item NOT NULL DEFAULT 'accepted',
--     quantity numeric NOT NULL,
--     difference numeric NOT NULL DEFAULT 0,
--     unit_price numeric NOT NULL,
--     subtotal numeric GENERATED ALWAYS AS (quantity * unit_price) STORED,
--     tax_total numeric NOT NULL DEFAULT 0,
--     PRIMARY KEY (id),
--     FOREIGN KEY (delivery_purchase_note_id) REFERENCES delivery_purchase_note(id) ON DELETE CASCADE,
--     FOREIGN KEY (store_product_id) REFERENCES product_per_store(id) ON DELETE CASCADE
-- );

-- CREATE TABLE file
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     file_type character varying NOT NULL,
--     file_url character varying NOT NULL,
--     created_at timestamp with time zone NOT NULL DEFAULT now(),
--     updated_at timestamp with time zone NOT NULL DEFAULT now(),
--     PRIMARY KEY (id)
-- );

-- CREATE TABLE file_per_entity
-- (
--     id uuid NOT NULL DEFAULT uuid_generate_v4(),
--     file_id uuid NOT NULL,
--     entity_name character varying NOT NULL,
--     entity_id uuid NOT NULL,
--     description character varying,
--     PRIMARY KEY (id),
--     FOREIGN KEY (file_id) REFERENCES file(id) ON DELETE CASCADE
-- );

---------- Movimientos de inventario --------------
CREATE TYPE movement_type AS ENUM
(
    'NEWINPUT',   -- Nuevo ingreso de inventario
    'TRANSFER',   -- Transferencia entre bodegas
    'WITHDRAWAL',  -- Retiro de inventario
    'WASTE'       -- Desperdicio o pérdida de inventario
);

CREATE TABLE IF NOT EXISTS product_movement
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    store_product_id uuid NOT NULL,
    observation character varying NOT NULL,
    quantity FLOAT NOT NULL,
    unit_cost FLOAT NOT NULL,
    total_cost FLOAT NOT NULL,
    -- product_state character varying NOT NULL,
    moved_from character varying,
    moved_to character varying,
    moved_at timestamp with time zone NOT NULL DEFAULT NOW(),
    moved_by character varying NOT NULL,
    movement_type movement_type NOT NULL,
    movement_doc_type character varying,
    movement_doc_reference character varying,
    purchase_id uuid NULL,
    inventory_unit character varying,
    stock_before FLOAT,
    stock_after FLOAT,
    PRIMARY KEY (id),
    FOREIGN KEY (store_product_id) REFERENCES product_per_store(id)
    -- FOREIGN KEY (purchase_id) REFERENCES purchase(id) ON DELETE SET NULL  -- Comentado: tabla purchase no existe
);

-- CREATE INDEX IF NOT EXISTS idx_product_movement_purchase_id ON product_movement(purchase_id);  -- Comentado temporalmente

--------------- Inventory Counts ----------------

CREATE TYPE inventory_count_status AS ENUM
(
    'un_assigned',
    'pending',
    'in_progress',
    'completed',
    'rejected'
);

CREATE TABLE IF NOT EXISTS public.inventory_count
(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    display_id character varying(20) GENERATED ALWAYS AS ('IC-' || upper(right(id::text, 5))) STORED,
    store_id UUID NOT NULL,
    company_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    created_by UUID NOT NULL,
    assigned_to UUID,
    status inventory_count_status NOT NULL,
    scheduled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    movement_track_id CHARACTER VARYING NOT NULL,
    metadata json,
    PRIMARY KEY (id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id)
);

CREATE TABLE IF NOT EXISTS public.inventory_count_item
(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    inventory_count_id UUID NOT NULL,
    store_product_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    scheduled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    incidence_image_url CHARACTER VARYING NULL,
    incidence_observation CHARACTER VARYING NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (inventory_count_id) REFERENCES inventory_count(id),
    UNIQUE (store_product_id, warehouse_id, scheduled_at)
);

--------------- Historial de precios por producto (plantilla) ----------------
CREATE TABLE IF NOT EXISTS price_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    previous_price NUMERIC NOT NULL,
    new_price NUMERIC NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_price_history_product_id ON price_history(product_id);
CREATE INDEX IF NOT EXISTS idx_price_history_created_at ON price_history(created_at DESC);

-- Trigger registros de tabla de precios
CREATE OR REPLACE FUNCTION fn_price_history_trigger()
RETURNS TRIGGER AS $$
BEGIN
    --registrar precio inicial
    IF TG_OP = 'INSERT' THEN
        IF NEW.cost_estimated > 0 THEN
            INSERT INTO price_history (product_id, previous_price, new_price)
            VALUES (NEW.id, 0, NEW.cost_estimated);
        END IF;
        RETURN NEW;
    END IF;
    
    --registrar solo si el precio cambió
    IF TG_OP = 'UPDATE' THEN
        IF OLD.cost_estimated IS DISTINCT FROM NEW.cost_estimated THEN
            INSERT INTO price_history (product_id, previous_price, new_price)
            VALUES (NEW.id, COALESCE(OLD.cost_estimated, 0), NEW.cost_estimated);
        END IF;
        RETURN NEW;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_price_history ON product;
CREATE TRIGGER trg_price_history
    AFTER INSERT OR UPDATE ON product
    FOR EACH ROW
    EXECUTE FUNCTION fn_price_history_trigger();

------------- Sistema de Solicitudes (Request) --------------

CREATE TABLE IF NOT EXISTS request_status
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name CHARACTER VARYING NOT NULL
);

CREATE TABLE IF NOT EXISTS doc_status
(
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name CHARACTER VARYING NOT NULL
);

CREATE TABLE IF NOT EXISTS request
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    display_id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    company_id UUID NOT NULL,
    store_id UUID NOT NULL,
    warehouse_id UUID,
    status_id INTEGER NOT NULL,
    request_kind INTEGER NOT NULL,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE,
    FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE SET NULL,
    FOREIGN KEY (status_id) REFERENCES request_status(id) ON DELETE RESTRICT,
    FOREIGN KEY (created_by) REFERENCES user_accounts(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS request_item
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_id UUID NOT NULL,
    requested_quantity NUMERIC NOT NULL,
    max_quantity NUMERIC NOT NULL,
    FOREIGN KEY (request_id) REFERENCES request(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_history
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_id UUID NOT NULL,
    status_id INTEGER NOT NULL,
    changed_by JSONB NOT NULL,
    observation CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (request_id) REFERENCES request(id) ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES request_status(id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS request_document
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_id UUID NOT NULL,
    doc_type INTEGER NOT NULL,
    doc_display_id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
    father_doc_id UUID,
    doc_status_id INTEGER NOT NULL,
    supplier_info JSONB NOT NULL,
    reception_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (request_id) REFERENCES request(id) ON DELETE CASCADE,
    FOREIGN KEY (father_doc_id) REFERENCES request_document(id) ON DELETE SET NULL,
    FOREIGN KEY (doc_status_id) REFERENCES doc_status(id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS request_doc_item
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_doc_id UUID NOT NULL,
    table_reference CHARACTER VARYING NOT NULL,
    reference_item_id UUID NOT NULL,
    item_data JSONB,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    receptioned_at TIMESTAMP WITH TIME ZONE,
    recepcioned_by UUID,
    reception_note CHARACTER VARYING,
    quantity NUMERIC DEFAULT 0,
    FOREIGN KEY (request_doc_id) REFERENCES request_document(id) ON DELETE CASCADE,
    FOREIGN KEY (recepcioned_by) REFERENCES user_accounts(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS request_doc_history
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_doc_id UUID NOT NULL,
    changed_by UUID NOT NULL,
    observation CHARACTER VARYING NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    left_display_position BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (request_doc_id) REFERENCES request_document(id) ON DELETE CASCADE,
    FOREIGN KEY (changed_by) REFERENCES user_accounts(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS request_doc_status_history
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    request_doc_id UUID NOT NULL,
    changed_by UUID NOT NULL,
    doc_status_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (request_doc_id) REFERENCES request_document(id) ON DELETE CASCADE,
    FOREIGN KEY (changed_by) REFERENCES user_accounts(id) ON DELETE SET NULL,
    FOREIGN KEY (doc_status_id) REFERENCES doc_status(id) ON DELETE RESTRICT
);





-- Necesario para permitir upsert con ON CONFLICT en transferencias de productos
ALTER TABLE warehouse_per_product 
ADD CONSTRAINT unique_store_product_warehouse 
UNIQUE (store_product_id, warehouse_id);
