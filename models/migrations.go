package models

import (
	"log"

	"github.com/alfredomagalhaes/controle-faturamento/config"
)

//Função para criar todas as tabelas a serem usadas no banco de dados
func InicializarTabelas() {

	log.Println("Iniciando criação das tabelas no banco de dados")

	config.MI.DB.AutoMigrate(&User{})
	config.MI.DB.AutoMigrate(&SN_Faixa{})
	config.MI.DB.AutoMigrate(&SN_Tabela{})
	config.MI.DB.AutoMigrate(&INSS_Faixa{})
	config.MI.DB.AutoMigrate(&INSS_Tabela{})
	config.MI.DB.AutoMigrate(&IRRF_Faixa{})
	config.MI.DB.AutoMigrate(&IRRF_Tabela{})
	config.MI.DB.AutoMigrate(&Faturamento{})
	config.MI.DB.AutoMigrate(&Fechamento{})

	log.Println("Finalizando criação das tabelas no banco de dados")

}
