package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	SessionCookieName = "session_id"

	ErrSessionNotFound = errors.New("session not found")
)

type SessionStore interface {
	// AddSession adds created session to session store and returns session id
	AddSession(sess string) (string, error)
	// RemoveSession removes session from session store
	RemoveSession(sessionID string) error
	// GetSession returns session from store and error if session was not found
	GetSession(sessionID string) (string, error)
}

type MemorySessionStore struct {
	sessions map[string]string
}

func (s *MemorySessionStore) AddSession(sess string) string {
	sessionID := uuid.New().String()
	s.sessions[sessionID] = sess

	return sessionID
}

func (s *MemorySessionStore) RemoveSession(sessionID string) {
	delete(s.sessions, sessionID)
}

func (s *MemorySessionStore) GetSession(sessionID string) (string, error) {
	if sess, ok := s.sessions[sessionID]; ok {
		return sess, nil
	} else {
		return "", ErrSessionNotFound
	}
}

type RedisSession struct {
	redis *redis.Client
}

func (s *RedisSession) AddSession(userID string) (string, error) {
	sessionID := uuid.New().String()
	sessionKey := getSessionKey(sessionID)
	ctx := context.Background()
	err := s.redis.Set(ctx, sessionKey, userID, 0).Err()
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (s *RedisSession) RemoveSession(sessionID string) error {
	ctx := context.Background()
	sessionKey := getSessionKey(sessionID)
	err := s.redis.Del(ctx, sessionKey).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisSession) GetSession(sessionID string) (string, error) {
	ctx := context.Background()
	sessionKey := getSessionKey(sessionID)
	session, err := s.redis.Get(ctx, sessionKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrSessionNotFound
		}
		return "", err
	}
	return session, nil
}

func NewRedisSession(redis *redis.Client) *RedisSession {
	return &RedisSession{
		redis: redis,
	}
}

func NewMemorySessionStore() *MemorySessionStore {
	return &MemorySessionStore{
		sessions: make(map[string]string),
	}
}

func getSessionKey(sessionID string) string {
	return fmt.Sprintf("session_%s", sessionID)
}
