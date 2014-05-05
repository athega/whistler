require "lita"

require "nokogiri"
require "open-uri"

module Lita
  module Handlers
    class Athega < Handler
      route /^athegian\s+(\w+)/, :athegian, command: true,
        help: { "athegian NAME" => "Retrieves image and position for employee" }

      def athegian(response)
        name = arg(response)

        if data = get_athegian_data(name)
          response.reply(data['medium_image_url'])
          response.reply(data['name'])
          response.reply(data['position'])
        else
          response.reply("Sorry, no such employee at Athega.")
        end
      end

      # http://open.spotify.com/track/659TMQbtxXIBE8AVWomlvx
      route /(http:\/\/open.spotify.com\/\w+\/\w+)/, :spotify, command: false,
        help: { "[Spotify URL]" => "Retrieves image and title from any Spotify URL" }

      def spotify(response)
        doc = Nokogiri::HTML open(arg(response))

        twitter_images = doc.css('meta[property="twitter:image"]')

        if twitter_images.any?
          response.reply twitter_images.first['content'] + "#.jpg"
        end

        descriptions = doc.css('meta[property="og:description"]')

        if descriptions.any?
          response.reply descriptions.first['content'].gsub(" on Spotify.", "")
        end
      end

      route /^aww$/, :aww,
        help: { "/aww" => "Random cute image" }

      def aww(response)
        url = get_random_reddit_field('aww/top.json?limit=10', :url)

        response.reply(url) if url
      end

      http.get "/", :whistler

      def whistler(request, response)
        response.headers["Content-Type"] = "text/html"

        logo_url = 'https://raw.github.com/athega/whistler/master/images/whistler.jpg'

        html = "<a href='https://github.com/athega/whistler'>" +
               "<img src='#{logo_url}' alt='Whistler' />" +
               "</a>"

        response.write(html)
      end

    private

      def arg(response)
        response.matches[0][0]
      end

      def get_random_reddit_field(path, field)
        url = reddit_url(path)

        if data = json_get(url)
          post = data['data']['children'].sample
          post['data'][field.to_s]
        end
      end

      def reddit_url(path)
        "http://api.reddit.com/r/#{path}"
      end

      def get_athegian_data(name)
        json_get("http://athega.se/api/employees/#{name.downcase}")
      end

      def json_get(url)
        response = http.get(url)

        MultiJson.load(response.body)
      end
    end

    Lita.register_handler(Athega)
  end
end
