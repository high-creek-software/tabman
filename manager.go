package tabman

import "fyne.io/fyne/v2/container"

type Manager[I comparable] struct {
	IDs  map[I]*container.TabItem
	Tabs map[*container.TabItem]I
}

func NewManager[I comparable]() *Manager[I] {
	return &Manager[I]{IDs: make(map[I]*container.TabItem), Tabs: make(map[*container.TabItem]I)}
}

func (m *Manager[I]) GetTabItem(id I) (*container.TabItem, bool) {
	ti, ok := m.IDs[id]
	return ti, ok
}

func (m *Manager[I]) AddTabItem(id I, tab *container.TabItem) {
	m.Tabs[tab] = id
	m.IDs[id] = tab
}

func (m *Manager[I]) RemoveTab(tab *container.TabItem) {
	id, ok := m.Tabs[tab]
	if !ok {
		return
	}
	delete(m.Tabs, tab)
	delete(m.IDs, id)
}
