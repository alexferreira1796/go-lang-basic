package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoringInterval = 3
const delay = 5
const fileNameSite = "sites.txt"
const fileNameLogs = "logs.txt"

func main() {
	showIntroduction()

	for {
		showMenu()
	
		command := readCommand()
		switch command {
			case 1:
				monitoring() 
			case 2:
				logs()
			case 0:
				exit()
			default:
				fmt.Println("Não conheço esse comando")
				os.Exit(-1)
		}
	}	
}

func returnNameVersion() (string, float64) {
	name := "Alex"
	version := 1.0
	return name, version
}

func showIntroduction() {
	_, version := returnNameVersion()
	// fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var command int
	// fmt.Scanf("%d", &command) // Preciso passar o modificador
	fmt.Scan(&command)
	fmt.Println("Tipo da variavel é", reflect.TypeOf(command))
	return command

	// & comercial => Espero um digito e vou colocar na variavel (& ta esperando esse lugar, endereço (ponteiro))
}

func monitoring() {
	fmt.Println("Monitorando...")

	// webSite := []string{"https://httpbin.org/status/400", "http://alura.com.br", "http://google.com.br", "http://youtube.com.br"}

	webSite := openFile()
	
	for i := 0; i < monitoringInterval; i++ {
		for i, site := range webSite {
			testWebSite(i, site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")

	// for i := 0; i < len(webSite); i++ {
	// 	site := webSite[i]
	// 	fmt.Println(site)
	// }

	// var i int
	// for i = 0; i < 10; i++ {
	// 	fmt.Println("Tipo ", reflect.TypeOf(webSite))
	// }

	// for i := 0; i < monitoringInterval; i++ {
	// 	var i int
	// 	var site string
	// 	for i, site = range webSite {
	// 		testWebSite(i, site)
	// 	}
	// 	time.Sleep(delay * time.Second)
	// }
	
}

func logs() {
	fmt.Println("Exibindo logs...")

	file, err := ioutil.ReadFile(fileNameLogs)
	showError(err)

	fmt.Printf(string(file))
}

func exit() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func testWebSite(indice int, site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro no site: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: [",indice,"]", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site: [",indice,"]", site, "- Erro:", resp.Status)
		registerLog(site, false)
	}
}

func openFile() []string {
	var sites []string

	// file, err := os.Open("sites.txt")
	// file, err := ioutil.ReadFile("sites.txt")
	file, err := os.Open(fileNameSite)
	showError(err)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line) // Limpa a quebra linha
		
		sites = append(sites, string(line))

		if err == io.EOF {
			break
		}
	
	}
	file.Close() // Boa prática, fechar o arquivo
	return sites
}

func showError(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
}

func registerLog(site string, status bool) {
	
	flags := os.O_RDWR | os.O_CREATE | os.O_APPEND

	file, err := os.OpenFile(fileNameLogs, flags, 0666)
	showError(err)

	now := getTimeNow()
	writeText := now + " - " + site + " - online: " + strconv.FormatBool(status) + "\n"
	file.WriteString(writeText)

	file.Close()	
}

func getTimeNow() string {
	timeNow := time.Now().Format("02/01/2006 15:04:05") // https://go.dev/src/time/format.go
	return timeNow
}

func exibeNomes() []string {
	// Arrays
	var webSite [4]string // Array é obrigátorio especificar o tamanho [4] string
	webSite[0] = "https://httpbin.org/status/200"
	webSite[1] = "http://alura.com.br"
	webSite[2] = "http://google.com.br"
	webSite[3] = "http://youtube.com.br"

	// Slices -> Abstração do array
	nomes := []string{"Alex", "Renato", "Luciano"}
	fmt.Println("Total:", len(nomes))
	
	nomes = append(nomes, "Aparecida") // Slice dobra a quantidade quando é adicionado mais um item
	fmt.Println("Capacidade:", cap(nomes))
	fmt.Println(len(nomes))

	return nomes
}