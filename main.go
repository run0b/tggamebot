package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lib/pq"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	localizer := CreateLocale()
	// пример локализации количества кнопок
	// fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "button", TemplateData: map[string]interface{}{"Count": 3}, PluralCount: 3}))

	c := CreateConfig()

	db := CreateDatabase(c)

	b, err := tb.NewBot(tb.Settings{
		Token:  c.TGKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// getSticker := func(setName string, n int) (*tb.Sticker, error) {
	// 	set, err := b.GetStickerSet(setName)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return &tb.Sticker{}, err
	// 	}
	// 	sticker := set.Stickers[n]
	// 	return &sticker, nil
	// }

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
	rK3Ore := tb.ReplyButton{Text: "Ресурсы"}
	rK3Chest := tb.ReplyButton{Text: "Сундуки"}
	rK3Class := tb.ReplyButton{Text: "Классы"}
	rK3back := tb.ReplyButton{Text: "Назад 🔙"}
	rK3 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK3Rules, rK3Des, rK3Ore, rK3Chest, rK3Class, rK3back},
	}

	rK4Warrior := tb.ReplyButton{Text: "Воин"}
	rK4Archer := tb.ReplyButton{Text: "Лучник"}
	rK4Wizard := tb.ReplyButton{Text: "Волшебник"}
	rK4Paladin := tb.ReplyButton{Text: "Паладин"}

	rK4back := tb.ReplyButton{Text: "Назад 🔙"}
	rK4 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK4Warrior, rK4Archer, rK4Wizard, rK4Paladin, rK4back},
	}

	rK5Dig := tb.ReplyButton{Text: "Копать руду"}
	rK5Сhop := tb.ReplyButton{Text: "Рубить дерево"}
	rK5Fight := tb.ReplyButton{Text: "Убивать монстров"}
	rK5Bag := tb.ReplyButton{Text: "Сумка"}
	rK5back := tb.ReplyButton{Text: "Назад 🔙"}
	rK5 := [][]tb.ReplyButton{
		[]tb.ReplyButton{rK5Dig, rK5Сhop, rK5Fight, rK5Bag, rK5back},
	}

	b.Handle(&rK1Stat, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Тут вы можете посмотреть общию статистику", &tb.ReplyMarkup{
			ReplyKeyboard:       rK2,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK3Ore, func(m *tb.Message) {
		b.Send(m.Sender, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "desc"}),
			&tb.ReplyMarkup{
				ReplyKeyboard:       rK3,
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

	b.Handle(&rK1Start, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Выберите класс,учтите у каждого класс разные навыки и бонуссы", &tb.ReplyMarkup{
			ReplyKeyboard:       rK4,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK4Warrior, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Вы выбрали Воина", &tb.ReplyMarkup{
			ReplyKeyboard:       rK5,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK4Archer, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Вы выбрали Лучника", &tb.ReplyMarkup{
			ReplyKeyboard:       rK5,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK4Wizard, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Вы выбрали Волшебника", &tb.ReplyMarkup{
			ReplyKeyboard:       rK5,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK4Paladin, func(m *tb.Message) {
		b.Send(m.Sender, ""+m.Sender.FirstName+", Вы выбрали Паладина", &tb.ReplyMarkup{
			ReplyKeyboard:       rK5,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle("/start", func(m *tb.Message) {
		err := db.AddUser(m.Chat.ID, m.Chat.FirstName)
		if err != nil {
			e, ok := err.(*pq.Error)
			if !ok || e.Code != "23505" {
				log.Println(err)
				return
			}
		}
		b.Send(m.Sender, m.Sender.FirstName+", Привет", &tb.ReplyMarkup{
			ReplyKeyboard:       rK1,
			ResizeReplyKeyboard: true,
		})
		// sticker, err := getSticker("Charande", 3)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		// _, err = b.Send(m.Sender, sticker)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }

		_, err = b.Send(m.Sender, "Перед началом игры советуем ознакомится с игрой.\nНажмите кнопку Информация ")
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
