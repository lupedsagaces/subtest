package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)


func printBanner() {
	banner := `
            _     _            _   
           | |   | |          | |  
  ___ _   _| |__ | |_ ___  ___| |_ 
 / __| | | | '_ \| __/ _ \/ __| __|
 \__ \ |_| | |_) | ||  __/\__ \ |_ 
 |___/\__,_|_.__/ \__\___||___/\__|
                                   
                                   
    Subdomain test    
		By: lupedsagaces      
	`
	fmt.Println(banner)
}

func main() {
	printBanner()
	// Pergunta pelo domínio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o domínio principal: ")
	dominio, _ := reader.ReadString('\n')
	dominio = dominio[:len(dominio)-1] // Remove o '\n'

	// Pergunta pelo arquivo com a lista de subadomínios
	fmt.Print("Digite o caminho para o arquivo de subdomínios: ")
	caminhoArquivo, _ := reader.ReadString('\n')
	caminhoArquivo = caminhoArquivo[:len(caminhoArquivo)-1] // Remove o '\n'

	// Abre o arquivo de subdomínios
	file, err := os.Open(caminhoArquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Abre ou cria o arquivo output.txt para salvar subdomínios válidos
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Erro ao criar output.txt:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	scanner := bufio.NewScanner(file)

	// Variável para verificar se algo foi salvo
	algumSalvo := false

	// Testa cada subdomínio
	for scanner.Scan() {
		subdominio := scanner.Text()
		url := "http://" + subdominio + "." + dominio

		// Faz uma requisição HTTP com timeout
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("[+] Subdomínio %s não encontrado.\n", url)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("[+] SUBDOMÍNIO %s ENCONTRADO.\n", url)
			// Salva o subdomínio válido no arquivo
			_, err := writer.WriteString(subdominio + "." + dominio + "\n")
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo:", err)
			} else {
				algumSalvo = true
			}
		} else {
			fmt.Printf(" Subdomínio %s retornou status %d.\n", url, resp.StatusCode)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	// Mostra a mensagem se houver algo salvo
	if algumSalvo {
		fmt.Println("=== Subdomínios salvos em output.txt ===")
	}
}
