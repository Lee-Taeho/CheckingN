package interfaces

// bson tag will tell golang to extract that specific field from mongodb into the variable
// json tag will be the way it is called from front end
type Tutor struct {
	Name            string     `bson:"name" json:"name"`
	Students        []*Student `bson:"students" json:"students"`
	FluentLanguages []string   `bson:"fluent_languages" json:"fluent_languages"`
	Courses         []*Course  `bson:"courses" json:"courses"`
}

type Student struct {
	Name string `bson:"name" json:"name"`
}

type Course struct {
	Name       string `bson:"name" json:"name"`
	Department string `bson:"department" json:"department"`
}
