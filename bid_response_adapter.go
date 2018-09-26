package twofive

import "sort"

// Creative is a convience wrapper around the bid response,
// which grabs the bid and ad markup from potential multiple bids
type Creative struct {
	Bid    float64
	Markup []byte
}

type creatives []Creative

func (c creatives) Len() int           { return len(c) }
func (c creatives) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c creatives) Less(i, j int) bool { return c[i].Bid > c[j].Bid }

// GetVastCreative parses the highest paying creative from a potential list
func (b BidResponse) GetVastCreative() *Creative {
	var c []Creative
	for _, s := range b.SeatBid {
		for _, b := range s.Bid {
			c = append(c, Creative{Bid: b.Price, Markup: []byte(b.Adm)})
		}
	}

	if len(c) == 0 {
		return nil
	}

	sort.Sort(creatives(c))
	return &c[0]
}
