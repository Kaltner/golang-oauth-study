package oauth

import (
	"strings"

	"github.com/google/uuid"
)

type StateManager struct {
	stateList map[string]OauthProvider
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateList: make(map[string]OauthProvider),
	}
}

func (s *StateManager) GenerateState(provider OauthProvider) string {
	state := strings.ReplaceAll(uuid.NewString(), "-", "")
	s.stateList[state] = provider
	return state
}

func (s *StateManager) FindState(state string) (OauthProvider, bool) {
	provider, ok := s.stateList[state]
	return provider, ok
}

func (s *StateManager) DeleteState(state string) {
	delete(s.stateList, state)
}
