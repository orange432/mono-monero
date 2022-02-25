package cache

import "html/template"

type AppCache struct {
	Templates        map[string]*template.Template
	LastGotTemplates int64
}
