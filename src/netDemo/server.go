package main


import(
	"fmt"
	"io"
	"net"
	"log"
	"time"
	"strconv"
)	

func  main()  {
	listener,_ := net.Listen("tcp","localhost:8088")
	// if (err != nil) {
	// 	log.Fatal(err)
	// }
	for{
		conn,_ := 
		listener.Accept()
		// if (err != nil){
		// 	fmt.Println("a Accept client err")
			 //log.Fatal("Accept")//会退出程序
			 log.Println("Accept")
		// }
		fmt.Println("a client connected")
		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	var i int = 0 
	for{
		i++
		//_,err := 
		io.WriteString(conn,"\n hello" +strconv.Itoa(i))
		
		// if (err != nil){
		// 	fmt.Println("a client err")
		// 	log.Fatal(err)
		// }
		time.Sleep(time.Second)
	}
	
}

