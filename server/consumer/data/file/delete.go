package file

type Delete struct {
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

type RemoveId struct {
	Ids []string `json:"ids"`
}