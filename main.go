package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

//go:embed castores/*
var embedFs embed.FS

var ipMap = make(map[string]int)
var totalCastores int

func main() {
	var err error

	totalCastores, err = ContarTotalFotosCastores(embedFs)

	log.Printf("Total de castores carregados: %d\n", totalCastores)

	if err != nil {
		log.Fatal(err)
		return
	}

	r := gin.Default()
	r.StaticFS("/castores", http.FS(embedFs))

	r.GET("", func(c *gin.Context) {
		ip := c.ClientIP()

		var castorAtual = ipMap[ip]

		if castorAtual > totalCastores {
			castorAtual = 0
		}

		log.Printf("IP %s acessou o castor #%d\n", ip, castorAtual)

		ipMap[ip] = castorAtual + 1

		var img = fmt.Sprintf("castores/%d.jpg", castorAtual)

		c.FileFromFS(img, http.FS(embedFs))
	})

	r.Run(":9001")
}

func ContarTotalFotosCastores(embedFs embed.FS) (total int, err error) {
	dir, err := fs.ReadDir(embedFs, "castores")
	if err != nil {
		log.Fatal("Erro ao abrir a pasta de castores: ", err)
		return
	}

	var nomes = NomesOrdenados(dir)

	if len(nomes) == 0 {
		err = fmt.Errorf("nenhuma foto de castor encontrada")
		return
	}

	fmt.Sscanf(nomes[len(nomes)-1], "%d.jpg", &total)

	return
}

func NomesOrdenados(entries []fs.DirEntry) []string {
	var names = make([]string, len(entries))

	for i, entry := range entries {
		names[i] = entry.Name()
	}

	slices.SortFunc(names, func(a, b string) int {
		var ai, bi int
		fmt.Sscanf(a, "%d.jpg", &ai)
		fmt.Sscanf(b, "%d.jpg", &bi)
		return ai - bi
	})

	return names
}
