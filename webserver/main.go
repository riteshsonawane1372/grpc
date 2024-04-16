package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/riteshsonawane/grpc/app"
	"google.golang.org/grpc"
)


type server struct {
 app.UnimplementedAppServiceServer 
}


func (s *server) AppCall(ctx context.Context, req *app.AppRequest) (app.AppResponse,error) {
  
  return app.AppResponse{Message: "App Server running..."},nil 
}


func main(){
  
  lis ,err:=net.Listen("tcp",":8081")
  if err!=nil{
    log.Fatal("Check port 8081",err)
  }

  sgrpc:= grpc.NewServer()
  app.RegisterAppServiceServer(sgrpc,&server{})
  fmt.Println("grpc server listening",lis.Addr())
  
  err= sgrpc.Serve(lis)
  if err!=nil{
    log.Fatal("Unable to start grpc server")
  }  
    
}
