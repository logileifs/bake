# bake
Bake is a modern take on Make

# install
## macos
`curl -L https://github.com/logileifs/bake/releases/download/0.1.0/bake-darwin-amd64 > ~/.local/bin/bake`  
`chmod +x ~/.local/bin/bake`  
## linux
`curl -L https://github.com/logileifs/bake/releases/download/0.1.0/bake-linux-amd64 > ~/.local/bin/bake`  
`chmod +x ~/.local/bin/bake`  

# usage
`$ cat ./recipes.yml`
```yaml
---
recipes:
  test: &test
    vars:
      BUILD_ID: shell(openssl rand -hex 3)
      TEST_DIR: ./tests/
    steps:
      - 'echo building new release with id: {{.BUILD_ID}}'
      - echo testing files under {{.TEST_DIR}}
```
## run the recipe **test** from *recipes.yml*
`$ bake test`
produces output:
```
building new release with id: 9e86af
testing files under ./tests/
```

`$ bake -f ./recipes.yml test`
does the same, just explicitly provides path to *recipes.yml*

`bake test FILE=tests/testother.py`
does the same and pass argument FILE=tests/testother.py to the **test** recipe

`bake -w on_save`
start the **on_save** watch defined in *recipes.yml*

`bake -w *`
start all watches defined in *recipes.yml*
