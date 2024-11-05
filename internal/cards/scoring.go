package cards

type FontHandBonus int
const (
	FHNone        = 0
	FHSixPair     = 1
	FHSevenPair   = 2
	FHEightPair   = 3
	FHNinePair    = 4
	FHTenPair     = 5
	FHJackPair    = 6
	FHQueenPair   = 7
	FHKingPair    = 8
	FHAcePair     = 9
	FHTwoTrips    = 10
	FHThreeTrips  = 11
	FHFourTrips   = 12
	FHFiveTrips   = 13
	FHSixTrips    = 14
	FHSevenTrips  = 15
	FHEightTrips  = 16
	FHNineTrips   = 17
	FHTenTrips    = 18
	FHJackTrips   = 19
	FHQueenTrips  = 20
	FHKingTrips   = 21
	FHAceTrips    = 22
)

type MiddleHandBonus int
const (
	MHNone          = 0
	MHThreeOfAKind  = 2
	MHStraight      = 4
	MHFlush         = 8
	MHFullHouse     = 12
	MHFourOfAKind   = 20
	MHStriaghtFlush = 30
	MHRoyalFlush    = 50
)
type BackHandBonus int
const (
	BHNone          = 0
	BHStraight      = 2
	BHFlush         = 4
	BHFullHouse     = 6
	BHFourOfAKind   = 10
	BHStriaghtFlush = 15
	BHRoyalFlush    = 25
)
