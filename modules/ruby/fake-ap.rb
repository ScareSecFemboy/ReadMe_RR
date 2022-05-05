require 'packetgen'
require 'colorize'

module FakeAP
    def FakeAP.clear
        puts "\x1b[H\x1b[2J\x1b[3J"
    end
    def FakeAP.reciever(interface, stars, ynonpa)
        interface = interface
        ssidname  = stars
        iface     = interface
        broadcast = "ff:ff:ff:ff:ff:ff"
        bssid     = "aa:aa:aa:aa:aa:aa"
        ssid      = stars
        pkt = PacketGen.gen('RadioTap').add('Dot11::Management', mac1: broadcast, mac2: bssid, mac3: bssid)
                                    .add('Dot11::Beacon', interval: 0x600, cap: 0x401)
        pkt.dot11_beacon.elements << {type: 'SSID', value: ssid}
        if ynonpa 
            pp pkt
        end
        if ynonpa == false
            puts "[ + ] Packet overview canceled"
        end

        bt = "Fake Beacon"
        puts "\tBeacon Type       SSID Name            Interface        Beacons sent "
        puts "\t---------------------------------------------------------------------"
        while true
            i = 0
            100000.times do
                i = i + 1
                pkt.to_w(iface)
                puts "\t\033[31m#{bt}        \033[36m#{ssid}          \033[31m#{iface}          \033[36m#{i}" 
            end
        end
    rescue Interrupt
        puts " [ + ] Interupt"
        exit!
    end  
end