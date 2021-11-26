package models

import (
	"math"

	"gorm.io/gorm"
)

type Paginacao struct {
	Limite       int    `json:"limite,omitempty;query:limit"`
	Pagina       int    `json:"pagina,omitempty;query:page"`
	Ordem        string `json:"ordem,omitempty;query:sort"`
	TotalLinha   int64  `json:"total_linhas"`
	TotalPaginas int    `json:"total_paginas"`
}

func (p *Paginacao) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Paginacao) GetLimit() int {
	if p.Limite == 0 {
		p.Limite = 10
	}
	return p.Limite
}

func (p *Paginacao) GetPage() int {
	if p.Pagina == 0 {
		p.Pagina = 1
	}
	if p.Pagina > 100 {
		p.Pagina = 100
	}
	return p.Pagina
}

func (p *Paginacao) GetSort() string {
	if p.Ordem == "" {
		p.Ordem = "Id desc"
	}
	return p.Ordem
}

func Paginar(value interface{}, Paginacao *Paginacao, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	Paginacao.TotalLinha = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(Paginacao.Limite)))
	Paginacao.TotalPaginas = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(Paginacao.GetOffset()).Limit(Paginacao.GetLimit()).Order(Paginacao.GetSort())
	}
}
