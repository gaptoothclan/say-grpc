package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%d", port))

	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	//pb.
}

// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }
