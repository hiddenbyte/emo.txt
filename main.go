package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/hiddenbyte/emo.txt/slack"
)

func main() {
	emoWriter := slack.NewAlphabetizer(os.Stdout)

	r := bufio.NewReader(os.Stdin)
	for {
		c, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		err = emoWriter.WriteRune(c)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	err := emoWriter.Flush()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
