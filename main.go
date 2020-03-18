package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	// "time"

	// "github.com/Braveheart7854/rabbitmqPool"
	// "github.com/saisaixv/utils/mq"
	// "github.com/saisaixv/utils/redis"
	"github.com/streadway/amqp"
)

type testTask struct{}

func (t *testTask) Run() {
	fmt.Println("hello world")
}

func main() {

	var b byte = 65
	ret := b & (1 << 1)
	fmt.Printf("result = %d\n", ret)
	fmt.Printf("result = %s\n", string(b))

	// crontab:=cron.NewCrontab()
	// task:=&testTask{}
	// if err:=crontab.AddByJob("1","* * * * *",task);err!=nil{
	// 	fmt.Printf("error to add crontab task:%s",err)
	// 	os.Exit(-1)
	// }

	// taskFunc:=func(){
	// 	fmt.Println("hello world")
	// }

	// if err:=crontab.AddByFunc("2","* * * * *",taskFunc);err!=nil{
	// 	fmt.Printf("error to add crontab task:%s",err)
	// 	os.Exit(-1)
	// }
	// crontab.Start()
	// select{}

	// redis.InitPool("192.168.10.168", 6379, "caton", 1, 30, 30, 30, 30)

	// err := redis.DoStrSet("a", "AAAA")
	// if err != nil {
	// 	fmt.Printf("set err: %s\n", err.Error())
	// }
	// redis.DoExpire("a", time.Duration(10)+time.Second)
	// go redis.Subscriber()

	// ret, err := redis.DoStrGet("a")
	// if err != nil {
	// 	fmt.Printf("get err: %s\n", err.Error())
	// }
	// fmt.Println(ret)

	// proxy := &Proxy{
	// 	Host:     "192.168.10.168",
	// 	Port:     8388,
	// 	Password: "caton",
	// 	Method:   "aes-256-cfb",
	// }

	// bArr, err := json.Marshal(proxy)
	// if err != nil {
	// 	fmt.Printf("Login:json.Marshal error = %s\n", err.Error())
	// }

	// ret := crypto.EncryptWithAES("abcdefghabcdefgh", string(bArr))
	// fmt.Println(ret)

	// ret = crypto.DecryptWithAES("abcdefghabcdefgh", ret)
	// fmt.Println(ret)

	// var url = "amqp://guest:guest@localhost:5672/"

	// conn, err := amqp.Dial(url)
	// if err != nil {
	// 	fmt.Printf("Dial: %s\n", err.Error())
	// 	return
	// }

	// defer conn.Close()

	// channel, err := conn.Channel()
	// if err != nil {
	// 	fmt.Printf("Channel: %s\n", err.Error())
	// 	return
	// }

	// if err := channel.ExchangeDeclare(
	// 	"auth",
	// 	"topic",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// ); err != nil {
	// 	fmt.Printf("Exchange Declare: %s\n", err.Error())
	// 	return
	// }

	// if err := channel.Confirm(false); err != nil {
	// 	fmt.Printf("Channel could not be put into confirm mode: %s\n", err.Error())
	// 	return
	// }

	// confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	// var body = "ABCD"

	// log.Printf("declared Exchange,publishing %dB body (%q)", len(body), body)
	// if err := channel.Publish(
	// 	"auth",
	// 	"auth",
	// 	false,
	// 	false,
	// 	amqp.Publishing{
	// 		Headers:         amqp.Table{},
	// 		ContentType:     "text/plain",
	// 		ContentEncoding: "",
	// 		Body:            []byte(body),
	// 		DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
	// 		Priority:        0,              // 0-9
	// 	},
	// ); err != nil {
	// 	fmt.Printf("Exchange Publish: %s\n", err.Error())
	// 	return
	// }

	// confirmOne(confirms)

	// c, err := mq.NewConsumer(
	// 	"amqp://guest:guest@localhost:5672/",
	// 	"auth",
	// 	"topic",
	// 	"authQueue",
	// 	"auth",
	// 	"auth_tag",
	// )
	// if err != nil {
	// 	log.Fatalf("%s", err)
	// }
	// time.Sleep(time.Duration(5) * time.Second)

	// log.Printf("shutting down")

	// if err := c.Shutdown(); err != nil {
	// 	log.Fatalf("error during shutdown: %s", err)
	// }

	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// channel, err := conn.Channel()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := channel.ExchangeDeclare(
	// 	"auth",
	// 	"topic",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// queue, err := channel.QueueDeclare(
	// 	"auth",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := channel.QueueBind(
	// 	queue.Name,
	// 	"auth",
	// 	"auth",
	// 	false,
	// 	nil,
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// deliveries, err := channel.Consume(
	// 	queue.Name,
	// 	"auth_tag",
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// go handle(deliveries)

	// waitSignal()
}
func waitSignal() {
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP)
	for sig := range sigChan {
		if sig == syscall.SIGHUP {
			// updatePasswd()
		} else {
			// is this going to happen?
			log.Printf("caught signal %v, exit", sig)
			os.Exit(0)
		}
	}
}

func handle(deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")

}

// func confirmOne(confirms <-chan amqp.Confirmation) {
// 	if confirmed := <-confirms; confirmed.Ack {
// 		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
// 	} else {
// 		log.Printf("failed delivert of delivery tag: %d", confirmed.DeliveryTag)
// 	}
// }

type Proxy struct {
	Host     string
	Port     int
	Password string
	Method   string
}
