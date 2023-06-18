package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"strings"
)

const (
	TICKET_PATH = "./resources/ticket.txt"
)

func ExtractByte(url string) []byte {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	return bytes
}
func ExtractTickets(file *os.File) {
	//var TICKETS []string

	resp := ExtractByte("https://www.nseindia.com/sitemap-stocks.xml")
	m := regexp.MustCompile(`=(.*?)\<`)
	res := m.FindAllString(string(resp), -1)

	//fmt.Println(string(resp))

	for i := 0; i < len(res); i++ {
		s := strings.ReplaceAll(res[i], "<", "")
		s = strings.ReplaceAll(s, "=", "")
		//fmt.Println(s)
		//TICKETS = append(TICKETS, s)
		_, _ = fmt.Fprintf(file, fmt.Sprintln(s))
	}
}

func ReadTicketFromFile() []string {
	if _, err := os.Stat(TICKET_PATH); err != nil {
		file, err := os.Create(TICKET_PATH)
		if err != nil {
			log.Fatal(err)
		}
		ExtractTickets(file)
		file.Close()
	}

	file, err := os.Open(TICKET_PATH)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func ReadLTP(ticket string) {

}
func main() {
	//var file *os.File
	s := ReadTicketFromFile()
	fmt.Println(s[0])

}
