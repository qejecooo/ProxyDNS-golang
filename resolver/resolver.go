package resolver

import (
	"context"
	"log"
	"net"

	//"os"
	"strconv"

	"example.com/dns/utils"
)

func ResolveDomainName(domain_name string, ip_version string) (string, error) {
	config, _ := utils.ParseJson()

	upstream_dns := config.UpstreamServer.IP + ":" + strconv.Itoa(config.UpstreamServer.Port)

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("tcp", upstream_dns)
		},
	}

	if config.BlackList[domain_name[:len(domain_name)-1]] != "" {
		log.Printf("Domain name is blacklisted\n")
		return config.BlackList[domain_name[:len(domain_name)-1]], nil
	}

	ips, err := resolver.LookupIP(context.Background(), ip_version, domain_name[:len(domain_name)-1])

	if err != nil {
		log.Printf("Error resolving domain name: %v\n", err)
		return "", err
	}

	return ips[0].String(), nil
}
