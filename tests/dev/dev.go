package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	err              error
	snapshot_main    int32         = 1024
	timeout_shot_cap time.Duration = 40 * time.Second
	handeler_pcap    *pcap.Handle
	clear_hex        = "\x1b[H\x1b[2J\x1b[3J"
	BLK              = "\033[0;30m"
	RED              = "\033[0;31m"
	GRN              = "\033[0;32m"
	YEL              = "\033[0;33m"
	BLU              = "\033[0;34m"
	MAG              = "\033[0;35m"
	CYN              = "\033[0;36m"
	WHT              = "\033[0;37m"
	BBLK             = "\033[1;30m"
	BRED             = "\033[1;31m"
	BGRN             = "\033[1;32m"
	BYEL             = "\033[1;33m"
	BBLU             = "\033[1;34m"
	BMAG             = "\033[1;35m"
	BCYN             = "\033[1;36m"
	BWHT             = "\033[1;37m"
	UBLK             = "\033[4;30m"
	URED             = "\033[4;31m"
	UGRN             = "\033[4;32m"
	UYEL             = "\033[4;33m"
	UBLU             = "\033[4;34m"
	UMAG             = "\033[4;35m"
	UCYN             = "\033[4;36m"
	UWHT             = "\033[4;37m"
	BLKB             = "\033[40m"
	REDB             = "\033[41m"
	GRNB             = "\033[42m"
	YELB             = "\033[43m"
	BLUB             = "\033[44m"
	MAGB             = "\033[45m"
	CYNB             = "\033[46m"
	WHTB             = "\033[47m"
	BLKHB            = "\033[0;100m\033[31m"
	REDHB            = "\033[0;101m"
	GRNHB            = "\033[0;102m"
	YELHB            = "\033[0;103m"
	BLUHB            = "\033[0;104m"
	MAGHB            = "\033[0;105m"
	CYNHB            = "\033[0;106m"
	WHTHB            = "\033[0;107m"
	HBLK             = "\033[0;90m"
	HRED             = "\033[0;91m"
	HGRN             = "\033[0;92m"
	HYEL             = "\033[0;93m"
	HBLU             = "\033[0;94m"
	HMAG             = "\033[0;95m"
	HCYN             = "\033[0;96m"
	HWHT             = "\033[0;97m"
	BHBLK            = "\033[1;90m"
	BHRED            = "\033[1;91m"
	BHGRN            = "\033[1;92m"
	BHYEL            = "\033[1;93m"
	BHBLU            = "\033[1;94m"
	BHMAG            = "\033[1;95m"
	BHCYN            = "\033[1;96m"
	BHWHT            = "\033[1;97m"
	now              = time.Now()
	formatDate       = now.Format("15:04:05")
	chunkType        string
	filename         string
	packets          = gopacket.Packet.Layer(s)
)

func errne(err error, color, msg string) {
	if err != nil {
		log.Fatal(err)
	}
}

func pcap_parser_OFFLINE_byte_NT(filenametakein string, payload1 string) {
	var handle *pcap.Handle
	handle, err = pcap.OpenOffline(filenametakein)
	errne(err, REDHB, "<RR6> Packet Parsing Module: Could not parse, open, or log file -> ")
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	counter := 0
	for packet := range packetSource.Packets() {
		counter++
		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().LayerContents()
			if strings.Contains(string(payload), payload1) {
				src := packet.NetworkLayer().NetworkFlow().Src()
				dst := packet.NetworkLayer().NetworkFlow().Dst()
				fmt.Println(RED, "------------------- TOP HEADER DATA ---------------------------")
				fmt.Println(MAG, "| Source      Address |> \033[49m", BLKHB, src, "\033[49m")
				fmt.Println(MAG, "| Destination Address |> \033[49m", BLKHB, dst, "\033[49m")
				fmt.Println(MAG, "| Detected on packet# |> \033[49m", BLKHB, counter, "\033[49m")
				fmt.Println(RED, "------------------- BOTTOM HEADER INF --------------------------")
				fmt.Println(MAG, "|", BLKHB, " Found Search Payload     |> \033[49m", BLUHB, string(payload1), "\033[49m")
				fmt.Printf(string(payload))
				fmt.Println("...........................................................................................")
				fmt.Print("\n\n\n\n")
			}
			//fmt.Println("CODE LINE: ", counter, " HAS PAYLOAD|> ", string(payload))
			//fmt.Println(packet)

		}
	}
}

func main() {
	filename := os.Args[1]
	//NTLMSSP
	//SSP
	pcap_parser_OFFLINE_byte_NT(filename, "NTLMv2")
}
