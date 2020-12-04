package octpass_metadata_go

import (
	"encoding/json"
	"testing"

	"github.com/cheekybits/is"
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

	is := is.New(t)
	var err error

	md, err := NewOctpassMetadata()
	is.NoErr(err)
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
	is.NoErr(err)

	err = json.Unmarshal(([]byte)(desiredJSONString), desired)
	is.NoErr(err)

	dj, err := json.Marshal(desired)
	is.NoErr(err)

	mdj, err := json.Marshal(md)
	is.NoErr(err)
	is.Equal(string(dj), string(mdj))
}
