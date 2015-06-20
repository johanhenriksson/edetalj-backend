package views

import (
    "os"
    "fmt"
    "github.com/flosch/pongo2"
)

type Theme struct {
    Name    string
    Path    string
    Parent  *Theme

    Templates   map[string]*Template
}

type Template struct {
    Theme       *Theme
    Name        string
    Path        string

    Template    *pongo2.Template
}

/* Instantiate new theme */
func NewTheme(name string) *Theme {
    return &Theme {
        Name: name,
        Path: name,
        Templates: make(map[string]*Template),
    }
}

func (t *Theme) File(name string) string {
    filename := fmt.Sprintf("templates/%s/%s", t.Path, name)
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        fmt.Printf("no such file or directory: %s", filename)
    }
    return filename
}

func (t *Theme) GetTemplate(name string) *Template {
    fmt.Println("Get template", name)
    if template, ok := t.Templates[name]; ok {
        fmt.Println("Template cached")
        return template
    }
    return t.LoadTemplate(name)
}

func (t *Theme) LoadTemplate(name string) *Template {
    filename := t.File(name)

    template := &Template {
        Theme: t,
        Name: name,
        Path: filename,

        Template: pongo2.Must(pongo2.FromFile(filename)),
    }
    t.Templates[name] = template
    fmt.Println("Loaded template", name, ":", filename)
    return template
}

func (t *Template) Render(context ViewContext) {
    fmt.Println("Render Template")
    fmt.Println(context.Writer)
    t.Template.ExecuteWriter(context.Vars, context.Writer)
}
