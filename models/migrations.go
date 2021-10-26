package models

import "github.com/alfredomagalhaes/controle-faturamento/config"

//Função para criar todas as tabelas a serem usadas no banco de dados
func InicializarTabelas() {
	config.MI.DB.AutoMigrate(&User{})
}
