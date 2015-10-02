package stockholmfoodtrucks

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPic(t *testing.T) {
	for i, tt := range []struct {
		slug string
		want string
	}{
		{"", ""},
		{"foo", "http://stockholmfoodtrucks.nu/img-dist/foo.png"},
		{"bar", "http://stockholmfoodtrucks.nu/img-dist/bar.png"},
	} {
		if got := (FoodTruck{Slug: tt.slug}).Pic(); got != tt.want {
			t.Fatalf(`[%d] ft.Pic() = %q, want %q`, i, got, tt.want)
		}
	}
}

func TestAll(t *testing.T) {
	bodyString := `<div class="trucks-list" role="main">
		<div class="truck truck-jeffreys-food-truck">
			<h2 class="truck-name"><a href="/jeffreys-food-truck/">Jeffrey’s Food Truck</a></h2>
			<ul class="posts">
				<li class="post old-post">
					<span class="content">Amigos tyvärr fanns det ingen plats för oss i <span class="location js-location" data-id="332" data-name="Hälsingegatan" data-type="street">hälsingegatan</span> så vi befinner oss just nu i kungsholmstorget</span>
          <span class="meta">
          	<a href="/jeffreys-food-truck/amigos-tyvarr-fanns-det-ingen-plats-for-oss-i-halsingegatan/" class="date" title="2015-09-04 11:14">4 dagar sedan</a>
            <a href="http://www.facebook.com/642730225762785_881233348579137" target="_blank">Visa inlägget</a>
          </span>
				</li>
			</ul>
		</div>
		<div class="truck truck-bun-bun-truck">
			<h2 class="truck-name"><a href="/bun-bun-truck/">Bun Bun Truck</a></h2>
			<ul class="posts">
				<li class="post old-post">
					<span class="content">Pling, hej söndag! Vi är lite möra (och glada) efter stöket på Popaganda igår men masar oss iväg till <span class="location js-location" data-id="1048" data-name="Hornstulls marknad" data-type="event" data-coordinates="{&quot;lat&quot;:&quot;59.3151364&quot;,&quot;lng&quot;:&quot;18.0308404&quot;}">Hornstulls Marknad</span> ändå! Ses!</span>
					<span class="meta">
          	<a href="/bun-bun-truck/pling-hej-sondag-vi-ar-lite-mora-och-glada-efter/" class="date" title="2015-08-30 10:42">9 dagar sedan</a>
            <a href="http://www.facebook.com/567605503274218_1046635032037927" target="_blank">Visa inlägget</a>
          </span>
				</li>
			</ul>
		</div>
	</div>`

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(bodyString))
		},
	))
	defer server.Close()

	c := NewClient(&http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}})

	trucks, err := c.All()
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if got, want := len(trucks), 2; got != want {
		t.Fatalf(`len(trucks) = %d, want %d`, got, want)
	}

	if got, want := trucks[0].Name, "Jeffrey’s Food Truck"; got != want {
		t.Fatalf(`trucks[0].Name = %q, want %q`, got, want)
	}

	if got, want := trucks[1].Name, "Bun Bun Truck"; got != want {
		t.Fatalf(`trucks[1].Name = %q, want %q`, got, want)
	}

	if got, want := trucks[0].Slug, "jeffreys-food-truck"; got != want {
		t.Fatalf(`trucks[0].Slug = %q, want %q`, got, want)
	}

	if got, want := trucks[1].Slug, "bun-bun-truck"; got != want {
		t.Fatalf(`trucks[1].Slug = %q, want %q`, got, want)
	}

	if got, want := trucks[0].TimeText, "4 dagar sedan"; got != want {
		t.Fatalf(`trucks[0].TimeText = %q, want %q`, got, want)
	}

	if got, want := trucks[1].TimeText, "9 dagar sedan"; got != want {
		t.Fatalf(`trucks[1].TimeText = %q, want %q`, got, want)
	}
}

func TestGet(t *testing.T) {
	bodyString := `<div class="single-truck truck-bun-bun-truck">
		<div class="bubble-inner">
    	<img src="https://fbcdn-photos-d-a.akamaihd.net/hphotos-ak-xpt1/v/t1.0-0/s130x130/1898063_1051737591527671_2838314519516257804_n.jpg?oh=05c259b2a5884e47590bd3c5724f7909&amp;oe=56A7F76D&amp;__gda__=1453776203_bc3977a9a1581d236b97d977b700c39d" alt="" class="image js-overlay-image" data-image="https://scontent.xx.fbcdn.net/hphotos-xpt1/v/t1.0-9/s720x720/1898063_1051737591527671_2838314519516257804_n.jpg?oh=679bd247c4065e6ba433d609f5acb8d1&amp;oe=56627C92">
      <strong>Det senaste från Bun Bun Truck:</strong><br>
      <span class="location js-location" data-id="445" data-name="Kungsbroplan" data-type="street">Kungsbroplan</span>
			Hej! Boka catering från Bun Bun till nästa festkonferensbröllopsmiddagsmingel! Bánh Mi - baguetter, Bún - nudlar eller Viet Tacos, i en budbil från vårt kök till dig. Mer info på <a href="http://bunbun.se/catering" target="_blank" rel="nofollow">http://bunbun.se/catering</a>, boka på <a href="mailto:info@bunbun.se">info@bunbun.se</a> eller 0721-91 92 93.
      <span class="meta">
      	<a href="/bun-bun-truck/hej-boka-catering-fran-bun-bun-till-nasta-festkonferensbrollopsmiddagsmingel-bnh/" class="date" title="2015-09-08 11:25">12 timmar sedan</a>
        <a href="https://www.facebook.com/bunbunsthlm/photos/a.647632215271546.1073741832.567605503274218/1051737591527671/?type=1" target="_blank">Visa inlägget</a>
      </span>
    </div>
		<div class="truck-content">
			<div class="col col-w8 main">
				<h1 class="main-title">Bun Bun Truck</h1>
				<p>Bun Bun Truck har specialiserat sig på baguetter och serverar idag fyra olika varianter av Bánh Mì:  citrongräsmarinerad kyckling, brässerad fläsksida, grillad tofu och klassikern med tre sorters kallskuret.</p>
				<p>Bánh Mì är en klassisk maträtt med rötterna i Vietnam som föddes när människor började fylla baguetter med inhemska specialister: koriander, gurka, pickles, majonäs med marinerat kött.</p>
				<p>Bun Bun Truck cirkulerar i Stockholm med omnejd – på gator och torg, festivaler och marknader.</p>
				<p><a href="https://www.facebook.com/bunbunsthlm" target="_blank" class="link-to facebook">Facebook</a></p>
				<p><a href="https://www.twitter.com/BunBunSthlm" target="_blank" class="link-to twitter">Twitter</a></p>
				<p><a href="http://instagram.com/bunbunsthlm" target="_blank" class="link-to instagram">Instagram</a></p>
				<p><a href="http://www.bunbun.se/" target="_blank" class="link-to web">Webbplats</a></p>
			</div>
			<div class="menu">
				<h2 class="menu-title">Meny</h2>
				<p><small>Menyerna kan skilja sig från dag till dag, så se detta som en hint om vad som finns.</small></p>
				<ul>
					<li>Citrongräsmarinerad kyckling</li>
					<li>Brässerad fläsksida</li>
					<li>Marinerad tofu</li>
					<li>Special combo</li>
					<li>Vietnamesisk pate, skinka, terrine på grishuvud</li>
				</ul>
			</div>
		</div>
	</div>`

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(bodyString))
		},
	))
	defer server.Close()

	c := NewClient(&http.Client{Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}})

	truck, err := c.Get("bun-bun-truck")
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	if got, want := truck.Slug, "bun-bun-truck"; got != want {
		t.Fatalf(`truck.Slug = %q, want %q`, got, want)
	}

	if got, want := truck.Name, "Bun Bun Truck"; got != want {
		t.Fatalf(`truck.Name = %q, want %q`, got, want)
	}

	if got, want := truck.Hex, "#bfecea"; got != want {
		t.Fatalf(`truck.Hex = %q, want %q`, got, want)
	}

	if got, want := len(truck.Menu), 5; got != want {
		t.Fatalf(`len(truck.Menu) = %d, want %d`, got, want)
	}

	if got, want := truck.Menu[2], "Marinerad tofu"; got != want {
		t.Fatalf(`truck.Menu[2] = %q, want %q`, got, want)
	}

	if got, want := truck.Facebook, "https://www.facebook.com/bunbunsthlm"; got != want {
		t.Fatalf(`truck.Facebook = %q, want %q`, got, want)
	}

	if got, want := truck.Instagram, "http://instagram.com/bunbunsthlm"; got != want {
		t.Fatalf(`truck.Instagram = %q, want %q`, got, want)
	}

	if got, want := truck.Twitter, "https://www.twitter.com/BunBunSthlm"; got != want {
		t.Fatalf(`truck.Twitter = %q, want %q`, got, want)
	}

	if got, want := truck.Web, "http://www.bunbun.se/"; got != want {
		t.Fatalf(`truck.Web = %q, want %q`, got, want)
	}

	if got, want := len(truck.Description), 3; got != want {
		t.Fatalf(`len(truck.Description) = %d, want %d`, got, want)
	}

	if got, want := truck.TimeText, "12 timmar sedan"; got != want {
		t.Fatalf(`truck.TimeText = %q, want %q`, got, want)
	}

	if truck.Location == nil {
		t.Fatalf(`truck.Location == nil`)
	}

	if got, want := truck.Location.Text, "Kungsbroplan"; got != want {
		t.Fatalf(`truck.Location.Text = %q, want %q`, got, want)
	}

	if got, want := truck.Location.Type, "street"; got != want {
		t.Fatalf(`truck.Location.Type = %q, want %q`, got, want)
	}
}

func TestNameToHex(t *testing.T) {
	for i, tt := range []struct {
		in   string
		want string
	}{
		{"Chilibussen", "#f2900c"},
		{"El Taco Truck", "#f38ab3"},
		{"Foo Bar", "#000000"},
		{"Punto Sur", "#b71c27"},
	} {
		if got := nameToHex(tt.in); got != tt.want {
			t.Fatalf(`[%d] nameToHex(%q) = %q, want %q`, i, tt.in, got, tt.want)
		}
	}
}
