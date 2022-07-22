package picbed

type Github struct {
	RepoName     string `json:"repo-name"`
	BranchName   string `json:"branch-name"`
	Token        string `json:"token"`
	StoragePath  string `json:"storage-path"`
	DefineDomain string `json:"define-domain"`
}

func (g *Github) CheckProperty() (propertyName string, isEmpty bool) {
	if g.Token == "" {
		return "Token", true
	}
	if g.RepoName == "" {
		return "RepoName", true
	}
	if g.BranchName == "" {
		return "BranchName", true
	}
	if g.StoragePath == "" {
		return "StoragePath", true
	}
	//if g.DefineDomain == "" {
	//	return "Token", true
	//}
	return "", false
}

func (g *Github) UploadFile(path string) string {
	// urlString:=fmt.Sprintf("https://api.github.com/%s/$s", g.)
	return ""
}
