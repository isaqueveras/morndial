package exemple

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/isaqueveras/morndial"
)

func Teste() {
	configs := []morndial.Morndial{
		{Name: "", Url: "", Insecure: false, Timeout: time.Minute * 1},
	}

	morndial.NewService("Pendencias", "localhost:3333", true, time.Second*10)

	_ = morndial.Get(uuid.UUID{})

	for _, service := range configs {
		service.Timeout *= time.Minute
		if erro := service.NewConnection(); erro != nil {
			log.Println("Erro ao abrir conexão com o sistema: " + service.Name)
			continue
		}
	}
}
