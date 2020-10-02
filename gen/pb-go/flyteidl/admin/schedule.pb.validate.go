// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/schedule.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _schedule_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on FixedRate with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FixedRate) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	// no validation rules for Unit

	return nil
}

// FixedRateValidationError is the validation error returned by
// FixedRate.Validate if the designated constraints aren't met.
type FixedRateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FixedRateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FixedRateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FixedRateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FixedRateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FixedRateValidationError) ErrorName() string { return "FixedRateValidationError" }

// Error satisfies the builtin error interface
func (e FixedRateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFixedRate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FixedRateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FixedRateValidationError{}

// Validate checks the field values on CronScheduleWithOffset with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CronScheduleWithOffset) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Schedule

	// no validation rules for Offset

	return nil
}

// CronScheduleWithOffsetValidationError is the validation error returned by
// CronScheduleWithOffset.Validate if the designated constraints aren't met.
type CronScheduleWithOffsetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CronScheduleWithOffsetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CronScheduleWithOffsetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CronScheduleWithOffsetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CronScheduleWithOffsetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CronScheduleWithOffsetValidationError) ErrorName() string {
	return "CronScheduleWithOffsetValidationError"
}

// Error satisfies the builtin error interface
func (e CronScheduleWithOffsetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCronScheduleWithOffset.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CronScheduleWithOffsetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CronScheduleWithOffsetValidationError{}

// Validate checks the field values on Schedule with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Schedule) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for KickoffTimeInputArg

	switch m.ScheduleExpression.(type) {

	case *Schedule_CronExpression:
		// no validation rules for CronExpression

	case *Schedule_Rate:

		if v, ok := interface{}(m.GetRate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ScheduleValidationError{
					field:  "Rate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Schedule_CronScheduleWithOffset:

		if v, ok := interface{}(m.GetCronScheduleWithOffset()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ScheduleValidationError{
					field:  "CronScheduleWithOffset",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ScheduleValidationError is the validation error returned by
// Schedule.Validate if the designated constraints aren't met.
type ScheduleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScheduleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScheduleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScheduleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScheduleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScheduleValidationError) ErrorName() string { return "ScheduleValidationError" }

// Error satisfies the builtin error interface
func (e ScheduleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchedule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScheduleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScheduleValidationError{}
