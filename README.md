# render-alt-delete

render-alt-delete is a TUI tool to quickly delete a large number of Render services.

```bash
$ go install github.com/scottnuma/render-alt-delete
$ RAD_RENDER_API_TOKEN=rnd_XXX render-alt-delete
```

# Config

The Render API endpoint defaults to api.render.com.

## Environment Variables

```bash
$ RAD_RENDER_API_TOKEN=rnd_XXX render-alt-delete

$ RAD_RENDER_API_TOKEN=rnd_XXX RAD_RENDER_API_ENDPOINT=api.render.com render-alt-delete
```

## Config File

The config file is read at `$HOME/.config/render-alt-delete/config.yaml`.

Sample Config files

```
render_api_token: rnd_XXX
render_api_endpoint: api.render.com
```

```
# profile sets the default profile to use
# this can be overriden by setting the RAD_PROFILE env var
profile: work

profiles:
    personal:
        render_api_token: rnd_XXX
    work:
        render_api_token: rnd_XXX
        render_api_endpoint: api.render.com
```

