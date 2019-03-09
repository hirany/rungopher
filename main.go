package main

import (
	"log"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/urfave/cli"
)

const (
	backColor = termbox.ColorDefault
	runeColor = termbox.ColorDefault
)

func main() {

	app := cli.NewApp()
	app.Name = "rungopher"
	app.Usage = "running gopher"
	app.Action = func(c *cli.Context) {
		err := runGopher()
		if err != nil {
			log.Fatal(err)
		}
	}
	app.Run(os.Args)
}

func runGopher() error {

	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	w, h := termbox.Size()

	var aa []rune
	// ʕ◔ϖ◔ʔ
	aa = append(aa, 'ʕ')
	aa = append(aa, '◔')
	aa = append(aa, 'ϖ')
	aa = append(aa, '◔')
	aa = append(aa, 'ʔ')

	for i := 0; i < w; i++ {
		termbox.Clear(backColor, backColor)
		for k := range aa {
			termbox.SetCell(w-i+k, h/2, rune(aa[k]), backColor, runeColor)
		}
		termbox.Flush()
		time.Sleep(30 * time.Millisecond)
	}
	return nil

}
