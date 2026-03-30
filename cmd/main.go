package main

import(
	"fmt"
	"flag"
	"log"
	"go_storage/internal/models"
	"go_storage/internal/parser"
	"go_storage/internal/storage"
)

func main(){
	// rep := storage.NewRepository[*models.Product]()

		fileNamePTR := flag.String("file", "data.txt", "Передайте адрес data файла")

	flag.Parse()
	// используется для того чтобы компилятор прочитал переданный флаг

	currentFile := *fileNamePTR

	var repos storage.productStorage = storage.

	fmt.Printf("Система настроена на чтение из файла: %s", currentFile)
	fmt.Printf("Чтение из файла (%s)", path)

	Products, err := parser.ParseProductsFromFile(path)

	if err != nil{
		log.Fatalf("Fatal Error: %v\n", err)
	}

	for _, product := range Products{
		rep.Add(product)
	}

	items := rep.GetAll()

	fmt.Printf("На склад отгружено %d товаров \n",len(items))

	for _, item := range items{
		fmt.Printf("-Name: %10s | SBIN: %20s | Годен до %s \n", item.Name, item.SBIN, item.DateToString())
	}

	fmt.Scanln()
}
