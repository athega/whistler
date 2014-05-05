require_relative "lib/lita/handlers/help"

Lita.configure do |config|
  # The name your robot will use.
  config.robot.name = "Whistler"
  config.robot.mention_name = "/"

  # The severity of messages to log
  config.robot.log_level = :info

  # An array of user IDs that are considered administrators. These users
  # the ability to add and remove other users from authorization groups.
  # What is considered a user ID will change depending on which adapter you use.
  # config.robot.admins = ["1", "2"]

  # The adapter you want to connect with.
  config.robot.adapter = :campfire

  config.adapter.tap do |a|
    a.subdomain      = ENV["CAMPFIRE_SUBDOMAIN"]
    a.apikey         = ENV["CAMPFIRE_APIKEY"]
    a.rooms          = String(ENV["CAMPFIRE_ROOMS"]).split(',')
    a.debug          = false
    a.tinder_options = { timeout: 30, user_agent: 'lita-campfire' }
  end

  config.redis.url = ENV["REDISCLOUD_URL"]
  config.http.port = ENV["PORT"]
end
