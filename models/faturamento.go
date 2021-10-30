package models

import (
	"errors"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

type Faturamento struct {
	Base
	Referencia    string  `json:"referencia" gorm:"type:varchar(6);index:idx_referencia"`
	ValorFaturado float32 `json:"valor_faturado"`
}

var ErrReferenciaJaCadastrada = errors.New("referencia já existe no banco de dados")

func (f *Faturamento) InserirFaturamento() error {

	pesq := Faturamento{}

	config.MI.DB.Where("referencia = ?", f.Referencia).First(&pesq)

	if pesq.Referencia != "" {
		return ErrReferenciaJaCadastrada
	}

	//TODO - Inserir validações
	result := config.MI.DB.Create(&f)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ObterFaturamento(id uuid.UUID) (Faturamento, error) {

	ret := Faturamento{}

	x := config.MI.DB.First(&ret, id)

	return ret, x.Error
}

func ObterTodosFaturamentos() ([]Faturamento, error) {

	ret := []Faturamento{}

	x := config.MI.DB.Find(&ret)

	return ret, x.Error
}

func (f *Faturamento) AtualizarFaturamento() error {

	//TODO - Inserir validações
	if f.Referencia != "" {
		result := config.MI.DB.Model(&f).Update("referencia", &f.Referencia)

		if result.Error != nil {
			return result.Error
		}
	}
	result := config.MI.DB.Model(&f).Update("valor_faturado", &f.ValorFaturado)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ApagarFaturamento(id uuid.UUID) error {

	err := config.MI.DB.Where("id = ?", id.String()).Delete(&Faturamento{})

	return err.Error
}
