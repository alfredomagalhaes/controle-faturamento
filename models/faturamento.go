package models

import (
	"errors"
	"sort"
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	uuid "github.com/satori/go.uuid"
)

type Faturamento struct {
	Base
	Referencia    string  `json:"referencia" gorm:"type:varchar(6);index:idx_referencia"`
	ValorFaturado float64 `json:"valor_faturado"`
}

var ErrReferenciaJaCadastrada = errors.New("referencia já existe no banco de dados")
var formatoData string = "20060102"
var formatoDataDB string = "2006-01-02"

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

	if x.Error == nil {
		sort.Slice(ret, func(i, j int) bool {
			return ret[i].Referencia > ret[j].Referencia
		})
	}
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

//Função para somar os faturamentos anteriores
//com base nos parâmetros de referencia (r) e delta (d) de meses anteriores
func SomarFaturamentosAnteriores(r string, d int) (float64, float64, error) {
	var somaFat float64
	var quantFat float64
	var anoMesInicial string
	var anoMesFinal string

	diaFormat, err := time.Parse(formatoData, r+"01")
	if err != nil {
		return 0, 0, err
	}
	diaFormat = diaFormat.AddDate(0, -d, 0)
	anoMesInicial = diaFormat.Format(formatoData)[:6]
	anoMesFinal = r

	rows, err := config.MI.DB.Model(&Faturamento{}).Select("sum(valor_faturado) as total, count(valor_faturado) as qtd").Where("referencia between ? and ?", anoMesInicial, anoMesFinal).Rows()
	if err != nil {
		return 0, 0, err
	}
	for rows.Next() {
		err := rows.Scan(&somaFat, &quantFat)
		if err != nil {
			return 0, 0, err
		}
	}

	return somaFat, quantFat, nil
}

func ObterFaturamentoMes(r string) (Faturamento, error) {

	fatRet := Faturamento{}

	retGr := config.MI.DB.Where("referencia = ?", r).First(&fatRet)

	if retGr.Error != nil {
		return Faturamento{}, retGr.Error
	}

	return fatRet, nil

}
