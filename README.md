# gha-trigger-set-env-action

GitHub Actions to set Environment variables for [gha-trigger](https://github.com/suzuki-shunsuke/gha-trigger)

## Example

```yaml
on:
  workflow_dispatch:
    inputs:
      data:
        required: true
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: suzuki-shunsuke/gha-trigger-set-env-action@main
        with:
          data: ${{inputs.data}}
```

## Inputs

- `data`: gha-trigger's workflow_dispatch event input `data`

## Outputs

Nothing.

## Set Environment variables

name | example | description
--- | --- | ---
GHA_ACTOR | | GITHUB_ACTOR
GHA_BASE_REF | | GITHUB_BASE_REF
GHA_EVENT_NAME | | GITHUB_EVENT_NAME
GHA_HEAD_REF | | GITHUB_HEAD_REF
GHA_HEAD_SHA | | 
GHA_REF | | GITHUB_REF
GHA_REF_NAME | | GITHUB_REF_NAME
GHA_REPOSITORY | | GITHUB_REPOSITORY
GHA_REPOSITORY_OWNER | | GITHUB_REPOSITORY_OWNER
GHA_SHA | | GITHUB_SHA

## LICENSE

[MIT](LICENSE)
