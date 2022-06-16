package handler

import pb "github.com/dtas-pm/mailer/proto"

type Server struct {
	pb.MailerServer
}

func NewServer() *Server {
	return &Server{}
}
