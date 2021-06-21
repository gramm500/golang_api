package models

type Config struct {
	MongoHost string `split_words:"true"`
	MongoDB   string `split_words:"true"`
}
