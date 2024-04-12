package erro

import (
	"errors"

)

var (
	ErrNotFound 		= errors.New("Item não encontrado")
	ErrUnmarshal 		= errors.New("Erro na conversão do JSON")
	ErrUnauthorized 	= errors.New("Erro de autorização")
	ErrServer		 	= errors.New("Erro não identificado")
	ErrHTTPForbiden		= errors.New("Requisição recusada")
	ErrInvalidId		= errors.New("Id invalido para a pesquisa, deve ser um numerico")
)
