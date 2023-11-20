package utils

import ("fmt")

// Go nao tem Private, Public, Protected (Modificadores de acesso), na verdade tem, mas nao
// eh como em outras linguagens. 
// o que temos é a visibilidade de pacotes. OU ele é visivel para o pacote ou nao é
// se o nome da funcao comeca com letra maiuscula, ela é publica
func Public()  {
	testingVisibilidade := "Public"
	fmt.Println(testingVisibilidade)
}
// se o nome da funcao comeca com letra minuscula, ela é privada
func private()  {
	testingVisibilidade := "Private"
	fmt.Println(testingVisibilidade)
}
// para usar a funcao privada em outro pacote, temos que usar a funcao publica
// que chama a funcao privada
func UsePrivate()  {
	private()
}
