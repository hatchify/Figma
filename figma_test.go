package figma

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	testFigma *Figma
)

func TestFigma_New(t *testing.T) {
	var (
		accessToken = os.Getenv("FIGMA_ACCESS_TOKEN")
		teamID      = os.Getenv("FIGMA_TEAM_ID")

		err error
	)

	if _, err = New(accessToken, teamID); err != nil {
		t.Fatal(err)
	}
}

func TestFigma_NewFromFile(t *testing.T) {
	var (
		fpath = os.Getenv("FILE_PATH")

		err error
	)

	if _, err = NewFromFile(fpath); err != nil {
		t.Fatal(err)
	}
}

func TestFigma_GetFile(t *testing.T) {
	var (
		teamID      = os.Getenv("FIGMA_TEAM_ID")
		accessToken = os.Getenv("FIGMA_ACCESS_TOKEN")
		fileKey     = os.Getenv("FIGMA_FILE_KEY")

		figma *Figma
		file  *EntryFile
		err   error
	)

	if figma, err = New(accessToken, teamID); err != nil {
		t.Fatal(err)
	}

	if file, err = figma.GetFile(fileKey); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("File: %+v\n%+v\n\n", file, file.Document)
}

func TestFigma_GetNodes(t *testing.T) {
	var (
		teamID      = os.Getenv("FIGMA_TEAM_ID")
		accessToken = os.Getenv("FIGMA_ACCESS_TOKEN")
		fileKey     = os.Getenv("FIGMA_FILE_KEY")
		nodeID1     = os.Getenv("FIGMA_NODE_ID1")
		nodeID2     = os.Getenv("FIGMA_NODE_ID2")

		nodes   *EntryNodes
		nodeIDs []string
		figma   *Figma
		err     error
	)

	nodeIDs = append(nodeIDs, nodeID1, nodeID2)

	if figma, err = New(accessToken, teamID); err != nil {
		t.Fatal(err)
	}

	if nodes, err = figma.GetNodes(fileKey, nodeIDs); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Nodes: %+v\n\n", nodes)
}

func TestFigma_GetStylesByTeamID(t *testing.T) {
	var (
		teamID      = os.Getenv("FIGMA_TEAM_ID")
		accessToken = os.Getenv("FIGMA_ACCESS_TOKEN")

		figma *Figma

		colors    []*TeamColor
		gradients []*TeamGradient
		texts     []*TeamText
		effects   []*TeamEffect

		err error
	)

	if figma, err = New(accessToken, teamID); err != nil {
		t.Fatal(err)
	}

	if colors, gradients, texts, effects, err = figma.GetStylesByTeamID(teamID); err != nil {
		t.Fatal(err)
	}

	for _, color := range colors {
		fmt.Printf("%+v\n%+v\n\n", color, color.Color)
	}

	for _, gradient := range gradients {
		fmt.Printf("%+v\n%+v\n%+v\n", gradient, gradient.GradientHandlePositions, gradient.GradientStops)
		for _, stop := range *gradient.GradientStops {
			color := stop.Color
			fmt.Printf("%+v\n", *color)
		}
		fmt.Print("\n\n")
	}

	for _, text := range texts {
		fmt.Printf("%+v\n%+v\n\n", text, *text.Style)
	}

	for _, effect := range effects {
		fmt.Printf("%+v\n%+v\n%+v%+v\n\n", effect, effect.Effect, effect.Effect.Color, effect.Effect.Offset)
	}
}

func ExampleNew() {
	var (
		accessToken = os.Getenv("FIGMA_ACCESS_TOKEN")
		teamID      = os.Getenv("FIGMA_TEAM_ID")

		err error
	)

	if _, err = New(accessToken, teamID); err != nil {
		log.Fatal(err)
	}
}

func ExampleNewFromFile() {
	var (
		fpath = os.Getenv("FILE_PATH")

		err error
	)

	if _, err = NewFromFile(fpath); err != nil {
		log.Fatal(err)
	}
}

func ExampleGetFile(t *testing.T) {
	var (
		fileKey = os.Getenv("FIGMA_FILE_KEY")
	)

	file, err := testFigma.GetFile(fileKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("File: %+v\n%+v\n\n", file, file.Document)
}

func ExampleGetNodes() {
	var (
		fileKey = os.Getenv("FIGMA_FILE_KEY")
		nodeID1 = os.Getenv("FIGMA_NODE_ID1")
		nodeID2 = os.Getenv("FIGMA_NODE_ID2")

		nodeIDs []string
	)

	nodeIDs = append(nodeIDs, nodeID1, nodeID2)

	if _, err := testFigma.GetNodes(fileKey, nodeIDs); err != nil {
		log.Fatal(err)
	}
}

func ExampleGetStylesByTeamID() {
	var (
		teamID = os.Getenv("FIGMA_TEAM_ID")
	)
	colors, gradients, texts, effects, err := testFigma.GetStylesByTeamID(teamID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Colors: %+v\n\nGradients: %+v\n\nTexts: %+v\n\nEffects: %+v", colors, gradients, texts, effects)
}
