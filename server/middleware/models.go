package middleware

// bson tag will tell golang to extract that specific field from mongodb into the variable
// json tag will be the way it is called from front end
type Tutor struct {
	Name            string     `bson:"name" json:"name"`
	Students        []*Student `bson:"students" json:"students"`
	FluentLanguages []string   `bson:"fluent_languages" json:"fluent_languages"`
	Courses         []*Course  `bson:"courses" json:"courses"`
}

type Student struct {
	FirstName string     `bson:"first_name" json:"first_name"`
	LastName  string     `bson:"last_name" json:"last_name"`
	LoginInfo *LoginInfo `bson:"login_info" json:"login_info"`
}

type LoginInfo struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type Course struct {
	Name       string `bson:"name" json:"name"`
	Department string `bson:"department" json:"department"`
}
