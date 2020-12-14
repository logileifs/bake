# bake
Bake is a modern take on Make

# usage
`bake test`
run the recipe **test** from *recipes.yml*

`bake -f ./recipes.yml test`
does the same, just explicitly provides path to *recipes.yml*

`bake test FILE=tests/testother.py`
does the same and pass argument FILE=tests/testother.py to the **test** recipe

`bake -w on_save`
start the **on_save** watch defined in *recipes.yml*

`bake -w *`
start all watches defined in *recipes.yml*