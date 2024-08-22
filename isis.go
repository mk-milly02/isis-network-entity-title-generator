package isis

import (
	"fmt"
	"log"
	"net/netip"
	"slices"
	"strconv"
	"strings"
)

func Validate(ipAdrr string) netip.Addr {
	ip, err := netip.ParseAddr(ipAdrr)
	if err != nil || !ip.Is4() {
		log.Fatalf("Invalid ipv4 address : %v", err)
	}
	return ip
}

func ConvertToNET(ipAdrr netip.Addr) string {
	octects := ipAdrr.As4()
	var net []string
	for i := 0; i < 4; i++ {
		o := octects[i]
		if o < 10 {
			net = append(net, fmt.Sprintf("00%s", strconv.Itoa(int(o))))
		} else if o > 9 && o < 100 {
			net = append(net, fmt.Sprintf("0%s", strconv.Itoa(int(o))))
		} else {
			net = append(net, strconv.Itoa(int(o)))
		}
	}
	charArr := strings.Split(strings.Join(net, ""), "")
	var res []string
	for v := range slices.Chunk(charArr, 4) {
		res = append(res, strings.Join(v, ""))
	}
	return strings.Join(res, ".")
}