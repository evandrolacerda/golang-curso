package main

import (
	"curso-golang-2/controllers/livro"

	"github.com/gin-gonic/gin"
)

func main() {

	roteador := gin.Default()

	roteador.GET("/teste", func(contexto *gin.Context) {

		contexto.JSON(200, gin.H{
			"mensagem": "Olá mundo!",
		})
	})

	roteador.POST("/livros", livro.CriarLivro)
	roteador.GET("/livros", livro.BuscarLivros)
	roteador.GET("/livros/:id", livro.BuscarLivro)
	roteador.PUT("/livros/:id", livro.AtualizarLivro)

	roteador.Run(":8040")

}
