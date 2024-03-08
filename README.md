# render-alt-delete

render-alt-delete is a TUI tool to quickly delete a large number of Render services.

```bash
$ go install github.com/scottnuma/render-alt-delete
$ RAD_RENDER_API_TOKEN=rnd_XXX render-alt-delete
```

# Config

## Endpoint

The Render API endpoint defaults to api.render.com.

```bash
$ RAD_RENDER_API_ENDPOINT=api.render.com render-alt-delete
```

## API Token

```bash
$ RAD_RENDER_API_TOKEN=rnd_XXX render-alt-delete
```

## Config File

The config file is read at `$HOME/.config/render-alt-delete/config.yaml`.


### Simple Config File
```
render_api_token: rnd_XXX
render_api_endpoint: api.render.com
```

### Multi-Profile Config File
```
# profile sets the default profile to use
profile: work

profiles:
    personal:
        render_api_token: rnd_XXX
    work:
        render_api_token: rnd_XXX
        render_api_endpoint: api.render.com
```

Override your set profile with `RAD_PROFILE`.
```bash
$ RAD_PROFILE=personal render-alt-delete
```

