package com
import (
  "github.com/miekg/dns"
)
//通过DNS查询域名的A记录，返回IP地址，可能是多个
func resolveFromDNS(domain string) ([]string, error) {
	answer := make([]string, 0)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	c := new(dns.Client)
	in, _, err := c.Exchange(m, dnsserver+":53")
	if err != nil {
		return answer, err
	}
	for _, ain := range in.Answer {
		if a, ok := ain.(*dns.A); ok {
			answer = append(answer, a.A.String())
		}
	}
	return answer, nil
}
