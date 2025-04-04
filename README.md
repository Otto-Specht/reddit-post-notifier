# Reddit post nofifier

This is a project made in go which polls a reddits user account for new posts and sends a discord message on a server.

## Configuration:
Create and configure a .env file like this:
```env
# .env file

#Configuration
USER_NAMES=user1,user2 # List of users to poll (without u/)

#Reddit API
REDDIT_CLIENT_ID=<id> # Reddit app client id
REDDIT_CLIENT_SECRET=<secret> # Reddit app client secret

#Disocrd APi
DISCORD_BOT_TOKEN=<token> # Discord bot token
DISCORD_CHANNEL_ID=<id> # Discord channel id

#Optional
LOG_LEVEL=DBG # DBG/INF/WRN/ERR
POLL_INTERVAL=10m # Poll interval (default 15m)
```