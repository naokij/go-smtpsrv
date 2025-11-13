package smtpsrv

import (
	"errors"
	"io"
	"net/mail"

	"github.com/emersion/go-smtp"
)

// A Session is returned after successful login.
type Session struct {
	conn     *smtp.Conn
	From     *mail.Address
	To       *mail.Address
	handler  HandlerFunc
	body     io.Reader
	auther   AuthFunc
	username *string
	password *string
}

// NewSession initialize a new session
func NewSession(conn *smtp.Conn, handler HandlerFunc, auther AuthFunc) *Session {
	return &Session{
		conn:    conn,
		handler: handler,
		auther:  auther,
	}
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	var err error
	s.From, err = mail.ParseAddress(from)

	// Extract authentication information from MailOptions if available
	if opts != nil && opts.Auth != nil {
		// The Auth field contains the authorization identity
		// For now, we store it as username (password would need to be handled via AuthSession)
		authIdentity := *opts.Auth
		s.username = &authIdentity
	}

	return err
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	var err error
	s.To, err = mail.ParseAddress(to)
	return err
}

func (s *Session) Data(r io.Reader) error {
	if s.handler == nil {
		return errors.New("internal error: no handler")
	}

	s.body = r

	c := Context{
		session: s,
	}

	return s.handler(&c)
}

func (s *Session) Reset() {
}

func (s *Session) Logout() error {
	return nil
}
