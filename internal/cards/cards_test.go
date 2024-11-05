package cards

import (
	"testing"
)

func TestCardNumToSuit(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected Suit
	}{
		{ "Ace of Spades",   0, Spades },
		{ "Two of Spades",   1, Spades },
		{ "King of Spades", 12, Spades },
		{ "Ace of Clubs",   39, Clubs  },
		{ "Joker valid",    53, Joker  },
		{ "Joker invalid", -69, Joker  }, // I could just throw an error, but I can't be bothered to right now
		                                  //   especially since this function isn't exposed externally
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cardNumToSuit(tc.num)
			if result != tc.expected {
				t.Errorf("cardNumToSuit(%d) = %v; expected %v", tc.num, result, tc.expected) 
			}
		})
	}
}
