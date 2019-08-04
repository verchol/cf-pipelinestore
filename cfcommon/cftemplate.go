package main

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/Masterminds/sprig"
)

type Step interface {
	GetStep(string)
}

const (
	//MainClone ...
	//{{- optional "title:" .Title | indent 4 -}}
	MainCloneTemplate = ` 
		{{printf "main_clone:"  | indent 2}}
		{{printf "type: git-clone" | indent 4 }}
		{{printf "repo:%s" .Repo | indent 4}}
		{{printf "revision:%s" .Revision | indent 4}}
		{{"git: cf_github"| indent 4 }}`

	DockerBuildTemplate = ` 
	{{.Name}}:
		type: {{.Type}}
		fail_fast: true
		image_name: {{.Image}}
		arguments:
			dockerfile: {{.Dockerfile}}

	`

	DockerPushTemplate = ` 
	{{.Name}}:
		type: {{.Type}}
		fail_fast: true
		arguments:
			dockerfile: {{.Dockerfile}}

	`

	CompositionTemplate = ` 
	{{.Name}}:
	    {{- required type: composition -}}
		title: {{.Title}}
		{{optional description: {{.Description}}
		
		working_directory: ${{a_clone_step}}
		composition:
			version: '2'
			services:
			db:
				image: postgres
		composition_candidates:
			test_service:
			image: ${{build_step}}
			command: gulp integration_test
			working_dir: /app
			environment:
				- key=value
			`
)

type StepMetadata struct {
	Name     string
	Template string
}

type MainClone struct {
	StepMetadata
	Title    string
	Repo     string
	Revision string
}
type DockerBuild struct {
	StepMetadata
	Dockerfile string
	Type       string
}
type Composition struct {
	StepMetadata
	arguments []interface{}
}

func DefaultRender(name string, stepTemplate string, values interface{}, writer io.Writer) {
	tmpl, err := template.New(name).Parse(stepTemplate)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(writer, values)
}
func (this *DockerBuild) Execute(writer io.Writer) {
	DefaultRender(this.Name, this.Template, this, writer)
}
func (this *Composition) Execute(writer io.Writer) {
	DefaultRender(this.Name, this.Template, this, writer)
}
func (this *MainClone) Execute(writer io.Writer) {
	tmpl := template.New(this.Name)

	funcs := template.FuncMap{
		"required": func(text string, value string) string {
			if len(value) > 0 {
				return text + ":" + value
			}
			return ""

		},
		"optional": func(text string, value string) string {
			if len(value) > 0 {
				return text + ":" + value
			}
			return "\n"

		},
	}

	tmpl, err := tmpl.Funcs(funcs).Funcs(sprig.FuncMap()).Parse(this.Template)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(writer, this)
	if err != nil {
		panic(err)
	}
}

func concatTemplates(templates map[string]string) (*template.Template, error) {
	var concat string
	for _, t := range templates {
		concat = concat + t
	}
	return template.New("template").Parse(concat)
}
func createTemplate(
	steps []string) {

	for _, step := range steps {
		fmt.Println(step)
		if step == "main_clone" {
			values := MainClone{StepMetadata{"main_clone1", MainCloneTemplate}, "",
				"t1", "t2"}
			values.Execute(os.Stdout)
		}
		if step == "docker_build" {
			values := DockerBuild{StepMetadata{"build1", DockerBuildTemplate},
				"./Dockerfile", "build",
			}
			values.Execute(os.Stdout)
		}
		if step == "composition" {
			args := make([]interface{}, 10)
			args[0] = "r"
			values := Composition{StepMetadata{"composition", MainCloneTemplate}, args}
			values.Execute(os.Stdout)
		}
		if step == "docker_push" {
			values := DockerBuild{StepMetadata{"build1", DockerBuildTemplate},
				"./Dockerfile", "build",
			}
			values.Execute(os.Stdout)
		}

		if step == "manual-approval" {
			values := MainClone{StepMetadata{"main_clone1", MainCloneTemplate}, "r1",
				"t1", "t2"}
			values.Execute(os.Stdout)
		}

	}

}
func main() {
	templates := make(map[string]string)
	templates["main_clone"] = MainCloneTemplate
	templates["docker_build"] = DockerBuildTemplate

	steps := []string{
		"main_clone",
		/*"composition",
		"docker_build",
		"docker_push",
		"composition",*/
	}
	/*tmpl, error := concatTemplates(templates)
	if error != nil {
		panic(error)
	}*/
	createTemplate(steps)
}
