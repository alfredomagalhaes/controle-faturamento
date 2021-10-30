package models

import (
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

type INSS_Tabela struct {
	Base
	DataInicial time.Time    `json:"data_inicial"`
	DataFinal   time.Time    `json:"data_final"`
	ValorLimite float64      `json:"valor_limite" gorm:"type:decimal(10,2)"`
	Faixas      []INSS_Faixa `gorm:"foreignKey:IDTabelaINSS" json:"faixa_valores"`
}

type INSS_Faixa struct {
	Base
	Sequencia    int       `json:"sequencia"`
	ValorInicial float64   `json:"valor_inicial" gorm:"type:decimal(10,2)"`
	ValorFinal   float64   `json:"valor_final" gorm:"type:decimal(10,2)"`
	Aliquota     float64   `json:"aliquota" gorm:"type:decimal(10,2)"`
	IDTabelaINSS uuid.UUID `json:"ID_tabela_INSS"`
}

func (in *INSS_Tabela) CriarTabelaINSS() error {
	//TODO - Inserir validações
	result := config.MI.DB.Create(&in)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (in *INSS_Tabela) AtualizarTabelaINSS() error {

	//TODO - Inserir validações
	var err error

	//Verifica se foi informado faixa de valores
	if in.Faixas != nil {
		err = validarSequenciaFaixaINSS(in)
	}

	if err != nil {
		return err
	}

	result := config.MI.DB.Save(&in)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ObterTabelaINSS(id uuid.UUID) (INSS_Tabela, error) {

	ret := INSS_Tabela{}

	x := config.MI.DB.First(&ret, id)

	return ret, x.Error
}

func ObterFaixasPorTabelaINSS(id uuid.UUID) ([]INSS_Faixa, error) {

	ret := []INSS_Faixa{}

	x := config.MI.DB.Where("id_tabela_INSS = ?", id).Find(&ret)

	return ret, x.Error
}

func validarSequenciaFaixaINSS(in *INSS_Tabela) error {

	//Busca as faixas atuais da tabela
	fxAtu, err := ObterFaixasPorTabelaINSS(in.ID)
	if err != nil {
		return err
	}

	for _, faixa := range fxAtu {
		for _, fxPar := range in.Faixas {
			if (fxPar.ValorInicial >= faixa.ValorInicial && fxPar.ValorInicial <= faixa.ValorFinal) ||
				(fxPar.ValorFinal >= faixa.ValorInicial && fxPar.ValorFinal <= faixa.ValorFinal) {
				return ErrFaixaJaExist
			}
		}
	}

	return nil
}

func ApagarFaixasINSS(id uuid.UUID) error {

	err := config.MI.DB.Where("id = ?", id.String()).Delete(&INSS_Faixa{})

	return err.Error
}

func ObterTabelaINSSVigente(r string) (INSS_Tabela, error) {
	var ret INSS_Tabela
	var faixas []INSS_Faixa
	var referencia string

	diaFormat, err := time.Parse(formatoData, r+"01")
	if err != nil {
		return ret, err
	}
	referencia = diaFormat.Format(formatoDataDB)

	x := config.MI.DB.Where(" ? between data_inicial and data_final ", referencia).Find(&ret)
	if x.Error != nil {
		return INSS_Tabela{}, x.Error
	}
	//Busca as faixas de valores referente
	x = config.MI.DB.Model(&INSS_Faixa{}).Where("id_tabela_inss = ? ", ret.ID.String()).Scan(&faixas)
	if x.Error != nil {
		return INSS_Tabela{}, x.Error
	}

	ret.Faixas = faixas
	return ret, nil
}
