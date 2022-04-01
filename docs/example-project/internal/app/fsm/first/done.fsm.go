// Code generated by fsm-generator. DO NOT EDIT.
package first

import (
	"time"

	"fsm-framework/fsm-engine/model"
)

var DoneState model.State = &DoneStateDeclaration{}

type DoneStateDeclaration struct {
}

func (s *DoneStateDeclaration) Name() string {
	return "FIRST_TX_DONE"
}

func (s *DoneStateDeclaration) EventType() string {
	return "first_tx_done_state_event"
}

func (s *DoneStateDeclaration) Queue() string {
	return "first_tx_done_state_event_queue"
}

func (s *DoneStateDeclaration) CanTransitIn(state model.State) bool {
	if state == DoneState {
		return true
	}

	return false
}

func (s *DoneStateDeclaration) MaxRetiesCount() int {
	return 3
}

func (s *DoneStateDeclaration) MinRetiesDelay() time.Duration {
	return 15 * time.Second // 15s
}

func (s *DoneStateDeclaration) CancellationTTL() time.Duration {
	return 1800 * time.Second // 30m0s
}

func (s *DoneStateDeclaration) FallbackState() model.State {
	return nil
}

func (s *DoneStateDeclaration) IsInitial() bool {
	return false
}

func (s *DoneStateDeclaration) IsSuccessFinal() bool {
	return true
}

func (s *DoneStateDeclaration) IsFailFinal() bool {
	return false
}

func (s *DoneStateDeclaration) Model() model.Model {
	return Model
}

func (s *DoneStateDeclaration) Service() Service {
	return Model.Service().(Service)
}
