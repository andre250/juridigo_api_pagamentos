package helpers

import (
	"github.com/juridigo/juridigo_api_pagamentos/config"
	"github.com/juridigo/juridigo_api_pagamentos/models"
)

var configuration models.Config

/*
InitiConfig - Inicializador de configurações
*/
func InitConfig() {

	configuration = config.GetConfig()
}
