package documentmanager

import (
	"fmt"
	"github.com/adaptive-scale/dockyard/asset"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

const OutputDir = "public"

type DocumentManager struct {
	location string
}

func New(location string) *DocumentManager {
	return &DocumentManager{location:location}
}

func (d *DocumentManager) ListAllFiles() (map[string]map[string]string, error) {
	return render(d.location, OutputDir)
}

type FileLocation struct {
	FileLocation string
	Render map[string]string
}

func render(location, root string) (map[string]map[string]string, error) {

	fileLoc := map[string]map[string]string{}
	renderedJS := map[string]string{}

	var currentPath string
	err := filepath.Walk(location,
		func(filePath string, info os.FileInfo, err error) error {

			if currentPath != path.Join(root, filepath.Dir(filePath)) {
				renderedJS = map[string]string{}
			}
			if err != nil {
				return err
			}

			if !info.IsDir() {
				if filepath.Ext(filePath) == ".md"{
					md, _ := ioutil.ReadFile(filePath)
					output := markdown.ToHTML(md, nil, nil)
					renderedJS[strings.ReplaceAll(info.Name(), filepath.Ext(info.Name()), "")] = string(output)
					join := path.Join(root, filepath.Dir(filePath))
					currentPath = join
					fileLoc[join] = renderedJS
				}
			}

			return nil
		})
	return fileLoc, err
}

type Content struct {
	index string
	id string
	content string
}

func (d *DocumentManager) GetMenu(menuAndContent map[string]string) (string,string, string) {
	fmt.Println("==> generating menu")

	start := "{"
	end := "}"


	var contents []Content
	var menu string
	var activated string
	for key, value := range menuAndContent {

		key := strings.Split(key, "_")

		if len(key) >= 2 {

			var content Content

			content.index = key[0]
			content.content = value
			content.id = strings.Join(key[1:], "_")

			contents = append(contents, content)
		}
	}

	sort.Slice(contents, func(i, j int) bool { return contents[i].index < contents[j].index })

	for _, content := range contents {
		if activated == "" {
			activated = content.id
		}
		start = start + "'" + content.id + "':" + "`" + content.content + "`,"
		menu = menu + `<li id="` + content.id + `" style="border: 0px; border-radius:6px;cursor: pointer;" class="list list-group-item" onclick="activate(this);">` + beautify(content.id) + `</li>
`
	}

	 start = start[:len(start)-1] + end

	 return start, menu, activated
}

func (d *DocumentManager) GenerateJS(start string, activated string) string {
	fmt.Println("==> generating content")

	return `var activated = '`+activated+`';
var advent = {id: '`+activated+`'};
var content=` + start

}

func (d *DocumentManager) GenerateIndexHTML(theme, branding, menu, js string) (string, error) {

	fmt.Println("==> generating template")

	css, err := asset.Asset("templates/style.css")
	if err != nil {
		return "", err
	}

	newTemplate, err := asset.Asset("templates/index_template.html")
	if err != nil {
		return "", err
	}

	theme = "templates/theme/" + theme + ".css"
	themeInfo, err := asset.Asset(theme)
	if err != nil {
		return "", err
	}

	tmpl := strings.ReplaceAll(string(newTemplate), "<MENU />", menu)
	tmpl = strings.ReplaceAll(tmpl, "<CSS />", string(css))
	tmpl = strings.ReplaceAll(tmpl, "<JS />", js)
	tmpl = strings.ReplaceAll(tmpl, "<Brand />", branding)
	tmpl = strings.ReplaceAll(tmpl, "<Theme />", string(themeInfo))

	fmt.Println("==> generating index.html length=", len(tmpl))

	return tmpl, nil
}

func beautify(key string) string {
	return strings.Title(strings.ReplaceAll(key, "_", " "))
}