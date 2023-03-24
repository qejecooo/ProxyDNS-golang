package resolver

import (
	"context"
	"log"
	"net"

	"example.com/dns/utils"
)

func ResolveDomainName(domain_name string) (string, error) {
	config, _ := utils.ParseJson()
	log.Printf("Upstream DNS: %s:%d\n", config.UpstreamServer.IP, config.UpstreamServer.Port)
	upstream_dns := config.UpstreamServer.IP + ":" + string(config.UpstreamServer.Port)

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial(protocol, upstream_dns)
		},
	}

	ips, err := resolver.LookupHost(context.Background(), domain_name)
	if err != nil {
		log.Printf("Error resolving domain name: %v\n", err)
		return "", err
	}

	return ips[0], nil
}
