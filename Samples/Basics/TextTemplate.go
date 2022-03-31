package main

import (
	"html/template"
	"os"
)

func main() {
	/*
		create a new template and parse
		its body from a string. Templates
		are a mix of static text and “actions”
		enclosed in {{...}} that are used to dynamically insert content
	*/
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	/*
		use the template.Must function
		to panic in case Parse returns an error.
		This is especially useful for templates
		initialized in the global scope
	*/
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))
	t1.Execute(os.Stdout, "@     #### some text")
	t1.Execute(os.Stdout, 5)
	/*
		“executing” the template we generate its text
		with specific values for its actions.
		The {{.}} action is replaced by the
		 value passed as a parameter to Execute
	*/
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))

	}

	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Micky Mouse",
	})
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
	t4.Execute(os.Stdout, "String")
}
