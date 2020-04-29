package bizMqLogSender

import "github.com/ciaolee87/go-server-util/bizRabbitMq"

var conn *bizRabbitMq.Connection
var queue *bizRabbitMq.Queue
var svId = ""

// sever : mq 주소 (아이디, 비번 포함)
// serverId : 서버 아이디
func InitLogger(server string, queueName string, serverId string) {
	conn = bizRabbitMq.NewConnection(server)
	queue = conn.NewBizQueue(queueName)
	svId = serverId
}
