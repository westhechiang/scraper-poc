package models

import "github.com/jinzhu/gorm"

// Product
type Product struct {
	gorm.Model
        Id int64 `json:"id"`
        Name string `json:"name"`
        "slug": "bubba-kush-lone-ranger-flower",
        "dispensary": {
            "id": 1053,
            "name": "Lightshade Peoria - Recreational",
            "slug": "lightshade-peoria-recreational",
            "enabled_delivery": false,
            "enabled_pickup": true,
            "lat": "39.7732103",
            "lng": "-104.84904749999998"
        },
        "description": null,
        "category": {
            "id": 1,
            "name": "Flowers",
            "slug": "flowers"
        },
        "minimum_price": null,
        "units": {
            "data": []
        },
        "images": {
            "data": [
                {
                    "id": 5993695,
                    "name": "flowers",
                    "aspectRatio": 1,
                    "blur": "https://cdn.greenrush.com/media/6/9/5/5993695/c/blur.jpg",
                    "tiny": "https://cdn.greenrush.com/media/6/9/5/5993695/c/tiny.jpg",
                    "small": "https://cdn.greenrush.com/media/6/9/5/5993695/c/small.jpg",
                    "listing": "https://cdn.greenrush.com/media/6/9/5/5993695/c/listing.jpg",
                    "large": "https://cdn.greenrush.com/media/6/9/5/5993695/c/large.jpg",
                    "zoom": "https://cdn.greenrush.com/media/6/9/5/5993695/c/zoom.jpg",
                    "medium": "https://cdn.greenrush.com/media/6/9/5/5993695/c/medium.jpg",
                    "thumb": "https://cdn.greenrush.com/media/6/9/5/5993695/c/thumb.jpg",
                    "square_medium": "https://cdn.greenrush.com/media/6/9/5/5993695/c/square_medium.jpg",
                    "square_small": "https://cdn.greenrush.com/media/6/9/5/5993695/c/square_small.jpg",
                    "listing_medium": "https://cdn.greenrush.com/media/6/9/5/5993695/c/listing_medium.jpg",
                    "listing_small": "https://cdn.greenrush.com/media/6/9/5/5993695/c/listing_small.jpg"
                }
            ]
        },
        "active": false,
        "review_rating": 0,
        "positive_effects": [],
        "negative_effects": [],
        "flavors": [],
        "uses": [],
        "updated_at": "2019-05-11T21:26:37+00:00",
        "created_at": "2019-05-11T19:49:56+00:00",
        "_links": {
            "web": {
                "href": "https://www.greenrush.com/dispensary/lightshade-peoria-recreational/menu/bubba-kush-lone-ranger-flower"
            },
            "dispensary": {
                "self": {
                    "href": "https://www.greenrush.com/dispensary/lightshade-peoria-recreational/menu"
                },
                "web": {
                    "href": "https://www.greenrush.com/dispensary/lightshade-peoria-recreational/menu"
                }
            }
        },
        "logo": {
            "data": {
                "id": 5993695,
                "name": "flowers",
                "aspectRatio": 1,
                "blur": "https://cdn.greenrush.com/media/6/9/5/5993695/c/blur.jpg",
                "tiny": "https://cdn.greenrush.com/media/6/9/5/5993695/c/tiny.jpg",
                "small": "https://cdn.greenrush.com/media/6/9/5/5993695/c/small.jpg",
                "listing": "https://cdn.greenrush.com/media/6/9/5/5993695/c/listing.jpg",
                "large": "https://cdn.greenrush.com/media/6/9/5/5993695/c/large.jpg",
                "zoom": "https://cdn.greenrush.com/media/6/9/5/5993695/c/zoom.jpg",
                "medium": "https://cdn.greenrush.com/media/6/9/5/5993695/c/medium.jpg",
                "thumb": "https://cdn.greenrush.com/media/6/9/5/5993695/c/thumb.jpg",
                "square_medium": "https://cdn.greenrush.com/media/6/9/5/5993695/c/square_medium.jpg",
                "square_small": "https://cdn.greenrush.com/media/6/9/5/5993695/c/square_small.jpg",
                "listing_medium": "https://cdn.greenrush.com/media/6/9/5/5993695/c/listing_medium.jpg",
                "listing_small": "https://cdn.greenrush.com/media/6/9/5/5993695/c/listing_small.jpg"
            }
        }
}
