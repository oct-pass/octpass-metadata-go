package octpass_metadata_go

import (
	"encoding/json"
	"testing"
)

var desiredJSONString = `
{
  "name": "Dave McPufflestein",
  "description": "Generic puff description. This really should be customized.",
  "image": "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png",
  "external_url": "https://cryptopuff.io/3",
  "attributes": {
    "level": 3,
    "weapon_power": 55,
    "personality": "sad",
    "stamina": 11.7
  }
}
`

type attributes struct {
	Level       int64   `json:"level"`
	WeaponPower int64   `json:"weapon_power"`
	Personality string  `json:"personality"`
	Stamina     float64 `json:"stamina"`
}

func TestERC721Metadata(t *testing.T) {

	md, err := NewOctpassMetadata()
	if err != nil {
		t.Errorf(err.Error())
	}
	tmp := *md
	desired := &tmp

	// string values
	md.Name = "Dave McPufflestein"
	md.Description = "Generic puff description. This really should be customized."
	md.Image = "https://storage.googleapis.com/opensea-prod.appspot.com/puffs/3.png"
	md.ExternalURL = "https://cryptopuff.io/3"

	a := attributes{
		Level:       3,
		WeaponPower: 55,
		Personality: "sad",
		Stamina:     11.7,
	}
	err = md.SetAttributes(a)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	if err != nil {
		t.Errorf(err.Error())
	}

	dj, err := json.Marshal(desired)
	if err != nil {
		t.Errorf(err.Error())
	}

	mdj, err := json.Marshal(md)
	if err != nil {
		t.Errorf(err.Error())
	}

	if string(dj) != string(mdj) {
		t.Errorf("got: %v\nwant: %v", mdj, dj)
	}
}
