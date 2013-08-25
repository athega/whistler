Whistler
=====

![Whistler avatar](https://github.com/athega/athega-hubot/raw/master/images/whistler.jpg)

[Whistler](http://starwars.wikia.com/wiki/Whistler), also called Xeno, was the astromech droid of Corran Horn.
It was the same R-series model as the legendary R2-D2.

_Whistler is based on [Lita](http://lita.io/). He's pretty cool._

Playing with Whistler
=====================

Deployment
==========

Campfire Variables
------------------

Create a separate user for your bot and get their token from the web UI.

    % heroku config:add HUBOT_CAMPFIRE_TOKEN="..."

Get the numeric ids of the rooms you want the bot to join, comma
delimited. If you want the bot to connect to `https://mysubdomain.campfirenow.com/room/42` 
and `https://mysubdomain.campfirenow.com/room/1024` then you'd add it like this:

    % heroku config:add HUBOT_CAMPFIRE_ROOMS="447885"

Add the subdomain hubot should connect to. If you web URL looks like
`http://mysubdomain.campfirenow.com` then you'd add it like this:

    % heroku config:add HUBOT_CAMPFIRE_ACCOUNT="athega"

Restarting Whistler
-------------------
You may want to get comfortable with `heroku logs` and `heroku restart`
if you're having issues.
