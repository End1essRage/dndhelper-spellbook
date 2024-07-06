package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetSpellInfo(c *gin.Context) {
	logrus.Info("hitted GetSpellInfo in handler")
	//Обработать входные данные
	spellName := c.Request.URL.Query().Get("name")
	//дернуть сервис
	spell, err := h.services.SpellBook.GetSpellInfo(spellName)
	if err != nil {
		logrus.Error("error in GetSpellInfo", err.Error())
	}
	//отправить ответ
	c.JSON(http.StatusOK, spell)
}
