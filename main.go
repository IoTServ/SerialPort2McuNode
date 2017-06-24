package main

import (
	"log"
	"flag"
	"github.com/tarm/serial"
	"time"
	"net"
	"fmt"
	"runtime"
)

var (
	Port	string
	Baud	int

	Host    string
	Id	string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	recv := make([]byte, 10240)
	buf := make([]byte, 10240)
	conn, err := net.DialTimeout("tcp", Host+":8001", time.Millisecond*700)
	if err != nil {
		fmt.Printf("连接服务器失败！")
	}
	conn.Write([]byte(Id))
	c := &serial.Config{Name: Port, Baud: Baud, ReadTimeout: time.Microsecond}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	go func(){
		for {
			n, err := conn.Read(recv)
			if err != nil {
				fmt.Printf("read err")
			}
			s.Write([]byte(string(recv[:n])+"\n"))
			//fmt.Printf("conn:"+string(recv[:n]))
		}


	}()
	go func(){
		for {
			time.Sleep(time.Second*100)
			conn.Write([]byte("<h1></h1>"))
		}
	}()
	for {
		n, err := s.Read(buf)
		if err!=nil{
			log.Fatal(err)
		}else{
			time.Sleep(time.Millisecond*300)
			conn.Write(buf[:n])
			//fmt.Printf(string(buf[:n]))
		}

	}

}


func init() {
	fmt.Printf("Usage: -B int Baud rate,\n -C string COM Port eg:COM3,\n -H string Host eg：eiot.club ,\n -I string ID ,your id setting\n")
	flag.StringVar(&Port, "C", "COM3", "COM Port eg:COM3 (串口号)")
	flag.IntVar(&Baud, "B", 115200, "Baud rate（波特率）")

	flag.StringVar(&Host, "H", "127.0.0.1", "Host eg：eiot.club（服务器域名或ip）")
	flag.StringVar(&Id, "I", "4567", "ID ,your id setting (自定义ID)eg:4567")
	flag.Parse()
}