# Extended Features

This part of documentation holds extended library features.

## Server Side Actions

Server Side Actions (SSA) very similar to component methods in traditional frontend frameworks.
Main difference - all actions are executed on server side, code placed only on server and client has only thin communication layer.
Frontend only recieves ready for use HTML markup.  

### SSA Installation

For using SSA, you need to define and register SSA handler, and include communication layer on target page.  

For creating SSA handler, you can use high-level SSA handler factory:

```go
func ssahandler() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ssc.SSAHandlerFactory(ssatemplate, map[string]interface{}{
            "internal:rw": rw,
            "internal:r":  r,
        })(rw, r)
    }
}
```

After creation of SSA handler, you need to register it under `/SSA/` route.

```go
...

mux.HandleFunc("/SSA/", ssahandler())

// In case of default http mux, use this
http.HandleFunc("/SSA/", ssahandler())

...
```

And now, we need to include thin communication layer, implemented with JS, into target page.  
This can be done with `dynamics` template function, provided by `ssc.Funcs()` function (check [Page rendering](/core-features/#page-rendering) for details)

```html
<html>
    <head>
        ...
    </head>
    <body>
        ...
        {{ dynamics }}
    </body>
</html>
```

### SSA Usage

Now you can implement `Actions` method to define own component's methods.  
This method must return `ssc.ActionMap`, map which holds your methods. Each method accepts dynamic arguments amount with `...interface{}`.
In the method you can modify component's state, dynamicaly create and initialize another components, etc.

Usage:

```go
...

func (c *ComponentExample) Actions() ssc.ActionMap {
    return ssc.ActionMap{
        "ExampleAction": func(args ...interface{}) {
            // Do what you want here
        }
    }
}

...
```

Now, let's see how to trigger an action in your component template.  
First, you need to include component attributes into your top-level node with `componentattrs` template function. This function accepts component as argument.
This includes different internal library data and component state. After that, you can use `action` template function to trigger an action. This function accepts multiple arguments: first arguments is always action name, all arguments after that will be passed as `...interface{}` to action arguments.  

> Please note, that you can use `action` template function only in event handlers, like `onclick="..."`.  

Usage:

```html
{{ define "ComponentExample" }}
<div {{ componentattrs . }}>
    <button onclick="{{ action 'ExampleAction' }}">Click Me</button>
</div>
{{ end }}
```

### SSA Lifecycle

SSA has own lifecycle, which is a bit different in comparison with page rendering  

- Creating request on client side with communication layer
- Extracting action data from request on server side
- Finding registered component type
- Creating component instance
- Triggering component's initialization method (if implemented)
- Populating component's state
- Calling action
- If new components where registed while action execution, do asynchronous operations for them (overall async process is the same as for page rendering)
- Rendering component and returning HTML to client side
- Morphing recieved HTML with component, or replacing in case of morph failure or explicit `ssa-replace` attribute

### SSA Notes

- You may have problems on morph stage. It requires correct HTML structure and may cause unexpected behavior in some cases. Use `ssa-replace` attribute in your top-level node to explicitly switch to HTML replacement mode

## Server Side State

::: danger
Not implemented yet. Check [issue](https://github.com/yuriizinets/ssceng/issues/28) state
:::

This feature is useful in case of large state payloads.
Instead of saving state inline as html tag, store state on server side and inject state hash as html tag.
Using this, you will decrease amount of data sent with SSA request and total HTML document size.

## Meta builder

::: warning
Not stable. In active development.
:::

Widget on page, that can be included with inisghts template function. Widget includes general page render information, like errors, overall lifecycle timings, etc. Also, widget includes list of each rendered component with lifecycle timings (init, async, afterasync, etc). On hover, highlights component on page.

## Insights

::: danger
Not implemented yet. Check [issue](https://github.com/yuriizinets/ssceng/issues/26) state
:::

Widget on page, that can be included with inisghts template function. Widget includes general page render information, like errors, overall lifecycle timings, etc. Also, widget includes list of each rendered component with lifecycle timings (init, async, afterasync, etc). On hover, highlights component on page.