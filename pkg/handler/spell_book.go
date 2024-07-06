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
	h.services.SpellBook.GetSpellInfo(spellName)
	//отправить ответ
	c.JSON(http.StatusOK, spellName)
}
