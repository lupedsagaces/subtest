# subtest
Subtest é uma ferrameta de validação de subdomínios. 
O programa lê uma lista de subdomínios de um arquivo txt e verifica se o mesmo retorna status 200.
Subdomínios válidos são salvos como output.txt


## Como instalar a ferramenta:
```bash

go install github.com/lupedsagaces/subtest@latest
```
## Como funciona:
Ao rodar o comando subtest, o script pedirá o domínio principal: ex: redbull.com
Depois, é solicitado o caminho e nome do arquivo que contém os subdomínios.
Após isso, é apenas aguardar o script validar.
