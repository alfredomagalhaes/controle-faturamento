package models

import (
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/config"
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
