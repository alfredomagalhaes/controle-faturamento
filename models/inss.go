package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type INSS_Tabela struct {
	Base
	DataInicial time.Time
	DataFinal   time.Time
	ValorLimite float32
	Faixas      []INSS_Faixa `gorm:"foreignKey:IDTabelaINSS"`
}

type INSS_Faixa struct {
	Base
	Sequencia    int
	ValorInicial float32
	ValorFinal   float32
	Aliquota     float32
	IDTabelaINSS uuid.UUID
}
