package wrapper

import (
	"sync"
	"time"

	"github.com/benbjohnson/clock"
)

// ChargeTimer measures charging time between start and stop events
type ChargeTimer struct {
	sync.Mutex
	clck clock.Clock

	charging bool
	start    time.Time
	duration time.Duration
}

// NewChargeTimer creates ChargeTimer for tracking duration between
// start and stop events
func NewChargeTimer() *ChargeTimer {
	return &ChargeTimer{
		clck: clock.New(),
	}
}

// StartCharge signals charge timer start
func (m *ChargeTimer) StartCharge() {
	m.Lock()
	defer m.Unlock()

	m.charging = true
	m.start = m.clck.Now()
	m.duration = 0
}

// StopCharge signals charge timer stop
func (m *ChargeTimer) StopCharge() {
	m.Lock()
	defer m.Unlock()

	m.charging = false
	m.duration = m.clck.Since(m.start)
}

// ChargingTime implements the api.ChargeTimer interface
func (m *ChargeTimer) ChargingTime() (time.Duration, error) {
	m.Lock()
	defer m.Unlock()

	if m.charging {
		return m.clck.Since(m.start), nil
	}
	return m.duration, nil
}
