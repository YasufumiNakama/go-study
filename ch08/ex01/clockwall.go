package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type clock struct {
	location string
	url      string
	time     string
}

func parseArgs() ([]*clock, error) {
	var clocks []*clock

	if len(os.Args) == 1 {
		return clocks, nil
	}

	for _, arg := range os.Args[1:] {
		// expected format example: NewYork=localhost:8010
		components := strings.Split(arg, "=")
		if len(components) != 2 {
			return nil, fmt.Errorf("NOT expected format: %q", arg)
		}
		c := &clock{location: components[0], url: components[1], time: ""}
		clocks = append(clocks, c)
	}

	return clocks, nil
}

func (c *clock) getTime() {
	conn, err := net.Dial("tcp", c.url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		c.time = string(bytes)
	}
}

func main() {
	clocks, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range clocks {
		go c.getTime()
	}

	times := make([]string, len(clocks))
	for {
		for i, c := range clocks {
			times[i] = fmt.Sprintf("%s: %s", c.location, c.time)
		}
		fmt.Printf("\r%s", strings.Join(times, ", "))
	}
}
