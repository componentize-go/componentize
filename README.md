# :rocket: Start componentizing your templates in Go

## Getting started:

 1. Add the library to your Go modules with the command
    ```bash
    $ go get github.com/componentize-go/componentize
    ```
 3. Call the `Default()` function to get a `html/template.FuncMap` that you can pass in to the `Funcs` method of a `html/template.Template` struct before parsing the template with the `Parse` method.
 4. Now you have access to the functions `KVM`, `UID` and `Array` inside of your templates. See the [examples](/examples) for more information about how to use these functions.

## Docs:

 - ## Note:
  All of the functions are only intended to use inside of a Go template, note that the functions are still exported but this is only for people that wanted to use my API in a different way. The recommended way is using either the `Default()` or `WithConfig(config Config)` constructors

 - ### `KVM`:

  **Signature:**
  ```go
  func KVM(key string, value any, otherMap ...map[string]any) (map[string]any, error)
  ```

  Template function very useful to use when you don't know how big will be your data (and in consequence your bind struct) to pass in to your template. So this function solve this problem by providing a mechanism to create Go maps "On The Way" inside of your templates so you can pass the maps inside of your components. You don't have worry anymore about giving arbitrary names to your bind structs and add more and more fields to it.

  **Quick usage:**
  ```go
  package main

  import (
      "html/template"
      "os"

      "github.com/componentize-go/componentize"
  )

  func main() {
      // Note that the name of the Key of the map must be matching with the "K" param of the `KVM` function
      myComponent := `<div>{{.content}}</div>`

      myIndexTmpl := `
          <div>
            {{template "my-component" KVM "content" "Hi I'm a reusable component"}}
            {{template "my-component" KVM "content" "Hi I'm another reusable component"}}
            {{template "my-component" KVM "content" "Hi I'm BOB"}}
          </div>
      `

      componentizeFuncs := componentize.Default()

      // it's as easy to integrate with your current project as just passing the FuncMap returned
      // by the constructors to the Funcs method and voilá you have componentize at your service
      tmpl, _ := template.New("index").Funcs(componentizeFuncs).Parse(myIndexTemplate)

      tmpl, _ = tmpl.New("my-component").Parse(myComponent)

      _ = tmpl.ExecuteTemplate(os.Stdout, "index", nil)
  }
  ```

 - ### `UID`:
  
  **Signature:**
  ```go
  func UID() (string, error)
  ```

  Template function that provides a way of generate Unique ID's. You may want to use this function when you want to add Javascript to your template, and you don't want think about an ID that is not repeated in the template

  **Quick usage:**
  ```go
  package main

  import (
      "html/template"
      "os"

      "github.com/componentize-go/componentize"
  )

  func main() {
      myTmpl := `
          {{$divId := UID}}
          {{$btnId := UID}}

          <div id="{{$divId}}">Hi I'm a div</div>

          <button id="{{$btnId}}" type="button">Change the content of the div</button>

          <script>
            let myBtn = document.querySelector("#{{$btnId}}")
            
            let myDiv = document.querySelector("#{{$divId}}")

            myBtn.addEventListener("click", function() {   
                myDiv.innerHTML = "My content has been changed by Javascript and Go :O"
            })
          </script>
      `

      // You must display this in a browser if you want to see if it works but I hope that the idea is understanded

      componentizeFuncs := componentize.Default()

      tmpl, _ := template.New("index").Funcs(componentizeFuncs).Parse(myTmpl)

      _ = tmpl.ExecuteTemplate(os.Stdout, "index", nil)
  }
  ```
 - ## `Array`:

   **Signature:**
   ```go
   func Array(args ...any) []any
   ```

   Template function that helps you create arrays inside of your Go templates

   **Quick Usage:**
   ```go
   package main

   import (
       "html/template"
       "os"

       "github.com/componentize-go/componentize"
   )

   func main() {
       myComponent := `
           <div class="user-info">{{.userName}} is {{.age}} years old</div>
       `
   
       myTmpl := `
           <div class="users-container">
             {{$users := Array (KVM "userName" "John47" | KVM "age" 47)
                               (KVM "userName" "Mary26" | KVM "age" 26)
                               (KVM "userName" "PeterParker22" | KVM "age" 22)
             }}
   
             {{range $user := $users}}
               {{template "user-info-component" $user}}
             {{end}}
           </div>
       `

       componentizeFuncs := componentize.Default()

       tmpl, _ := template.New("index").Funcs(componentizeFuncs).Parse(myTmpl)
   
       tmpl, _ := tmpl.New("user-info-component").Parse(myComponent)

       tmpl.ExecuteTemplate(os.Stdout, "index", nil)
   }
   ```
