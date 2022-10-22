package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	templateUrlCdn    = "https://cdn.apicep.com/file/apicep/%s.json"
	templateUrlViaCep = "http://viacep.com.br/ws/%s/json/"
)

type Response struct {
	processTime  time.Time
	requestUrl   string
	responseBody string
}

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("Erro ao executar comando, utilizar como o exemplo seguinte:\ngo run main.go XXXXX-XXX")
	}

	cep := os.Args[1]

	match, _ := regexp.MatchString("^\\d{2}\\d{3}[-]\\d{3}$", cep)

	if !match {
		log.Fatalf("Erro no formato do cep, utilizar no seguinte formato:\nEx.: XXXXX-XXX")
	}

	ch := make(chan Response)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	startTime := time.Now()

	go apiCall(ctx, fmt.Sprintf(templateUrlViaCep, cep), ch)
	go apiCall(ctx, fmt.Sprintf(templateUrlCdn, cep), ch)

	select {
	case <-time.After(time.Second):
		log.Println("Tempo de 1s excedido, nenhuma Api retornou no tempo esperado")
	case response := <-ch:
		log.Println("Api Retornou em menos de 1 segundo")
		log.Printf("Url da api mais rapida: %s\n", response.requestUrl)
		log.Printf("Tempo de processamento foi de %d ms\n", response.processTime.Sub(startTime).Milliseconds())
		log.Printf("Dados da resposta: %s\n", response.responseBody)
	}
}

func apiCall(ctx context.Context, url string, response chan<- Response) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response <- Response{
		processTime:  time.Now(),
		requestUrl:   url,
		responseBody: string(b),
	}
}
