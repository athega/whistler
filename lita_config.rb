require "lita-athega"

Lita.configure do |config|
  # The name your robot will use.
  config.robot.name = "Whistler"

  # The severity of messages to log
  config.robot.log_level = :info

  # An array of user IDs that are considered administrators. These users
  # the ability to add and remove other users from authorization groups.
  # What is considered a user ID will change depending on which adapter you use.
  # config.robot.admins = ["1", "2"]

  # The adapter you want to connect with.
  config.robot.adapter = :campfire

  config.adapter.subdomain = ENV["CAMPFIRE_SUBDOMAIN"]
  config.adapter.apikey    = ENV["CAMPFIRE_APIKEY"]
  config.adapter.rooms     = ENV["CAMPFIRE_ROOMS"].split(',')
  config.adapter.debug     = false

  config.redis.url         = ENV["REDISCLOUD_URL"]
  config.http.port         = ENV["PORT"]
end
