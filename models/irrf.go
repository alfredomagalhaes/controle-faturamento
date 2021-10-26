package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type IRRF_Tabela struct {
	Base
	DataInicial time.Time
	DataFinal   time.Time
	ValorLimite float32
	Faixas      []IRRF_Faixa `gorm:"foreignKey:IDTabelaIRRF"`
}

type IRRF_Faixa struct {
	Base
	Sequencia    int
	ValorInicial float32
	ValorFinal   float32
	Aliquota     float32
	IDTabelaIRRF uuid.UUID
}
