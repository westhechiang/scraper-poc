package scrapers

type GreenRushImage struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	AspectRatio   uint   `json:"aspectRatio"`
	Blur          string `json:"blur"`
	Tiny          string `json:"tiny"`
	Small         string `json:"small"`
	Listing       string `json:"listing"`
	Large         string `json:"large"`
	Zoom          string `json:"zoom"`
	Medium        string `json:"medium"`
	Thumb         string `json:"thumb"`
	SquareMedium  string `json:"square_medium"`
	SquareSmall   string `json:"square_small"`
	ListingMedium string `json:"listing_medium"`
	ListingSmall  string `json:"listing_small"`
}

type GreenRushImages struct {
	Data []GreenRushImage `json:"data"`
}

type GreenRushDispensary struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	EnabledDelivery bool   `json:"enabled_delivery"`
	EnabledPickup   bool   `json:"enabled_pickup"`
	Lat             string `json:"lat"`
	Lng             string `json:"lng"`
}

type GreenRushCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GreenRushLink struct {
	Href string `json:"href"`
}

type GreenRushDispensaryLink struct {
	Self struct {
		Href GreenRushLink
	} `json:"self"`
	Web struct {
		Href GreenRushLink
	} `json:"web"`
}

type GreenRushLinks struct {
	Web        GreenRushLink           `json:"web"`
	Dispensary GreenRushDispensaryLink `json:"dispensary"`
}

type GreenRushUnits struct {
	Data []string `json:"data"`
}

type GreenRushProduct struct {
	ID              uint                `json:"id"`
	Name            string              `json:"name"`
	Slug            string              `json:"slug"`
	Dispensary      GreenRushDispensary `json:"dispensary"`
	Description     string              `json:"description"`
	Category        GreenRushCategory   `json:"category"`
	MinimumPrice    uint                `json:"minimum_price"`
	Units           GreenRushUnits      `json:"units"`
	Images          GreenRushImages     `json:"images"`
	Active          bool                `json:"active"`
	ReviewRating    uint                `json:"review_rating"`
	PositiveEffects []string            `json:"positive_effects"`
	NegativeEffects []string            `json:"negative_effects"`
	Flavors         []string            `json:"flavors"`
	Uses            []string            `json:"uses"`
	UpdatedAt       string              `json:"updated_at"`
	CreatedAt       string              `json:"created_at"`
	Links           GreenRushLinks      `json:"_links"`
	Logo            GreenRushImages     `json:"logo"`
}

// GreenRushProductsAPIResponsee type
type GreenRushProductAPIResponse struct {
	GreenRushProduct GreenRushProduct `json:"data"`
}
