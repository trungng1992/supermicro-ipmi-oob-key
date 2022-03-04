package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var key = "8544E3B47ECA58F9583043F8"

func hex2bin(s string) string {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", decoded)
}

func main() {
	mac := flag.String("m", "", "The mac address of BMC")

	flag.Parse()

	if mac == nil {
		log.Fatalf("Please input the mac address")
	}

	// Check mac is valid
	match, _ := regexp.MatchString(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`, *mac)

	if !match {
		log.Fatalf("Please input the valid mac address")
	}

	editedMac := strings.Replace(*mac, ":", "", -1)
	editedMac = strings.Replace(editedMac, "-", "", -1)
	// Calculate SHA1 with key
	h := hmac.New(sha1.New, []byte(hex2bin(key)))
	h.Write([]byte(hex2bin(editedMac)))

	licenseKey := hex.EncodeToString(h.Sum(nil))[0:24]

	result := ""
	s := 0
	keylen := len([]rune(licenseKey))
	for i := 0; i < keylen; i++ {
		s++
		result += string(licenseKey[i])
		if s == 4 && i != keylen-1 {
			result += " "
			s = 0
		}
	}

	fmt.Printf("The license key: %s\n", result)
}
