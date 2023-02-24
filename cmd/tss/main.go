package main

import (
	"bridge/network/common"
	"bridge/network/p2p"
	"bridge/network/tss"
	"bufio"
	"flag"
	"log"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client/input"
)

var (
	tssAddr    string
	baseFolder string
)

func main() {

	// Parse the cli into configuration structs
	tssConf, p2pConf := readconfig()

	// Read stdin for the private key
	inBuf := bufio.NewReader(os.Stdin)
	priKeyBytes, err := input.GetPassword("input node secret key:", inBuf)
	if err != nil {
		log.Fatalf("error in get the secret key: %s\n", err.Error())
	}
	priKey, err := common.GetPriKey(priKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	tss := tss.NewTss()
	if err = tss.Start(); err != nil {
		panic(err)
	}

}

func readconfig() (*tss.Config, *p2p.Config) {
	tssConfig := &tss.Config{}
	p2pConfig := &p2p.Config{}
	// we setup the configure for the general configuration
	flag.StringVar(&tssAddr, "tss-port", "127.0.0.1:8080", "tss port")
	//flag.StringVar(&logLevel, "loglevel", "info", "Log Level")
	flag.StringVar(&baseFolder, "home", "", "home folder to store the keygen state file")

	// we setup the Tss parameter configuration
	flag.DurationVar(&tssConfig.KeyGenTimeout, "gentimeout", 30*time.Second, "keygen timeout")
	flag.DurationVar(&tssConfig.KeySignTimeout, "signtimeout", 30*time.Second, "keysign timeout")
	flag.DurationVar(&tssConfig.PreParamTimeout, "preparamtimeout", 5*time.Minute, "pre-parameter generation timeout")
	flag.BoolVar(&tssConfig.EnableMonitor, "enablemonitor", true, "enable the tss monitor")

	// we setup the p2p network configuration
	flag.StringVar(&p2pConfig.RendezvousString, "rendezvous", "Asgard",
		"Unique string to identify group of nodes. Share this with your friends to let them connect with you")
	flag.IntVar(&p2pConfig.Port, "p2p-port", 6668, "listening port local")
	flag.StringVar(&p2pConfig.ExternalIP, "external-ip", "", "external IP of this node")
	flag.Var(&p2pConfig.BootstrapPeers, "peer", "Adds a peer multiaddress to the bootstrap list")
	flag.Parse()
	return tssConfig, p2pConfig
}
