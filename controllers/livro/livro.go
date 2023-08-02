package livro

import (
	"curso-golang-2/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SalvarLivroRequisicao struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

type Livro struct {
	Id     int32  `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func CriarLivro(contexto *gin.Context) {
	var requisicao SalvarLivroRequisicao

	erro := contexto.ShouldBindJSON(&requisicao)

	if erro != nil {
		contexto.JSON(400, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	conexao, erro := database.Conectar()

	if erro != nil {
		contexto.JSON(500, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	defer conexao.Close()

	query := "INSERT INTO livros(titulo, autor) OUTPUT Inserted.ID VALUES( ?, ? )"

	resultado, erro := conexao.Query(query, requisicao.Titulo, requisicao.Autor)

	if erro != nil {
		contexto.JSON(500, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	for resultado.Next() {
		var id int32

		resultado.Scan(&id)

		fmt.Printf("ID Inserido = %d", id)

		livroResultado, erro := conexao.Query("SELECT id, titulo, autor FROM livros WHERE id = ?", id)

		if erro != nil {
			contexto.JSON(http.StatusInternalServerError, gin.H{
				"erro": erro.Error(),
			})
			return
		}

		var livroInserido Livro

		for livroResultado.Next() {
			livroResultado.Scan(&livroInserido.Id, &livroInserido.Titulo, &livroInserido.Autor)
		}

		contexto.JSON(http.StatusCreated, livroInserido)

	}

}

func BuscarLivros(contexto *gin.Context) {

	conexao, erro := database.Conectar()

	if erro != nil {
		contexto.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	defer conexao.Close()

	query := "SELECT * FROM livros"

	resultado, erro := conexao.Query(query)

	if erro != nil {
		contexto.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	var livros []Livro

	for resultado.Next() {
		var livro Livro

		resultado.Scan(&livro.Id, &livro.Titulo, &livro.Autor)

		livros = append(livros, livro)
	}

	contexto.JSON(http.StatusOK, livros)

}

func BuscarLivro(contexto *gin.Context) {

	id := contexto.Param("id")

	conexao, erro := database.Conectar()

	if erro != nil {
		contexto.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	defer conexao.Close()

	query := "SELECT TOP 1 id, autor, titulo FROM livros WHERE id = ?"

	resultado, _ := conexao.Query(query, id)

	var livro Livro

	if resultado.Next() {
		resultado.Scan(&livro.Id, &livro.Autor, &livro.Titulo)
		contexto.JSON(http.StatusOK, livro)

		return
	} else {
		contexto.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Livro não encontrado",
		})
		return
	}

}

func AtualizarLivro(contexto *gin.Context) {

	id := contexto.Param("id")

	conexao, erro := database.Conectar()

	if erro != nil {
		contexto.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	defer conexao.Close()

	query := "SELECT TOP 1 id, autor, titulo FROM livros WHERE id = ?"

	resultado, _ := conexao.Query(query, id)

	var livro Livro

	if resultado.Next() {
		resultado.Scan(&livro.Id, &livro.Autor, &livro.Titulo)

	} else {
		contexto.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Livro não encontrado",
		})
		return
	}

	queryUpdate := "UPDATE livros SET titulo = ?, autor = ? WHERE id = ?"

	var requisicao SalvarLivroRequisicao

	erro = contexto.ShouldBindJSON(&requisicao)

	if erro != nil {
		contexto.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	_, erro = conexao.Exec(queryUpdate, requisicao.Titulo, requisicao.Autor, id)

	if erro != nil {
		contexto.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})
		return
	}

	livro.Autor = requisicao.Autor
	livro.Titulo = requisicao.Titulo

	contexto.JSON(http.StatusOK, livro)

}
