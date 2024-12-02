package main

import (
	"fmt"
	"regexp"
)

type Ad struct {
	Title       string
	Описание string
}

func main() {
	ads := []Ad{
		{
			Title:       "Куплю велосипед MeRiDa",
			Описание: "Куплю велосипед meriDA в хорошем состоянии.",
		},
		{
			Title:       "Продам ВаЗ 2101",
			Описание: "Продам ваз 2101 в хорошем состоянии.",
		},
		{
			Title:       "Продам БМВ",
			Описание: "Продам бМв в хорошем состоянии.",
		},
		{
			Title:       "Продам macBook pro",
			Описание: "Продам macBook PRO в хорошем состоянии.",
		},
	}

	ads = censorAds(ads, map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	})

	for _, ad := range ads {
		fmt.Println(ad.Title)
		fmt.Println(ad.Описание)
		fmt.Println()
	}
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	replaceText := func(txt string, censor map[string]string) string {
		for i, d := range censor {
			re := regexp.MustCompile("(?i)" + i)
			txt = re.ReplaceAllString(txt, d)
		}

		return txt
	}

	for i := range ads {
		ads[i].Title = replaceText(ads[i].Title, censor)
		ads[i].Описание = replaceText(ads[i].Описание, censor)
	}

	return ads
}
