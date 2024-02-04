package user

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"path"
	"path/filepath"
	"text/template"
	"time"

	"github.com/gclenz/tinybookingapi/internal/app/infra/email"
)

func NewRequestAuthenticationCode(emailService email.IEmailService, repository Repository) *RequestAuthenticationCode {
	return &RequestAuthenticationCode{
		emailService: emailService,
		repository:   repository,
	}
}

type RequestAuthenticationCode struct {
	emailService email.IEmailService
	repository   Repository
}

func (rac *RequestAuthenticationCode) Execute(email string, ctx context.Context) error {
	user, err := rac.repository.FindByEmail(email, ctx)
	if err != nil {
		return ErrUserNotFound
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%d", rnd.Uint32())[:6]
	codeExpiration := time.Now().Add(time.Minute * 30).UTC()

	err = rac.repository.UpdateCode(user.Email, code, codeExpiration, ctx)
	if err != nil {
		return err
	}

	msg := generateEmailContent(code)

	// TODO: add idempotency and send to background queue
	go func() {
		err = rac.emailService.Send(user.Email, msg, ctx)
		if err != nil {
			slog.Error(err.Error())
		}
	}()

	return nil
}

func generateEmailContent(code string) bytes.Buffer {
	p, _ := filepath.Abs(path.Join("internal", "app", "infra", "email", "templates", "authcode.html"))
	t, _ := template.ParseFiles(p)

	var content bytes.Buffer

	mimeHeaders := `MIME-version: 1.0
Content-Type: text/html; charset=UTF-8`
	content.Write([]byte(fmt.Sprintf("Subject: Your login code \n%s\n\n", mimeHeaders)))

	t.Execute(&content, struct {
		Code string
	}{
		Code: code,
	})

	return content
}
