package open_api_qq

import (
	"fim_server/utils/stores/logs"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
)

func ReceiveEmail(e EmailConfig) {
	c, err := client.DialTLS("imap.qq.com:993", nil)
	if err != nil {
		logs.Fatal(err)
	}
	defer c.Logout()
	
	if err := c.Login(e.ReceiveUser, e.Code); err != nil {
		logs.Fatal(err)
	}
	
	// 选择收件箱
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		logs.Fatal(err)
	}
	logs.Info(mbox.Name, mbox.Flags)
	
	// 设置搜索条件
	seqset := new(imap.SeqSet)
	seqset.AddRange(1, mbox.Messages)
	
	// 获取邮件
	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, imap.FetchBody}, messages)
	}()
	for msg := range messages {
		logs.Info("* " + msg.Envelope.Subject)
		
		// 获取正文
		section := &imap.BodySectionName{}
		r := msg.GetBody(section)
		if r == nil {
			logs.Fatal("Server didn't return message body")
		}
		
		mr, err := mail.CreateReader(r)
		if err != nil {
			logs.Fatal(err)
		}
		
		// 读取邮件的每个部分
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				logs.Fatal(err)
			}
			
			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				b, _ := io.ReadAll(p.Body)
				logs.Info("Got text", string(b))
			case *mail.AttachmentHeader:
				filename, _ := h.Filename()
				logs.Info("Got attachment", filename)
			}
		}
	}
	
	if err := <-done; err != nil {
		logs.Fatal(err)
	}
}
