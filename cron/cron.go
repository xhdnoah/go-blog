package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
	"github.com/xhdnoah/go-blog/models"
)

func main() {
	log.Println("Starting ...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("RUn models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for range t1.C {
		t1.Reset(time.Second * 10)
	}
}
