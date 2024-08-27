package entry

type UserEntry struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      string `json:"created_on"`
	LastModifiedOn string `json:"last_modified_on"`
	Version        int    `json:"version"`
}
