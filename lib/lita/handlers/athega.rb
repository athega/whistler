require "lita"

module Lita
  module Handlers
    class Athega < Handler
      # coderwall <coderwall username> - Returns coder achievements from coderwall.com
      # hn.top - refer to the top item on hn
      # hn[i] - refer to the ith item on hn
      # hubot 9gag me - Returns a random meme image
      # hubot <keyword> tweet - Returns a link to a tweet about <keyword>
      # hubot abstract <topic> - Prints a nice abstract of the given topic
      #   hubot eval me <lang> <code> - evaluate <code> and show the result
      # hubot hn top <N> - get the top N items on hacker news (or your favorite RSS feed)
      #   hubot md5|sha|sha1|sha256|sha512|rmd160 me <string> - Generate hash of <string>
      # hubot mustache me <query> - Searches Google Images for the specified query and mustaches it.
      #   hubot mustache me <url> - Adds a mustache to the specified URL.
      #   hubot reddit (me) <reddit> [limit] - Lookup reddit topic
      # hubot roll <x>d<y> - roll x dice, each of which has y sides
      # hubot roll dice - Roll two six-sided dice
      # hubot ruby|rb <script> - Evaluate one line of Ruby script
      # ship it - Display a motivation squirrel
      # yoda pic - Returns a random quote
      # yoda quote - Returns a random yoda quote
      # hubot xkcd - The latest XKCD comic
      # hubot xkcd <num> - XKCD comic <num>

      route /^\/athegian\s+(\w+)/, :athegian,
        help: { "/athegian NAME" => "Retrieves image and position for employee" }

      def athegian(response)
        name = response.matches[0][0]

        if data = get_athegian_data(name)
          response.reply(data['medium_image_url'])
          response.reply(data['name'])
          response.reply(data['position'])
        else
          response.reply("Sorry, no such employee at Athega.")
        end
      end

      route /^\/xkcd\s?(\d+)?/, :xkcd, help: {
        "/xkcd"     => "The latest XKCD comic",
        "/xkcd NUM" => "XKCD comic <num>"
      }

      def xkcd(response)
        num = response.matches[0][0]

        data = num ? get_xkcd_data(num) : get_latest_xkcd_data

        if data
          response.reply(data['img'])
          response.reply(data['title'])
          response.reply(data['alt'])
        end
      end

      route /^\/aww$/, :aww,
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

      def get_latest_xkcd_data
        json_get('http://xkcd.com/info.0.json')
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
        MultiJson.load(http.get(url).body)
      rescue
        nil
      end
    end

    Lita.register_handler(Athega)
  end
end
