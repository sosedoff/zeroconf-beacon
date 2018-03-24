package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/grandcat/zeroconf"
)

var (
	name     string
	protocol string
	domain   string
	port     int
	txt      string
)

func init() {
	flag.StringVar(&name, "name", "", "Service name")
	flag.StringVar(&protocol, "protocol", "_http._tcp", "Service protocol")
	flag.StringVar(&domain, "domain", "local.", "Service domain")
	flag.IntVar(&port, "port", 80, "Service port")
	flag.StringVar(&txt, "txt", "", "Service text records: foo=bar,hello=world")
	flag.Parse()

	if name == "" {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println("cant get hostname:", err)
		}
		if hostname == "" {
			log.Fatal("host is not set")
		}
		name = hostname
	}
}

func main() {
	log.Printf(
		"starting zeroconf service name=%v protocol=%v domain=%v port=%v\n",
		name, protocol, domain, port,
	)
	defer log.Println("stopping zeroconf service")

	server, err := zeroconf.Register(
		name,
		protocol,
		domain,
		port,
		strings.Split(txt, ","),
		nil,
	)
	if err != nil {
		log.Fatal("zeroconf server error:", err)
	}
	defer server.Shutdown()

	stop := make(chan bool)
	<-stop
}
