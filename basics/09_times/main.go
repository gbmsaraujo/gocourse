package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	p := fmt.Println
	var dataEspecifica time.Time

	p("A data e hora atual é:", truncateToSeconds(now))

	// 1. Carrega a localização de São Paulo (que representa o Horário de Brasília, UTC-3)
	localizacaoSP, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		fmt.Println("Erro ao carregar a localização:", err)
		return // Encerra o programa se não conseguir carregar o fuso
	}

	// 2. Criando Data específica

	dataEspecifica = time.Date(1993, time.February, 15, 11, 11, 32, 0, localizacaoSP)

	p("A data específica é:", dataEspecifica)

	// 3. Calculos e durações

	umaHora := time.Hour

	daquiAUmaHora := now.Add(umaHora)
	duasHorasAtrás := now.Add(umaHora * -2)

	p("Daqui a uma hora será:", truncateToSeconds(daquiAUmaHora))
	p("A duas horas atrás era: ", truncateToSeconds(duasHorasAtrás))

	// para saber a subtração entre horas

	compromissoDeAmanha := time.Date(2025, time.July, 30, 8, 00, 00, 0, localizacaoSP)

	p("Quanto falta para o compromisso de amanhã? ", compromissoDeAmanha.Sub(now).Truncate(time.Second))

	// Usando format

	// nosso molde :Mon Jan 2 15:04:05 MST 2006

	layoutBrasileiro := "02/01/2006 15:04:05"

	p("Agora usando format com layout Brasil: ", now.Format(layoutBrasileiro))

	//Parsing

	dateTimeDb := "07-29-2025T21:04:37"

	parsingDate,err := time.Parse("01-02-2006T15:04:05", dateTimeDb)

	if err != nil {
		p("Erro ao fazer parsing: ", err)
		return
	}

	p("O parsing é: ", parsingDate)
}

func truncateToSeconds(t time.Time) time.Time {
	return t.Truncate(time.Second)
}
