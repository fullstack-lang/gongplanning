package main

import (
	"flag"
	"log"
	"strconv"

	gongplanning_models "github.com/fullstack-lang/gongplanning/go/models"
	gongplanning_stack "github.com/fullstack-lang/gongplanning/go/stack"
	gongplanning_static "github.com/fullstack-lang/gongplanning/go/static"

	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"

	gongsvg_fullstack "github.com/fullstack-lang/gongsvg/go/fullstack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gongplanning: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongplanning_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := gongplanning_stack.NewStack(r, "gongplanning", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// create stacks for diagrammer
	gongsvgStage, _ := gongsvg_fullstack.NewStackInstance(r, gongplanning_models.StackNameDefault.ToString())
	gongtreeStage, _ := gongtree_fullstack.NewStackInstance(r, gongplanning_models.StackNameDefault.ToString())
	_ = gongsvgStage
	_ = gongtreeStage

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
