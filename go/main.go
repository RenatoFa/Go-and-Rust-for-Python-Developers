package main

import "fmt"

func main() {
	banner()
	fizzbuzz()
	hw()
	media()
	Now()
	MostCommon()
	CharFrequency()
	println(CollatzStep(4))
	println(CollatzStep(5))
	root, extension := SplitExt("app.py")
	fmt.Println(root, extension)
	values := []float64{2, 4, 8}
	mean, err := Mean(values)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Println(mean)

	user := User{"elliot", Viewer}

	// No Go, os parâmetros são sempre passados por valor, ou seja, uma cópia do valor é feita.
	// Para permitir que uma função modifique o valor original, passamos um ponteiro para esse valor.
	// Para isso, usamos o operador "&" ao passar o argumento, que obtém o endereço da variável,
	// e na função que recebe o ponteiro, usamos o operador "*" para acessar o valor ao qual o ponteiro aponta.
	Promoter(&user, Admin)

	fmt.Printf("%#v\n", user)
}
