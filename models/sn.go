package models

import (
	"errors"
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

var ErrFaixaJaExist = errors.New("faixa de valor já cadastrada")

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
	ValorInicial   float64   `json:"valor_inicial" gorm:"type:decimal(10,2)"`
	ValorFinal     float64   `json:"valor_final" gorm:"type:decimal(10,2)"`
	Aliquota       float64   `json:"aliquota" gorm:"type:decimal(10,2)"`
	ParcelaDeducao float64   `json:"parcela_deducao" gorm:"type:decimal(10,2)"`
	IDTabelaSN     uuid.UUID `json:"ID_tabela_SN"`
}

func (s *SN_Tabela) CriarTabelaSN() error {
	//TODO - Inserir validações
	result := config.MI.DB.Create(&s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *SN_Tabela) AtualizarTabelaSN() error {

	//TODO - Inserir validações
	var err error

	//Verifica se foi informado faixa de valores
	if s.Faixas != nil {
		err = validarSequenciaFaixa(s)
	}

	if err != nil {
		return err
	}

	result := config.MI.DB.Save(&s)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ObterTabelaSN(id uuid.UUID) (SN_Tabela, error) {

	ret := SN_Tabela{}

	x := config.MI.DB.First(&ret, id)

	return ret, x.Error
}

func ObterFaixasPorTabelaSN(id uuid.UUID) ([]SN_Faixa, error) {

	ret := []SN_Faixa{}

	x := config.MI.DB.Where("id_tabela_sn = ?", id).Find(&ret)

	return ret, x.Error
}

func validarSequenciaFaixa(s *SN_Tabela) error {

	//Busca as faixas atuais da tabela
	fxAtu, err := ObterFaixasPorTabelaSN(s.ID)
	if err != nil {
		return err
	}

	for _, faixa := range fxAtu {
		for _, fxPar := range s.Faixas {
			if (fxPar.ValorInicial >= faixa.ValorInicial && fxPar.ValorInicial <= faixa.ValorFinal) ||
				(fxPar.ValorFinal >= faixa.ValorInicial && fxPar.ValorFinal <= faixa.ValorFinal) {
				return ErrFaixaJaExist
			}
		}
	}

	return nil
}

func ApagarFaixasSN(id uuid.UUID) error {

	err := config.MI.DB.Where("id = ?", id.String()).Delete(&SN_Faixa{})

	return err.Error
}

func ObterTabelaSNVigente(r string) (SN_Tabela, error) {
	var ret SN_Tabela
	var faixas []SN_Faixa
	var referencia string

	diaFormat, err := time.Parse(formatoData, r+"01")
	if err != nil {
		return ret, err
	}
	referencia = diaFormat.Format(formatoDataDB)

	x := config.MI.DB.Where(" ? between data_inicial and data_final ", referencia).Find(&ret)
	if x.Error != nil {
		return SN_Tabela{}, x.Error
	}
	//Busca as faixas de valores referente
	x = config.MI.DB.Model(&SN_Faixa{}).Where("id_tabela_sn = ? ", ret.ID.String()).Scan(&faixas)
	if x.Error != nil {
		return SN_Tabela{}, x.Error
	}

	ret.Faixas = faixas
	return ret, nil
}
