package main

import (
	"fmt"
	"time"

	"github.com/claudiu/gocron"
)

func main() {
	gocron.Start()
	gocron.Every(1).Day().Do(task)
	//gocron.Every(5).Seconds().Do(task)
	//gocron.Every(10).Seconds().Do(vijay)

	//time.Sleep(20 * time.Second)

	//gocron.Clear()
	//fmt.Println("All task removed")
	for {
	}
}
func task() {
	fmt.Println("I am runnning task.", time.Now())
}
func vijay() {
	fmt.Println("I am runnning vijay.", time.Now())
}
func test(stop chan bool) {
	time.Sleep(20 * time.Second)
	gocron.Clear()
	fmt.Println("All task removed")
	close(stop)
}

/*db, err := database.New("postgres", "postgresql://root@localhost:26257/a_word_a_day?sslmode=disable")
defer db.Client.Close()
if err != nil {
	fmt.Println(err)
}

_w := "Heyhey"
_m := "A Wish"
_t := "noun"
_u := "jitenp@outlook.com"
word := &models.Word{Word: _w, Meaning: _m, Type: _t, UpdatedBy: _u}
//word := &models.Word{}
//word.Word = "Heyhey"
err = db.InsertWord(word)
if err != nil {
	fmt.Println(err)
}

_id := "44a5bdbe-50f9-450d-b86c-b042cce9b99e"
word = &models.Word{ID: _id}
//*word.ID = "44a5bdbe-50f9-450d-b86c-b042cce9b99e"
//db.Client.Model(word).Updates(models.Word{Meaning: "Everyday"})
err = db.UpdateWord(word, map[string]interface{}{"word": "Tupuk Tupuk"})
if err != nil {
	fmt.Println(err)
}
//fmt.Println(db.Find())

fmt.Println(db.FindWordByWord("Tupuk Tupuk"))

fmt.Println(db.FindSentencesByWord("Tupuk Tup"))*/
