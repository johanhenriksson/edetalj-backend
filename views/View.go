package views

import (
    "net/http"
//    "github.com/johanhenriksson/edetalj-backend/api"
)

type View struct {
    Pattern     string
    Template    *Template

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
    view.Template.Render(context)
}
