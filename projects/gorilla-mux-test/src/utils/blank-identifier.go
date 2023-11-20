package utils

import (
 	"net/http"
	"io/ioutil"
)

func TestingBlankIdentifier() string {
	var CrawlerResult string
	// _ é o blank identifier, ele ignora o valor
	// se nao usarmos a variavel, o compilador vai reclamar
	// mas se usarmos o blank identifier, ele ignora o valor
	// e nao reclama
	// se nao usarmos o blank identifier, o compilador vai reclamar

	// criando uma especie de web crawler improvisado para pegar o html do google
	// a funcao http.Get retorna dois valores, o primeiro é o resultado da requisicao 
	// e o segundo é um erro, como nao vamos usar o erro, usamos o blank identifier
	result, _ := http.Get("http://google.com")
	// o mesmo acontece com o ioutil.ReadAll, ele retorna dois valores, o primeiro é o
	// conteudo da pagina e o segundo é um erro, como nao vamos usar o erro, usamos o
	// blank identifier
	body, _ := ioutil.ReadAll(result.Body)
	// fechando o body da requisicao
	result.Body.Close()
	// atribuindo o conteudo da pagina para a variavel global CrawlerResult
	CrawlerResult = string(body)

	// retornando o conteudo da pagina
	return CrawlerResult
	
}