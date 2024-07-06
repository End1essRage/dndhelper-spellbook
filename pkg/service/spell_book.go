package service

import (
	client "github.com/end1essrage/dndhelper-spellbook/pkg/client"
	"github.com/sirupsen/logrus"
)

type SpellBookService struct {
	client *client.Client
}

func NewSpellBookService(client *client.Client) *SpellBookService {
	return &SpellBookService{client: client}
}

func (s *SpellBookService) GetSpellInfo(spellName string) (client.Spell, error) {
	//add log
	logrus.Info("SpellBookService - GetSpellInfo")
	spell, err := s.client.GetSpellInfo(spellName)

	return spell, err
}
