/* session serviceのclientの実装例です。 */

package main

import (
	"context"
	"fmt"
	"log"
	"session/session"
	"time"

	"google.golang.org/grpc"
)

var serverAddr string = "localhost:9090"
var opts []grpc.DialOption

func main() {

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal("err: %w", err)
	}

	client := session.NewSessionClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	status, err := client.SetSession(ctx, &session.SessionRequest{StatusID: "2.3", UserID: "fdaoisjgg"})
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return
	}
	fmt.Printf("request success!\n status: %+v\n", status)
}
