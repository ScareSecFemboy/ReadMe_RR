package PPARSER

import (
	"bytes"
	"fmt"
	"strings"

	v "main/modg/colors"
	opc "main/modg/copt"
	"main/modg/system-runscript"
	ec "main/modg/warnings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/pflag"
)

var (
	rr6f          opc.RR6_options
	err           error
	handeler_pcap *pcap.Handle
	flags         = pflag.FlagSet{SortFlags: false}
)

// pcap file parser

// dumping PCAP parser, short [BYTE]
func Pcap_parser_OFFLINE_byte(filenametakein string, payload1 string) {
	var handle *pcap.Handle
	handle, err = pcap.OpenOffline(filenametakein)
	ec.Warning_advanced("<RR6> IEEE-802.11 Packet Parsing Module: Could not parse, open, or log file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	counter := 0
	for packet := range packetSource.Packets() {
		counter++

		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().LayerContents()
			if bytes.Contains(payload, []byte(payload1)) {
				src := packet.NetworkLayer().NetworkFlow().Src()
				dst := packet.NetworkLayer().NetworkFlow().Dst()
				fmt.Println(v.RED, "------------------- TOP HEADER DATA ---------------------------")
				fmt.Println(v.MAG, "| Source      Address |> \033[49m", v.BLKHB, src, "\033[49m")
				fmt.Println(v.MAG, "| Destination Address |> \033[49m", v.BLKHB, dst, "\033[49m")
				fmt.Println(v.MAG, "| Detected on packet# |> \033[49m", v.BLKHB, counter, "\033[49m")
				fmt.Println(v.RED, "------------------- BOTTOM HEADER INF --------------------------")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Search Payload     |> \033[49m", v.BLUHB, string(payload1), "\033[49m")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Information Element|> \033[49m", v.RED, string(payload), "\033[49m")
			}
		}
	}
}

// dumping PCAP parser, short
func Pcap_parser_OFFLINE(filenametakein string, payload1 string) {
	var handle *pcap.Handle
	handle, err = pcap.OpenOffline(filenametakein)
	ec.Warning_advanced("<RR6> IEEE-802.11 Packet Parsing Module: Could not parse, open, or log file -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	counter := 0
	for packet := range packetSource.Packets() {
		counter++

		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().LayerContents()
			if strings.Contains(string(payload), string(payload1)) {
				src := packet.NetworkLayer().NetworkFlow().Src()
				dst := packet.NetworkLayer().NetworkFlow().Dst()
				fmt.Println(v.RED, "------------------- TOP HEADER DATA ---------------------------")
				fmt.Println(v.MAG, "| Source      Address |> \033[49m", v.BLKHB, src, "\033[49m")
				fmt.Println(v.MAG, "| Destination Address |> \033[49m", v.BLKHB, dst, "\033[49m")
				fmt.Println(v.MAG, "| Detected on packet# |> \033[49m", v.BLKHB, counter, "\033[49m")
				fmt.Println(v.RED, "------------------- BOTTOM HEADER INF --------------------------")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Search Payload     |> \033[49m", v.BLUHB, string(payload1), "\033[49m")
				fmt.Println(v.MAG, "|", v.BLKHB, " Found Information Element|> \033[49m", v.RED, string(payload), "\033[49m")
			}
		}
	}
}

// simple pcap parsing
func Parser(filename string) {
	handeler_pcap, err = pcap.OpenOffline(filename)
	ec.Warning_advanced("<RR6> IEEE-802.11 Packet Parsing Module: Could not open pcap listener offline, something went wrong ->", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handeler_pcap.Close()
	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}

// FTP sniffing OFFLINE, parse through a PCAP file

func Ftp_sniff_OFFLINE_PCAP(filenametakein string, cside bool) {
	if cside {
		handeler_pcap, err = pcap.OpenOffline(rr6f.Filepath_general)
	} else {
		handeler_pcap, err = pcap.OpenOffline(filenametakein)
	}

	defer handeler_pcap.Close()

	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	for packet := range packetSource.Packets() {
		if packet.ApplicationLayer() != nil {
			payload := packet.ApplicationLayer().Payload()
			dst := packet.NetworkLayer().NetworkFlow().Dst()
			if bytes.Contains(payload, []byte("USER")) {
				fmt.Println(v.RED, "[Ether-Shark] [FTP_Authentication] Found FTP Username \t| ", v.BLU, system.FormatDate, "\t | ")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Print(v.MAG, "| Destination -> ", v.BLU, dst, "\n")
				fmt.Print(v.MAG, "| Payload     -> ", v.BLU, string(payload), "\n")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Println("\n\n", packet)
			} else if bytes.Contains(payload, []byte("PASS")) {
				fmt.Println(v.RED, "[Ether-Shark] [FTP_Authentication] Found FTP Password \t| ", v.BLU, system.FormatDate, "\t | ")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Print(v.MAG, "| Destination -> ", v.BLU, dst, "\n")
				fmt.Print(v.MAG, "| Payload     -> ", v.BLU, string(payload), "\n")
				fmt.Print(v.MAG, "---------------------------------------------------------------------------------------\n")
				fmt.Println("\n\n", packet)
			}
		}
	}
}

// OSPF authentication

func OSPF_OFFLINE_Parsing(filename string) {
	handle, err := pcap.OpenOffline(rr6f.Filepath_general)
	ec.Warning_advanced("<RR6>  Network Parsing Module: Could not open network capture offline for file, something went wrong -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ospf := packet.Layer(layers.LayerTypeOSPF)
		if nil != ospf {
			ospf, _ := ospf.(*layers.OSPFv2)
			switch {
			case ospf.AuType == 1:
				fmt.Println(v.RED, "[RR6]", v.YEL, "[Net Module] [  Authentication ]", v.BLU, "\t | ", ospf.Authentication, "|", v.BLU, system.FormatDate)
				fmt.Println(v.RED, "[RR6]", v.YEL, "[Net Module] \033[36m[  \033[31mAuth_Version\033[36m   ]", v.BLU, "\t | ", ospf.AuType, "\t\t        |", v.BLU, system.FormatDate)
			case ospf.AuType == 2:
				fmt.Println(v.RED, "[RR6]", v.YEL, "[Net Module] [  Authentication ]", v.BLU, "\t | ", ospf.Authentication, "|", v.BLU, system.FormatDate)
				fmt.Println(v.RED, "[RR6]", v.YEL, "[Net Module] \033[36m[  \033[31mMD5  Version\033[36m   ]", v.BLU, "\t | ", ospf.AuType, "\t\t  |", v.BLU, system.FormatDate)
			}
		}
	}
}

func Open_parse_BSSID(pcapfile string) {
	handeler_pcap, err = pcap.OpenOffline(pcapfile)
	ec.Warning_advanced("<RR6> Network Module: Could not parse packet listener and make or open an offline handeler", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer handeler_pcap.Close()
	packetSource := gopacket.NewPacketSource(handeler_pcap, handeler_pcap.LinkType())
	for packet := range packetSource.Packets() {
		dot11 := packet.Layer(layers.LayerTypeDot11)
		if nil != dot11 {
			dot11, _ := dot11.(*layers.Dot11)
			if dot11.Address3 != nil {
				fmt.Println(v.BLKHB, "BSSID | ", dot11.Address3, " | Flags -> ", dot11.Flags, "\033[49m\033[31m")
			}
		}
		dot11info := packet.Layer(layers.LayerTypeDot11InformationElement)
		if nil != dot11info {
			dot11info, _ := dot11info.(*layers.Dot11InformationElement)
			if dot11info.ID == layers.Dot11InformationElementIDSSID {
				fmt.Printf(" \033[0;100mSSID | %q\n", dot11info.Info)
				fmt.Print("\033[49m")
			}
		}
		fmt.Printf("\n")
	}
}
