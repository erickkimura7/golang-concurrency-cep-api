# Desafio Go Api

Duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como executar programa

1. Clonar repositorio
```
git clone https://github.com/erickkimura7/golang-concurrency-cep-api
```
2. Entrar na pasta
```
cd golang-concurrency-cep-api
```
3. Rodar o arquivo main.go com o CEP como parametro
```
go run main.go XXXXX-XXX(CEP)
```
