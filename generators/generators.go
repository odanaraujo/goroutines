package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// a generators encapsula uma camada de goroutines dentro dela.
// google i/o 2012 - concurrency patterns

// chan <- channel somente-leitura

func title(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, err := io.ReadAll(resp.Body)

			if err != nil {
				fmt.Printf("erro ao ler o corpo da resposta %s: %s\n", url, err)
				return
			}

			r, err := regexp.Compile("<title>(.*?)<\\/title>")
			if err != nil {
				fmt.Printf("Erro ao compilar a expressão regular: %s\n", err)
				return
			}

			ch <- r.FindStringSubmatch(string(html))[1]

		}(url)
	}

	return ch
}

func main() {
	t1 := title("https://www.youtube.com", "https://www.google.com")
	t2 := title("https://www.mercadolivre.com.br", "https://www.cod3r.com.br/")

	fmt.Printf("first: %s | %s\n", <-t1, <-t2)
	fmt.Printf("seconds: %s | %s", <-t1, <-t2)
}
