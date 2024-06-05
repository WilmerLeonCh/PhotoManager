# Photo-Manager
✨ A console application to manage photos from the JSONPlaceHolder API and notify Slack when a photo is created or deleted.

# Demo
![demo](demo.gif)

## Built with
- [Go](https://golang.org/)
- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Slack API](https://api.slack.com/)
- [JSONPlaceHolder API](https://jsonplaceholder.typicode.com/)

## Usage:
- Main Option List:
    - <kbd>Up</kbd> <kbd>Down</kbd> to switch between options
    - <kbd>Enter</kbd> to select an option
    - <kbd>Ctrl + C</kbd> <kbd>Esc</kbd> to exit the application
- Create:
    - <kbd>Up</kbd> <kbd>Down</kbd> to switch between fields
    - <kbd>Enter</kbd> to save the photo (title, url *required*)
    - <kbd>Ctrl + C</kbd> <kbd>Esc</kbd> cancel action and return to the main menu
- List & Search:
    - <kbd>Up</kbd> <kbd>Down</kbd> to switch between rows
    - <kbd>Ctrl + Left</kbd> <kbd>Ctrl + Right</kbd> to switch between pagination
    - <kbd>Enter</kbd> to search by input title value
    - <kbd>Ctrl + C</kbd> <kbd>Esc</kbd> cancel action and return to the main menu
- Update:
    - <kbd>Up</kbd> <kbd>Down</kbd> to switch between fields
    - <kbd>Enter</kbd> to save the photo (id, title, url *required*)
    - <kbd>Ctrl + C</kbd> <kbd>Esc</kbd> cancel action and return to the main menu
- Delete:
    - <kbd>Enter</kbd> to delete the photo (id *required*)
    - <kbd>Ctrl + C</kbd> <kbd>Esc</kbd> cancel action and return to the main menu

## Slack configuration:
- Create a Slack app: https://api.slack.com/apps
- Enable incoming webhooks: https://api.slack.com/messaging/webhooks
- Add new webhook to workspace and pick a channel that the app will post to, then select authorize
- Add the webhook URL to the environment variables
- Add some username and the channel name picked to the environment variables
- Use your incoming webhook URL to request

## Environment variables:
- `PHOTO_MANAGER_SLACK_WEBHOOK_URL`: Slack  incoming webhook URL
- `PHOTO_MANAGER_SLACK_USERNAME`: Slack username
- `PHOTO_MANAGER_SLACK_CHANNEL_NAME`: Slack channel name

## Description problem
EDteam needs to query the Photo API information found in: https://jsonplaceholder.typicode.com/

We need a **console application** made in Go that allows us to manage the information from said API. Therefore, when the application starts, a menu will appear where the action to be performed will be selected (create, consult, list, modify, delete).

It's very important to us to know when a person creates or deletes a photo, so we need **Slack to be notified when those actions happen**. We leave the format of the notification to your imagination.
