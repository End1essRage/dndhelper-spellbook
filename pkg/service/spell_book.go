package service

import (
	client "github.com/end1essrage/dndhelper-spellbook/pkg/client"
	"github.com/sirupsen/logrus"
)

type SpellBookService struct {
	//тут нужен клиент для dnd5api
	client *client.Client
}

func NewSpellBookService(client *client.Client) *SpellBookService {
	return &SpellBookService{client: client}
}

func (s *SpellBookService) GetSpellInfo(spellName string) (string, error) {
	//add log
	logrus.Info("SpellBookService - GetSpellInfo")
	s.client.GetSpellInfo(spellName)

	return "GetSpellInfo", nil
}
