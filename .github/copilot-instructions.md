# Sofia Backend - Guía para Agentes de IA

## Arquitectura General

Este es un backend Go que sigue **Clean Architecture** con una jerarquía estricta de dependencias:

```
API (controllers) → Facades → Services → Repositories (data)
                         ↓
                    External Services
```

- **domain/**: Lógica de negocio pura (services, facades, models, ports)
- **infrastructure/**: Implementaciones concretas (data repos, external-services)
- **api/v1/**: Capa HTTP (controllers, DTOs, recipes)
- **shared/**: Utilidades transversales (validación, errores, tokens)

### Flujo de Datos Típico

1. **Controller** recibe request HTTP → valida con `recipe` (input) → llama **Facade**
2. **Facade** coordina múltiples **Services** y **External Services**
3. **Service** ejecuta lógica de negocio → llama **Port** (interfaz)
4. **Repository** (en `infrastructure/data/`) implementa **Port** → accede DB

## Convenciones Críticas del Proyecto

### Nomenclatura de Archivos
- `*.facade.go` en `domain/facades/`
- `*.service.go` en `domain/services/`
- `*.port.go` en `domain/ports/` (interfaces)
- `*.repo.go` en `infraestructure/data/` (implementaciones)
- `*.model.go` en `domain/models/`
- `*.dto.go` en `api/v1/dto/` (output)
- `*.recipe.go` en `api/v1/recipe/` (input con validaciones)

### Patrón de Dependencias
Cada capa tiene un archivo `build.go` que construye el contenedor de dependencias:
- `infraestructure/data/build.go` → `DataContainer`
- `infraestructure/external-services/build.go` → `ExternalServicesContainer`
- `domain/services/build.go` → `ServiceContainer`
- `domain/facades/build.go` → `FacadeContainer`
- `main.go` orquesta la inyección completa

**Ejemplo:** Agregar un nuevo servicio requiere:
1. Crear servicio en `domain/services/`
2. Definir port en `domain/ports/`
3. Implementar repo en `infraestructure/data/`
4. Actualizar `build.go` en cada capa
5. Crear facade si es endpoint público

### Sistema de Permisos (PowerChecker)
Los servicios embeben `PowerChecker` para control de acceso:

```go
type ProductService struct {
    PowerChecker  // Embebido
    productRepo ports.PortProduct
}

// Uso en métodos
if !s.EveryPower(ctx, POWER_TEMPLATE_PRODUCT_CREATE) {
    return nil, types.ThrowPower("Sin permisos...")
}
```

- `EveryPower(ctx, ...powers)`: Usuario DEBE tener TODOS los permisos
- `SomePower(ctx, ...powers)`: Usuario debe tener AL MENOS UNO
- Permisos vienen del contexto Gin (`shared.UserPowersKeys()`)

### Manejo de Errores
Usa tipos de error específicos en `shared/error.shared.go`:
- `DomainError`: Errores de lógica de negocio
- `DataError`: Errores de persistencia
- `PowerError`: Errores de permisos
- `ValidationErrorItem`: Errores de validación estructural

En controllers, usa `shared.FormatValidationErrors()` para validaciones y `shared.SendValidationErrorResponse()`.

### Validaciones en Recipes
Los `recipe` usan tags de `binding` con mensajes personalizados:

```go
type RecipeProduct struct {
    Name string `form:"name" binding:"required" errMsg:"El nombre es obligatorio"`
    Codes string `form:"codes" binding:"required" errMsg:"Los códigos son obligatorios"`
}
```

Valida en controller con `gctx.ShouldBindUri(&params)` o `gctx.ShouldBind()`.

## Configuración y Entorno

### Setup de Desarrollo
1. Copiar `config/config.example.json` → `config/config.json`
2. Levantar servicios: `docker compose up -d` (PostgreSQL, Redis, Meilisearch)
3. Ejecutar con hot-reload: `air` (recomendado) o `go run main.go`

### Bases de Datos
- **PostgreSQL**: Schema en `migrations/schema.sql`, seed en `migrations/seed.sql`
- **Redis**: Para caché (via `CacheService`)
- **Meilisearch**: Para búsqueda (via `SearchService`)

**Reiniciar DB desde cero:**
```bash
docker compose down -v
docker compose up -d
```

### Servicios Externos (infraestructure/external-services/)
- `BucketService`: GCP Cloud Storage para imágenes
- `MailerService`: Resend API para emails
- `CacheService`: Redis wrapper
- `SearchService`: Meilisearch wrapper
- `SSEservice`: Server-Sent Events para notificaciones

## Patrones de Implementación

### Crear Nuevo Endpoint
1. Define modelo en `domain/models/*.model.go`
2. Define port (interfaz) en `domain/ports/*.port.go`
3. Implementa repo en `infraestructure/data/*.repo.go`
4. Actualiza `infraestructure/data/build.go`
5. Crea servicio en `domain/services/*.service.go`
6. Actualiza `domain/services/build.go`
7. Crea facade en `domain/facades/*.facade.go`
8. Actualiza `domain/facades/build.go`
9. Crea recipe (input) en `api/v1/recipe/*.recipe.go`
10. Crea DTO (output) en `api/v1/dto/*.dto.go`
11. Crea controller en `api/v1/controllers/*.controller.go`
12. Registra rutas en `api/v1/server/server.go`

### Controller Típico
```go
func (ctrl *ProductController) RegisterRoutes(rg *gin.RouterGroup) {
    r := rg.Group("/products")
    r.GET("codes", ctrl.getAllCodes)      // Rutas específicas primero
    r.GET(":id", ctrl.getById)            // Parámetros después
}

func (ctrl *ProductController) getById(gctx *gin.Context) {
    type pathParams struct {
        ID string `uri:"id" binding:"required,uuid"`
    }
    var params pathParams
    if err := gctx.ShouldBindUri(&params); err != nil {
        errors := shared.FormatValidationErrors(err, &params)
        shared.SendValidationErrorResponse(gctx, errors)
        return
    }
    // Lógica...
}
```

### Subida de Archivos
Usa `BucketService` en servicios que manejan imágenes (ver `ProductService.SaveImageToBucket`).

## Contexto de Ejecución

- **Framework HTTP**: Gin (router en `api/v1/server/`)
- **ORM**: sqlx (sin ORM completo, queries manuales)
- **Migraciones**: SQL manual en `migrations/`
- **Dependencias clave**: `go.mod` incluye Gin, JWT, Redis, Meilisearch, Resend

## Notas Importantes

- **No hay tests automáticos** actualmente en el proyecto
- **Middleware de autenticación** está en `api/v1/api-services/auth.middleware.go`
- **CORS configurado** solo en modo debug (ver `server.go`)
- **Live reload** con Air (archivo `.air.toml` esperado en raíz si se usa)
