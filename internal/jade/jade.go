package jade

import (
	"database/sql"
	"time"
)

type Page struct {
	ID          string       `db:"id" orm:"primaryKey"`
	Name        string       `db:"name"`
	Url         string       `db:"url"`
	Status      string       `db:"status"`
	VerticalID  string       `db:"vertical_id"`
	PublicID    string       `db:"public_id"`
	CountryID   string       `db:"country_id"`
	CountryName string       `db:"country_name"`
	RegionID    string       `db:"region_id"`
	RegionName  string       `db:"region_name"`
	CityID      string       `db:"city_id"`
	CityName    string       `db:"city_name"`
	StarRating  float64      `db:"star_rating"`
	Version     int          `db:"version"`
	CreatedDate time.Time    `db:"created_date"`
	CreatedBy   string       `db:"created_by"`
	UpdatedDate sql.NullTime `db:"updated_date"`
	UpdatedBy   string       `db:"updated_by"`
}

type SeoTag struct {
	ID              string       `db:"id" orm:"primaryKey"`
	PageID          string       `db:"page_id"`
	PageTitle       string       `db:"page_title"`
	SlugId          string       `db:"slug_id"`
	SlugEn          string       `db:"slug_en"`
	TitleTagId      string       `db:"title_tag_id"`
	TitleTagEn      string       `db:"title_tag_en"`
	DescriptionId   string       `db:"description_id"`
	DescriptionEn   string       `db:"description_en"`
	HeaderId        string       `db:"h1_id"`
	HeaderEn        string       `db:"h1_en"`
	RobotIndexId    string       `db:"robot_index_id"`
	RobotIndexEn    string       `db:"robot_index_en"`
	RobotFollowId   string       `db:"robot_follow_id"`
	RobotFollowEn   string       `db:"robot_follow_en"`
	CanonicalId     string       `db:"canonical_id"`
	CanonicalEn     string       `db:"canonical_en"`
	OgTitleTagId    string       `db:"og_title_tag_id"`
	OgTitleTagEn    string       `db:"og_title_tag_en"`
	OgDescriptionId string       `db:"og_description_id"`
	OgDescriptionEn string       `db:"og_description_en"`
	OgImageUrlId    string       `db:"og_image_url_id"`
	OgImageUrlEn    string       `db:"og_image_url_en"`
	CreatedDate     time.Time    `db:"created_date"`
	CreatedBy       string       `db:"created_by"`
	UpdatedDate     sql.NullTime `db:"updated_date"`
	UpdatedBy       string       `db:"updated_by"`
}

type Content struct {
	ID           string `db:"id" orm:"primaryKey"`
	PageID       string `db:"page_id"`
	SectionName  string `db:"section_name"`
	SectionType  string `db:"section_type"`
	SectionOrder string `db:"section_order"`
	Content      string `db:"content"`
	CreatedDate  string `db:"created_date"`
	CreatedBy    string `db:"created_by"`
	UpdatedDate  string `db:"updated_date"`
	UpdatedBy    string `db:"updated_by"`
}
