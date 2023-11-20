package utils

// contantes sao variaveis que nao mudam de valor 
// muito comum usar constantes para definir valores padroes
// mesmo conceito de constantes em outras linguagens

func TestingConstants() string {
	// definindo uma constante
	const stringTest string = "Constante de teste"

	// retornando o conteudo da pagina
	return stringTest
}

func TestingMultipleContants() []string {
	// é possivel definir varias constantes de uma vez só
	// ao definir os nomes das constantes em go, devemos lembrar que nao é possivel
	// escrever o nome da constante com todas as letras maiusculas, pois isso
	// significa que a constante é publica, e isso pode gerar conflito com outras
	// bibliotecas
	// tente definir como "metodoPrivado" ou "MetodoPublico". Sempre camelCase para privado e PascalCase para publico
	const (
		stringTest string = "Test"
		stringTest2 string = "Test2"
		stringTest3 string = "Test3"
	)
	return  []string{stringTest, stringTest2, stringTest3}
}