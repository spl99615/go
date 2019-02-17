package main

import (
	"bufio"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	//fmt.Println(message.Greeting("Surya", "Hello"))
	f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Unable to open /etc/motd")
		os.Exit(1)
	}

	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting:")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := message.Greeting(name, phrase)
	_, err = f.Write([]byte(m))
	//err := ioutil.WriteFile("/etc/motd", []byte(m), 0644)
	if err != nil {
		fmt.Println("Unable to write /etc/motd")
		os.Exit(1)
	}
}
