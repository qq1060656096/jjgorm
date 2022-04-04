package jjgorm

import (
	"encoding/json"
	"sync"
)

var defaultManager = NewManager()

type Manager struct {
	conList map[string]*Connection
	rwMutex sync.RWMutex
}

func (m *Manager) Add(name string, conf Config)  {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	if m.noLockExist(name) {
		return
	}
	m.conList[name] = NewConnection(conf)
}

func (m *Manager) Change(name string, conf Config)  {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	m.conList[name] = NewConnection(conf)
}

func (m *Manager) Remove(name string)  {
	delete(m.conList, name)
}



func (m *Manager) Get(name string) *Connection {
	defer m.rwMutex.RUnlock()
	m.rwMutex.RLock()
	con, ok := m.conList[name]
	if ok {
		return con
	}
	return nil
}

func (m *Manager) Exist(name string) bool {
	defer m.rwMutex.RUnlock()
	m.rwMutex.RLock()
	_, ok := m.conList[name]
	if ok {
		return true
	}
	return false
}

func (m *Manager) noLockExist(name string) bool {
	_, ok := m.conList[name]
	if ok {
		return true
	}
	return false
}

func (m *Manager) Len() int {
	return len(m.conList)
}

func (m *Manager) String() string {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	if m.conList == nil {
		return ""
	}
	if m.Len() < 1{
		return ""
	}
	type tmpType struct {
		Name  string
		HasDb bool
		Conf Config
	}
	l := make([]tmpType, 0, len(m.conList))
	for name, conn := range m.conList {
		tmp := tmpType{
			Name: name,
			Conf: conn.conf,
		}
		if conn.db != nil {
			tmp.HasDb = true
		}
		l = append(l, tmp)
	}
	b, _:= json.Marshal(l)
	return string(b)
}

func NewManager() *Manager {
	return &Manager{
		conList: make(map[string]*Connection, 0),
	}
}

func DefaultManager() *Manager {
	return defaultManager
}
