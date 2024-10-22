package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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
	data, err := fileHead("/Users/renato.asiss/Desktop/Captura.png", 8)

	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println(data)
	}

	data_type, err := FileType("/Users/renato.asiss/Desktop/Captura.png")

	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println(data_type)
	}

	location, err := NewLocation(32.5252837, 34.9427434)

	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Printf("%#v\n", location)

	location.Move(32.0641339, 34.8742343)

	fmt.Printf("%#v\n", location)

	car, err := NewCar("gopher", 32.5253837, 34.9427434)

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	car.Move(32.0641339, 34.8742343)

	fmt.Printf("%#v\n", car)

	items := []Mover{&Location{32.3477669, 34.9160405}, &Car{ID: "gopher",
		Location: Location{Lat: 32.5253837, Lng: 34.9427434}}}

	moveAll(items, 32.0641334, 34.8742343)

	for _, item := range items {
		fmt.Printf("%#v\n", item)
	}

	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.0, 2.0))
	fmt.Println(Add("G", "O"))

	fmt.Println(AddTwo(1, 2))
	fmt.Println(AddTwo(1.0, 2.0))
	fmt.Println(AddTwo("O", "I"))

	vms := []Coster{
		&VM{
			StartTime: time.Date(2022, time.April, 12, 17, 30, 0, 0, time.UTC),
			EndTime:   time.Date(2022, time.April, 12, 19, 54, 0, 0, time.UTC),
		},
		Spot{
			VM: VM{
				StartTime: time.Date(2022, time.April, 13, 9, 46, 0, 0, time.UTC),
				EndTime:   time.Date(2022, time.April, 15, 12, 7, 0, 0, time.UTC),
			},
		},
	}

	fmt.Println((TotalCost(vms)))

	for i := 0; i < 5; i++ {
		go worker(i)
	}
	fmt.Print("main")
	time.Sleep(time.Second)
	chain()

	wait_group()

	url := "https://www.linkedin.com/learning"

	CheckUrl(url)

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error: %s", err)
	}
}
