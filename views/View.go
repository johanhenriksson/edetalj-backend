package views

import (
    "fmt"
    "net/http"
//    "github.com/johanhenriksson/edetalj-backend/api"
)

type View struct {
    Template *Template

}

type ViewContext struct {
    Writer  http.ResponseWriter
    Vars    map[string]interface{}
}

func NewView(template *Template) *View {
    return &View {
        Template: template,
    }
}

func (view View) Render(context ViewContext) {
    fmt.Println("View Render", context)
    view.Template.Render(context)
}
