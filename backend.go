package smtpsrv

import (
	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct {
	handler HandlerFunc
	auther  AuthFunc
}

func NewBackend(auther AuthFunc, handler HandlerFunc) *Backend {
	return &Backend{
		handler: handler,
		auther:  auther,
	}
}

// NewSession creates a new SMTP session from the connection.
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	// Note: Authentication is now handled by the Conn/Session interface
	// We create an anonymous session here. If authentication is required,
	// it should be handled through the session's Auth method if needed.
	return NewSession(c, bkd.handler, bkd.auther), nil
}
