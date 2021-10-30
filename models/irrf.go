package models

import (
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

type IRRF_Tabela struct {
	Base
	DataInicial       time.Time    `json:"data_inicial"`
	DataFinal         time.Time    `json:"data_final"`
	DeducaoDependente float64      `json:"deducao_depente" gorm:"type:decimal(10,2)"`
	Faixas            []IRRF_Faixa `gorm:"foreignKey:IDTabelaIRRF" json:"faixa_valores"`
}

type IRRF_Faixa struct {
	Base
	Sequencia      int       `json:"sequencia"`
	ValorInicial   float64   `json:"valor_inicial" gorm:"type:decimal(10,2)"`
	ValorFinal     float64   `json:"valor_final" gorm:"type:decimal(10,2)"`
	Aliquota       float64   `json:"aliquota" gorm:"type:decimal(10,2)"`
	ParcelaDeducao float64   `json:"parcela_deducao" gorm:"type:decimal(10,2)"`
	IDTabelaIRRF   uuid.UUID `json:"ID_tabela_IR"`
}

func (ir *IRRF_Tabela) CriarTabelaIRRF() error {
	//TODO - Inserir validações
	result := config.MI.DB.Create(&ir)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ir *IRRF_Tabela) AtualizarTabelaIR() error {

	//TODO - Inserir validações
	var err error

	//Verifica se foi informado faixa de valores
	if ir.Faixas != nil {
		err = validarSequenciaFaixaIR(ir)
	}

	if err != nil {
		return err
	}

	result := config.MI.DB.Save(&ir)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ObterTabelaIRRF(id uuid.UUID) (IRRF_Tabela, error) {

	ret := IRRF_Tabela{}

	x := config.MI.DB.First(&ret, id)

	return ret, x.Error
}

func ObterFaixasPorTabelaIR(id uuid.UUID) ([]IRRF_Faixa, error) {

	ret := []IRRF_Faixa{}

	x := config.MI.DB.Where("id_tabela_irrf = ?", id).Find(&ret)

	return ret, x.Error
}

func validarSequenciaFaixaIR(ir *IRRF_Tabela) error {

	//Busca as faixas atuais da tabela
	fxAtu, err := ObterFaixasPorTabelaIR(ir.ID)
	if err != nil {
		return err
	}

	for _, faixa := range fxAtu {
		for _, fxPar := range ir.Faixas {
			if (fxPar.ValorInicial >= faixa.ValorInicial && fxPar.ValorInicial <= faixa.ValorFinal) ||
				(fxPar.ValorFinal >= faixa.ValorInicial && fxPar.ValorFinal <= faixa.ValorFinal) {
				return ErrFaixaJaExist
			}
		}
	}

	return nil
}

func ApagarFaixasIR(id uuid.UUID) error {

	err := config.MI.DB.Where("id = ?", id.String()).Delete(&IRRF_Faixa{})

	return err.Error
}

func ObterTabelaIRVigente(r string) (IRRF_Tabela, error) {
	var ret IRRF_Tabela
	var faixas []IRRF_Faixa
	var referencia string

	diaFormat, err := time.Parse(formatoData, r+"01")
	if err != nil {
		return ret, err
	}
	referencia = diaFormat.Format(formatoDataDB)

	x := config.MI.DB.Where(" ? between data_inicial and data_final ", referencia).Find(&ret)
	if x.Error != nil {
		return IRRF_Tabela{}, x.Error
	}
	//Busca as faixas de valores referente
	x = config.MI.DB.Model(&IRRF_Faixa{}).Where("id_tabela_irrf = ? ", ret.ID.String()).Scan(&faixas)
	if x.Error != nil {
		return IRRF_Tabela{}, x.Error
	}

	ret.Faixas = faixas
	return ret, nil
}
