package file

type Create struct {
	Id         string `json:"id"`
	ProjectId  string `json:"projectId"`
	Name       string `json:"name"`
	Folder     string `json:"folder"`
	FolderPath string `json:"folderPath"`
	Path       string `json:"path"`
	Url        string `json:"url"`
	CreatedUp  string `json:"createdUp"`
	UpdateUp   string `json:"updateUp"`
}
