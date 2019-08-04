package main

import (
	"os"
	"testing"
	"text/template"
)

func getComposition() string {
	return ""
}
func getFreeStyle() string {
	return ""
}
func getBuild() {

}
func getPush() {}
func getManualApproval() {

}

type Values struct {
	Stages []string
	Steps  []string
}

func TestSteps(t *testing.T) {
	step1 := `
     step1: 
      image: alpine 
        commands:
         - echo works
	`
	step2 :=
		` 
      step1: 
        image: alpine 
        commands:
          - echo works
	`
	values := Values{[]string{
		"stage1",
		"stage2",
		"stage3"},
		[]string{step1, step2},
	}
	tmpl, err := template.
		ParseFiles("cf.tpl")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, values)
	if err != nil {
		panic(err)
	}
}
