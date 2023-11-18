package password_test

import (
	. "hugobde.dev/amaretti/password"

	"testing"
)

func TestContainsLowerCaseChar(t *testing.T) {
	inputStr := ""
	if ContainsLowerCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "A"
	if ContainsLowerCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if ContainsLowerCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "a"
	if ContainsLowerCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "z"
	if ContainsLowerCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "abcdefghijklmnopqrstuvwxyz"
	if ContainsLowerCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "!@#$%^&*()_-+={[:;\"'.>,</?~`]}"
	if ContainsLowerCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "!@#$%^&*()_a-+={[:;\"'.>,</?~`]}"
	if ContainsLowerCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "AAAAbAAAA"
	if ContainsLowerCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}
}

func TestContainsUpperCaseChar(t *testing.T) {

	inputStr := ""
	if ContainsUpperCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "A"
	if ContainsUpperCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if ContainsUpperCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "a"
	if ContainsUpperCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "z"
	if ContainsUpperCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "abcdefghijklmnopqrstuvwxyz"
	if ContainsUpperCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "!@#$%^&*()_-+={[:;\"'.>,</?~`]}"
	if ContainsUpperCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "!@#$%^&*()_a-+={[:;\"'.>,</?~`]}"
	if ContainsUpperCaseChar(inputStr) != false {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "!@#$%^&*()_A-+={[:;\"'.>,</?~`]}"
	if ContainsUpperCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}

	inputStr = "aaaaBaaaa"
	if ContainsUpperCaseChar(inputStr) != true {
		t.Errorf("Failed on \"%s\"", inputStr)
	}
}

func TestContainsDigit(t *testing.T) {}

func TestVerifyPassword(t *testing.T) {
	passwordStr := "Bloubliboulga123*"

	password, _ := CreatePassword(passwordStr)

	if !VerifyPassword(passwordStr, password) {
		t.Fail()
	}
}
