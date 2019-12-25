package rssreader

import (
	"errors"
	"testing"
)

func TestErrorList(t *testing.T) {
	// case 0: test for no errors:
	errs := &errorList{}

	if errs.Error() != "" {
		t.Fatal("failed to ensure error conversion")
	}
	if errs.ErrorOrNil() != nil {
		t.Fatal("failed to ensure error list returns nil in case of no errors")
	}

	// case 1: test for present single error:
	errs = &errorList{
		errors.New("fail1"),
	}

	if errs.Error() != "fail1" {
		t.Fatal("failed to ensure error conversion")
	}
	if errs.ErrorOrNil() == nil {
		t.Fatal("failed to ensure error list returns correct object in case of any errors")
	}

	// case 2: test for present multiple errors:
	errs = &errorList{
		errors.New("fail1"),
		errors.New("fail2"),
		errors.New("fail3"),
	}

	if errs.Error() != "fail1\nfail2\nfail3" {
		t.Fatal("failed to ensure error conversion")
	}
	if errs.ErrorOrNil() == nil {
		t.Fatal("failed to ensure error list returns correct object in case of any errors")
	}
}
