package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"gtihub.com/tripathysagar/nseTools/crawling"
)

const (
	TICKET_PATH = "./resources/ticket.txt"
)

func ExtractTickets(file *os.File) {
	//var TICKETS []string

	resp := crawling.ExtractByte("https://www.nseindia.com/sitemap-stocks.xml")
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
	args := os.Args[1:]

	fmt.Println(args)

	if len(args) == 2 || len(args) == 1 {
		switch args[0] {
		case "-h":
			fmt.Printf(" tag: description\n -t: with ticket to show current ticket data \n -p: with portfolio file to show current data of portfolio \n")
		case "-t":
			if args[1] != "" {
				fmt.Printf("getting detatils for the ticket : %s", args[1])
			} else {
				fmt.Println("expected a ticket to be entered")
			}
		case "-p":
			if args[1] != "" {
				fmt.Printf("getting detatils for the portfolio : %s", args[1])
			} else {
				fmt.Println("expected a portfolio to be entered")
			}
		default:
			fmt.Println("not a valid option. please run the cmd with -h tag to choose a valid option")
		}

	}
}
