package breakpoint

import "sort"

type Manager struct {
	breakpoints map[uint16]struct{}
	hitPoint    uint16
}

func NewManager() *Manager {
	return &Manager{
		breakpoints: make(map[uint16]struct{}, 0),
	}
}

func (m *Manager) AddBreakpoint(point uint16) bool {
	_, exists := m.breakpoints[point]
	if !exists {
		m.breakpoints[point] = struct{}{}
	}
	return exists
}

func (m *Manager) Remove(point uint16) {
	delete(m.breakpoints, point)
}

func (m *Manager) CheckAndHitBreakpoint(point uint16) bool {
	if m.hitPoint == point {
		m.hitPoint = 0
		return true
	}
	_, exists := m.breakpoints[point]
	if exists {
		m.hitPoint = point
	}
	return exists
}

func (m *Manager) IsBreakPoint(point uint16) bool {
	_, exists := m.breakpoints[point]
	return exists
}

func (m *Manager) GetBreakpoints() []uint16 {
	breakpoints := make([]uint16, len(m.breakpoints))
	i := 0
	for bp := range m.breakpoints {
		breakpoints[i] = bp
		i++
	}
	sort.Slice(breakpoints, func(i, j int) bool { return breakpoints[i] > breakpoints[j] })
	return breakpoints
}

func (m *Manager) Clear() {
	m.breakpoints = make(map[uint16]struct{}, 2)
	m.hitPoint = 0
}
