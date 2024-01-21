package terrors_test

import (
	"fmt"
	"testing"

	"github.com/Dionid/go-boiler/pkg/terrors"
)

func TestXxx(t *testing.T) {
	err := terrors.NewPrivateError("some message")

	switch any(err).(type) {
	case terrors.BaseErrorSt:
		fmt.Println("BaseErrorSt")
	case *terrors.BaseErrorSt:
		fmt.Println("Pointer BaseErrorSt")
	}

	switch any(err).(type) {
	case terrors.PrivateError:
		fmt.Println("PrivateError")
	case *terrors.PrivateError:
		fmt.Println("Pointer PrivateError")
	}

	_, ok := any(err).(terrors.Error)
	if ok {
		fmt.Println("OK interface BaseError")
	}

	switch any(err).(type) {
	case terrors.Error:
		fmt.Println("interface BaseError")
	case *terrors.Error:
		fmt.Println("interface Pointer BaseError")
	}

	fmt.Println("Done")
}
