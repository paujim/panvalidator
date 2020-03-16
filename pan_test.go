package panvalidator

import "testing"

func TestAmexPANs(t *testing.T) {
	validPans := []int{378282246310005, 371449635398431, 378734493671000}
	for _, pan := range validPans {
		ok, cardType := IsValid(pan)
		if !ok {
			t.Error("Valid Amex PAN validated as invalid [%n]", pan)
		}
		if cardType != Amex {
			t.Error("Valid Amex PAN validated as non Amex")
		}
	}

	invalidPans := []int{378282000310005, 371449000398431, 378734000671000}
	for _, pan := range invalidPans {
		ok, cardType := IsValid(pan)
		if ok {
			t.Error("Invalid Amex PAN validated as valid [%n]", pan)
		}
		if cardType != Amex {
			t.Error("Invalid Amex PAN validated as non Amex")
		}
	}
}

func TestDiscoverPANs(t *testing.T) {
	validPans := []int{6011000990139424, 6011111111111117}
	for _, pan := range validPans {
		ok, cardType := IsValid(pan)
		if !ok {
			t.Error("Valid Discover PAN validated as invalid [%n]", pan)
		}
		if cardType != Discover {
			t.Error("Valid Discover PAN validated as non Discover")
		}
	}

	invalidPans := []int{6011000000009424, 6011110000011117}
	for _, pan := range invalidPans {
		ok, cardType := IsValid(pan)
		if ok {
			t.Error("Invalid Discover PAN validated as valid [%n]", pan)
		}
		if cardType != Discover {
			t.Error("Invalid Discover PAN validated as non Discover")
		}
	}
}

func TestMasterCardPANs(t *testing.T) {
	validPans := []int{5105105105105100, 5111010030175156, 5553042241984105}
	for _, pan := range validPans {
		ok, cardType := IsValid(pan)
		if !ok {
			t.Error("Valid MasterCard PAN validated as invalid", pan)
		}
		if cardType != MasterCard {
			t.Error("Valid MasterCard PAN validated as non MasterCard", pan)
		}
	}

	invalidPans := []int{5105105105105106, 5111010000075156, 5553040000084105}
	for _, pan := range invalidPans {
		ok, cardType := IsValid(pan)
		if ok {
			t.Error("Invalid MasterCard PAN validated as valid", pan)
		}
		if cardType != MasterCard {
			t.Error("Invalid MasterCard PAN validated as non MasterCard", pan)
		}
	}
}

func TestVisaPANs(t *testing.T) {
	validPans := []int{4408041234567893, 4111111111111111, 4012888888881881}
	for _, pan := range validPans {
		ok, cardType := IsValid(pan)
		if !ok {
			t.Error("Valid Visa PAN validated as invalid [%n]", pan)
		}
		if cardType != Visa {
			t.Error("Valid Visa PAN validated as non Visa")
		}
	}

	invalidPans := []int{4111111111111, 4408041230000893, 4012888880000881}
	for _, pan := range invalidPans {
		ok, cardType := IsValid(pan)
		if ok {
			t.Error("Invalid Visa PAN validated as valid [%n]", pan)
		}
		if cardType != Visa {
			t.Error("Invalid Visa PAN validated as non Visa")
		}
	}
}

func TestInvalidPANs(t *testing.T) {
	invalidPans := []int{9111111111111111}
	for _, pan := range invalidPans {
		ok, cardType := IsValid(pan)
		if ok {
			t.Error("Invalid PAN validated as valid [%n]", pan)
		}
		if cardType != Unknown {
			t.Error("Unknown PAN")
		}
	}
}
