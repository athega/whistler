require "lita"

module Lita
  module Handlers
    class Athega < Handler
      route /athegian\s+(\w+)/, :athegian, command: true, help: {
        "athegian NAME" => "Retrieves image and position from the Athega API"
      }

      def athegian(response)
        name = response.matches[0][0]
        response.reply("Athegian: '#{name}'")
      end
    end

    Lita.register_handler(Athega)
  end
end
