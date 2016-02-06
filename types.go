package clearbit

type Enrichment struct {
	Person  *Person  `json:"person"`
	Company *Company `json:"company"`
}

type Person struct {
	ID         *string     `json:"id"`
	Name       *Name       `json:"name"`
	Email      *string     `json:"email"`
	Gender     *string     `json:"gender"`
	Location   *string     `json:"location"`
	TimeZone   *string     `json:"timeZone"`
	UtcOffset  *int        `json:"utcOffset"`
	Geo        *Geo        `json:"geo"`
	Bio        *string     `json:"bio"`
	Site       *string     `json:"site"`
	Avatar     *string     `json:"avatar"`
	Employment *Job        `json:"employment"`
	Facebook   *Facebook   `json:"facebook"`
	Github     *Github     `json:"github"`
	Twitter    *Twitter    `json:"twitter"`
	Linkedin   *Linkedin   `json:"linkedin"`
	Googleplus *Googleplus `json:"googleplus"`
	Aboutme    *Aboutme    `json:"aboutme"`
	Gravatar   *Gravatar   `json:"gravatar"`
	Fuzzy      *bool       `json:"fuzzy"`
}

type Company struct {
	ID            *string     `json:"id"`
	Name          *string     `json:"name"`
	LegalName     *string     `json:"legalName"`
	Domain        *string     `json:"domain"`
	DomainAliases []string    `json:"domainAliases"`
	URL           *string     `json:"url"`
	Site          *Site       `json:"site"`
	Category      *Category   `json:"category"`
	Tags          []string    `json:"tags"`
	Description   *string     `json:"description"`
	FoundedDate   *string     `json:"foundedDate"`
	Location      *string     `json:"location"`
	TimeZone      *string     `json:"timeZone"`
	UtcOffset     *int        `json:"utcOffset"`
	Geo           *Geo        `json:"geo"`
	Logo          *string     `json:"logo"`
	Facebook      *Facebook   `json:"facebook"`
	Linkedin      *Linkedin   `json:"linkedin"`
	Twitter       *Twitter    `json:"twitter"`
	Crunchbase    *Crunchbase `json:"crunchbase"`
	EmailProvider *bool       `json:"emailProvider"`
	Type          *string     `json:"type"`
	Ticker        *string     `json:"ticker"`
	Phone         *string     `json:"phone"`
	Metrics       *Metrics    `json:"metrics"`
	Tech          []string    `json:"tech"`
}

type Name struct {
	FullName   *string `json:"fullName"`
	GivenName  *string `json:"givenName"`
	FamilyName *string `json:"familyName"`
}

type Geo struct {
	StreetNumber *string  `json:"streetNumber"`
	StreetName   *string  `json:"streetName"`
	SubPremise   *string  `json:"subPremise"`
	City         *string  `json:"city"`
	PostalCode   *string  `json:"postalCode"`
	State        *string  `json:"state"`
	StateCode    *string  `json:"stateCode"`
	Country      *string  `json:"country"`
	CountryCode  *string  `json:"countryCode"`
	Lat          *float64 `json:"lat"`
	Lng          *float64 `json:"lng"`
}

type Job struct {
	Domain    *string `json:"domain"`
	Name      *string `json:"name"`
	Title     *string `json:"title"`
	Role      *string `json:"role"`
	Seniority *string `json:"seniority"`
}

type Github struct {
	Handle    *string `json:"handle"`
	ID        *int    `json:"id"`
	Avatar    *string `json:"avatar"`
	Company   *string `json:"company"`
	Blog      *string `json:"blog"`
	Followers *int    `json:"followers"`
	Following *int    `json:"following"`
}

type Twitter struct {
	Handle    *string `json:"handle"`
	ID        *int    `json:"id"`
	Bio       *string `json:"bio"`
	Followers *int    `json:"followers"`
	Following *int    `json:"following"`
	Statuses  *int    `json:"statuses"`
	Favorites *int    `json:"favorites"`
	Location  *string `json:"location"`
	Site      *string `json:"site"`
	Avatar    *string `json:"avatar"`
}

type Facebook struct {
	Handle *string `json:"handle"`
}

type Googleplus struct {
	Handle *string `json:"handle"`
}

type Linkedin struct {
	Handle *string `json:"handle"`
}

type Crunchbase struct {
	Handle *string `json:"handle"`
}

type Aboutme struct {
	Handle *string `json:"handle"`
	Bio    *string `json:"bio"`
	Avatar *string `json:"avatar"`
}

type Gravatar struct {
	Handle  *string  `json:"handle"`
	Urls    []string `json:"urls"`
	Avatar  *string  `json:"avatar"`
	Avatars []Avatar `json:"avatars"`
}

type Avatar struct {
	URL  *string `json:"url"`
	Type *string `json:"type"`
}

type Site struct {
	URL             *string `json:"url"`
	Title           *string `json:"title"`
	H1              *string `json:"h1"`
	MetaDescription *string `json:"metaDescription"`
	MetaAuthor      *string `json:"metaAuthor"`
}

type Category struct {
	Sector        *string `json:"sector"`
	IndustryGroup *string `json:"industryGroup"`
	Industry      *string `json:"industry"`
	SubIndustry   *string `json:"subIndustry"`
}

type Metrics struct {
	AlexaUsRank     *int     `json:"alexaUsRank"`
	AlexaGlobalRank *int     `json:"alexaGlobalRank"`
	GoogleRank      *int     `json:"googleRank"`
	Employees       *int     `json:"employees"`
	MarketCap       *float64 `json:"marketCap"`
	Raised          *float64 `json:"raised"`
	AnnualRevenue   *float64 `json:"annualRevenue"`
}
