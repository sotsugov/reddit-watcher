# Reddit Watcher

[![Current Release](https://img.shields.io/github/tag/sotsugov/reddit-watcher.svg)](https://img.shields.io/github/tag/sotsugov/reddit-watcher.svg)  [![CircleCI](https://circleci.com/gh/sotsugov/reddit-watcher/tree/master.svg?style=shield)](https://circleci.com/gh/sotsugov/reddit-watcher/tree/master) [![Coverage Status](https://coveralls.io/repos/github/sotsugov/reddit-watcher/badge.svg?branch=master)](https://coveralls.io/github/sotsugov/reddit-watcher?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/sotsugov/reddit-watcher)](https://goreportcard.com/report/github.com/sotsugov/reddit-watcher) [![Docker Pulls](https://img.shields.io/docker/pulls/sotsugov/reddit-watcher.svg)](https://hub.docker.com/r/sotsugov/reddit-watcher/) [![Chat Now](https://img.shields.io/badge/telegram-MechKeyBot-blue.svg?style=social)](https://telegram.me/MechKeyBot)

Watches specific subreddits (`/r/mechmarket`) for items matching a specific pattern.

## Running the Bot

Running the bot is easy!  You just need a Telegram token that you can get from the [BotFather](https://core.telegram.org/bots#3-how-do-i-create-a-bot).

Once you have that, start up your server with the following command:

```bash
docker run -v `pwd`/config:/config sotsugov/reddit-watcher:latest --token ${TELEGRAM_TOKEN}
```

In this example, I'm running the container with settings being saved to a local directory.

## Using the Bot

The bot responds to private or group messages that look like a command (start with a `/`).

### Notification

The most basic usage is to monitor for posts that match your keywords.  The following commands will subscribe (or unsubscribe) you on new posts matching your keywords.  If you leave the keyword empty, it defaults to `*` which is ALL posts.

#### `/selling <keyword>`

Look for items matching that keyword that are being sold.  Sold means the listing includes "cash" or "paypal" in the "want" field.

#### `/buying <keyword>`

Look for items matching that keyword that are being bought.  Sold means the listing includes "cash" or "paypal" in the "have" field.

#### `/vendor <keyword>`

Look for items matching that keyword posted by a vendor.

#### `/artisan <keyword>`

Look for items matching that keyword posted by an artisan.

#### `/groupbuy <keyword>`

Look for items matching that keyword posted as a group buy.

#### `/interestcheck <keyword>`

Look for items matching that keyword posted as an interest check.

#### `/giveaway <keyword>`

Look for items matching that keyword posted as a giveaway.

### Other

#### `/help`

Replies with a simple help message listing all the available commands.

#### `/items`

Outputs a list of your keywords and the number of matches found so far.

#### `/stats`

Outputs interesting information about the current bot.
