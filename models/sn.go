package models

import (
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

type SN_Tabela struct {
	Base
	DataInicial     time.Time  `json:"data_inicial"`
	DataFinal       time.Time  `json:"data_final"`
	PercentualFolha float64    `json:"target_folha"`
	Faixas          []SN_Faixa `gorm:"foreignKey:IDTabelaSN" json:"faixa_valores"`
}

type SN_Faixa struct {
	Base
	Sequencia      int       `json:"sequencia"`
	ValorInicial   float32   `json:"valor_inicial"`
	ValorFinal     float32   `json:"valor_final"`
	Aliquota       float32   `json:"aliquota"`
	ParcelaDeducao float32   `json:"parcela_deducao"`
	IDTabelaSN     uuid.UUID `json:"ID_tabela_SN"`
}

func (s *SN_Tabela) CriarTabelaSN() error {
	result := config.MI.DB.Create(&s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ObterTabelaSN(id uuid.UUID) (SN_Tabela, error) {

	ret := SN_Tabela{}

	//pesq := SN_Tabela{}
	//pesq.ID = id
	//config.MI.DB.Where(pesq).Find(&ret)
	x := config.MI.DB.First(&ret, id)

	return ret, x.Error
}
