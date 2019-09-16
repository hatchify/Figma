# Figma 
Figma is an SDK for Figma design software

## Usage

### New
```go
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
```

### NewFromFile
```go
func ExampleNewFromFile() {
	var (
		fpath = os.Getenv("FILE_PATH")

		err error
	)

	if _, err = NewFromFile(fpath); err != nil {
		log.Fatal(err)
	}
}
```
### GetFile 
``` go
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
```
### GetNodes
```go
func ExampleGetNodes() {
	var (
		fileKey = os.Getenv("FIGMA_FILE_KEY")
		nodeID1 = os.Getenv("FIGMA_NODE_ID1")
		nodeID2 = os.Getenv("FIGMA_NODE_ID2")

		nodeIDs []string
		err     error
	)

	nodeIDs = append(nodeIDs, nodeID1, nodeID2)

	if _, err = testFigma.GetNodes(fileKey, nodeIDs); err != nil {
		log.Fatal(err)
	}
}
```

## GetStylesByTeamID
```go
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
```
