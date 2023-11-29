package model

type Args struct {
	AppsPort        string `mapstructure:"SERVET_PORT" validate:"required"`
	PostgreHost     string `mapstructure:"DB_HOST" validate:"required"`
	PostgreUser     string `mapstructure:"DB_USER" validate:"required"`
	PostgrePassword string `mapstructure:"DB_PASSWORD" validate:"required"`
	PostgreDb       string `mapstructure:"DB_NAME" validate:"required"`
	PostgrePort     string `mapstructure:"DB_PORT" validate:"required"`
}
