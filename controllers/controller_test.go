package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup das rotas para teste
func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.TestMode)
	rotas := gin.Default()
	return rotas
}

func TestSaudacao(t *testing.T) {
	// Arrange
	r := SetupRotasDeTeste()
	r.GET("/:nome", Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()

	// Act
	r.ServeHTTP(resposta, req)

	// Assert
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"message":"Olá, seja bem-vindo(a) gui"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}
