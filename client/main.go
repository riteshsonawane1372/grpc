package main

import (
	"context"
	"fmt"
	"log"
	

	"github.com/riteshsonawane/grpc/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
  
)



func main(){

  
  client,err:=grpc.Dial("127.0.0.1:8081",grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err!=nil{
    log.Fatal("Unable to make a connection with the server")
  }

  defer client.Close()
  
  appclient:= app.NewAppServiceClient(client)
  var req app.AppRequest
  r,_:=appclient.AppInit(context.Background(),&req)
  
  fmt.Println("Printing Response from gRPC server",r.GetMessage())
}
