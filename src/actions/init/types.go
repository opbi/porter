package actions

type Answers struct {
	ResourceType     string
	ResourceName     string
	OwnerType        string
	OwnerGithub      string
	ReleasePublic    bool
	ReleaseNamespace string
}

type PorterRc struct {
	Resource struct {
		Type        string
		Name        string
		Language    string
		Description string
		Keywords    string
	}
	Owner struct {
		Type             string
		Team             string
		Name             string
		Email            string
		Github           string
		ReleasePublic    bool
		ReleaseNamespace string
	}
	Config struct {
		Tempalte        string
		LanguageVersion string
		BuildTool       string
		CodeTester      string
		TypeChecker     string
		CodeLinter      string
		CommitHooks     string
		CommitLinter    string
		ReleaseConfig   string
		ContainerConfig string
		DeployConfig    string
		ReleaseCI       string
		CoverageCI      string
		CodeQualityCI   string
		DepsUpgrader    string
		License         string
	}
}
