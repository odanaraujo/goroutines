package main

import (
	"fmt"
	"time"
)

func speak(pessoa string, text string, qtdSpeak int) {
	for i := 0; i < qtdSpeak; i++ {
		fmt.Printf("%s %s (interation %d)\n", pessoa, text, i+1)
	}
}

func main() {
	//speak("Maria", "por que você não fala comigo?", 3)
	//speak("João", "Só posso falar depois de você", 2)

	go speak("Maria", "ei...", 3)
	go speak("João", "oi...", 2)
	time.Sleep(time.Second * 2)
}
