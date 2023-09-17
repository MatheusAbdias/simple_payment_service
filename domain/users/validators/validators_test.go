package pkg

import "testing"

func TestValidateCPF(t *testing.T) {

	testCases := []struct {
		Name  string
		CPF   string
		Valid bool
	}{
		{"Valid CPF", "09471172020", true},
		{"Invalid CPF", "11111111111", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			valid := validateCPF(testCase.CPF)
			if valid != testCase.Valid {
				t.Errorf("Expected %v, got %v", testCase.Valid, valid)
			}
		})
	}
}

func TestValidateCNPJ(t *testing.T) {
	testCases := []struct {
		Name  string
		CNPJ  string
		Valid bool
	}{
		{"Valid CNPJ", "33782010000185", true},
		{"Invalid CNPJ", "11111111111111", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			valid := validateCNPJ(testCase.CNPJ)
			if valid != testCase.Valid {
				t.Errorf("Expected %v, got %v", testCase.Valid, valid)
			}
		})
	}
}
