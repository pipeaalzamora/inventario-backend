package ports

type PortMailer interface {
	Send(to []string, subject string, body string) error
}
