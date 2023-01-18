package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Doarakko/go-yugioh/yugioh"
	"github.com/mattn/go-sixel"
)

func main() {
	client := yugioh.NewClient()

	card, _, err := client.RandomCards.One()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(card.Images[0].URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	err = sixel.NewEncoder(os.Stdout).Encode(img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(card.Images[0].URL)
}
