package health

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)



func (c Client) VerifyIdCard(ctx context.Context, cid, pid, bp1no string) (*ResponseCheckIdCard, error) {
	ctx, span := c.tracer.Start(ctx, "health.CheckIdCard", oteltrace.WithAttributes(
		attribute.String("cid", cid),
		attribute.String("pid", pid),
		attribute.String("bp1no", bp1no),
	))
	defer span.End()

	url := fmt.Sprintf("%s/api/v1/verify/thai-id-card", c.url)
	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}
	reqBody, _ := sonic.Marshal(fiber.Map{
		"cid":        cid,
		"pid":        pid,
		"bp1no":      bp1no,
		"client_id":  c.clientId,
		"secret_key": c.secretKey,
	})

	response, err := c.http.Post(ctx, url, headers, bytes.NewBuffer(reqBody), 10)
	if err != nil {
		return nil, errors.New("service unavailable")
	}

	var responseService ResponseCheckIdCard
	if err := sonic.Unmarshal(response.Body, &responseService); err != nil {
		return nil, err
	}
	return &responseService, nil
}
