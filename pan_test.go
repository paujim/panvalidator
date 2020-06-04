package panvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPANs(t *testing.T) {
	assert := assert.New(t)
	var tests = []struct {
		input        int
		expected     bool
		expectedCard CardType
	}{
		{378282246310005, true, Amex},
		{371449635398431, true, Amex},
		{378734493671000, true, Amex},
		{378282000310005, false, Amex},
		{371449000398431, false, Amex},
		{378734000671000, false, Amex},
		{6011000990139424, true, Discover},
		{6011111111111117, true, Discover},
		{6011000000009424, false, Discover},
		{6011110000011117, false, Discover},
		{5105105105105100, true, MasterCard},
		{5111010030175156, true, MasterCard},
		{5553042241984105, true, MasterCard},
		{5105105105105106, false, MasterCard},
		{5111010000075156, false, MasterCard},
		{5553040000084105, false, MasterCard},
		{4408041234567893, true, Visa},
		{4111111111111111, true, Visa},
		{4012888888881881, true, Visa},
		{4111111111111, false, Visa},
		{4408041230000893, false, Visa},
		{4012888880000881, false, Visa},
		{9111111111111111, false, Unknown},
	}
	for _, test := range tests {
		ok, card := IsValid(test.input)
		assert.Equal(ok, test.expected)
		assert.Equal(card, test.expectedCard)
	}
}
