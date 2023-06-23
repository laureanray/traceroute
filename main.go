package main

import (
	"fmt"
	"net"
	"net/netip"
	"strings"
)

func main() {
  fmt.Println("Hello, World!")

  url := "laureanray.com:33434"

  conn, err := net.Dial("udp", url)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Printf("Conn: %T \n" , conn)
  
  udpAddr, _ := net.ResolveUDPAddr("udp", url)
  
  fmt.Println(udpAddr);


  udpConn, ok := conn.(*net.UDPConn)

  if ok {
    remoteAddr := udpConn.RemoteAddr()
    // bytesWritten, _ := udpConn.WriteToUDP([]byte("Hello from UDP client"), udpAddr)
    parsedAddr := netip.MustParseAddr(strings.Split(remoteAddr.String(), ":")[0])
    addrPort := netip.AddrPortFrom(parsedAddr, 33434)
    bytesWritten, err := udpConn.WriteToUDPAddrPort([]byte("Hello from UDP client"), addrPort)

    if err != nil {
      fmt.Println(err)
    }

    fmt.Println(bytesWritten)
  }
}

