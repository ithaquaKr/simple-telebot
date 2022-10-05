package main

import (
	"log"

	libgiphy "github.com/sanzaru/go-giphy"
	tele "gopkg.in/telebot.v3"
)

// const (
// 	GIPHY_API_KEY = "I1gzX2emsI7DXllHhTff1RSimeqESuzr"
// 	// TELEGRAM_TOKEN = "5772285972:AAGfog9OMQm1a0N0ePsEnlPKL8Rl8mRPcp4"
// )

var (
	// apiKey = os.Getenv("I1gzX2emsI7DXllHhTff1RSimeqESuzr")
	// token  = os.Getenv("5772285972:AAGfog9OMQm1a0N0ePsEnlPKL8Rl8mRPcp4")

	menu = &tele.ReplyMarkup{ResizeKeyboard: true}

	btnSlap  = menu.Text("/slap")
	btnPunch = menu.Text("/punch")
	btnKick  = menu.Text("/kick")
)

func getSlap() string {
	giphy := libgiphy.NewGiphy("I1gzX2emsI7DXllHhTff1RSimeqESuzr")

	gif, err := giphy.GetRandom("slap")
	if err != nil {
		log.Printf("failed to retrive gif, reason=%s\n", err)
	}

	return gif.Data.Bitly_gif_url
}
func getPunch() string {
	giphy := libgiphy.NewGiphy("I1gzX2emsI7DXllHhTff1RSimeqESuzr")

	gif, err := giphy.GetRandom("punch")
	if err != nil {
		log.Printf("failed to retrive gif, reason=%s\n", err)
	}

	return gif.Data.Bitly_gif_url
}

func getKick() string {
	giphy := libgiphy.NewGiphy("I1gzX2emsI7DXllHhTff1RSimeqESuzr")

	gif, err := giphy.GetRandom("kick")
	if err != nil {
		log.Printf("failed to retrive gif, reason=%s\n", err)
	}

	return gif.Data.Bitly_gif_url
}

func main() {
	pref := tele.Settings{
		Token: "5772285972:AAGfog9OMQm1a0N0ePsEnlPKL8Rl8mRPcp4",
	}

	menu.Reply(
		menu.Row(btnSlap),
		menu.Row(btnPunch),
		menu.Row(btnKick),
	)

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// menu.Reply(
	// 	menu.Row(btnCats),
	// )

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi there, I'm a gif bot. Press the /slap /kick /punch command to attack your friend!", menu)
	})

	b.Handle("/slap", func(c tele.Context) error {
		c.Notify("typing")
		gif := &tele.Animation{File: tele.FromURL(getSlap())}
		return c.Send(gif.FileURL)
	})

	b.Handle("/punch", func(c tele.Context) error {
		c.Notify("typing")
		gif := &tele.Animation{File: tele.FromURL(getPunch())}
		return c.Send(gif.FileURL)
	})

	b.Handle("/kick", func(c tele.Context) error {
		c.Notify("typing")
		gif := &tele.Animation{File: tele.FromURL(getKick())}
		return c.Send(gif.FileURL)
	})

	b.Start()
}
