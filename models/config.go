package models

/*
App - Modelo para inicialização do App
*/
type App struct {
	Port   string `json:"port"`
	Secret string `json:"secret"`
}

/*
Database - Modelo para inicializaçao do Database
*/
type Database struct {
	Path     string `json:"path"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/*
Config - Model responsável por controlar configurações do Microsserviço
*/
type Config struct {
	Version  string   `json:"version"`
	App      App      `json:"app"`
	Database Database `json:"database"`
}
