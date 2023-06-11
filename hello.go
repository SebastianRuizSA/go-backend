package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	x := "hola" // Se usa cuando el tipo no esta especificado
	fmt.Println(x)
	read_json()
}

func read_json(){
	file, err := os.Open("data/dito.json")
	if err != nil{
		fmt.Println("ERROR", err)
		return
	}
	defer file.Close()

	fmt.Println("Archivo abierto \n ------")
	//scanner := bufio.NewScanner(file)

	//for scanner.Scan() {
		// fmt.Println(scanner.Text()) // Retorna todo el texo, esto no es bueno si quiero ver un JSON en estructura.
	//	text = scanner.Text()
	//}
	// Maps es lo mismo que diccionarios

	byteValue, _ := ioutil.ReadAll(file)

	var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

	json_keys := Keys(result)
	fmt.Println(json_keys)

    //fmt.Println(result["abilities"]) // Como defino una estructura(de formato JSON) si la desconozco???

}

func Keys(m map[string]interface{}) []string {
	//const largo = len(m) // Raro, len() puede entregar un valor constante o variable dependiendo de los inputs.
    //keys := make([]string, len(m))
	keys := make([]string, 0, len(m)) // Otra forma de escribir este codigo, el 0 es lenght inicial y len() el final.
    //i := 0
    //for k := range m {
    //    keys[i] = k
    //    i++
    //}
	for key := range m {
		keys = append(keys, key)
	}
    return keys
}