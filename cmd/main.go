package main

import (
	"context"
	"crawler/cmd/models"
	"crawler/configs"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func CurrentTime() string {
	year, month, day := time.Now().Date()
	current_time := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	return current_time
}

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	db, err := configs.InitDB(false)
	if err != nil {
		panic(err)
	}

	fortunes := []models.Fortune{}
	zooKeywords := [...]string{"쥐띠", "소띠", "호랑이띠", "토끼띠", "용띠", "뱀띠", "말띠", "양띠", "원숭이띠", "닭띠", "개띠", "돼지띠"}
	constellKeywords := [...]string{"물병자리", "물고기자리", "양자리", "황소자리", "쌍둥이자리", "게자리", "사자자리", "처녀자리", "천칭자리", "전갈자리", "사수자리", "염소자리"}
	constellDates := [...]string{"1월 20일 ~ 2월 18일", "2월 19일 ~ 3월 20일", "3월 21일 ~ 4월 19일", "4월 20일 ~ 5월 20일", "5월 21일 ~ 6월 21일", "6월 22일 ~ 7월 22일", "7월 23일 ~ 8월 22일", "8월 23일 ~ 9월 23일", "9월 24일 ~ 10월 22일", "10월 23일 ~ 11월 22일", "11월 23일 ~ 12월 24일", "12월 25일 ~ 1월 19일"}

	for _, keyword := range zooKeywords {
		var result string
		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=`+keyword),
			chromedp.Text(`div#yearFortune > div.infors > div.detail > p.text._cs_fortune_text`, &result),
		)
		if err != nil {
			log.Fatal(err)
		}

		fortunes = append(fortunes, models.Fortune{
			Name:    keyword,
			Content: result,
			DueDate: CurrentTime(),
		})
	}

	for index, keyword := range constellKeywords {
		var result string
		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=`+keyword),
			chromedp.Text(`div#yearFortune > div.infors > div.detail > p.text._cs_fortune_text`, &result),
		)
		if err != nil {
			log.Fatal(err)
		}

		fortunes = append(fortunes, models.Fortune{
			Name:    keyword,
			Content: constellDates[index] + " " + result,
			DueDate: CurrentTime(),
		})
	}

	db.Create(&fortunes)
	log.Printf("Complate fortunes : %d\n", len(fortunes))
}
