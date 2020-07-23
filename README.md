# Share HvH Server via Discord Bot

Load lua and press K to send server information to the discord bot.

## Build
```bash
go build .
```

## Usage
Create the config.json
```json
{
    "discordbot": {
        "token": "",
        "email": "",
        "password": "",
        "useToken": true
    },
    "shareServer": {
        "channelId": ""
    },
    "apikey": ""
}
```

You can use a token or just login with your own discord credentials. 

Specify a channel on a server, to get the channel id, enable developer mode and press right click on a channel and copy the channel id.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
