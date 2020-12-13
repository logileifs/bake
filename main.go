package main

import "os"
import "fmt"
import "bytes"
import "strings"
import "os/exec"
import "io/ioutil"
import "text/template"
import "path/filepath"
import "gopkg.in/yaml.v3"
import "github.com/thatisuday/clapper"

func exec_command(command string) {
	//fmt.Println("exec_command: ", command)
	var args = strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func do_recipe(recipe interface{}) {
	//fmt.Println("do_recipe ", recipe)
	var steps = recipe.(map[string]interface{})["steps"]
	//fmt.Print("steps: ")
	//fmt.Println(steps)
	var vars = recipe.(map[string]interface{})["vars"]
	//fmt.Println("vars: ", vars)
	for _, step := range steps.([]interface{}) {
		buf := &bytes.Buffer{}
		tmpl, err := template.New("").Parse(step.(string))
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(buf, vars)
		if err != nil {
			panic(err)
		}
		step_formatted := buf.String()
		exec_command(step_formatted)
	}
}

func start_watch() {
	fmt.Println("starting watch")
}

func main() {
	// create a new registry
	registry := clapper.NewRegistry()

	// register the root command
	rootCommand, _ := registry.Register("") // fake - never used
	rootCommand.AddArg("recipe_name", "")
	rootCommand.AddFlag("file", "f", false, "./recipes.yml")
	rootCommand.AddFlag("watch", "w", true, "false")

	// parse command-line arguments
	command, err := registry.Parse(os.Args[1:])

	// check for error
	if err != nil {
		fmt.Printf("error => %#v\n", err)
		return
	}

	recipe_name := command.Args["recipe_name"].Value
	file := command.Flags["file"].Value
	if file == "" {
		file = "./recipes.yml"
	}
	watch_arg := command.Flags["watch"].Value
	var watch = false
	if watch_arg == "" {
		watch = false
	} else {
		watch = true
	}
	//fmt.Println("recipe: ", recipe_name)
	//fmt.Println("file: ", file)
	//fmt.Println("watch: ", watch)
	if watch == true {

	}

	path, _ := filepath.Abs(file)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}
	var v interface{}
	err = yaml.Unmarshal(data, &v)
	//fmt.Print("yaml: ")
	//fmt.Println(v)
	var recipes = v.(map[string]interface{})["recipes"]
	//fmt.Print("recipes: ")
	//fmt.Println(recipes)
	var recipe = recipes.(map[string]interface{})[recipe_name]
	do_recipe(recipe)
}
