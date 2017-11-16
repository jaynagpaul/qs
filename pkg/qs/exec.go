package qs

import (
	"fmt"
	"reflect"

	"gopkg.in/AlecAivazis/survey.v1"
)

func execTemplate(tmpl Template, defaultName string) error {
	var name string
	var folder string

	if defaultName == "" {
		defaultName = "QS Template"
	}

	if s, ok := tmpl["TemplateName"].(string); ok {
		name = s
	} else {
		name = defaultName
	}
	if s, ok := tmpl["TemplateFolder"].(string); ok {
		folder = s
	} else {
		folder = "."
	}

	delete(tmpl, "TemplateName")
	delete(tmpl, "TemplateFolder")

	var questions []*survey.Question

	for key, val := range tmpl {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Slice, reflect.Array:

			// Convert slice of interface to slice of strings
			options := []string{}
			s := reflect.ValueOf(val)
			arr := make([]interface{}, s.Len())

			for i := 0; i < s.Len(); i++ {
				arr[i] = s.Index(i).Interface()
			}
			for _, v := range arr {
				options = append(options, toString(v))
			}

			if len(options) > 0 {
				questions = append(questions, &survey.Question{
					Name: key,
					Prompt: &survey.Select{
						Message: key,
						Options: options,
						Default: options[0],
					},
					Validate: survey.Required,
				})
			} else {
				questions = append(questions, &survey.Question{
					Name: key,
					Prompt: &survey.Select{
						Message: key,
						Options: options,
					},
					Validate: survey.Required,
				})
			}
		default:
			questions = append(questions, &survey.Question{
				Name: key,
				Prompt: &survey.Input{
					Message: key,
					Default: toString(val),
				},
				Validate: survey.Required,
			})
		}
	}

	var m map[string]interface{}
	m = tmpl

	// It requires map[string]interface{}, not type Template
	survey.Ask(questions, &m)

	// TODO Execute templates
	fmt.Println(name, folder)
	return nil
}

// Run all imports
// TODO, check for cyclic imports
func execImports(imports []string) error {
	for _, imp := range imports {
		if p, err := Get(imp); err == nil {
			if err := Run(p); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func toString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
