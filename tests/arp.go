package main

// putting it all into one main file as adding onto easy make and call

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/macs"
	"github.com/google/gopacket/pcap"
	"github.com/theckman/yacspin"
)

var (
	now        = time.Now()
	formatDate = now.Format("15:04:05")
	BLK        = "\033[0;30m"
	RED        = "\033[0;31m"
	GRN        = "\033[0;32m"
	YEL        = "\033[0;33m"
	BLU        = "\033[0;34m"
	MAG        = "\033[0;35m"
	CYN        = "\033[0;36m"
	WHT        = "\033[0;37m"
	BBLK       = "\033[1;30m"
	BRED       = "\033[1;31m"
	BGRN       = "\033[1;32m"
	BYEL       = "\033[1;33m"
	BBLU       = "\033[1;34m"
	BMAG       = "\033[1;35m"
	BCYN       = "\033[1;36m"
	BWHT       = "\033[1;37m"
	UBLK       = "\033[4;30m"
	URED       = "\033[4;31m"
	UGRN       = "\033[4;32m"
	UYEL       = "\033[4;33m"
	UBLU       = "\033[4;34m"
	UMAG       = "\033[4;35m"
	UCYN       = "\033[4;36m"
	UWHT       = "\033[4;37m"
	BLKB       = "\033[40m"
	REDB       = "\033[41m"
	GRNB       = "\033[42m"
	YELB       = "\033[43m"
	BLUB       = "\033[44m"
	MAGB       = "\033[45m"
	CYNB       = "\033[46m"
	WHTB       = "\033[47m"
	BLKHB      = "\033[0;100m"
	REDHB      = "\033[0;101m"
	GRNHB      = "\033[0;102m"
	YELHB      = "\033[0;103m"
	BLUHB      = "\033[0;104m"
	MAGHB      = "\033[0;105m"
	CYNHB      = "\033[0;106m"
	WHTHB      = "\033[0;107m"
	HBLK       = "\033[0;90m"
	HRED       = "\033[0;91m"
	HGRN       = "\033[0;92m"
	HYEL       = "\033[0;93m"
	HBLU       = "\033[0;94m"
	HMAG       = "\033[0;95m"
	HCYN       = "\033[0;96m"
	HWHT       = "\033[0;97m"
	BHBLK      = "\033[1;90m"
	BHRED      = "\033[1;91m"
	BHGRN      = "\033[1;92m"
	BHYEL      = "\033[1;93m"
	BHBLU      = "\033[1;94m"
	BHMAG      = "\033[1;95m"
	BHCYN      = "\033[1;96m"
	BHWHT      = "\033[1;97m"
)

// spinner config

var (
	cfg = yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[54],
		Suffix:          "",
		SuffixAutoColon: true,
		Message:         "Being slow on purpose, wait a danmn minute",
		StopCharacter:   "[Ether-Shark] Stat: Spinner to hold your fucking horses is done",
		StopColors:      []string{"fgGreen"},
	}

	second_config = yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[54],
		Suffix:          "",
		SuffixAutoColon: true,
		Message:         "\tStarting the ether_shark framework",
		StopCharacter:   "",
		StopColors:      []string{"fgGreen"},
	}
)

func spinner() {
	fmt.Print("\n\n")
	spinner, err := yacspin.New(cfg)
	if err != nil {
		panic(err)
	}
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Message("Sending out ARP frame...")
	spinner.Stop()
}

func ce(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func localaddr() {
	ifaces, err := net.Interfaces()
	ce(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		ce(err)
		for _, a := range addrs {
			fmt.Println(RED, "<RR6> Net Module: Found Device  | ", MAG, i.Name, RED, "\t | Addr:", a)
		}
	}
}

func scan(iface *net.Interface) error {

	go handelreturncon(make(chan os.Signal, 1))
	var addr *net.IPNet
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					addr = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}
	if addr == nil {
		fmt.Println(RED, "<RR6> Net Module: "+REDHB+"Unstable Net\033[49m\033[31m  | ", BLU, formatDate, RED, "\t | Addr:", addr)
	} else if addr.IP[0] == 127 {
		fmt.Println(RED, "<RR6> Net Module: "+REDHB+"Skipping LO:\033[49m\033[31m  | ", BLU, formatDate, RED, "\t | Addr:", REDHB, addr)
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		fmt.Println(RED, "<RR6> Net Module: "+REDHB+"Netmask err\033[49m\033[31m  | ", BLU, formatDate, RED, "\t | Addr:", REDHB)
	}
	fmt.Println(RED, "<RR6> Net Module: "+MAGHB+"Using Range:\033[49m\033[31m ", "| ", BLU, formatDate, "\t | Range: ", addr)
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	ce(err)
	defer handle.Close()
	stop := make(chan struct{})
	go ARPR(handle, iface, stop, true)
	time.Sleep(1 * time.Second)
	defer close(stop)
	for {

		if err := ARPW(handle, iface, addr); err != nil {
			fmt.Println("[DATA]->[ERROR] An error has occured during the following write of packets-> ", iface.Name, err)
			return err
		}
	}
}

func ARPR(handle *pcap.Handle, iface *net.Interface, stop chan struct{}, OUI bool) {
	go handelreturncon(make(chan os.Signal, 1))
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {

				continue
			}
			fmt.Println(RED, "<RR6> Net Module:", BLKHB, "Address\033[49m\033[31m | ", BLU, net.IP(arp.SourceProtAddress), "| ", BLKHB, "MAC\t\033[49m| ", MAG, net.HardwareAddr(arp.SourceHwAddress), "\t| ")
			//binary.BigEndian.Uint16(net.HardwareAddr(arp.SourceHwAddress))
		}
	}
}

// parse through a file of macs to identify their addresses
func tracer(filename string) string {
	content, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
			if mac, err := net.ParseMAC(scanner.Text()); err == nil {
				prefix := [3]byte{
					mac[0],
					mac[1],
					mac[2],
				}
				manufacturer, ok := macs.ValidMACPrefixMap[prefix]
				if ok {
					fmt.Print(manufacturer)
				}
			}
		}
	}
	return ""
}

func ARPW(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) error {
	go handelreturncon(make(chan os.Signal, 1))
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // ff:ff:ff:ff:ff:ff
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	for _, ip := range ips(addr) {
		arp.DstProtAddress = []byte(ip)
		gopacket.SerializeLayers(buf, opts, &eth, &arp)
		if err := handle.WritePacketData(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	network := num & mask
	broadcast := network | ^mask
	for network++; network < broadcast; network++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network)
		out = append(out, net.IP(buf[:]))
	}
	return
}

func arpmain() {
	localaddr()
	go handelreturncon(make(chan os.Signal, 1))
	var err error
	ifaces, err := net.Interfaces()
	ce(err)
	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)
		go func(iface net.Interface) {
			defer wg.Done()
			if err := scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}
		}(iface)
		go handelreturncon(make(chan os.Signal, 1))
	}
	go handelreturncon(make(chan os.Signal, 1))
	wg.Wait()
}

func handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("\nDetected Interupt.....")
			os.Exit(1)
		case os.Kill:
			fmt.Println("\n\n\tKILL received")
			os.Exit(1)
		}
	}
}

func main() {
	arpmain()
}
