package util

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func ValidateIDCardNum(ctx context.Context, cid string) bool {
	_, span := otel.Tracer("util").Start(ctx, "ValidateIDCardNum",
		oteltrace.WithAttributes(attribute.String("cid", cid)),
	)
	defer span.End()
	cid = strings.ReplaceAll(cid, "-", "")
	cid = strings.ReplaceAll(cid, " ", "")
	match, _ := regexp.MatchString(`^\d{13}$`, cid)
	if !match {
		return false
	}

	// Check the checksum (digit at index 12)
	sum := 0
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(cid[i]))
		sum += digit * (13 - i)
	}
	checkDigit := (11 - (sum % 11)) % 10
	lastDigit, _ := strconv.Atoi(string(cid[12]))
	return checkDigit == lastDigit
}
