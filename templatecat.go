package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"strings"
	"io/ioutil"
	"text/template"
)

func ReadPath(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(file)
	return string(b), err
}

func main () {
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr, "TemplateCat expects a template from STDIN and an argument list of path@variable_name references. Output is written to STDOUT.\n")
		fmt.Fprintf(os.Stderr, "Example: templatecat a@first_file b@second_file < template\n")
		fmt.Fprintf(os.Stderr, "All instances of 'a' (as {{ .a }} Golang template style) within the template will be replaced with the content of first_file, and likewise for 'b' and second_file.\n")
	}

	flag.Parse()
	values := make(map[string]string)

	var err error
	for _, varname := range flag.Args() {
		pair := strings.Split(varname, "@")

		if len(pair) != 2 {
			log.Fatalf("Could not parse varname: %s", varname)
		}

		values[pair[0]], err = ReadPath(pair[1])

		if err != nil {
			log.Fatal(err)
		}
	}

	b, _ := ioutil.ReadAll(os.Stdin)
	base := template.New("base")
	t, err := base.Parse(string(b))

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(os.Stdout, values)
}
