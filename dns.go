package gobase

import (
	"context"
	"net"
	"net/http"
	"time"
)

func OverrideDns(ip string) {

	var (
		resolver_port  = "53"
		resolver_ip    = ip + ":" + resolver_port
		resolver_proto = "udp"

		timeout_ms = time.Millisecond * 500
	)

	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: false,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: timeout_ms,
				}
				return d.DialContext(ctx, resolver_proto, resolver_ip)
			},
		},
	}

	dialContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.DialContext(ctx, network, addr)
	}

	http.DefaultTransport.(*http.Transport).DialContext = dialContext

}
