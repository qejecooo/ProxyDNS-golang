# <b>DNSProxy</b><br/><br/>

This is a DNS Proxy server that supports black list and written in GoLang using "github.com/miekg/dns" library.
This proxy server supports both TCP and UDP. 
It supports package types such as:<br/>
<b>A</b> - IPv4 address<br/>
<b>AAAA</b> - IPv6 address<br/>

# <b>How does it work?</b><br/>
<br/>
User sends a DNS request to the proxy server.<br/><br/>
This request is handled by handler, where the request is checked against the black list.<br/>
If it is, the second check is made to see if we need to resolve the request to a different IP address.<br/><br/>
If the IP address where to resolve is in config file, we return the response packet with the new IP address.<br/>
If not, we return the response packet with no IP address.<br/>
If domain in not in black list, we send the request to the upstream DNS Server and return the response packet.<br/>
<br/>

The first thing to do is to specify upstream DNS server in config file.<br/>
Then you can specify black list file in config file(if needed).<br/>
<br/>

To run the server: 
<pre>
go build -o *app name*
./*app name* -p *port you want to use*
</pre>

In this particular case the server will listen on all addresses,
upstreaming the queries to configfile upstream server address. Only UDP datagrams will be accepted.<br/>

# <b>Flags:</b><br/>
<b>-p, --port</b>: Local proxy port (default: 53)<br/>
<b>-a, --address</b>: Local proxy address (default: all)<br/>
<b>-t, --protocol</b>: Protocol to use (default: udp)<br/>

# <b>How to test your DNSProxy server?</b><br/>
You can use dig command to test your DNSProxy server.<br/>
For example:<br/>
<b>dig @<your_dns_server_ip> -p <your_dns_port> google.com</b><br/>

Or change your DNS server to your DNSProxy server IP address and port in your network settings.<br/>
Then you can test it by opening your browser and going to google.com.<br/>
To do this open by text editor your <b>/etc/resolv.conf</b> file and add the following line:<br/>
<b>nameserver <your_dns_server_ip></b><br/>