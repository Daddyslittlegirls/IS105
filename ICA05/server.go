package main
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"time"
	"strconv"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

type SiteData struct {
	Name string
	Description string
	MapData string
}

//var count int = 0

func main() {

	m := martini.Classic()

	m.Use( render.Renderer(render.Options{
		IndentJSON: true, // so we can read it..
	}))

	m.Get("/", func(r render.Render, x *http.Request) {
		place := string(x.FormValue("place"))
		place = strings.Replace(place, " ", "+", -1)

		if len(place) <= 0 {place = "UIA+Kristiansand"}
		r.HTML(200, "index", SiteData{"Gruppe 6", "ICA05", strconv.Itoa(count), getServerIP(), place})
	})





}
