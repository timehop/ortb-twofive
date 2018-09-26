package twofive

// BidResponse object is the top-level bid response object (i.e., the unnamed outer JSON object). The id attribute
// is a reflection of the bid request ID for logging purposes. Similarly, bidid is an optional response
// tracking ID for bidders. If specified, it can be included in the subsequent win notice call if the bidder
// wins. At least one seatbid object is required, which contains at least one bid for an impression. Other
// attributes are optional
type BidResponse struct {
	ID         string         `json:"id"`
	SeatBid    []Seatbid      `json:"seatbid"`
	BidID      string         `json:"bidid"`
	Cur        string         `json:"cur"`
	CustomData string         `json:"customdata"`
	NBR        int            `json:"nbr"`
	Ext        BidResponseExt `json:"ext"`
}

// BidResponseExt ...
type BidResponseExt struct{}

// Seatbid -> A bid response can contain multiple SeatBid objects, each on behalf of a different bidder seat and each
// containing one or more individual bids. If multiple impressions are presented in the request, the group
// attribute can be used to specify if a seat is willing to accept any impressions that it can win (default) or if
// it is only interested in winning any if it can win them all as a group
type Seatbid struct {
	Bid   []Bid      `json:"bid"`
	Seat  string     `json:"seat"`
	Group int        `json:"group"`
	Ext   SeatbidExt `json:"ext"`
}

// SeatbidExt ...
type SeatbidExt struct{}

// Bid -> A SeatBid object contains one or more Bid objects, each of which relates to a specific impression in the
// bid request via the impid attribute and constitutes an offer to buy that impression for a given price.
type Bid struct {
	ID             string   `json:"id"`
	ImpID          string   `json:"impid"`
	Price          float64  `json:"price"`
	NURL           string   `json:"nurl"`
	BURL           string   `json:"burl"`
	LURL           string   `json:"lurl"`
	Adm            string   `json:"adm"`
	Adid           string   `json:"adid"`
	Adomain        []string `json:"adomain"`
	Bundle         string   `json:"bundle"`
	IURL           string   `json:"iurl"`
	Cid            string   `json:"cid"`
	Crid           string   `json:"crid"`
	Tactic         string   `json:"tactic"`
	Cat            []string `json:"cat"`
	Attr           []int    `json:"attr"`
	API            int      `json:"api"`
	Protocol       int      `json:"protocol"`
	QagMediaRating int      `json:"qagmediarating"`
	Language       string   `json:"language"`
	DealID         string   `json:"dealid"`
	W              int      `json:"w"`
	H              int      `json:"h"`
	WRatio         int      `json:"wratio"`
	HRatio         int      `json:"hratio"`
	Exp            int      `json:"exp"`
	Ext            BidExt   `json:"ext"`
}

// BidExt ...
type BidExt struct{}
