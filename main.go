package main

import (
	"curso-golang-2/controllers/livro"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func logger() gin.HandlerFunc {
	return func(contexto *gin.Context) {
		ipDoCliente := contexto.ClientIP()
		agora := time.Now()
		metodo := contexto.Request.Method
		url := contexto.Request.URL.Path

		userAgent := contexto.Request.UserAgent()

		log.Printf("[%s] %s - %s %s - User Agent: %s", agora.Format("2006-01-02 15:04:05"), metodo, url, ipDoCliente, userAgent)

		contexto.Next()
	}
}

func checarJwt() gin.HandlerFunc {
	return func(contexto *gin.Context) {
		authorization := contexto.GetHeader("Authorization")

		godotenv.Load()

		jwtSecret := os.Getenv("JWT_SECRET")

		token, erro := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
			}

			return []byte(jwtSecret), nil
		})

		if erro != nil {
			contexto.JSON(401, gin.H{
				"erro": erro.Error(),
			})
			contexto.Abort()
		}

		fmt.Print(token)

		fmt.Print(authorization)
		fmt.Print(jwtSecret)
	}
}

func init() {
	godotenv.Load()
}

func main() {

	roteador := gin.Default()

	roteador.SetTrustedProxies([]string{"127.0.0.1"})

	roteador.Use(logger())

	roteador.GET("/teste", func(contexto *gin.Context) {

		contexto.JSON(200, gin.H{
			"mensagem": "Olá mundo!",
		})
	})

	rotasProjegidas := roteador.Group("/admin")

	rotasProjegidas.Use(checarJwt())

	rotasProjegidas.POST("/livros", livro.CriarLivro)
	rotasProjegidas.GET("/livros", livro.BuscarLivros)
	rotasProjegidas.GET("/livros/:id", livro.BuscarLivro)
	rotasProjegidas.PUT("/livros/:id", livro.AtualizarLivro)
	rotasProjegidas.DELETE("/livros/:id", livro.DeletarLivro)

	roteador.Run()

}
