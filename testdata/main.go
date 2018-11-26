package testdatamain

// User is application user.
// datastore-gen-doc
type User struct {
	ID   int64
	Name string
}

// Profile is user profile.
type Profile struct {
	UserID int64
	Score  int
	Age    int
	Desc   string
}

func main() {

}
