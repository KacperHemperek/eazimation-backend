package auth

import (
	"errors"
	"github.com/google/uuid"
)

var (
	SessionCookieName = "session_id"

	SessionNotFoundErr = errors.New("session not found")
)

type SessionUser struct {
	Email  string `json:"email"`
	UserID string `json:"userId"`
}

type SessionStore interface {
	// AddSession adds created session to session store and returns session id
	AddSession(sess *SessionUser) string
	// RemoveSession removes session from session store
	RemoveSession(sessionID string)
	// GetSession returns session from store and error if session was not found
	GetSession(sessionID string) (*SessionUser, error)
}

type MemorySessionStore struct {
	sessions map[string]*SessionUser
}

func (s *MemorySessionStore) AddSession(sess *SessionUser) string {
	sessionID := uuid.New().String()
	s.sessions[sessionID] = sess

	return sessionID
}

func (s *MemorySessionStore) RemoveSession(sessionID string) {
	delete(s.sessions, sessionID)
}

func (s *MemorySessionStore) GetSession(sessionID string) (*SessionUser, error) {
	if sess, ok := s.sessions[sessionID]; ok {
		return sess, nil
	} else {
		return nil, SessionNotFoundErr
	}
}

func NewMemorySessionStore() *MemorySessionStore {
	return &MemorySessionStore{
		sessions: make(map[string]*SessionUser),
	}
}

func NewSessionUser(e, id string) *SessionUser {
	return &SessionUser{
		Email:  e,
		UserID: id,
	}
}
