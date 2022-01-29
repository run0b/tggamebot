package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	c := CreateConfig()
	b, err := tb.NewBot(tb.Settings{
		Token:  c.TGKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	getSticker := func(setName string, n int) (*tb.Sticker, error) {
		set, err := b.GetStickerSet(setName)
		if err != nil {
			log.Println(err)
			return &tb.Sticker{}, err
		}
		sticker := set.Stickers[n]
		return &sticker, nil
	}

	rK1Start := tb.ReplyButton{Text: "Начать"}
	rK1Info := tb.ReplyButton{Text: "👀 Информация"}
	rK1 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK1Start, rK1Info},
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, m.Sender.FirstName+", Привет", &tb.ReplyMarkup{
			ReplyKeyboard:       rK1,
			ResizeReplyKeyboard: true,
		})
		sticker, err := getSticker("Charande", 3)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = b.Send(m.Sender, sticker)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = b.Send(m.Sender, "Радуйся хуть кто-то на тебя дрочит.")
		if err != nil {
			log.Println(err)
			return
		}

	})

	b.Handle(tb.OnSticker, func(m *tb.Message) {
		spew.Dump(m)
	})

	b.Start()
}
