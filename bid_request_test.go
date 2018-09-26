package twofive

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/marcsantiago/govalidator"
)

func TestBidRequest(t *testing.T) {

	staticBidRequest, err := ioutil.ReadFile("./test_data/static_bid_request.json")
	if err != nil {
		t.Fatal(err)
	}

	videoBidRequest, err := ioutil.ReadFile("./test_data/video_bid_request.json")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		bidRequest []byte
	}{
		{
			name:       "Static Bid Response",
			bidRequest: staticBidRequest,
		},
		{
			name:       "Video Bid Response",
			bidRequest: videoBidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Request
			err = json.Unmarshal(tt.bidRequest, &r)
			if err != nil {
				t.Error(err)
			}
			// everything should validate
			_, err = govalidator.ValidateStruct(r)
			if err != nil {
				t.Error(err)
			}

			r.Ext.APIKey = ""
			_, err = govalidator.ValidateStruct(r)
			if err == nil {
				t.Errorf("should have failed, r.ext.api_key is missing")
			}

			r.App.Name = ""
			_, err = govalidator.ValidateStruct(r)
			if err == nil {
				t.Errorf("should have failed, app.name is missing")
			}

			r.App.Bundle = ""
			_, err = govalidator.ValidateStruct(r)
			if err == nil {
				t.Errorf("should have failed, app.bundle is missing")
			}

			r.App.Domain = ""
			_, err = govalidator.ValidateStruct(r)
			if err == nil {
				t.Errorf("should have failed, app.domain is missing")
			}

			r.App.StoreURL = ""
			_, err = govalidator.ValidateStruct(r)
			if err == nil {
				t.Errorf("should have failed, app.storeurl is missing")
			}

			r.Regs.Coppa = 3
			if err == nil {
				t.Errorf("should have failed, r.regs.coppa value can only be 0 or 1")
			}

			r.Regs.Ext.GDPR = 3
			if err == nil {
				t.Errorf("should have failed, r.regs.ext.gdpr value can only be 0 or 1")
			}

			r.App.Publisher.Name = ""
			if err == nil {
				t.Errorf("should have failed, r.app.publisher.name is missing")
			}

			r.App.Publisher.Domain = ""
			if err == nil {
				t.Errorf("should have failed, r.app.publisher.domain is missing")
			}

			r.Imp[0].ID = ""
			if err == nil {
				t.Errorf("should have failed, r.imp[0].id is missing")
			}

			r.User.Gender = "foo"
			if err == nil {
				t.Errorf("should have failed, r.user.gender foo is not a valid type off gender input")
			}

			if r.Imp[0].Banner != nil {
				r.Imp[0].Banner.W = 0
				if err == nil {
					t.Errorf("should have failed, r.imp[0].w cannot be 0 or missing")
				}

				r.Imp[0].Banner.H = 0
				if err == nil {
					t.Errorf("should have failed, r.imp[0].h cannot be 0 or missing")
				}
			}

			if r.Imp[0].Video != nil {
				r.Imp[0].Video.W = 0
				if err == nil {
					t.Errorf("should have failed, r.imp[0].w cannot be 0 or missing")
				}

				r.Imp[0].Video.H = 0
				if err == nil {
					t.Errorf("should have failed, r.imp[0].h cannot be 0 or missing")
				}
			}

			r.Device.Ua = ""
			if err == nil {
				t.Errorf("should have failed, r.device.ua is missing")
			}

			r.Device.IP = ""
			if err == nil {
				t.Errorf("should have failed, r.device.ip is missing")
			}

			r.Device.Make = ""
			if err == nil {
				t.Errorf("should have failed, r.device.make is missing")
			}

			r.Device.Ifa = ""
			if err == nil {
				t.Errorf("should have failed, r.device.ifa is missing")
			}

			if r.Device.Geo != nil {
				r.Device.Geo.Country = "US"
				if err == nil {
					t.Errorf("should have failed, r.device.geo.country is not alpha 3")
				}

				r.Device.Geo.Country = ""
				if err == nil {
					t.Errorf("should have failed, r.device.geo.country is missing")
				}
			}

		})
	}
}
