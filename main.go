package main

import (
	"bufio"
	"fmt"
	"go-email-checker/utils"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Could not read from input %v", err)
	}
}
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	utils.ErrorLogger(err)

	if len(mxRecord) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	utils.ErrorLogger(err)

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarc_records, err := net.LookupTXT("_dmarc." + domain)
	utils.ErrorLogger(err)

	for _, record := range dmarc_records {
		if strings.HasPrefix(record, "v=dmarc1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("Domain=%v\nhasMx=%v\nhasSPF=%v\nspfRecord%v\nhasDMARC=%v\ndmarcRecord=%v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
