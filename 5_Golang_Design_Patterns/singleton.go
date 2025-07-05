package main

import "sync"

type Manager interface {
	GetConfig() bool
}

type manager struct {
	config bool
}

var instance *manager

var once sync.Once

func Instance() Manager {
	once.Do(func() {
		instance = &manager{
			config: true,
		}
	})
	return instance
}

func (m *manager) GetConfig() bool {
	return false
}
