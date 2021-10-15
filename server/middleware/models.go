package middleware

type Tutor struct {
	Uuid            int        `bson:"uuid" json:"uuid"`
	Name            string     `bson:"name" json:"name"`
	Students        []*Student `bson:"students" json:"students"`
	FluentLanguages []string   `bson:"fluent_languages" json:"fluent_languages"`
	Courses         []*Course  `bson:"courses" json:"courses"`
}

type Student struct {
	Uuid      int    `bson:"uuid" json:"uuid"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
}

type GoogleUser struct {
	Id            string `json:"id" bson:"id"`
	Email         string `json:"email" bson:"email"`
	VerifiedEmail bool   `json:"verified_email" bson:"verified_email"`
	FirstName     string `json:"given_name" bson:"first_name"`
	LastName      string `json:"family_name" bson:"last_name"`
	PictureLink   string `json:"picture" bson:"picture_link"`
	Locale        string `json:"locale" bson:"locale"`
}

type LoginRequest struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type Course struct {
	Name       string `bson:"name" json:"name"`
	Department string `bson:"department" json:"department"`
}

type Cookie struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	ExpiresAt int64  `json:"expires_at"`
}

// "first_name": "Mcclovin", "last_name":"", "email":"jonahhill@gmail.com", "password":"dartmouth"
