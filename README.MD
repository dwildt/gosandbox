Sites que uso de referência / reference websites: 
* https://golang.org/doc/install
* https://golang.org/doc/code.html#Workspaces
* https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go
* https://github.com/diegogaulke/go-workshop

Para iniciar um worskpace, lembrar de convenções (linux):
* workspace $HOME/go/src/github.com/dwildt
* artigo sobre workspaces https://golang.org/doc/code.html#Workspaces
* lembrar de: 
* export GOPATH=$HOME/go
* PATH=$PATH:$GOPATH/bin
* export PATH

No cloud9 precisei fazer estes "go get" para poder rodar os testes unitários:
* go get gotest.tools  
* go get github.com/google/go-cmp/cmp  
* go get github.com/pkg/errors  

Objetivo é contribuir com o repositório do Diego Gaulke, criando um capítulo sobre automação de testes / objective is to help Diego's repository, building a chapter about automated tests. 
* https://github.com/diegogaulke/go-workshop

Na pasta ch14-fizzbuzz se faz uso gotest.tools. Para instalar manualmente: 
* go get gotest.tools/v3
* ref https://github.com/gotestyourself/gotest.tools

Na pasta ch15-webfizzbuzz se faz uso do gin-gonic. 
* go get -github.com/gin-gonic/gin
* ref https://github.com/gin-gonic/gin
* exemplo de execução: /fb/5 -> buzz 

Sobre módulos. Para ativar uso de módulos, modificar o bash_profile (mac os no meu caso):
* export GO111MODULE=on  
