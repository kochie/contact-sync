package proxycurl

type Person struct {
	AccomplishmentCourses       []interface{} `json:"accomplishment_courses"`
	AccomplishmentHonorsAwards  []interface{} `json:"accomplishment_honors_awards"`
	AccomplishmentOrganisations []interface{} `json:"accomplishment_organisations"`
	AccomplishmentPatents       []interface{} `json:"accomplishment_patents"`
	AccomplishmentProjects      []struct {
		Description string `json:"description"`
		EndsAt      struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"ends_at"`
		StartsAt struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"starts_at"`
		Title string      `json:"title"`
		URL   interface{} `json:"url"`
	} `json:"accomplishment_projects"`
	AccomplishmentPublications []interface{} `json:"accomplishment_publications"`
	AccomplishmentTestScores   []interface{} `json:"accomplishment_test_scores"`
	Activities                 []interface{} `json:"activities"`
	Articles                   []interface{} `json:"articles"`
	BackgroundCoverImageURL    interface{}   `json:"background_cover_image_url"`
	BirthDate                  interface{}   `json:"birth_date"`
	Certifications             []struct {
		Authority     string      `json:"authority"`
		DisplaySource interface{} `json:"display_source"`
		EndsAt        interface{} `json:"ends_at"`
		LicenseNumber interface{} `json:"license_number"`
		Name          string      `json:"name"`
		StartsAt      interface{} `json:"starts_at"`
		URL           interface{} `json:"url"`
	} `json:"certifications"`
	City            string      `json:"city"`
	Connections     interface{} `json:"connections"`
	Country         string      `json:"country"`
	CountryFullName string      `json:"country_full_name"`
	Education       []struct {
		DegreeName  string      `json:"degree_name"`
		Description interface{} `json:"description"`
		EndsAt      struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"ends_at"`
		FieldOfStudy             string      `json:"field_of_study"`
		LogoURL                  string      `json:"logo_url"`
		School                   string      `json:"school"`
		SchoolLinkedinProfileURL interface{} `json:"school_linkedin_profile_url"`
		StartsAt                 struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"starts_at"`
	} `json:"education"`
	Experiences []struct {
		Company                   string      `json:"company"`
		CompanyLinkedinProfileURL string      `json:"company_linkedin_profile_url"`
		Description               interface{} `json:"description"`
		EndsAt                    interface{} `json:"ends_at"`
		Location                  interface{} `json:"location"`
		LogoURL                   string      `json:"logo_url"`
		StartsAt                  struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"starts_at"`
		Title string `json:"title"`
	} `json:"experiences"`
	Extra struct {
		FacebookProfileID interface{} `json:"facebook_profile_id"`
		GithubProfileID   interface{} `json:"github_profile_id"`
		TwitterProfileID  interface{} `json:"twitter_profile_id"`
	} `json:"extra"`
	FirstName      string        `json:"first_name"`
	FullName       string        `json:"full_name"`
	Gender         interface{}   `json:"gender"`
	Groups         []interface{} `json:"groups"`
	Headline       string        `json:"headline"`
	Industry       interface{}   `json:"industry"`
	InferredSalary struct {
		Max interface{} `json:"max"`
		Min interface{} `json:"min"`
	} `json:"inferred_salary"`
	Interests        []interface{} `json:"interests"`
	Languages        []string      `json:"languages"`
	LastName         string        `json:"last_name"`
	Occupation       string        `json:"occupation"`
	PeopleAlsoViewed []struct {
		Link     string      `json:"link"`
		Location interface{} `json:"location"`
		Name     string      `json:"name"`
		Summary  string      `json:"summary"`
	} `json:"people_also_viewed"`
	PersonalEmails         []interface{} `json:"personal_emails"`
	PersonalNumbers        []interface{} `json:"personal_numbers"`
	ProfilePicURL          string        `json:"profile_pic_url"`
	PublicIdentifier       string        `json:"public_identifier"`
	Recommendations        []string      `json:"recommendations"`
	SimilarlyNamedProfiles []interface{} `json:"similarly_named_profiles"`
	Skills                 []interface{} `json:"skills"`
	State                  string        `json:"state"`
	Summary                string        `json:"summary"`
	VolunteerWork          []struct {
		Cause                     string      `json:"cause"`
		Company                   string      `json:"company"`
		CompanyLinkedinProfileURL string      `json:"company_linkedin_profile_url"`
		Description               string      `json:"description"`
		EndsAt                    interface{} `json:"ends_at"`
		LogoURL                   string      `json:"logo_url"`
		StartsAt                  struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"starts_at"`
		Title string `json:"title"`
	} `json:"volunteer_work"`
}
