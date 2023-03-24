package handler

import (
	"fmt"
	"log"

	"example.com/dns/resolver"
	"github.com/miekg/dns"
)

func HandleDNS(writer dns.ResponseWriter, reply *dns.Msg) {
	message := new(dns.Msg)
	message.SetReply(reply)
	message.Compress = false

	switch reply.Opcode {
	case dns.OpcodeQuery:
		ParseDNSQuery(message)
	}

	writer.WriteMsg(message)
}

func ParseDNSQuery(query *dns.Msg) {
	log.Printf("DNS query received")
	for _, q := range query.Question {
		log.Printf("Query for %s\n", q.Name)
		ip, _ := resolver.ResolveDomainName(q.Name)
		log.Printf("IP: %s\n", ip)
		switch q.Qtype {
		case dns.TypeA:
			ip := ip
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					query.Answer = append(query.Answer, rr)
				}
			}
		}
	}
}
