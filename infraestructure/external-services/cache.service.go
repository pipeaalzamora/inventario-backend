package externalservices

import (
	"context"
	"encoding/json"
	"reflect"
	"sofia-backend/types"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService(redisClient *redis.Client) *CacheService {
	return &CacheService{
		client: redisClient,
		ctx:    context.TODO(),
	}
}

func (s *CacheService) AddKeyValue(key string, data interface{}, ex time.Duration) error {
	marshalled, err := json.Marshal(data)
	if err != nil {
		return types.ThrowData("Error al serializar los datos del caché")
	}

	stsCmd := s.client.Set(s.ctx, s.key(key), marshalled, ex)
	if err := stsCmd.Err(); err != nil {
		return types.ThrowData("Error al guardar la clave en el caché")
	}
	return nil
}

func (s *CacheService) GetWithKey(key string, st interface{}) error {
	stsCmd := s.client.Get(s.ctx, s.key(key))
	if err := stsCmd.Err(); err != nil {
		if err == redis.Nil {
			return err
		}
		return types.ThrowData("Error al obtener la clave del caché")
	}

	bytes, err := stsCmd.Bytes()
	if err != nil {
		return types.ThrowData("Error al leer los bytes del caché")
	}

	err = json.Unmarshal(bytes, st)
	if err != nil {
		return types.ThrowData("Error al deserializar los datos del caché")
	}

	return nil
}

func (s *CacheService) GetWithPrefix(prefix string, st interface{}) error {
	var cursor uint64
	var allBytes [][]byte

	for {
		keys, newCursor, err := s.client.Scan(s.ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			return types.ThrowData("Error al escanear las claves del caché")
		}
		cursor = newCursor

		if len(keys) > 0 {
			vals, err := s.client.MGet(s.ctx, keys...).Result()
			if err != nil {
				return types.ThrowData("Error al obtener los valores del caché")
			}

			for _, val := range vals {
				if val == nil {
					continue
				}
				strVal, ok := val.(string)
				if !ok {
					continue
				}
				allBytes = append(allBytes, []byte(strVal))
			}
		}

		if cursor == 0 {
			break
		}
	}

	// Deserializar todos los JSON en elementos del slice
	sliceValue := reflect.ValueOf(st)
	if sliceValue.Kind() != reflect.Ptr {
		return types.ThrowData("Se esperaba un puntero a un slice")
	}
	sliceValue = sliceValue.Elem()
	elemType := sliceValue.Type().Elem()

	for _, b := range allBytes {
		elemPtr := reflect.New(elemType)
		err := json.Unmarshal(b, elemPtr.Interface())
		if err != nil {
			return types.ThrowData("Error al deserializar los datos del caché")
		}
		sliceValue.Set(reflect.Append(sliceValue, elemPtr.Elem()))
	}

	return nil
}

func (s *CacheService) DeleteByKey(key string) error {
	stsCmd := s.client.Del(s.ctx, s.key(key))
	if err := stsCmd.Err(); err != nil {
		return types.ThrowData("Error al eliminar la clave del caché")
	}
	return nil
}

func (s *CacheService) DeleteByPrefix(prefix string) error {
	var cursor uint64

	for {
		// Use SCAN to find keys matching the prefix
		keys, newCursor, err := s.client.Scan(s.ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			return types.ThrowData("Error al escanear las claves del caché")
		}
		cursor = newCursor

		if len(keys) > 0 {
			if err := s.client.Del(s.ctx, keys...).Err(); err != nil {
				return types.ThrowData("Error al eliminar las claves del caché")
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (s *CacheService) key(key string) string {
	// return fmt.Sprintf("%s", key)
	return key
}
