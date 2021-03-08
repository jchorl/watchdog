package email

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	pb "github.com/jchorl/watchdog/proto"
)

type Client struct {
	SendGridAPIKey string
	Domain         string
	FromEmail      string
	ToEmail        string
}

func (e *Client) SendErrorEmail(ctx context.Context, err error) {
	subject := "Watchdog is down"
	body := fmt.Sprintf("Watchdog is down. Error: %s", err)
	if err := e.sendEmail(subject, body); err != nil {
		glog.Errorf("sending email: %s", err)
	}
}

func (e *Client) SendServiceDownEmail(ctx context.Context, watch *pb.Watch) {
	subject := fmt.Sprintf("%s is down", watch.Name)
	lastSeen := time.Unix(watch.LastSeen, 0)
	body := fmt.Sprintf(
		`%s is down and was last seen %s.
The frequency is set to %s.

Visit https://%s/remove?name=%s to disable the alert.`,
		watch.GetName(), lastSeen, watch.GetFrequency().String(), e.Domain, watch.GetName())
	if err := e.sendEmail(subject, body); err != nil {
		glog.Errorf("sending email: %s", err)
	}
}

func (e *Client) sendEmail(subject, body string) error {
	from := mail.NewEmail("Watchdog Notifications", e.FromEmail)
	to := mail.NewEmail("", e.ToEmail)
	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(e.SendGridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	glog.Infof("sent email status_code=%v body=%v", response.StatusCode, response.Body)
	return nil
}
