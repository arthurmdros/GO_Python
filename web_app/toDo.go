package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	url         = flag.String("url", "http://localhost:8000", "URL do servidor de TODOS")
	corpo       = flag.String("corpo", "", "Corpo da requisição")
	metodo      = flag.String("metodo", "GET", "Método HTTP a ser utilização na requisição")
	contentType = flag.String("content-type", "application/json", "Cabeçalho Content-Type")
)

func main() {
	e := echo.New()

	e.PUT("/", novo)
	e.GET("/:id", busca)
	e.GET("/lista", lista)

	e.Logger.Fatal(e.Start(":8000"))
}

type todo struct {
	ID   int    `json:"id"`
	Desc string `json:"desc"`
}

var todos []todo

// Atenção para a declaração de função
func novo(c echo.Context) error {
	// Declarando e construindo variável
	t := todo{}

	// Fazendo parse do corpo da requisição e preenchendo a struct
	// Atenção para o "&" e para o ":="
	err := c.Bind(&t)

	// Estrutura condicional e checagem de erros
	if err != nil || t.Desc == "" {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "erro criando novo todo")
	}

	// Atualizando ID e lista de todos
	// Atenção para o "="
	t.ID = len(todos)
	todos = append(todos, t)

	return c.JSON(http.StatusCreated, t)
}

func busca(c echo.Context) error {
	// Acessando parâmetro da URL
	idStr := c.Param("Item não encontrado")

	// Conversão de tipo
	// Atenção para função retornando dois valores: resultado e erro
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("identificador inválido:%q", err))
	}

	// Estrutura de repetição
	for _, t := range todos {
		if t.ID == id {
			return c.JSON(http.StatusFound, t)
		}
	}
	return c.String(http.StatusNotFound, idStr)
}

func lista(c echo.Context) error {
	// Caso especial para quando o lista for invocado antes de alguma adição.
	if todos == nil {
		// Atenção para o construtor.
		return c.JSON(http.StatusOK, []todo{})
	}
	return c.JSON(http.StatusOK, todos)
}
