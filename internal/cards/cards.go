package cards

type Card = int

const CardsPerSuit int = 13
const TotalCards int = 54

type Suit int
const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
	Joker
)

type Rank int
const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var CardSymbols = []rune{
	'🂡', '🂢', '🂣', '🂤', '🂥', '🂦', '🂧', '🂨', '🂩', '🂪', '🂫', '🂭', '🂮', // Spades
	'🂱', '🂲', '🂳', '🂴', '🂵', '🂶', '🂷', '🂸', '🂹', '🂺', '🂻', '🂽', '🂾', // Hearts
	'🃁', '🃂', '🃃', '🃄', '🃅', '🃆', '🃇', '🃈', '🃉', '🃊', '🃋', '🃍', '🃎', // Diamonds
	'🃑', '🃒', '🃓', '🃔', '🃕', '🃖', '🃗', '🃘', '🃙', '🃚', '🃛', '🃝', '🃞', // Clubs
	'🃏', '🃏',                                                      // Joker
}

func cardNumToSuit(n int) Suit {
	if n >= 0 && n <= (CardsPerSuit*1)-1 {
		return Spades
	} else if n >= CardsPerSuit*1 && n <= (CardsPerSuit*2)-1 {
		return Hearts
	} else if n >= CardsPerSuit*2 && n <= (CardsPerSuit*3)-1 {
		return Diamonds
	} else if n >= CardsPerSuit*3 && n <= (CardsPerSuit*4)-1 {
		return Clubs
	} else {
		return Joker
	}
}


