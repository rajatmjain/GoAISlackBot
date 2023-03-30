GoSlackAIBot
============

GoSlackAIBot is a chatbot built using the Go programming language and the Slacker library for interacting with the Slack API. It uses OpenAI's GPT-3 language model to provide responses to questions asked in Slack.

Getting Started
---------------

To get started with GoSlackAIBot, you will need to have a Slack bot token and app token, as well as an API key for OpenAI's GPT-3 language model. These can be obtained by following the respective documentation for each service.

Once you have obtained the necessary credentials, you can create a .env file in the root directory of the project and set the following variables:

```.env
SLACK_BOT_TOKEN=<your Slack bot token>
SLACK_APP_TOKEN=<your Slack app token> 
API_KEY=<your OpenAI GPT-3 API key>
```


After setting the environment variables, you can run the bot by running `go run main.go` in the root directory of the project.

Usage
-----

To use GoSlackAIBot, simply invite it to a channel in Slack and ask it a question using the command `<@bot-name> <question>`. The bot will then use the GPT-3 language model to generate a response to your question.

Contributing
------------

If you would like to contribute to GoSlackAIBot, please feel free to fork the repository and submit a pull request. Contributions are always welcome!
