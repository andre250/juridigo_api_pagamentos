package config

import (
	"github.com/juridigo/juridigo_api_pagamentos/models"
)

/*
Devel - Responsável pode difinir confirgurações de ambiente de desenvolvimento
*/
func devel() {
	globaConfig = models.Config{
		App: models.App{
			Port:   "3041",
			Secret: "JUR1d1G00S3cr377",
		},
		Version: "0.0.1",
		Database: models.Database{
			Path:     "mongodb://<dbuser>:<dbpassword>@ds257314.mlab.com:57314/juridevel",
			User:     "juridigo",
			Password: "jur1digo",
			Database: "juridevel",
		},
	}
}
