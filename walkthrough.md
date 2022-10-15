# My first Go web app

## Initialise project

- Create _go.mod_ to use modules, run: `go mod init gowebapp`
- Create _main.go_, make sure it has `package main` and `func main() {}`
- Import `net/http` to enable web app functionalities
- Inside `main` function, write the first handler function and listen to port `8080`

## Add HTML templates

- Create _templates_ folder, and create template file _home.page.tmpl_
- Write normal HTML file inside the template file, then parse it with `template.ParseFiles()`
- render the parsed template with `parsedTemplate.Execute()`

## Use multiple packages

- Create empty folders _cmd/web_, _pkg/handlers_ and _pkg/render_
- Move _main.go_ into _cmd/web_, move _handlers.go_ into _pkg/handlers_, move _render.go_ into _pkg/render_
- In _render.go_, update first line to `package render`, rename the `renderTemplate` with capital camel case: `RenderTemplate`
- In _handlers.go_, update first line to `package handlers`, import `gowebapp/pkg/render`, replace the `renderTemplate` with `render.RenderTemplate`
- In _main.go_, import `gowebapp/pkg/handlers`, replace `Home` with `handlers.Home`, replace `About` with `handlers.About`
- To run the refactored app, run `go run cmd/web/*.go`

## Use third party package

- Use **"pat"** to handle routes, run `go get github.com/bmizerany/pat`
- This will create _go.sum_ and update _go.mod_
