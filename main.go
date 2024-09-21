package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mattn/go-sixel"
)

type CardImage struct {
	ID         int    `json:"id"`
	URL        string `json:"image_url"`
	URLSmall   string `json:"image_url_small"`
	URLCropped string `json:"image_url_cropped"`
}

type Card struct {
	ID     int         `json:"id"`
	Name   string      `json:"name"`
	Images []CardImage `json:"card_images"`
}

type CardResponse struct {
	Data []Card `json:"data"`
}

func getYuGiOhCardByRandom() (Card, error) {
	resp, err := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php?num=1&offset=0&sort=random&cachebust")
	if err != nil {
		return Card{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Card{}, err
	}

	var cardResponse CardResponse
	if err := json.Unmarshal(body, &cardResponse); err != nil {
		return Card{}, err
	}

	if len(cardResponse.Data) == 0 {
		return Card{}, fmt.Errorf("card not found")
	}

	return cardResponse.Data[0], nil
}

func main() {
	card, err := getYuGiOhCardByRandom()
	if err != nil {
		log.Fatal(err)
	}

	if len(card.Images) == 0 {
		log.Fatal("card has no images")
	}

	resp, err := http.Get(card.Images[0].URL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
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
