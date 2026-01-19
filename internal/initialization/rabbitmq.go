package initialization

import (
	"fmt"

	"github.com/InstaySystem/is_v1-be/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(cfg *config.Config) (*amqp091.Connection, error) {
	protocol := "amqp"
	if cfg.RabbitMQ.UseSSL {
		protocol += "s"
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		protocol,
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
		cfg.RabbitMQ.Vhost,
	)

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("message queue - %w", err)
	}

	return conn, nil
}

