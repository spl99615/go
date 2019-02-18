package main

import (
	"bufio"
	"flag"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	var name string
	var greeting string
	var prompt bool
	var preview bool

	flag.StringVar(&name, "name", "", "name for the greeting")
	flag.StringVar(&greeting, "greeting", "", "greeting for the name")
	flag.BoolVar(&prompt, "prompt", false, "prompt value to input name & greeting")
	flag.BoolVar(&preview, "preview", false, "preview value to input message with out writing to /etc/motd")

	flag.Parse()

	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Println("name:", name)
		fmt.Println("greeting:", greeting)
		fmt.Println("prompt:", prompt)
		fmt.Println("preview:", preview)
		os.Exit(0)
	}

	if prompt {
		name, greeting = renderprompt()
	}

	m := message.Greeting(name, greeting)

	if preview {
		fmt.Println(m)
	} else {
		f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Unable to open /etc/motd")
			os.Exit(1)
		}

		defer f.Close()
		_, err = f.Write([]byte(m))
		//err := ioutil.WriteFile("/etc/motd", []byte(m), 0644)
		if err != nil {
			fmt.Println("Unable to write /etc/motd")
			os.Exit(1)
		}

	}
}

func renderprompt() (name, phrase string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting:")
	phrase, _ = reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name:")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}
