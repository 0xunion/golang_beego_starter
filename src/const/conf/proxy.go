package conf

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

	"github.com/0xunion/exercise_back/src/routine"
)

type Proxy struct {
	Address  string
	Port     int
	Protocol string
	UseWay   string
}

const (
	// use for multi thread task
	PROXY_USEWAY_DISTRIBUTION = "distribution"
	// use for network aboard
	PROXY_USEWAY_ABOART = "aboard"
)

var (
	// Proxy
	proxyEnable = false
	proxies     = []Proxy{}
)

func init() {
	// Proxy
	proxyEnable = getConf("proxyenable") == "true"
	if !proxyEnable {
		return
	}

	proxies_list_str := getConf("proxies")
	if proxies_list_str == "" {
		return
	}

	proxies_list := strings.Split(proxies_list_str, ",")
	for _, proxy_str := range proxies_list {
		proxy := strings.Split(proxy_str, ":")
		if len(proxy) != 4 {
			routine.Panic("[Proxy] Proxy setting syntax error: " + proxy_str)
		}

		// https://xxx.xxx:xx:a
		port, err := strconv.Atoi(proxy[2])
		if err != nil {
			routine.Panic("[Proxy] Proxy port is not a number: " + proxy[1])
		}

		protocol := proxy[0]
		if protocol != "http" && protocol != "https" && protocol != "socks5" {
			routine.Panic("[Proxy] Proxy protocol is not support: " + protocol)
		}

		useway := proxy[3]
		switch useway {
		case "a":
			useway = PROXY_USEWAY_ABOART
		case "d":
			useway = PROXY_USEWAY_DISTRIBUTION
		default:
			routine.Panic("[Proxy] Proxy useway is not support: " + useway)
		}

		address := strings.TrimLeft(proxy[1], "/")
		proxies = append(proxies, Proxy{
			Address:  address,
			Port:     port,
			Protocol: protocol,
			UseWay:   useway,
		})
	}
}

func getSupportProxy(protocol []string, useway string) []Proxy {
	var supportProxies []Proxy
	for _, proxy := range proxies {
		for _, p := range protocol {
			if proxy.Protocol == p && proxy.UseWay == useway {
				supportProxies = append(supportProxies, proxy)
			}
		}
	}
	return supportProxies
}

func GetAboardProxy(protocol []string) []Proxy {
	return getSupportProxy(protocol, PROXY_USEWAY_ABOART)
}

func GetDistributionProxy(protocol []string) []Proxy {
	return getSupportProxy(protocol, PROXY_USEWAY_DISTRIBUTION)
}

type proxyHolder struct {
	current_idx int
}

func NewProxyHolder() *proxyHolder {
	return &proxyHolder{
		current_idx: -1,
	}
}

func (p *proxyHolder) NextProxy(proxies []Proxy) Proxy {
	if p.current_idx == -1 {
		p.current_idx = rand.Intn(len(proxies))
		return proxies[p.current_idx]
	}

	p.current_idx += 1
	if p.current_idx >= len(proxies) {
		p.current_idx = 0
	}

	return proxies[p.current_idx]
}

// WalkStartWithRandom walk through all proxies with random start
// if cb return error, the walk will continue until all proxies are walked
// if all proxies failed, WalkStartWithRandom will return an error
func (p *proxyHolder) WalkStartWithRandom(useway string, protocol []string, cb func(proxy Proxy) error) error {
	proxies := getSupportProxy(protocol, useway)
	if len(proxies) == 0 {
		return errors.New("no proxy found")
	}

	start_idx := rand.Intn(len(proxies))
	for i := 0; i < len(proxies); i++ {
		idx := (start_idx + i) % len(proxies)
		err := cb(proxies[idx])
		if err == nil {
			return nil
		}
	}

	return errors.New("all proxies failed")
}

func IsProxyEnable() bool {
	return proxyEnable
}
