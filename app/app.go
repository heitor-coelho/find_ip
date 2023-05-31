package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		}}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na Internet",
			Flags:  flags,
			Action: buscaIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na Internet",
			Flags:  flags,
			Action: buscaServidores,
		},
	}
	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		fmt.Println("Não foi possível encontrar o endereço: ", host)
	} else {
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		fmt.Println("Não foi possível encontrar o servidor: ", host)
	} else {
		for _, servidor := range servidores {
			fmt.Println(servidor.Host)
		}
	}
}
