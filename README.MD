Sites que uso de referência / reference websites: 
* https://golang.org/doc/install
* https://golang.org/doc/code.html#Workspaces
* https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go
* https://github.com/diegogaulke/go-workshop

Para iniciar um worskpace, lembrar de convenções (linux):
* workspace $HOME/go/src/github.com/dwildt
* artigo sobre workspaces https://golang.org/doc/code.html#Workspaces
* lembrar do export export GOPATH=$HOME/go
* PATH=$PATH:$GOPATH/bin
* export PATH

Objetivo é contribuir com o repositório do Diego Gaulke, criando um capítulo sobre automação de testes / objective is to help Diego's repository, building a chapter about automated tests. 
* https://github.com/diegogaulke/go-workshop

Na pasta ch14-fizzbuzz se faz uso gotest.tools. Para instalar manualmente: 
* go get gotest.tools/v3
* ref https://github.com/gotestyourself/gotest.tools
  
Sobre módulos. Para ativar uso de módulos, modificar o bash_profile (mac os no meu caso):
* export GO111MODULE=on  
