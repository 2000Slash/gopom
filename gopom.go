package gopom

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// https://maven.apache.org/ref/3.6.3/maven-model/maven.html#class_dependency

func Parse(path string) (*Project, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	var project Project

	err = xml.Unmarshal(b, &project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

type Project struct {
	XMLName              xml.Name             `xml:"project"`
	ModelVersion         string               `xml:"modelVersion"`
	Parent               Parent               `xml:"parent"`
	GroupID              string               `xml:"groupId"`
	ArtifactID           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	Packaging            string               `xml:"packaging"`
	Name                 string               `xml:"name"`
	Description          string               `xml:"description"`
	URL                  string               `xml:"url"`
	InceptionYear        string               `xml:"inceptionYear"`
	Organization         Organization         `xml:"organization"`
	Licenses             []License            `xml:"licenses>license"`
	Developers           []Developer          `xml:"developers>developer"`
	Contributors         []Contributor        `xml:"contributors>contributor"`
	MailingLists         []MailingList        `xml:"mailingLists>mailingList"`
	Prerequisites        Prerequisites        `xml:"prerequisites"`
	Modules              []string             `xml:"modules>module"`
	SCM                  Scm                  `xml:"scm"`
	IssueManagement      IssueManagement      `xml:"issueManagement"`
	CIManagement         CIManagement         `xml:"ciManagement"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
	Dependencies         []Dependency         `xml:"dependencies>dependency"`
	Repositories         []Repository         `xml:"repositories>repository"`
	PluginRepositories   []PluginRepository   `xml:"pluginRepositories>pluginRepository"`
	Build                Build                `xml:"build"`
	Reporting            Reporting            `xml:"reporting"`
	// TODO Profiles
}

type Parent struct {
	GroupID      string `xml:"groupId"`
	ArtifactID   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

type Organization struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type License struct {
	Name         string `xml:"name"`
	URL          string `xml:"url"`
	Distribution string `xml:"distribution"`
	Comments     string `xml:"comments"`
}

type Developer struct {
	ID              string   `xml:"id"`
	Name            string   `xml:"name"`
	Email           string   `xml:"email"`
	URL             string   `xml:"url"`
	Organization    string   `xml:"organization"`
	OrganizationURL string   `xml:"organizationUrl"`
	Roles           []string `xml:"roles>role"`
	Timezone        string   `xml:"timezone"`
	// TODO Implement properties/key=value
}

type Contributor struct {
	Name            string   `xml:"name"`
	Email           string   `xml:"email"`
	URL             string   `xml:"url"`
	Organization    string   `xml:"organization"`
	OrganizationURL string   `xml:"organizationUrl"`
	Roles           []string `xml:"roles>role"`
	Timezone        string   `xml:"timezone"`
	// TODO Implement properties/key=value
}

type MailingList struct {
	Name          string   `xml:"name"`
	Subscribe     string   `xml:"subscribe"`
	Unsubscribe   string   `xml:"unsubscribe"`
	Post          string   `xml:"post"`
	Archive       string   `xml:"archive"`
	OtherArchives []string `xml:"otherArchives>otherArchive"`
}

type Prerequisites struct {
	Maven string `xml:"maven"`
}

type Scm struct {
	Connection          string `xml:"connection"`
	DeveloperConnection string `xml:"developerConnection"`
	Tag                 string `xml:"tag"`
	URL                 string `xml:"url"`
}

type IssueManagement struct {
	System string `xml:"system"`
	URL    string `xml:"url"`
}

type CIManagement struct {
	System    string     `xml:"system"`
	URL       string     `xml:"url"`
	Notifiers []Notifier `xml:"notifiers>notifier"`
}

type Notifier struct {
	Type          string `xml:"type"`
	SendOnError   bool   `xml:"sendOnError"`
	SendOnFailure bool   `xml:"sendOnFailure"`
	SendOnSuccess bool   `xml:"sendOnSuccess"`
	SendOnWarning bool   `xml:"sendOnWarning"`
	Address       string `xml:"address"`
	// TODO add Configuration field
}

type DependencyManagement struct {
	Dependencies []Dependency `xml:"dependencies"`
}

type Dependency struct {
	GroupID    string      `xml:"groupId"`
	ArtifactID string      `xml:"artifactId"`
	Version    string      `xml:"version"`
	Type       string      `xml:"type"`
	Classifier string      `xml:"classifier"`
	Scope      string      `xml:"scope"`
	SystemPath string      `xml:"systemPath"`
	Exclusions []Exclusion `xml:"exclusions"`
	Optional   string      `xml:"optional"`
}

type Exclusion struct {
	ArtifactID string `xml:"artifactID"`
	GroupID    string `xml:"groupId"`
}

type Repository struct {
	UniqueVersion bool             `xml:"uniqueVersion"`
	Releases      RepositoryPolicy `xml:"releases"`
	Snapshots     RepositoryPolicy `xml:"snapshots"`
	ID            string           `xml:"id"`
	Name          string           `xml:"name"`
	URL           string           `xml:"URL"`
	Layout        string           `xml:"layout"`
}

type RepositoryPolicy struct {
	Enabled        string `xml:"enabled"`
	UpdatePolicy   string `xml:"updatePolicy"`
	ChecksumPolicy string `xml:"checksumPolicy"`
}

type PluginRepository struct {
	Releases  RepositoryPolicy `xml:"releases"`
	Snapshots RepositoryPolicy `xml:"snapshots"`
	ID        string           `xml:"id"`
	Name      string           `xml:"name"`
	URL       string           `xml:"URL"`
	Layout    string           `xml:"layout"`
}

type Build struct {
	SourceDirectory       string           `xml:"sourceDirectory"`
	ScriptSourceDirectory string           `xml:"scriptSourceDirectory"`
	TestSourceDirectory   string           `xml:"testSourceDirectory"`
	OutputDirectory       string           `xml:"outputDirectory"`
	TestOutputDirectory   string           `xml:"testOutputDirectory"`
	Extensions            []Extension      `xml:"extensions"`
	DefaultGoal           string           `xml:"defaultGoal"`
	Resources             []Resource       `xml:"resources"`
	TestResource          []Resource       `xml:"testResources"`
	Directory             string           `xml:"directory"`
	FinalName             string           `xml:"finalName"`
	Filters               string           `xml:"filters>filter"`
	PluginManagement      PluginManagement `xml:"pluginManagement"`
	Plugins               []Plugin         `xml:"plugins"`
}

type Extension struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"ArtifactId"`
	Version    string `xml:"version"`
}

type Resource struct {
	TargetPath string   `xml:"targetPath"`
	Filtering  string   `xml:"filtering"`
	Directory  string   `xml:"directory"`
	Includes   []string `xml:"includes>include"`
	Excludes   []string `xml:"excludes>exclude"`
}

type PluginManagement struct {
	Plugins []Plugin `xml:"plugins"`
}

type Plugin struct {
	GroupID      string            `xml:"groupId"`
	ArtifactID   string            `xml:"artifactId"`
	Version      string            `xml:"version"`
	Extensions   string            `xml:"extensions"`
	Executions   []PluginExecution `xml:"executions"`
	Dependencies []Dependency      `xml:"dependencies"`
	Inherited    string            `xml:"inherited"`
}

type PluginExecution struct {
	ID        string   `xml:"id"`
	Phase     string   `xml:"phase"`
	Goals     []string `xml:"goals>goal"`
	Inherited string   `xml:"inherited"`
}

type Reporting struct {
	ExcludeDefaults string   `xml:"excludeDefaults"`
	OutputDirectory string   `xml:"outputDirectory"`
	Plugins         []Plugin `xml:"plugins"`
}
