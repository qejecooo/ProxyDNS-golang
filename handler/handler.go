package handler

import (
	"fmt"
	"log"

	"example.com/dns/resolver"

	//	"example.com/dns/utils"
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
	var dns_type string
	var ip_type string

	log.Printf("DNS query received")
	for _, q := range query.Question {
		log.Printf("Query for %s\n", q.Name)

		switch q.Qtype {
		case dns.TypeA:
			dns_type = "A"
			ip_type = "ip4"
		case dns.TypeAAAA:
			dns_type = "AAAA"
			ip_type = "ip6"
		}

		ip, _ := resolver.ResolveDomainName(q.Name, ip_type)
		if ip != "" {
			rr, err := dns.NewRR(fmt.Sprintf("%s %s %s", q.Name, dns_type, ip))
			if err == nil {
				query.Answer = append(query.Answer, rr)
			}
		}
	}
}
