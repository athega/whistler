require "lita"

module Lita
  module Handlers
    class Athega < Handler
      # coderwall <coderwall username> - Returns coder achievements from coderwall.com
      # haters - Returns a random haters gonna hate url
      # hn.top - refer to the top item on hn
      # hn[i] - refer to the ith item on hn
      # hubot 9gag me - Returns a random meme image
      # hubot <keyword> tweet - Returns a link to a tweet about <keyword>
      # hubot abstract <topic> - Prints a nice abstract of the given topic
      # hubot animate me <query> - The same thing as `image me`, except adds a few parameters to try to return an animated GIF instead.
      #   hubot eval me <lang> <code> - evaluate <code> and show the result
      # hubot gem whois <gemname> - returns gem details if it exists
      # hubot hn top <N> - get the top N items on hacker news (or your favorite RSS feed)
      # hubot image me <query> - The Original. Queries Google Images for <query> and returns a random top result.
      #   hubot md5|sha|sha1|sha256|sha512|rmd160 me <string> - Generate hash of <string>
      # hubot mustache me <query> - Searches Google Images for the specified query and mustaches it.
      #   hubot mustache me <url> - Adds a mustache to the specified URL.
      #   hubot reddit (me) <reddit> [limit] - Lookup reddit topic
      # hubot roll <x>d<y> - roll x dice, each of which has y sides
      # hubot roll dice - Roll two six-sided dice
      # hubot ruby|rb <script> - Evaluate one line of Ruby script
      # hubot there's a gem for <that> - Returns a link to a gem on rubygems.org
      # hubot xkcd - The latest XKCD comic
      # hubot xkcd <num> - XKCD comic <num>
      # ship it - Display a motivation squirrel
      # yoda pic - Returns a random quote
      # yoda quote - Returns a random yoda quote



      route /^\/athegian\s+(\w+)/, :athegian,
        help: { "/athegian NAME" => "Retrieves image and position for employee" }

      def athegian(response)
        name = response.matches[0][0]

        if data = get_athegian_data(name)
          response.reply(data['medium_image_url'])
          response.reply(data['name'])
          response.reply(data['position'])
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
