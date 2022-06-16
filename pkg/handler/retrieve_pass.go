package handler

import (
	"context"
	pb "github.com/dtas-pm/mailer/proto"
)

func (s *Server) RetrievePass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {

	//А здесь ответим false

	return &pb.MsgReply{Sent: false}, nil
}
