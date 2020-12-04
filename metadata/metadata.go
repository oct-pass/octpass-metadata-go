package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OctpassMetadata struct {
	// Basic category
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	NFTID       string `json:"nft_id"`
	NFTClass    string `json:"nft_class"`
	NFTType     string `json:"nft_type,omitempty"`
	Symbol      string `json:"symbol"`
	SymbolImage string `json:"symbol_image,omitempty"`

	ExternalURL  string        `json:"external_url,omitempty"`
	Localization *Localization `json:"localization,omitempty"`
	Octpass      *Octpass      `json:"octpass"`
	Converted    bool          `json:"conerted,omitempty"`

	// Contents category
	Contents []*ContentsElem `json:"contents"`
	License  *License        `json:"license"`

	// Property
	Attributes *json.RawMessage `json:"attributes"` // EIPs#1071
	Extras     *json.RawMessage `json:"extras,omitempty"`
}

type Localization struct {
	URI     string   `json:"uri"`     // ex) https://www.mycryptoheroes.net/metadata/hero/50010001_{locale}.json
	Default string   `json:"default"` // ex) en
	Locales []string `json:"locales"` // ex) ["en", "ja"]
}

type Octpass struct {
	Version string `json:"version"` // 1.0
	API     string `json:"api,omitempty"`
}

type ContentsElem struct {
	URI    string `json:"uri"`
	Format string `json:"format"`
}

type License struct {
	Copyright string   `json:"copyright,omitempty"`
	URI       string   `json:"uri,omitempty"`
	Contact   string   `json:"contact,omitempty"`
	Type      string   `json:"type"`
	Usecase   *Usecase `json:"usecase"`
}

type Usecase struct {
	Reference   string        `json:"reference"` // allow/disallow
	Trade       string        `json:"trade"`     // allow/disallow
	Lock        string        `json:"lock"`      // allow/disallow
	TradeShares []*TradeShare `json:"trade_shares,omitempty"`
}

type TradeShare struct {
	Percentage string `json:"percentage"`
	Account    string `json:"account"`
}

// FetchOctpassMetadata returns *OctpassMetadata fetch Metadata from tokenURI
func FetchOctpassMetadata(tokenURI string) (*OctpassMetadata, error) {
	return FetchOctpassMetadataWithContext(context.TODO(), tokenURI)
}

// FetchOctpassMetadataWithContext returns *OctpassMetadata fetch Metadata from tokenURI with Context
func FetchOctpassMetadataWithContext(ctx context.Context, tokenURI string) (*OctpassMetadata, error) {
	client := new(http.Client)
	req, err := http.NewRequestWithContext(ctx, "GET", tokenURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status: %d, msg: %s", resp.StatusCode, string(body))
	}

	return UnmarshalERC721Metadata(body)
}

// NewOctpassMetadata returns *NewOctpassMetadata
func NewOctpassMetadata() (*OctpassMetadata, error) {
	return new(OctpassMetadata), nil
}

// UnmarshalOctpassMetadata unmarshals OctpassMetadata
func UnmarshalERC721Metadata(data []byte) (*OctpassMetadata, error) {
	var r = &OctpassMetadata{}
	err := json.Unmarshal(data, r)
	return r, err
}

// Marshal marshals OctpassMetadata
func (e *OctpassMetadata) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// SetAttributes convert to JSON.RawMessage and set to Attributes
func (e *OctpassMetadata) SetAttributes(attributes interface{}) error {
	byte, err := json.Marshal(attributes)
	if err != nil {
		return err
	}
	raw := json.RawMessage(byte)
	e.Attributes = &raw
	return nil
}
