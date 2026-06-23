CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS company_branding (
    company_id UUID PRIMARY KEY,
    app_name character varying NOT NULL,
    logo_url character varying,
    favicon_url character varying,
    primary_color character varying NOT NULL DEFAULT '#2563eb',
    accent_color character varying NOT NULL DEFAULT '#16a34a',
    email_from_name character varying NOT NULL,
    support_email character varying,
    custom_domain character varying,
    subdomain character varying,
    welcome_text character varying,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS company_branding_assets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID NOT NULL,
    asset_type character varying NOT NULL,
    file_name character varying NOT NULL,
    file_url character varying NOT NULL,
    mime_type character varying NOT NULL,
    file_size BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_company_branding_assets_company_id ON company_branding_assets(company_id);

CREATE TABLE IF NOT EXISTS company_email_templates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID NOT NULL,
    template_key character varying NOT NULL,
    subject character varying NOT NULL,
    body_html TEXT NOT NULL,
    body_text TEXT,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (company_id, template_key),
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS company_import_jobs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID NOT NULL,
    import_type character varying NOT NULL,
    file_name character varying NOT NULL,
    file_url character varying NOT NULL,
    mime_type character varying NOT NULL,
    status character varying NOT NULL DEFAULT 'previewed',
    total_rows integer NOT NULL DEFAULT 0,
    headers JSONB NOT NULL DEFAULT '[]'::jsonb,
    sample_rows JSONB NOT NULL DEFAULT '[]'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_company_import_jobs_company_id ON company_import_jobs(company_id);
