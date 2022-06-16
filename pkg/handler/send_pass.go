package handler

import (
	"context"
	"fmt"
	pb "github.com/dtas-pm/mailer/proto"
	"gopkg.in/gomail.v2"
	"io"
	"strings"
)

func (s *Server) SendPass(ctx context.Context, in *pb.MsgRequest) (*pb.MsgReply, error) {

	//from := os.Getenv("FromEmailAddr")
	//password := os.Getenv("SMTPpwd")
	//
	//toEmail := os.Getenv("ToEmailAddr")
	//to := []string{toEmail}
	//
	//host := "smtp.mail.ru"
	//port := "587"
	//address := host + ":" + port
	//// message
	//subject := "Hello"
	//body := "first message"
	//message := []byte(subject + body)
	emails := in.ToEmail
	files := in.Files
	person := in.Person
	//fmt.Println(emails, files, person)
	for _, value := range emails {
		m := gomail.NewMessage()
		m.SetHeader("From", "samarec1812@mail.ru")
		m.SetHeader("To", value)
		m.SetHeader("Subject", "Рассылка заданий")
		body := strings.Builder{}
		//for _, file := range files {
		//	body.WriteString(fmt.Sprintf("Hello <b>Bob</b> and <i>Cora</i>! %s \n", file))
		//}
		body.WriteString(fmt.Sprintf("С уважением, %s %s. для связи: %s", person.Role, person.Name, person.FromEmail))
		m.SetBody("text/html", body.String())
		for i := range files {
			m.Attach(files[i].Name, gomail.SetCopyFunc(func(i int) func(w io.Writer) error {
				return func(w io.Writer) error {
					_, err := w.Write(files[i].Data)
					return err
				}
			}(i)))
		}
		//	m.Attach("files/Kondratev_PDP.docx")
		//	m.Attach("files/Batyrev_PDP.docx")

		d := gomail.NewDialer("smtp.mail.ru", 587, "samarec1812@mail.ru", "q7FgSnfYZiRyG3aupNe5")
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("err: ", err)
			return &pb.MsgReply{Sent: false}, err
		}
		fmt.Println("go check your email")
	}

	//auth := smtp.PlainAuth("", from, password, host)

	//err := smtp.SendMail(address, auth, from, to, message)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return &pb.MsgReply{Sent: false}, err
	//}

	////В переменную m считываем MsgRequest(смотрим в mail.proto, чтобы вспомнить, что это).
	//
	//m := Message{From: fmt.Sprintf("%s <%s>", cnf.servicename, cnf.from), To: in.To, Code: in.Code, tplname: passtpl}
	//
	////А вот дальше нам надо записать полученное сообщение в очередь.
	////Сделать нам это нужно в неблокирующем стиле, для этого используем select.
	//
	//select {
	//case queue <- m: //Пишем в канал, если он заблокирован, выполняем default-ветку
	//
	//default:
	//	//Отвечаем на rpc-запрос false-ем
	//	//Таким образом клиент узнает, что сообщение по каким-то причинам не принято и сможет обработать ситуацию
	//	return &pb.MsgReply{Sent: false}, nil
	//}

	//Ну, а если все хорошо,  отвечаем клиенту true
	return &pb.MsgReply{Sent: true}, nil
}
