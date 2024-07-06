package service

import (
	client "github.com/end1essrage/dndhelper-spellbook/pkg/client"
)

type SpellBook interface {
	GetSpellInfo(spellName string) (client.Spell, error)
}

type Service struct {
	SpellBook
}

// Сюда надо клиент прокинуть
func NewService(client *client.Client) *Service {
	return &Service{
		SpellBook: NewSpellBookService(client),
	}
}
