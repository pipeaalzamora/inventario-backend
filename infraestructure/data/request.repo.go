package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type requestRepo struct {
	db *sqlx.DB
}

func NewRequestRepo(db *sqlx.DB) ports.PortRequest {
	return &requestRepo{db: db}
}

func (r *requestRepo) CreateRequest(request *models.ModelRequest) (*models.ModelRequest, error) {
	// Obtener status_id y request_kind del modelo
	statusID := 1 // Estado por defecto "Creada"
	if request.Status.Id != 0 {
		statusID = request.Status.Id
	}

	requestKind := 1 // Tipo por defecto
	if request.RequestKind.Id != 0 {
		requestKind = request.RequestKind.Id
	}

	// Crear entity para la request principal
	entity := &entities.EntityRequest{
		CompanyID:   request.CompanyId,
		StoreID:     request.StoreId,
		WarehouseID: nil, // Se asigna después si existe
		StatusID:    statusID,
		RequestKind: requestKind,
		CreatedBy:   request.CreatedBy.Id,
	}

	// Manejar warehouse_id nullable
	if request.WarehouseId != nil && *request.WarehouseId != "" {
		entity.WarehouseID = request.WarehouseId
	}

	// Iniciar transacción
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Insertar request principal
	err := tx.QueryRowx(`
		INSERT INTO 
			request (company_id, store_id, warehouse_id, status_id, request_kind, created_by, created_at, updated_at) 
		VALUES 
			($1, $2, null, $3, $4, $5, NOW(), NOW()) 
		RETURNING id, display_id, created_at, updated_at
		`,
		entity.CompanyID,
		entity.StoreID,
		entity.StatusID,
		entity.RequestKind,
		entity.CreatedBy,
	).Scan(&entity.ID, &entity.DisplayID, &entity.CreatedAt, &entity.UpdatedAt)

	if err != nil {
		return nil, types.ThrowData("Error al insertar la solicitud" + err.Error())
	}

	// Preparar changed_by para el historial
	changedBy := entities.RequestHistoryChangedBy{
		Name:             request.CreatedBy.UserName,
		OrganizationName: "", // Se puede obtener del company si es necesario
	}

	// Insertar historial inicial
	_, err = tx.Exec(
		`INSERT INTO request_history (request_id, status_id, changed_by, observation, created_at) 
		VALUES ($1, $2, $3, $4, NOW())`,
		entity.ID,
		entity.StatusID,
		changedBy,
		"Solicitud creada",
	)

	if err != nil {
		return nil, types.ThrowData("Error al insertar el historial de la solicitud")
	}

	// Commit de la transacción
	err = tx.Commit()
	if err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	// Mapear entity a modelo y retornar
	result := r.toModel(entity, []entities.EntityRequestItem{}, &request.CreatedBy)
	return result, nil
}

// toModel convierte las entidades a modelo de dominio
func getStatusName(statusID int) string {
	switch statusID {
	case 1:
		return "Creada"
	case 2:
		return "Con conflicto"
	case 3:
		return "Aprobada"
	case 4:
		return "Cancelada"
	case 5:
		return "Finalizada"
	default:
		return ""
	}
}

func getRequestKindName(kindID int) string {
	switch kindID {
	case 1:
		return "Proveedor"
	case 2:
		return "Empresa"
	case 3:
		return "Tienda"
	default:
		return ""
	}
}

func (r *requestRepo) toModel(entity *entities.EntityRequest, items []entities.EntityRequestItem, createdBy *models.ModelRequestUser) *models.ModelRequest {
	modelItems := make([]models.ModelRequestItem, 0)
	for _, item := range items {
		modelItems = append(modelItems, models.ModelRequestItem{
			Id:                item.ID,
			RequestedQuantity: item.RequestedQuantity,
			RequestRestriction: models.ModelRequestRestriction{
				MaxQuantity: item.MaxQuantity,
			},
		})
	}

	return &models.ModelRequest{
		Id:          entity.ID,
		DisplayId:   strconv.Itoa(entity.DisplayID),
		CompanyId:   entity.CompanyID,
		StoreId:     entity.StoreID,
		WarehouseId: entity.WarehouseID,
		Status: models.ModelRequestStatus{
			Id:   entity.StatusID,
			Name: getStatusName(entity.StatusID),
		},
		RequestKind: models.ModelRequestKind{
			Id:   entity.RequestKind,
			Name: getRequestKindName(entity.RequestKind),
		},
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
		CreatedBy:      *createdBy,
		Items:          modelItems,
		RequestHistory: []models.ModelRequestHistory{},
		DocsTree:       []models.ModelRequestDoc{},
	}
}
