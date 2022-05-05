/*

Package and module dedicated to sending out ARP requests

and responding with the clients MAC, IP, OUI, and range

this is a remodified version of the offical ARP script written by the gopacket developers
i thought i should edit this alot and seperate the code to where its readable, and designed
a bit more better since it uses modules and will rewrite the packets from other modules
*/

package arp

import (
	"fmt"
	"log"
	"net"
	"time"

	ARP_CONSTANTS "main/modg/scripts/IEEE-802.11/IEEE-802.11-c"
	ARP_ERRORS "main/modg/scripts/IEEE-802.11/recon/functions/errors"
	ARP_READERS "main/modg/scripts/IEEE-802.11/recon/functions/reader"
	ARP_WRITERS "main/modg/scripts/IEEE-802.11/recon/functions/writer"

	"github.com/google/gopacket/pcap"
)

func scan(iface *net.Interface) error {
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					ARP_CONSTANTS.Addresses = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}
	if err := ARP_ERRORS.Sanity(ARP_CONSTANTS.Addresses); err != nil {
		fmt.Println(err)
	}
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()
	stop := make(chan struct{})
	go ARP_READERS.Output_Packet_Information(handle, iface, stop)
	defer close(stop)
	for {
		if err := ARP_WRITERS.WRITER(handle, iface, ARP_CONSTANTS.Addresses); err != nil {
			log.Printf("error writing packets on %v: %v", iface.Name, err)
			return err
		}
		time.Sleep(ARP_CONSTANTS.Waiter * time.Second)
	}
}

func Return_Values_andCall() {
	interfaces, e := net.Interfaces()
	if e != nil {
		log.Fatal(e)
	} else {
		for _, inter := range interfaces {
			ARP_CONSTANTS.Wait_group.Add(1)
			go func(inter net.Interface) {
				defer ARP_CONSTANTS.Wait_group.Done()
				if err := scan(&inter); err != nil {
					log.Printf("Interface %v | %v ", inter.Name, err)
				}
			}(inter)
		}
	}
	ARP_CONSTANTS.Wait_group.Wait()
}
