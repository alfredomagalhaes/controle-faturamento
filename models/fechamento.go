package models

import (
	"errors"
	"fmt"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

var ErrReferenciaFechamentoJaCadastrado = errors.New("já existe um registro para essa referente e tipo no sistema")
var tiposFechamento = initTipo()

type Fechamento struct {
	Base
	Referencia    string  `json:"referencia" gorm:"type:varchar(6);index:idx_referencia"`
	Tipo          string  `json:"tipo" gorm:"type:varchar(1)"`
	ValorFaturado float32 `json:"valor_faturado"`
	AliquotaDAS   float32 `json:"aliquota_das"`
	ValorDAS      float32 `json:"valor_das"`
	AliquotaINSS  float32 `json:"aliquota_inss"`
	ValorINSS     float32 `json:"valor_inss"`
	AliquotaIRRF  float32 `json:"aliquota_irrf"`
	ValorIRRF     float32 `json:"valor_irff"`
}

func initTipo() map[string]string {

	tipos := make(map[string]string)
	tipos["P"] = "Previsto"
	tipos["R"] = "Realizado"

	return tipos
}

func (f *Fechamento) InserirFechamento() error {

	//Valida os tipos de fechamentos que podem ser incluidos
	if err := validarTipos(*f); err != nil {
		return err
	}
	//Valida os percentuais informado
	if err := validarPercentuais(*f); err != nil {
		return err
	}

	pesq := Fechamento{}

	config.MI.DB.Where("referencia = ? and tipo = ?", f.Referencia, f.Tipo).First(&pesq)

	if pesq.Referencia != "" {
		return ErrReferenciaFechamentoJaCadastrado
	}

	//TODO - Inserir validações
	result := config.MI.DB.Create(&f)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func validarTipos(f Fechamento) error {
	var msgErro string
	if tiposFechamento[f.Tipo] == "" {

		for c, v := range tiposFechamento {
			if msgErro != "" {
				msgErro += ", "
			}
			msgErro += fmt.Sprintf("%s - %s", c, v)

		}
		return fmt.Errorf("tipo de fechamento não suportado, utilize uma das opções: %s", msgErro)
	}
	return nil

}

func validarPercentuais(f Fechamento) error {
	//Verifica se algum dos percentuais é maior que 100%
	if f.AliquotaDAS > 100 {
		return errors.New("aliquota da DAS não pode ser maior que 100")
	}
	if f.AliquotaINSS > 100 {
		return errors.New("aliquota do INSS não pode ser maior que 100")
	}
	if f.AliquotaIRRF > 100 {
		return errors.New("aliquota do IR não pode ser maior que 100")
	}
	//Verifica se os percentuais são menores que zero
	if f.AliquotaDAS < 0 {
		return errors.New("aliquota da DAS não pode ser negativa")
	}
	if f.AliquotaINSS < 0 {
		return errors.New("aliquota do INSS não pode ser negativa")
	}
	if f.AliquotaIRRF < 0 {
		return errors.New("aliquota do IR não pode ser negativa")
	}

	return nil
}
func ObterFechamento(id uuid.UUID) (Fechamento, error) {

	ret := Fechamento{}

	x := config.MI.DB.First(&ret, id)

	return ret, x.Error
}

func ObterTodosFechamentos() ([]Fechamento, error) {

	ret := []Fechamento{}

	x := config.MI.DB.Find(&ret)

	return ret, x.Error
}

func (f *Fechamento) AtualizarFechamento() error {

	//TODO - Inserir validações
	//Valida os tipos de fechamentos que podem ser incluidos
	if err := validarTipos(*f); err != nil {
		return err
	}
	//Valida os percentuais informado
	if err := validarPercentuais(*f); err != nil {
		return err
	}
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

func ApagarFechamento(id uuid.UUID) error {

	err := config.MI.DB.Where("id = ?", id.String()).Delete(&Fechamento{})

	return err.Error
}
