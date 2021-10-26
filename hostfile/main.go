package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	hostsfile "github.com/kevinburke/hostsfile/lib"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	f, err := os.Open("/etc/hosts")
	checkError(err)
	h, err := hostsfile.Decode(f)
	checkError(err)

	local, err := net.ResolveIPAddr("ip", "127.0.0.1")
	checkError(err)
	// Necessary for sites like facebook & gmail that resolve ipv6 addresses,
	// if your network supports ipv6
	ip6, err := net.ResolveIPAddr("ip", "::1")
	checkError(err)
	h.Set(*local, "www.facebook.com")
	h.Set(*ip6, "www.facebook.com")
	h.Set(*local, "news.ycombinator.com")
	h.Set(*ip6, "news.ycombinator.com")

	// Write to a temporary file and then atomically copy it into place.
	tmpf, err := ioutil.TempFile("/tmp", "hostsfile-temp")
	checkError(err)

	err = hostsfile.Encode(tmpf, h)
	checkError(err)

	err = os.Chmod(tmp.Name(), 0644)
	checkError(err)

	err = os.Rename(tmp.Name(), "/etc/hosts")
	checkError(err)
	fmt.Println("done")
}
