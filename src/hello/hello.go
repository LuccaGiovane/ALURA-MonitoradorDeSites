package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {

		exibeMenu()
		comando := leComando()

		switch comando {

		case 1:

			iniciarMonitoramento()

		case 2:

			fmt.Println("Exibindo Logs...")
			imprimeLogs()

		case 0:

			fmt.Println("Saindo...")
			os.Exit(0)

		default:

			fmt.Println("Opção inválida!")
			os.Exit(-1)

		}

	}

}

func exibeIntroducao() {

	nome := "Lucca"
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("[1] - Iniciar Monitoramento.")
	fmt.Println("[2] - Exibir Logs.")
	fmt.Println("[0] - Sair do programa.")
}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	/*sites := []string{ //slice -> tipo um vetor soq melhor
		"https://random-status-code.herokuapp.com/",
		"https://quizlet.com/explanations/textbook-solutions/calculus-early-transcendentals-10th-edition-9780470647691",
		"https://vdocuments.mx/calculo-vol-2-12-ed-george-thomas-solutions.html",
		"https://sistemas.uel.br/portaldoestudante/index",
	}*/

	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Site[", i+1, "] ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(" Status: Online")
		registraLog(site, true)
	} else {
		fmt.Println("Status: Offline.")
		fmt.Println("Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	fmt.Println(sites)
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
