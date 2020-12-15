# source this file to enable bash completion of recipes
# bake b<tab> => bake build
_bake_complete()
{
	COMPREPLY=($(compgen -W "$(bake -c)" -- "${COMP_WORDS[1]}"))
}

complete -F _bake_complete bake
