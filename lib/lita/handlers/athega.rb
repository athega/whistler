require "lita"

module Lita
  module Handlers
    class Athega < Handler
      route /athegian\s+(\w+)/,
        :athegian, command: true, help: {
          "athegian NAME" => "Retrieves image and position from the Athega API"
        }

      def athegian(response)
        name = response.matches[0][0]

        if data = get_athegian_data(name)
          response.reply(data['medium_image_url'])
          response.reply(data['name'])
          response.reply(data['position'])
        end
      end

      route /^\/aww$/, :aww, command: false, help: {
        "/aww" => "Random cute image"
      }

      def aww(response)
        url = get_random_reddit_field('aww/top.json?limit=10', :url)

        response.reply(url) if url
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
