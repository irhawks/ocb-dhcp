package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"log"
)

type Config struct {
	Network struct {
		Port     int
		Username string
		Password string
	}
}

var config_path, key_pem, cert_pem, data_dir string
var server_ip string

func init() {
	flag.StringVar(&config_path, "config_path", "/etc/rebar-dhcp.conf", "Path to config file")
	flag.StringVar(&key_pem, "key_pem", "/etc/dhcp-https-key.pem", "Path to key file")
	flag.StringVar(&cert_pem, "cert_pem", "/etc/dhcp-https-cert.pem", "Path to cert file")
	flag.StringVar(&data_dir, "data_dir", "/var/cache/rebar-dhcp", "Path to store data")
	flag.StringVar(&server_ip, "server_ip", "", "Server IP to return in packets")
}

func main() {
	flag.Parse()

	var cfg Config
	cerr := gcfg.ReadFileInto(&cfg, config_path)
	if cerr != nil {
		log.Fatal(cerr)
	}

	fe := NewFrontend(data_dir, cert_pem, key_pem, cfg)

	err := StartDhcpHandlers(fe.DhcpInfo, server_ip)
	if err != nil {
		log.Fatal(err)
	}
	fe.RunServer(true)
}