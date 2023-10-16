package openseamodels

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	// The unique public blockchain identifier for the wallet.
	Address common.Address `json:"address"`
	// The OpenSea account's username.
	Username string `json:"username"`
	// The OpenSea account's image url.
	ProfileImageUrl string `json:"profile_image_url"`
	// The OpenSea account's banner url.
	BannerImageUrl string `json:"banner_image_url"`
	// Personal website for the OpenSea user.
	Website             string                `json:"website"`
	SocialMediaAccounts []*SocialMediaAccount `json:"social_media_accounts"`
	// The OpenSea account's bio.
	Bio string `json:"bio"`
	// Date the account was first added to OpenSea.
	JoinedDate time.Time `json:"joined_date"`
}

type SocialMediaAccount struct {
	// The social media platform, e.g. twitter or instagram
	Platform string `json:"platform"`
	// The username for the social media platform
	Username string `json:"username"`
}
