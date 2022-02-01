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
		b.Send(m.Sender, m.Sender.FirstName+", Вы вернулись на главную страницу", &tb.ReplyMarkup{
			ReplyKeyboard:       rK1,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&rK1Info, func(m *tb.Message) {
		b.Send(m.Sender, m.Sender.FirstName+", Тут вы можете ознакомится с механикой игры", &tb.ReplyMarkup{
			ReplyKeyboard:       rK3,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	})

	// rK4Warrior := tb.ReplyButton{Text: "Воин"}
	// rK4Archer := tb.ReplyButton{Text: "Лучник"}
	// rK4Wizard := tb.ReplyButton{Text: "Волшебник"}
	// rK4Paladin := tb.ReplyButton{Text: "Паладин"}

	// rK4back := tb.ReplyButton{Text: "Назад 🔙"}
	// rK4 := [][]tb.ReplyButton{
	// 	[]tb.ReplyButton{rK4Warrior, rK4Archer, rK4Wizard, rK4Paladin, rK4back},
	// }

	var (
		rK4        = &tb.ReplyMarkup{}
		rK4Warrior = rK4.Data("Воин", "warrior")
		rK4Archer  = rK4.Data("Лучник", "archer")
		rK4Wizard  = rK4.Data("Маг", "wizard")
		rK4Paladin = rK4.Data("Паладин", "paladin")
	)

	rK4.Inline(
		rK4.Row(rK4Warrior),
		rK4.Row(rK4Archer),
		rK4.Row(rK4Wizard),
		rK4.Row(rK4Paladin),
	)

	// b.Handle(tb.OnQuery, func(q *tb.Query) { spew.Dump(q) })

	// b.Handle(&rK4Warrior, func(c *tb.Callback) {
	// 	spew.Dump(c)
	// 	b.Respond(c, &tb.CallbackResponse{})
	// })

	b.Handle(&rK1Start, func(m *tb.Message) {
		class, err := db.GetUserClass(m.Sender.ID)
		if err != nil {
			log.Println(err)
			return
		}

		if class != -1 {
			className := ""
			switch class {
			case 0:
				className = "Воин"
			case 1:
				className = "Лучник"
			case 2:
				className = "Маг"
			case 3:
				className = "Паладин"
			}
			b.Send(m.Sender, "Ваш класс "+className, &tb.ReplyMarkup{
				ReplyKeyboard:       rK5,
				ResizeReplyKeyboard: true,
				// InlineKeyboard: inlineKeys,
			})
			return
		}
		b.Send(m.Sender, m.Sender.FirstName+", Выберите класс,учтите у каждого класс разные навыки и бонуссы", rK4)
	})

	selectClass := func(c *tb.Callback, classID int) {
		class, err := db.GetUserClass(c.Sender.ID)
		if err != nil {
			log.Println(err)
			return
		}

		if class != -1 {
			b.Send(c.Sender, "Соси хуй, клас изменить нельзя!")
			return
		}

		err = db.SetUserClass(c.Sender.ID, classID)
		if err != nil {
			log.Println(err)
			return
		}

		message := ""
		switch classID {
		case 0:
			message = c.Sender.FirstName + ", Вы выбрали Воина"
		case 1:
			message = c.Sender.FirstName + ", Вы выбрали Лучника"
		case 2:
			message = c.Sender.FirstName + ", Вы выбрали Мага"
		case 3:
			message = c.Sender.FirstName + ", Вы выбрали Паладина"
		}

		b.Send(c.Sender, message, &tb.ReplyMarkup{
			ReplyKeyboard:       rK5,
			ResizeReplyKeyboard: true,
			// InlineKeyboard: inlineKeys,
		})
	}

	b.Handle(&rK4Warrior, func(c *tb.Callback) {
		selectClass(c, 0)
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&rK4Archer, func(c *tb.Callback) {
		selectClass(c, 1)
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&rK4Wizard, func(c *tb.Callback) {
		selectClass(c, 2)
		b.Respond(c, &tb.CallbackResponse{})
	})

	b.Handle(&rK4Paladin, func(c *tb.Callback) {
		selectClass(c, 3)
		b.Respond(c, &tb.CallbackResponse{})
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
