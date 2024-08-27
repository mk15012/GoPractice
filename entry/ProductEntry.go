package entry

type ProductEntry struct {
	Description    string `json:"description"`
	Quantity       int    `json:"quantity"`
	CreatedBy      string `json:"created_by"`
	CreatedOn      string `json:"created_on"`
	LastModifiedOn string `json:"last_modified_on"`
	Version        int    `json:"version"`
}
