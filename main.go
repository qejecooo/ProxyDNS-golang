package main

import (
	"log"
	"os"
	"strconv"

	"example.com/dns/handler"
	"github.com/akamensky/argparse"
	"github.com/miekg/dns"
)

func main() {
	parser := argparse.NewParser("dns", "DNS server")

	port := parser.Int("p", "port", &argparse.Options{Required: true, Help: "Port where the server will be running"})
	address := parser.String("a", "address", &argparse.Options{Required: false, Default: "127.0.0.1", Help: "Address where the server will be running"})
	protocol := parser.Selector("t", "protocol", []string{"udp", "tcp"}, &argparse.Options{Required: false, Default: "udp", Help: "Protocol to use"})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	dns.HandleFunc(".", handler.HandleDNS)
	server := &dns.Server{Addr: *address + ":" + strconv.Itoa(*port), Net: *protocol}
	log.Printf("Starting at %s, using %s", server.Addr, server.Net)

	os.Setenv("port", strconv.Itoa(*port))
	os.Setenv("address", *address)
	os.Setenv("protocol", *protocol)

	err = server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
