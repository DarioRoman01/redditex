package env

// define config properties via env variables
// the env-default variables ARE ONLY FOR DEVELOPMENT
type Properties struct {
	Port         string `env:"MY_APP_PORT" env-default:"1323"`
	Host         string `env:"HOST" env-default:"localhost"`
	DBHost       string `env:"DB_HOST" env-default:"localhost"`
	DBPort       string `env:"DB_PORT" env-default:"5432"`
	DBUser       string `env:"DB_USER" env-default:"postgres"`
	DBPassword   string `env:"DB_PASSWORD" env-default:"admin123"`
	DBName       string `env:"DB_NAME" env-default:"redditex"`
	RedisHost    string `env:"REDIS_HOST" env-default:"localhost"`
	RedisPort    string `env:"REDIS_PORT" env-default:"6379"`
	CookieSecret string `env:"COOKIE_SECRET" env-default:"abrakadabra"`
	EmailApiKey  string `env:"EMAIL_API_KEY"`
	InboxId      string `env:"INBOX_EMAIL_ID"`
}

var Cfg Properties
