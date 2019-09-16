package figma

// EntryUser is the user entry
type EntryUser struct {
	// ID is the id of a user
	ID string `json:"id"`
	// Handle is the name of a user
	Handle string `json:"handle"`
	// ImgURL is the url link to the user's profile image
	ImgURL string `json:"img_url"`
}
