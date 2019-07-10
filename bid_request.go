package twofive

// Header required in RTB requests
const (
	Header  = "x-openrtb-version"
	Version = "2.5"

	EncodingHeader = "Content-Encoding"
	EncodingValue  = "gzip"
)

// Request openRTB 2.5 spec
type Request struct {
	ID      string     `json:"id"                valid:"-"`
	Imp     []Imp      `json:"imp"               valid:"required"`
	App     App        `json:"app"               valid:"required"`
	Device  Device     `json:"device"            valid:"required"`
	Format  Format     `json:"format"            valid:"required"` // this is not part of the spec, adding this here for convience allows h and width to be passed without the video/banner object to backwards support the GET
	User    User       `json:"user"              valid:"required"`
	Test    int        `json:"test,omitempty"    valid:"range(0|1),optional"`
	At      int        `json:"at,omitempty"      valid:"range(1|2),optional"` // 1 for first price, 2 for second price auctions defualts to 2
	Tmax    int        `json:"tmax,omitempty"    valid:"-"`
	WSeat   []string   `json:"wseat,omitempty"   valid:"-"`
	BSeat   []string   `json:"bseat,omitempty"   valid:"-"`
	AllImps int        `json:"allimps,omitempty" valid:"-"`
	Cur     []string   `json:"cur,omitempty"     valid:"-"`
	Wlang   []string   `json:"wlang,omitempty"   valid:"-"` // white list of languages in alpha 2, defualts no restriction
	Bcat    []string   `json:"bcat,omitempty"    valid:"-"`
	BAdv    []string   `json:"badv,omitempty"    valid:"-"`
	BApp    []string   `json:"bapp,omitempty"    valid:"-"`
	Source  *Source    `json:"source,omitempty"  valid:"-"`
	Regs    Regs       `json:"regs"              valid:"optional"` // optional because of GDPR
	Ext     RequestExt `json:"ext,omitempty"     valid:"required"`
}

// RequestExt used to communicate the publishers api key
type RequestExt struct {
	APIKey    string `json:"api_key"    valid:"uuidv4,required"`
	SessionID string `json:"session_id" valid:"required"`
}

// Source describes the nature and behavior of the entity that is the source of the bid request
// upstream from the exchange. The primary purpose of this object is to define post-auction or upstream
// decisioning when the exchange itself does not control the final decision. A common example of this is
// header bidding, but it can also apply to upstream server entities such as another RTB exchange, a
// mediation platform, or an ad server combines direct campaigns with 3rd party demand in decisioning.
type Source struct {
	FD     int        `json:"fd,omitempty"     valid:"range(0|1),optional"`
	TID    int        `json:"tid,omitempty"    valid:"-"`
	PChain string     `json:"pchain,omitempty" valid:"-"`
	Ext    *SourceExt `json:"ext,omitempty"    valid:"-"`
}

// SourceExt ...
type SourceExt struct{}

// Regs object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission’s
// regulations for the United States Children’s Online Privacy Protection Act (“COPPA”).
type Regs struct {
	Coppa int      `json:"coppa" valid:"range(0|1),optional"`
	Ext   *RegsExt `json:"ext"   valid:"optional"`
}

// RegsExt being used for GDPR
type RegsExt struct {
	GDPR int `json:"gdpr" valid:"range(0|1),optional"`
}

// Imp describes an ad placement or impression being auctioned. A single bid request can include
// multiple Imp objects, a use case for which might be an exchange that supports selling all ad positions on
// a given page. Each Imp object has a required ID so that bids can reference them individually.
type Imp struct {
	ID                   string   `json:"id,omitempty"                   valid:"required"`
	Metric               []Metric `json:"metric,omitempty"               valid:"optional"`
	Banner               *Banner  `json:"banner,omitempty"               valid:"optional"`
	Video                *Video   `json:"video,omitempty"                valid:"optional"`
	Audio                *Audio   `json:"audio,omitempty"                valid:"optional"`
	Native               *Native  `json:"native,omitempty"               valid:"optional"`
	PMP                  *PMP     `json:"pmp,omitempty"                  valid:"optional"`
	DisplayManager       string   `json:"displaymanager,omitempty"       valid:"-"`
	DisplayManagerServer string   `json:"displaymanagerserver,omitempty" valid:"-"`
	Instl                int      `json:"instl"                          valid:"range(0|1)"` // 0 = not interstitial, 1 = interstitial
	TagID                string   `json:"tagid,omitempty"                valid:"-"`
	BidFloor             float64  `json:"bidfloor"                       valid:"-"`
	BidFloorCur          string   `json:"bidfloorcur,omitempty"          valid:"-"` // defaults to USD on no send
	Secure               int      `json:"secure"                         valid:"range(0|1)"`
	Ext                  *ImpExt  `json:"ext,omitempty"                  valid:"-"`
}

// ImpExt ...
type ImpExt struct {
	APS           []APS  `json:"aps,omitempty"             valid:"-"`
	GoogleID      string `json:"google_id,omitempty"       valid:"-"` // a tempary condition (experiment) to determine if google is participating in the Nimbus auction
	FacebookAppID string `json:"facebook_app_id,omitempty" valid:"-"` // needed for pubs that have FB hybrid SDK solution in thier stack
	Position      string `json:"position,omitempty"        valid:"-"` // flexible optional field for publishers to track on ad position performance
	Viewability   int    `json:"viewability,omitempty"     valid:"-"` // for demand
}

// Metric is associated with an impression as an array of metrics. These metrics can offer insight into
// the impression to assist with decisioning such as average recent viewability, click-through rate, etc. Each
// metric is identified by its type, reports the value of the metric, and optionally identifies the source or
// vendor measuring the value.
// Attribute Type
type Metric struct {
	Type   string     `json:"type,omitempty"   valid:"required"`
	Value  string     `json:"value,omitempty"  valid:"required"`
	Vendor string     `json:"vendor,omitempty" valid:"_"`
	Ext    *MetricExt `json:"ext,omitempty"    valid:"_"`
}

// MetricExt ...
type MetricExt struct{}

// Banner represents the most general type of impression. Although the term “banner” may have very
// specific meaning in other contexts, here it can be many things including a simple static image, an
// expandable ad unit, or even in-banner video (refer to the Video object in Section 3.2.4 for the more
// generalized and full featured video ad units). An array of Banner objects can also appear within the
// Video to describe optional companion ads defined in the VAST specification.
type Banner struct {
	BidFloor *float64   `json:"bidfloor,omitempty" valid:"-"`
	BAttr    []int      `json:"battr,omitempty"    valid:"-"`
	Format   []Format   `json:"format,omitempty"   valid:"optional"`
	W        int        `json:"w,omitempty"        valid:"-"`
	H        int        `json:"h,omitempty"        valid:"-"`
	ID       string     `json:"id,omitempty"       valid:"optional"`
	Pos      int        `json:"pos,omitempty"      valid:"range(0|7),optional"`            // 0,1,2,3,4,5,6,7 -> Unknown,Above the Fold,DEPRECATED - May or may not be initially visible depending on screen size/resolution.,Below the Fold,Header,Footer,Sidebar,Full Screen
	API      []int      `json:"api,omitempty"      valid:"inintarr(1|2|3|4|5|6),optional"` // 3,5,6 -> mraid1, 2, and 3
	Ext      *BannerExt `json:"ext,omitempty"      valid:"-"`
}

// BannerExt ...
type BannerExt struct{}

// Video object represents an in-stream video impression. Many of the fields are non-essential for minimally
// viable transactions, but are included to offer fine control when needed. Video in OpenRTB generally
// assumes compliance with the VAST standard. As such, the notion of companion ads is supported by
// optionally including an array of Banner objects (refer to the Banner object in Section 3.2.3) that define
// these companion ads.
type Video struct {
	BidFloor       *float64  `json:"bidfloor,omitempty"       valid:"-"`
	Mimes          []string  `json:"mimes,omitempty"          valid:"-"`
	Minduration    int       `json:"minduration"              valid:"-"`
	Maxduration    int       `json:"maxduration,omitempty"    valid:"-"`
	Protocols      []int     `json:"protocols,omitempty"      valid:"inintarr(2|3|5|6),optional"` // 1,2,3,4,5,6,7,8,9,10 -> VAST 1.0,VAST 2.0,VAST 3.0,VAST 1.0 Wrapper,VAST 2.0 Wrapper,VAST 3.0 Wrapper,VAST 4.0,VAST 4.0 Wrapper,DAAST 1.0,DAAST 1.0 Wrapper
	W              int       `json:"w,omitempty"              valid:"-"`
	H              int       `json:"h,omitempty"              valid:"-"`
	StartDelay     int       `json:"startdelay"               valid:"-"`
	Placement      int       `json:"placement,omitempty"      valid:"range(1|5),optional"`            // 1,2,3,4,5 -> In-Stream, In-Banner, In-Article, In-Feed - Found in content, social, or product feeds, Interstitial/Slider/Floating
	Linearity      int       `json:"linearity,omitempty"      valid:"range(1|2),optional"`            // 1,2 -> linear, non linear
	Playbackmethod []int     `json:"playbackmethod,omitempty" valid:"inintarr(1|2|3|4|5|6),optional"` // 1,2,3,4,5,6 - > Initiates on Page Load with Sound On, Initiates on Page Load with Sound Off by Default, Initiates on Click with Sound On, Initiates on Mouse-Over with Sound On, Initiates on Entering Viewport with Sound On, Initiates on Entering Viewport with Sound Off by Default
	Skip           int       `json:"skip"                     valid:"range(0|1),optional"`            // 0 no 1 yes
	Delivery       []int     `json:"Delivery,omitempty"       valid:"range(0|3),optional"`            // 0,1,2,3 -> Unknown, Professionally Produced, Prosumer, User Generated (UGC)
	Pos            int       `json:"pos,omitempty"            valid:"range(0|7),optional"`            // 0,1,2,3,4,5,6,7 -> Unknown,Above the Fold,DEPRECATED - May or may not be initially visible depending on screen size/resolution.,Below the Fold,Header,Footer,Sidebar,Full Screen
	API            []int     `json:"api,omitempty"            valid:"inintarr(1|2|3|4|5|6),optional"` // current version of the IMA player doesn't support vpaid 2. So pass in the value 1 or nothing -> refer to list 5.12
	MinBitRate     int       `json:"minbitrate,omitempty"     valid:"-"`
	MaxBitRate     int       `json:"maxbitrate,omitempty"     valid:"-"`
	Boxingallowed  int       `json:"boxingallowed,omitempty"  valid:"range(0|1),optional"`
	Ext            *VideoExt `json:"ext,omitempty"            valid:"-"`
}

// VideoExt ...
type VideoExt struct{}

// Audio object represents an audio type impression. Many of the fields are non-essential for minimally
// viable transactions, but are included to offer fine control when needed. Audio in OpenRTB generally
// assumes compliance with the DAAST standard. As such, the notion of companion ads is supported by
// optionally including an array of Banner objects (refer to the Banner object in Section 3.2.3) that define
// these companion ads.
type Audio struct {
	Mimes         []string  `json:"mimes,omitempty"         valid:"required"`
	Minduration   int       `json:"minduration"             valid:"-"`
	Maxduration   int       `json:"maxduration,omitempty"   valid:"-"`
	Protocols     []int     `json:"protocols,omitempty"     valid:"-"`
	StartDelay    int       `json:"startdelay"              valid:"-"`
	Sequence      int       `json:"sequence,omitempty"      valid:"-"`
	Battr         []int     `json:"battr,omitempty"         valid:"-"`
	MaxExtended   int       `json:"maxextended,omitempty"   valid:"-"`
	MinBitRate    int       `json:"minbitrate,omitempty"    valid:"-"`
	MaxBitRate    int       `json:"maxbitrate,omitempty"    valid:"-"`
	Delivery      []int     `json:"delivery,omitempty"      valid:"-"`
	CompanionAd   []Banner  `json:"companionad,omitempty"   valid:"-"`
	API           []int     `json:"api,omitempty"           valid:"-"`
	CompanionType []int     `json:"companiontype,omitempty" valid:"-"`
	Maxseq        int       `json:"maxseq,omitempty"        valid:"-"`
	Feed          int       `json:"feed,omitempty"          valid:"range(1|3),optional"` // 1,2,3 -> Music Service, FM/AM Broadcast, Podcast
	Stitched      int       `json:"stitched,omitempty"      valid:"range(0|1),optional"`
	Nvol          int       `json:"nvol,omitempty"          valid:"range(0|4),optional"` // 0,1,2,3,4 -> None, Ad Volume Average Normalized to Content, Ad Volume Peak Normalized to Content, Ad Loudness Normalized to Content, Custom Volume Normalization
	Ext           *AudioExt `json:"ext,omitempty"           valid:"-"`
}

// AudioExt ...
type AudioExt struct{}

// Native object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
type Native struct {
	Request string     `json:"request"         valid:"required"`
	Ver     string     `json:"ver,omitempty"   valid:"-"`
	API     []int      `json:"api,omitempty"   valid:"-"`
	Battr   []int      `json:"battr,omitempty" valid:"-"`
	Ext     *NativeExt `json:"ext,omitempty"   valid:"-"`
}

// NativeExt ...
type NativeExt struct{}

// Format object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	W      int        `json:"w,omitempty"      valid:"-"`
	H      int        `json:"h,omitempty"      valid:"-"`
	WRatio int        `json:"wratio,omitempty" valid:"-"`
	HRatio int        `json:"hratio,omitempty" valid:"-"`
	WMin   int        `json:"wmin,omitempty"   valid:"-"`
	Ext    *FormatExt `json:"ext,omitempty"    valid:"-"`
}

// FormatExt ...
type FormatExt struct{}

// App object should be included if the ad supported content is a non-browser application (typically in
// mobile) as opposed to a website. A bid request must not contain both an App and a Site object. At a
// minimum, it is useful to provide an App ID or bundle, but this is not strictly required
type App struct {
	ID            string    `json:"id,omitempty"         valid:"-"`
	Name          string    `json:"name"                 valid:"required"`
	Bundle        string    `json:"bundle"               valid:"required"`
	Domain        string    `json:"domain"               valid:"required"`
	StoreURL      string    `json:"storeurl"             valid:"required"`
	Cat           []string  `json:"cat,omitempty"        valid:"-"`
	SectionCat    []string  `json:"sectioncat,omitempty" valid:"-"`
	PageCat       []string  `json:"pagecat,omitempty"    valid:"-"`
	Ver           string    `json:"ver,omitempty"        valid:"-"`
	PrivacyPolicy int       `json:"privacypolicy"        valid:"-"`                   // no policy 0 policy 1
	Paid          int       `json:"paid"                 valid:"range(0|1),optional"` // free 0 paid 1
	Publisher     Publisher `json:"publisher"            valid:"required"`
	Content       *Content  `json:"content,omitempty"    valid:"-"`
	Keywords      []string  `json:"keywords,omitempty"   valid:"-"`
	Ext           *AppExt   `json:"ext,omitempty"        valid:"-"`
}

// AppExt ...
type AppExt struct{}

// Publisher object describes the publisher of the media in which the ad will be displayed. The publisher is
// typically the seller in an OpenRTB transaction.
type Publisher struct {
	ID     string        `json:"id,omitempty"  valid:"-"`
	Name   string        `json:"name"          valid:"required"`
	Cat    []string      `json:"cat,omitempty" valid:"-"`
	Domain string        `json:"domain"        valid:"required"`
	Ext    *PublisherExt `json:"ext,omitempty" valid:"-"`
}

// PublisherExt ...
type PublisherExt struct{}

// APS is the response Object the APS sdk generates
// this should just pass through the information the clients sent
type APS struct {
	AmznB     interface{} `json:"amzn_b"`
	AmznVid   interface{} `json:"amzn_vid"`
	AmznH     interface{} `json:"amzn_h"`
	Amznp     interface{} `json:"amznp"`
	Amznrdr   interface{} `json:"amznrdr"`
	Amznslots interface{} `json:"amznslots"`
	Dc        interface{} `json:"dc"`
}

// Content object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher’s general content.
type Content struct {
	ID                 string      `json:"id,omitempty"                 valid:"-"`
	Episode            int         `json:"episode,omitempty"            valid:"-"`
	Title              string      `json:"title,omitempty"              valid:"-"`
	Series             string      `json:"series,omitempty"             valid:"-"`
	Season             string      `json:"season,omitempty"             valid:"-"`
	Artist             string      `json:"artist,omitempty"             valid:"-"`
	Genre              string      `json:"genre,omitempty"              valid:"-"`
	Album              string      `json:"album,omitempty"              valid:"-"`
	ISRC               string      `json:"isrc,omitempty"               valid:"-"`
	Producer           *Producer   `json:"producer,omitempty"           valid:"-"`
	URL                string      `json:"url,omitempty"                valid:"-"`
	Cat                []string    `json:"cat,omitempty"                valid:"-"`
	Prodq              int         `json:"prodq,omitempty"              valid:"range(0|3),optional"` // 0,1,2,3 -> Unknown, Professionally Produced, Prosumer, User Generated (UGC)
	Context            int         `json:"context,omitempty"            valid:"range(1|7),optional"` // 1,2,3,4,5,6,7 -> Video (i.e., video file or stream such as Internet TV broadcasts),Game (i.e., an interactive software game),Music (i.e., audio file or stream such as Internet radio broadcasts),Application (i.e., an interactive software application),Text (i.e., primarily textual document such as a web page, eBook, or news article),Other (i.e., none of the other categories applies),Unknown
	ContentRating      string      `json:"contentrating,omitempty"      valid:"-"`
	UserRating         string      `json:"userrating,omitempty"         valid:"-"`
	QagMediaRating     int         `json:"qagmediarating,omitempty"     valid:"range(1|3),optional"` // 1,2,3 -> All Audiences, Everyone Over 12, Mature Audiences
	Keywords           []string    `json:"keywords,omitempty"           valid:"-"`
	LiveStream         int         `json:"livestream,omitempty"         valid:"range(0|1),optional"` // 0 = not live, 1 = live
	SourceRelationShip int         `json:"sourcerelationship,omitempty" valid:"range(0|1),optional"` // 0 = indirect, 1 = direct
	Len                int         `json:"len,omitempty"                valid:"-"`
	Language           string      `json:"language,omitempty"           valid:"-"`
	Embeddable         int         `json:"embeddable,omitempty"         valid:"-"`
	Data               []Data      `json:"data,omitempty"               valid:"-"`
	Ext                *ContentExt `json:"ext,omitempty"                valid:"-"`
}

// ContentExt ...
type ContentExt struct{}

// Producer object defines the producer of the content in which the ad will be shown. This is particularly useful
// when the content is syndicated and may be distributed through different publishers and thus when the
// producer and publisher are not necessarily the same entity.
type Producer struct {
	ID     string       `json:"id,omitempty"  valid:"-"`
	Name   string       `json:"name"          valid:"required"`
	Cat    []string     `json:"cat,omitempty" valid:"-"`
	Domain string       `json:"domain"        valid:"required"`
	Ext    *ProducerExt `json:"ext,omitempty" valid:"-"`
}

// ProducerExt ...
type ProducerExt struct{}

// Device object provides information pertaining to the device through which the user is interacting. Device
// information includes its hardware, platform, location, and carrier data. The device can refer to a mobile
// handset, a desktop computer, set top box, or other digital device.
type Device struct {
	Ua             string     `json:"ua"                        valid:"required"`
	Geo            *Geo       `json:"geo,omitempty"             valid:"optional"`
	Dnt            int        `json:"dnt"                       valid:"range(0|1),optional"` // 0 = tracking is unrestricted, 1 = tracking is restricted
	Lmt            int        `json:"lmt"                       valid:"range(0|1),optional"` // 0 = tracking is unrestricted, 1 = tracking must be limited by commericial guidelines
	IP             string     `json:"ip"                        valid:"ipv4,required"`
	IPv6           string     `json:"ipv6,omitempty"            valid:"ipv6,optional"`
	DeviceType     int        `json:"device_type,omitempty"     valid:"-"`
	Make           string     `json:"make,omitempty"            valid:"in(Apple|Android|apple|android),required"`
	Model          string     `json:"model,omitempty"           valid:"-"`
	OS             string     `json:"os,omitempty"              valid:"-"`
	OSV            string     `json:"osv,omitempty"             valid:"-"`
	HWV            string     `json:"hwv,omitempty"             valid:"-"`
	H              int        `json:"h,omitempty"               valid:"-"`
	W              int        `json:"w,omitempty"               valid:"-"`
	PPI            int        `json:"ppi,omitempty"             valid:"-"`
	PXRatio        float64    `json:"px_ratio,omitempty"        valid:"-"`
	JS             int        `json:"js,omitempty"              valid:"-"`
	GeoFetch       int        `json:"geo_fetch,omitempty"       valid:"range(0|1),optional"` // 0 = no, 1 = yes
	FlashVer       string     `json:"flash_ver,omitempty"       valid:"-"`
	Language       string     `json:"language,omitempty"        valid:"-"`
	Carrier        string     `json:"carrier,omitempty"         valid:"-"`
	ConnectionType int        `json:"connectiontype,omitempty" valid:"-"`
	Ifa            string     `json:"ifa"                       valid:"required"`
	Ext            *DeviceExt `json:"ext,omitempty"             valid:"-"`
}

// DeviceExt ...
type DeviceExt struct{}

// Geo object encapsulates various methods for specifying a geographic location. When subordinate to a
// Device object, it indicates the location of the device which can also be interpreted as the user’s current
// location. When subordinate to a User object, it indicates the location of the user’s home base (i.e., not
// necessarily their current location).
type Geo struct {
	Lat       float64 `json:"lat,omitempty"       valid:"-"`
	Lon       float64 `json:"lon,omitempty"       valid:"-"`
	Type      int     `json:"type,omitempty"      valid:"range(1|3),optional"`    // 1,2,3 -> GPS/Location Services, IP Address, User provided (e.g., registration data)
	IPService int     `json:"ipservice,omitempty" valid:"range(1|4),optional"`    // 1,2,3,4 -> ip2location, Neustar (Quova), MaxMind, NetAcuity (Digital Element)
	Country   string  `json:"country,omitempty"   valid:"ISO3166Alpha3,optional"` // alpha 3
	City      string  `json:"city,omitempty"      valid:"-"`
	Ext       *GeoExt `json:"ext,omitempty"       valid:"-"`
}

// GeoExt ...
type GeoExt struct{}

// User object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID         string   `json:"id,omitempty"          valid:"-"`
	Age        int      `json:"age,omitempty"` // outside of spec, but i'm allowing for it
	BuyerUID   string   `json:"buyeruid,omitempty"    valid:"-"`
	YOB        int      `json:"yob,omitempty"         valid:"-"`
	Gender     string   `json:"gender,omitempty"      valid:"in(M|F|O|Male|male|Female|female),optional"`
	Keywords   string   `json:"keywords,omitempty"    valid:"-"`
	CustomData string   `json:"custom_data,omitempty" valid:"-"`
	Geo        *Geo     `json:"geo,omitempty"         valid:"-"`
	Data       []Data   `json:"data,omitempty"        valid:"-"`
	Ext        *UserExt `json:"ext"                   valid:"optional"` // can't be blank because of GDPR
}

// UserExt being used for GDPR
type UserExt struct {
	Consent    string `json:"consent,omitempty" valid:"base64rawstring,optional"`
	Age        int    `json:"age,omitempty" valid:"_"`                           // age used incase year of birth can't be passed and age can
	DidConsent int    `json:"did_consent,omitempty" valid:"range(0|1),optional"` // extention for apps to indicate a user accepted gdpr
}

// Data and Segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string    `json:"id,omitempty"      valid:"-"`
	Name    string    `json:"name,omitempty"    valid:"-"`
	Segment []Segment `json:"segment,omitempty" valid:"-"`
	Ext     *DataExt  `json:"ext,omitempty"     valid:"-"`
}

// DataExt ...
type DataExt struct{}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string      `json:"id,omitempty"    valid:"-"`
	Name  string      `json:"name,omitempty"  valid:"-"`
	Value string      `json:"value,omitempty" valid:"-"`
	Ext   *SegmentExt `json:"ext,omitempty"   valid:"-"`
}

// SegmentExt ...
type SegmentExt struct{}

// PMP object is the private marketplace container for direct deals between buyers and sellers that may
// pertain to this impression. The actual deals are represented as a collection of Deal objects. Refer to
// Section 7.3 for more deta
type PMP struct {
	PrivateAuction int     `json:"private_auction" valid:"range(0|1),optional"` // 0 = all bids accepted, 1 = bids are restricted to the dleads specified and the terms thereof
	Deals          []Deal  `json:"deals,omitempty" valid:"optional"`
	Ext            *PMPExt `json:"ext,omitempty"   valid:"-"`
}

// PMPExt ...
type PMPExt struct{}

// Deal object constitutes a specific deal that was struck a priori between a buyer and a seller. Its presence
// with the Pmp collection indicates that this impression is available under the terms of that deal. Refer to
// Section 7.3 for more details.
type Deal struct {
	ID          string   `json:"id"                    valid:"required"`
	BidFloor    float64  `json:"bidfloor,omitempty"    valid:"-"`
	BidFloorCur float64  `json:"bidfloorcur,omitempty" valid:"-"`
	At          int      `json:"at,omitempty"          valid:"range(1|3),optional"` // 1 = first price, 2 = second price, 3 = value passed in the bid floor is the agreeded upon deal price
	WSeat       []string `json:"wseat,omitempty"       valid:"-"`
	WAdomain    []string `json:"wadomain,omitempty"    valid:"-"`
	Ext         *DealExt `json:"ext,omitempty"         valid:"-"`
}

// DealExt ...
type DealExt struct{}
