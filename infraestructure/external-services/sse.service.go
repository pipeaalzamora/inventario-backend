package externalservices

import (
	"fmt"
	"log"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sync"
	"time"
)

// Basado en la implementación del siguiente enlace:
// https://gist.github.com/SubCoder1/3a700149b2e7bb179a9123c6283030ff

type SSEservice struct {
	message      chan models.NotificationModel
	totalClients map[string]chan models.NotificationModel
	mu           sync.RWMutex
}

func NewSSEService() ports.PortSee {
	sse := &SSEservice{
		message:      make(chan models.NotificationModel),
		totalClients: make(map[string]chan models.NotificationModel),
	}
	go sse.listen()
	go sse.cleanupClients()
	return sse
}

func (s *SSEservice) SendMessage(notification models.NotificationModel) {
	s.message <- notification
}

func (s *SSEservice) GetClients() []models.ClientModel {
	s.mu.RLock()
	defer s.mu.RUnlock()

	clients := make([]models.ClientModel, len(s.totalClients))
	i := 0
	for id, ch := range s.totalClients {
		clients[i] = models.ClientModel{ID: id, Channel: ch}
		i++
	}
	return clients
}

func (s *SSEservice) AddClient(client models.ClientModel) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("Adding client %s\n", client.ID)
	s.totalClients[client.ID] = client.Channel
}

func (s *SSEservice) RemoveClient(client models.ClientModel) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("Removing client %s\n", client.ID)
	if ch, exists := s.totalClients[client.ID]; exists {
		close(ch)
		delete(s.totalClients, client.ID)
	}
}

func (s *SSEservice) BroadcastMessage(notification models.NotificationModel) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for id, ch := range s.totalClients {
		if ch != nil {
			notification.To = id
			ch <- notification
		} else {
			log.Printf("Channel for client %s is nil. Skipping broadcast.", id)
		}
	}
}

func (s *SSEservice) GetClient(userID string) chan models.NotificationModel {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.totalClients[userID]
}

func (s *SSEservice) listen() {
	for notification := range s.message {
		s.mu.RLock()
		ch, exists := s.totalClients[notification.To]
		s.mu.RUnlock()

		if exists && ch != nil {
			select {
			case ch <- notification:
				// Message sent successfully
			case <-time.After(2 * time.Second):
				// Timeout - client probably disconnected
				log.Printf("Timeout sending to client %s, removing client", notification.To)
				s.RemoveClient(models.ClientModel{ID: notification.To, Channel: ch})
			}
		} else {
			log.Printf("No client channel found for ID %s", notification.To)
		}
	}
}

func (s *SSEservice) cleanupClients() {
	// Esta función no se si es necesaria, pero se deja para limpiar canales nil
	// y evitar fugas de memoria.
	// Se ejecuta cada 5 segundos para revisar los canales de los clientes.
	log.Println("Starting client cleanup routine")
	for {
		s.mu.Lock()
		for id, ch := range s.totalClients {
			if ch == nil {
				log.Printf("Removing nil channel for client %s", id)
				delete(s.totalClients, id)
			}
		}
		s.mu.Unlock()

		time.Sleep(5 * time.Second) // dormir fuera del lock para evitar bloqueos
	}
}
