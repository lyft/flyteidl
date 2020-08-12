// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/project.proto

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
var _project_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Domain with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Domain) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// DomainValidationError is the validation error returned by Domain.Validate if
// the designated constraints aren't met.
type DomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DomainValidationError) ErrorName() string { return "DomainValidationError" }

// Error satisfies the builtin error interface
func (e DomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DomainValidationError{}

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Project) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	for idx, item := range m.GetDomains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectValidationError{
					field:  fmt.Sprintf("Domains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Description

	if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "Labels",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on Projects with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Projects) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProjects() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectsValidationError{
					field:  fmt.Sprintf("Projects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProjectsValidationError is the validation error returned by
// Projects.Validate if the designated constraints aren't met.
type ProjectsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectsValidationError) ErrorName() string { return "ProjectsValidationError" }

// Error satisfies the builtin error interface
func (e ProjectsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjects.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectsValidationError{}

// Validate checks the field values on ProjectListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectListRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProjectListRequestValidationError is the validation error returned by
// ProjectListRequest.Validate if the designated constraints aren't met.
type ProjectListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectListRequestValidationError) ErrorName() string {
	return "ProjectListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectListRequestValidationError{}

// Validate checks the field values on ProjectRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectRegisterRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectRegisterRequestValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectRegisterRequestValidationError is the validation error returned by
// ProjectRegisterRequest.Validate if the designated constraints aren't met.
type ProjectRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectRegisterRequestValidationError) ErrorName() string {
	return "ProjectRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectRegisterRequestValidationError{}

// Validate checks the field values on ProjectRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectRegisterResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProjectRegisterResponseValidationError is the validation error returned by
// ProjectRegisterResponse.Validate if the designated constraints aren't met.
type ProjectRegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectRegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectRegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectRegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectRegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectRegisterResponseValidationError) ErrorName() string {
	return "ProjectRegisterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectRegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectRegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectRegisterResponseValidationError{}

// Validate checks the field values on ProjectUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectUpdateRequestValidationError{
				field:  "Project",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectUpdateRequestValidationError is the validation error returned by
// ProjectUpdateRequest.Validate if the designated constraints aren't met.
type ProjectUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectUpdateRequestValidationError) ErrorName() string {
	return "ProjectUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectUpdateRequestValidationError{}

// Validate checks the field values on ProjectUpdateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectUpdateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProjectUpdateResponseValidationError is the validation error returned by
// ProjectUpdateResponse.Validate if the designated constraints aren't met.
type ProjectUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectUpdateResponseValidationError) ErrorName() string {
	return "ProjectUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectUpdateResponseValidationError{}
