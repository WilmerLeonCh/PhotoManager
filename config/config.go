package config

type Config struct {
	PhotoManagerSlackWebHookURL  string `env:"PHOTO_MANAGER_SLACK_WEBHOOK_URL,required=true"`
	PhotoManagerSlackUserName    string `env:"PHOTO_MANAGER_SLACK_USERNAME,required=true"`
	PhotoManagerSlackChannelName string `env:"PHOTO_MANAGER_SLACK_CHANNEL_NAME,required=true"`
}
