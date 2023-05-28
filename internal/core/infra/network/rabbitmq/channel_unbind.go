package rabbitmq

func (c *Channel) Unbind(exchange, queue, routingKey string) error {
	err := c.QueueUnbind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange
		nil,        // arguments
	)

	if err != nil {
		return err
	}

	return nil
}
