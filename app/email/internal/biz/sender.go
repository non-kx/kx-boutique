package biz

import (
	"context"
	"kx-boutique/app/email/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// SendMailRequest.
type SendMailRequest struct {
	To       string
	Template string
	Subject  string
	Payload  interface{}
}

// Mailer.
type Mailer interface {
	SendMail(ctx context.Context, req *SendMailRequest) error
}

// SenderUsecase is a Sender usecase.
type SenderUsecase struct {
	m           Mailer
	templateDir string
	log         *log.Helper
}

// NewSenderUsecase new a Sender usecase.
func NewSenderUsecase(m Mailer, conf *conf.Data, logger log.Logger) *SenderUsecase {
	return &SenderUsecase{m: m, templateDir: conf.Smtp.TemplatesDir, log: log.NewHelper(logger)}
}

// SendMail send business mail with req payload.
func (uc *SenderUsecase) SendMail(ctx context.Context, req *SendMailRequest) error {
	uc.log.WithContext(ctx).Infof("Send email: %v", req)
	return uc.m.SendMail(ctx, &SendMailRequest{})
}
