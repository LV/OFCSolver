package cards

type ThreeCardHand = [3]Card
type FiveCardHand  = [5]Card

type HandStrength int
const (
	HighCard HandStrength = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type Hand struct {
	HandStrength HandStrength
	TopRank Rank
	KickerRank Rank
}

func handThreeCards(cards ThreeCardHand) Hand {
	// stub
	return Hand{ HighCard, Two, Two }
}

func handFiveCards(cards FiveCardHand) Hand {
	// stub
	return Hand{ HighCard, Two, Two }
}

func isFormerHandWeakerOrEqualToLatterHand(former interface{}, latter interface{}) bool {
	var formerHand Hand
	var latterHand Hand

	switch former := former.(type) {
	case ThreeCardHand:
		formerHand = handThreeCards(former)
	case FiveCardHand:
		formerHand = handFiveCards(former)
	}

	switch latter := latter.(type) {
	case ThreeCardHand:
		latterHand = handThreeCards(latter)
	case FiveCardHand:
		latterHand = handFiveCards(latter)
	}

	if formerHand.HandStrength > latterHand.HandStrength {
		return false
	}

	if formerHand.HandStrength < latterHand.HandStrength {
		return true
	}

	if formerHand.TopRank > latterHand.TopRank {
		return false
	}
	
	if formerHand.TopRank < latterHand.TopRank {
		return true
	}

	if formerHand.KickerRank > latterHand.KickerRank {
		return false
	}

	return true // the other cases are that formerHand.KickerRank <= latterHand.KickerRank both of which should be true
}
