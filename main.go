package main

import "os"
import "fmt"
import "time"
import "bytes"
import "strings"
import "os/exec"
import "io/ioutil"
import "text/template"
import "path/filepath"
import "gopkg.in/yaml.v3"
import "github.com/thatisuday/clapper"

var version string = "development"
var commit string = "nil"
var build string = "nil"
var date string = time.Now().Format("02-01-2006T15:04:05")

func exec_command(command string) {
	//fmt.Println("exec_command: ", command)
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func eval_in_shell(statement string) string {
	//fmt.Println("eval_in_shell: ", statement)
	cmd := exec.Command("sh", "-c", statement)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return strings.Trim(string(out), "\n")
}

func evaluate_vars(vars map[string]interface{}) map[string]interface{} {
	for key, value := range vars {
		var has_prefix = strings.HasPrefix(value.(string), "shell(")
		var has_suffix = strings.HasSuffix(value.(string), ")")
		if has_prefix == true && has_suffix == true {
			//fmt.Println(value, " has prefix and suffix")
			var to_eval = strings.Split(
				strings.Split(value.(string), "shell(")[1],
				")",
			)[0]
			vars[key] = eval_in_shell(to_eval)
		}
	}
	return vars
}

func handle_step(step string, vars map[string]interface{}) string {
	buf := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(step)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(buf, vars)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func do_recipe(recipe interface{}) {
	//fmt.Println("do_recipe ", recipe)
	var steps = recipe.(map[string]interface{})["steps"]
	//fmt.Print("steps: ")
	//fmt.Println(steps)
	var vars = recipe.(map[string]interface{})["vars"]
	//fmt.Println("vars: ", vars)
	for _, step := range steps.([]interface{}) {
		vars = evaluate_vars(vars.(map[string]interface{}))
		step_formatted := handle_step(step.(string), vars.(map[string]interface{}))
		//fmt.Println("step_formatted: ", step_formatted)
		exec_command(step_formatted)
	}
}

func start_watch() {
	fmt.Println("starting watch")
}

func print_recipes_and_exit(recipes interface{}) {
	for key, _ := range recipes.(map[string]interface{}) {
		fmt.Print(key, " ")
	}
	os.Exit(0)
}

func main() {
	// create a new registry
	registry := clapper.NewRegistry()

	// register the root command
	rootCommand, _ := registry.Register("") // fake - never used
	rootCommand.AddArg("recipe_name", "")
	rootCommand.AddFlag("file", "f", false, "./recipes.yml")
	rootCommand.AddFlag("watch", "w", true, "false")
	rootCommand.AddFlag("comp", "c", true, "false")

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

	comp_arg := command.Flags["comp"].Value
	var comp = false
	if comp_arg == "" {
		comp = false
	} else {
		comp = true
	}
	//fmt.Println("recipe: ", recipe_name)
	//fmt.Println("file: ", file)
	//fmt.Println("watch: ", watch)
	if watch == true {
		fmt.Println("watch = true")
	}

	path, _ := filepath.Abs(file)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}


	if (recipe_name != "") || (comp) {
		var v interface{}
		err = yaml.Unmarshal(data, &v)
		//fmt.Print("yaml: ")
		//fmt.Println(v)
		var recipes = v.(map[string]interface{})["recipes"]
		if comp {
			print_recipes_and_exit(recipes)
		}
		var recipe = recipes.(map[string]interface{})[recipe_name]
		do_recipe(recipe)
	} else {
		fmt.Println("Version: ", version)
		fmt.Println("Commit: ", commit)
		fmt.Println("Build id: ", build)
		fmt.Println("Build date: ", date)
	}
}
