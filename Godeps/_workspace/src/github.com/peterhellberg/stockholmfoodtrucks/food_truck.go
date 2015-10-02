package stockholmfoodtrucks

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/PuerkitoBio/goquery"
)

// ErrFoodTruckNotFound is returned when no food truck was found
var ErrFoodTruckNotFound = fmt.Errorf(`Food truck not found`)

// FoodTruck contains the information for a food truck
type FoodTruck struct {
	Name        string    `json:"name"`
	Text        string    `json:"text"`
	TimeText    string    `json:"time_text"`
	Time        time.Time `json:"time"`
	Location    *Location `json:"location"`
	Slug        string    `json:"slug"`
	Hex         string    `json:"hex"`
	Facebook    string    `json:"facebook"`
	Instagram   string    `json:"instagram"`
	Twitter     string    `json:"twitter"`
	Web         string    `json:"web"`
	Description []string  `json:"description"`
	Menu        []string  `json:"menu"`
}

// Pic returns the link to the food truck image
func (ft FoodTruck) Pic() string {
	if ft.Slug == "" {
		return ""
	}

	return fmt.Sprintf("http://stockholmfoodtrucks.nu/img-dist/%s.png", ft.Slug)
}

// Location contains the location information for a food truck
type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
	Type string `json:"type"`
}

// All returns a slice of FoodTruck
func (c *Client) All() ([]FoodTruck, error) {
	doc, err := c.NewDocument("")
	if err != nil {
		return nil, err
	}

	return c.FoodTrucks(doc)
}

// Get returns a FoodTruck
func (c *Client) Get(slug string) (FoodTruck, error) {
	doc, err := c.NewDocument("/" + slug + "/")
	if err != nil || !doc.Find(".single-truck").Is(".single-truck") {
		return FoodTruck{}, ErrFoodTruckNotFound
	}

	return c.FoodTruck(doc, slug)
}

// FoodTruck extracts food truck from goquery document
func (c *Client) FoodTruck(doc *goquery.Document, slug string) (FoodTruck, error) {
	t := doc.Find(".single-truck")

	truckName := t.Find(".truck-content .main-title").Text()
	truckText := ""

	for _, n := range t.Find(".truck-post .bubble-inner *").Nodes {
		if n.Type == html.TextNode && n.PrevSibling != nil && n.PrevSibling.DataAtom == atom.Br {
			truckText += strings.Trim(n.Data, "\n ") + "\n"
		}
	}

	truckMenu := []string{}

	t.Find(".truck-content .menu ul li").Each(func(i int, s *goquery.Selection) {
		truckMenu = append(truckMenu, s.Text())
	})

	truckFacebook := t.Find(".truck-content .main a.facebook").AttrOr("href", "")
	truckInstagram := t.Find(".truck-content .main a.instagram").AttrOr("href", "")
	truckTwitter := t.Find(".truck-content .main a.twitter").AttrOr("href", "")
	truckWeb := t.Find(".truck-content .main a.web").AttrOr("href", "")

	truckDescription := []string{}

	t.Find(".truck-content .main p").Each(func(i int, s *goquery.Selection) {
		if !s.Children().HasClass("link-to") {
			truckDescription = append(truckDescription, s.Text())
		}
	})

	truckTime, _ := time.Parse("2006-01-02 15:04", t.Find(".meta a").First().AttrOr("title", ""))
	truckTimeText := t.Find(".meta a").First().Text()

	foodTruck := FoodTruck{
		Slug:        slug,
		Name:        truckName,
		Hex:         nameToHex(truckName),
		Text:        truckText,
		Menu:        truckMenu,
		Facebook:    truckFacebook,
		Instagram:   truckInstagram,
		Twitter:     truckTwitter,
		Web:         truckWeb,
		Description: truckDescription,
		Time:        truckTime,
		TimeText:    truckTimeText,
	}

	location := t.Find(".bubble-inner .location").First()

	if id, exists := location.Attr("data-id"); exists {
		if n, exists := location.Attr("data-name"); exists {
			if t, exists := location.Attr("data-type"); exists {
				foodTruck.Location = &Location{
					ID:   id,
					Name: n,
					Type: t,
					Text: location.Text(),
				}
			}
		}
	}

	return foodTruck, nil
}

// FoodTrucks extracts slice of food trucks from goquery document
func (c *Client) FoodTrucks(doc *goquery.Document) ([]FoodTruck, error) {
	foodTrucks := []FoodTruck{}

	doc.Find(".trucks-list .truck").Each(func(i int, s *goquery.Selection) {
		truckName := s.Find(".truck-name").Text()
		truckSlug := strings.Trim(s.Find(".truck-name a").AttrOr("href", ""), "/")

		post := s.Find(".posts .post").First()
		truckText := post.Find(".content").Text()
		truckTime, _ := time.Parse("2006-01-02 15:04", post.Find(".meta a").First().AttrOr("title", ""))
		truckTimeText := post.Find(".meta a").First().Text()

		foodTruck := FoodTruck{
			Name:     truckName,
			Slug:     truckSlug,
			Hex:      nameToHex(truckName),
			Text:     truckText,
			Time:     truckTime,
			TimeText: truckTimeText,
		}

		location := post.Find(".content .location").First()

		if id, exists := location.Attr("data-id"); exists {
			if n, exists := location.Attr("data-name"); exists {
				if t, exists := location.Attr("data-type"); exists {
					foodTruck.Location = &Location{
						ID:   id,
						Name: n,
						Type: t,
						Text: location.Text(),
					}
				}
			}
		}

		foodTrucks = append(foodTrucks, foodTruck)
	})

	return foodTrucks, nil
}

func nameToHex(name string) string {
	if hex, found := map[string]string{
		"Boardwalk Streetfood": "#bd0a10",
		"Bon Coin":             "#4a4a4a",
		"Bun Bun Truck":        "#bfecea",
		"Chilibussen":          "#f2900c",
		"Curbside Sthlm":       "#3b3b3b",
		"El Taco Truck":        "#f38ab3",
		"Foodtruck Odjuret":    "#ff5e2c",
		"Frankie’s Coffee":     "#98c6e2",
		"Fred’s Food Truck":    "#c9285b",
		"Frick & Hagberg":      "#2064c0",
		"Funky Chicken":        "#b3b8bc",
		"Grillmobilen":         "#005ad9",
		"Indian Street Food":   "#ff5e2c",
		"Jeffrey’s Food Truck": "#0dd15d",
		"Kantarellkungen":      "#ee3942",
		"Köftebilen":           "#c0c3c3",
		"Punto Sur":            "#b71c27",
		"Rolling Street Food":  "#a3be79",
		"Silvias":              "#cb0000",
		"SOOK Streetfood":      "#1196cc",
		"SWAT street food":     "#ff7200",
		"Van Helleputte":       "#1b194b",
	}[name]; found {
		return hex
	}

	return "#000000"
}
