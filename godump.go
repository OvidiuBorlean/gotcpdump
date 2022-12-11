package main

import (
   "fmt"
   "log"
   "github.com/google/gopacket"
   "github.com/google/gopacket/pcap"
)

var iface = "eth0"
var snaplen = int32(1600)
var promisc = false
var timeout = pcap.BlockForever
var filter = ""
var devFound = false

func main() {
  devices, err := pcap.FindAllDevs()
  if err != nil {
     log.Panicln(err)
  }
  for _, device := range devices {
      if device.Name == iface {
         devFound = true
      }
  }
  if !devFound {
     log.Panicf("Device named '%s' do not exist\n", iface)
     }
  handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
  if err != nil {
     log.Panicln(err)
     }
  defer handle.Close()
  if err := handle.SetBPFFilter(filter); err != nil {
     log.Panicln(err)
     }
  source := gopacket.NewPacketSource(handle, handle.LinkType())
 for packet := range source.Packets() {
    fmt.Println(packet)
    fmt.Println("------------------------------")
    }
}
