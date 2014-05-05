module Lita
  module Handlers
    class Help < Handler
      def help_command(route, command, description)
        command = "#{name}#{command}" if route.command?

        "#{command} - #{description}"
      end
    end
  end
end
