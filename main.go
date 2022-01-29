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
	rK1Stat := tb.ReplyButton{Text: "Статистика"}
	rK1Info := tb.ReplyButton{Text: "👀 Информация"}
	rK1 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK1Start, rK1Info, rK1Stat},
	}

	rK2Stat := tb.ReplyButton{Text: "Руды добыто  ⛏ "}
	rK2Kill := tb.ReplyButton{Text: "Убийства 🏹"}
	rK2Xp := tb.ReplyButton{Text: "Навыки 💱"}
	rK2Money := tb.ReplyButton{Text: "Балланс 💲"}
	rK2back := tb.ReplyButton{Text: "Назад 🔙"}
	rK2 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK2Stat, rK2Kill, rK2Xp, rK2Money, rK2back},
	}

	rK3Rules := tb.ReplyButton{Text: "Правила "}
	rK3Des := tb.ReplyButton{Text: "Описание"}
	rK3Ore := tb.ReplyButton{Text: "Виды руд"}
	rK3Chest := tb.ReplyButton{Text: "Сундуки"}
	rK3back := tb.ReplyButton{Text: "Назад 🔙"}
	rK3 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK3Rules, rK3Des, rK3Ore, rK3Chest, rK3back},
	}

	b.Handle(&rK1Stat, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Тут вы можете посмотреть общию статистику", &tb.ReplyMarkup{
			ReplyKeyboard:       rK2,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK2back, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Вы вернулись на главную страницу", &tb.ReplyMarkup{
			ReplyKeyboard:       rK1,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK1Info, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Тут вы можете ознакомится с механикой игры", &tb.ReplyMarkup{
			ReplyKeyboard:       rK3,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

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
