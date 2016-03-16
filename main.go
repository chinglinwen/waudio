package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	ivona "github.com/jpadilla/ivona-go"
)

var (
	voice = ivona.Voice{"Brian", "en-GB", "Male"}
)

func main() {

	var inputFile string
	flag.StringVar(&inputFile, "f", "", "data file")

	var key, sec string
	flag.StringVar(&key, "k", "", "IVONA_ACCESS_KEY")
	flag.StringVar(&sec, "s", "", "IVONA_SECRET_KEY")

	h := flag.Bool("help", false, "show help")
	v := flag.Bool("v", false, "version")
	flag.Parse()
	if *h == true {
		fmt.Println(os.Args[0], "convert english text to audio through ivona service")
		fmt.Printf("Usage: \n\t%v inputFile [outputFile]\n", os.Args[0])
		os.Exit(0)
	}
	if *v == true {
		fmt.Println("Version 1.0.0, by Chingnlin Wen, Mar 16 2016")
		os.Exit(0)

	}
	log.Println("Starting...")

	log.Println("Read data from file", inputFile)

	in, err := os.Open(inputFile)
	check(err)
	defer in.Close()

	outputFile := inputFile + ".mp3"
	out, err := os.Create(outputFile)
	check(err)
	defer out.Close()

	client := ivona.New(key, sec)

	scanner := bufio.NewScanner(in)
	var buf []byte
	for scanner.Scan() {
		buf = scanner.Bytes()
		if string(buf) == "" {
			continue
		}
		options := ivona.NewSpeechOptions(string(buf))
		options.Voice = &voice
		result, err := client.CreateSpeech(options)
		if err != nil {
			log.Println("for:", string(buf))
			log.Fatal(err)
		}

		n, err := out.Write(result.Audio)
		check(err)
		log.Println("bytes", n, "Done")
	}
	log.Println("all Done")
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
