# prompter

POC templating engine for prompts

## usage

### cli

```
 go run main.go -t "hello {{var: User}}. Items: {{join: Items, \", \"}}." -v User="My username" -j Items="1,2,3"
```

### code

```
template := `hello {{var: User}}. Items: {{join: Items, ", "}}.`

data := map[string]interface{}{
            "User":    "Some username",
            "Items":   []int{1, 2, 3},
        }

engine := engine.NewEngine()

result, err := engine.Render(template, data)
```
