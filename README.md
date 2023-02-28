# Telegram Cloner
This is a simple project that uses [TDLight Fork of Telegram Bot API Server](https://tdlight-team.github.io/tdlight-telegram-bot-api/#/) to fullfill a very simple purpose:
forwarding transparently messages from one chat to another, behaving like a real user.

## Requirements
- Docker
- Values to fill `.env` file

## Usage
- `make core`
- `make run $command`

Available commands are:
- `configure`: Setups the core to access and behave like user
- `sync`: Starts a sync process