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

    private

      def get_athegian_data(name)
        name = name.split(' ').first.downcase

        base_url = "http://athega.se/api/employees/"
        http_response = http.get("#{base_url}#{name}")
        MultiJson.load(http_response.body)
      rescue
        nil
      end
    end

    Lita.register_handler(Athega)
  end
end
