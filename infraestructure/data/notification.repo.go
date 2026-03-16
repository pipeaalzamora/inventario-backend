package data

import (
	"context"
	"encoding/json"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type NotificationRepo struct {
	db *sqlx.DB
}

func NewNotificationRepo(db *sqlx.DB) ports.PortNotification {
	return &NotificationRepo{
		db: db,
	}
}

func (r *NotificationRepo) GetNotificationsByUserID(userID string, page int, size int, filter *map[string]interface{}) ([]models.NotificationModel, int, error) {
	var notifications []entities.EntityNotification
	query := `
		SELECT id, from_user, to_user, read_at, send_at, notification_type, payload
		FROM notifications
	`
	queryFilter := " WHERE to_user = $1"
	if filter != nil {
		jsonData, err := json.Marshal(*filter)
		if err != nil {
			return nil, 0, types.ThrowData("Error al procesar los filtros")
		}
		var filter entities.FilterNotification
		if err := json.Unmarshal(jsonData, &filter); err != nil {
			return nil, 0, types.ThrowData("Error al procesar los filtros")
		}

		if filter.WasRead != nil {
			queryFilter += " AND read_at IS " + map[bool]string{true: "NOT NULL", false: "NULL"}[*filter.WasRead]
		}
	}

	query += queryFilter + " ORDER BY send_at DESC LIMIT $2 OFFSET $3"

	err := r.db.SelectContext(context.TODO(), &notifications, query, userID, size, (page-1)*size)
	if err != nil {
		return nil, 0, err
	}

	var total int
	err = r.db.GetContext(context.TODO(), &total, "SELECT COUNT(*) FROM notifications"+queryFilter, userID)
	if err != nil {
		return nil, 0, types.ThrowData("Error al contar las notificaciones")
	}

	notificationModels := r.toModelMap(notifications)

	return notificationModels, total, nil
}

func (r *NotificationRepo) CreateNotification(notification *models.NotificationModel) (*models.NotificationModel, error) {
	bytesPayload, err := json.Marshal(notification.Payload)
	if err != nil {
		return nil, fmt.Errorf("invalid payload format: %w", err)
	}

	// Convertimos modelo a entidad para poder usar NamedExec con mapeo de tags `db`
	entity := entities.EntityNotification{
		From:             notification.From,
		To:               notification.To,
		ReadAt:           nil,
		SendAt:           time.Now(),
		NotificationType: notification.NotificationType,
		Payload:          bytesPayload,
	}

	query := `
		INSERT INTO notifications (from_user, to_user, read_at, send_at, notification_type, payload)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, from_user, to_user, read_at, send_at, notification_type, payload
	`

	var created entities.EntityNotification
	err = r.db.QueryRow(
		query,
		entity.From,
		entity.To,
		entity.ReadAt,
		entity.SendAt,
		entity.NotificationType,
		entity.Payload,
	).Scan(&created.ID, &created.From, &created.To, &created.ReadAt, &created.SendAt, &created.NotificationType, &created.Payload)
	if err != nil {
		return nil, types.ThrowData("Error al crear la notificación")
	}

	return r.toModel(&created), nil
}

func (r *NotificationRepo) UpdateRead(userId string, notificationIDs []string, wasRead bool) error {
	query := `
		UPDATE notifications
		SET read_at = CASE WHEN $3 THEN NOW() ELSE NULL END
		WHERE id = ANY($1) AND to_user = $2
	`
	_, err := r.db.ExecContext(context.TODO(), query, pq.Array(notificationIDs), userId, wasRead)
	if err != nil {
		return types.ThrowData("Error al actualizar el estado de lectura de la notificación")
	}
	return nil
}

func (r *NotificationRepo) toModel(e *entities.EntityNotification) *models.NotificationModel {

	// Convertir payload de JSON bruto a map[string]interface{}
	var payload map[string]interface{}
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		// Si hay un error en la conversión, asignar un mapa vacío o manejar el error según sea necesario
		payload = map[string]interface{}{}
	}
	return &models.NotificationModel{
		ID:               e.ID,
		From:             e.From,
		To:               e.To,
		ReadAt:           e.ReadAt,
		SendAt:           e.SendAt,
		NotificationType: e.NotificationType,
		Payload:          payload,
	}
}

func (r *NotificationRepo) toModelMap(es []entities.EntityNotification) []models.NotificationModel {
	result := make([]models.NotificationModel, len(es))
	for i, e := range es {
		result[i] = *r.toModel(&e)
	}
	return result
}
