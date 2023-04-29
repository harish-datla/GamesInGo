package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
)


func loadMaze(fp *os.File) []string{
	scanner := bufio.NewScanner(fp)
	var maze []string
	for scanner.Scan(){
		newLine := scanner.Text()
		maze = append(maze,newLine)
	}
	return maze
}

func loadMazeFromTxtFile() ([]string,error) {
	mazeTxtFilePath := "maze_ascii.txt"

	//openFile
	fp,err := os.Open(mazeTxtFilePath)
	if err != nil {
		return nil,err
	}
	defer fp.Close()

	//createMaze and return it
	return loadMaze(fp)
}


func printScreen(maze){
	for _,line := range maze{
		fmt.Println(line)
	}
}

func main(){
	// Standard steps in game, more or less

	// 1 - Initialize game
 
	// 2 - Load and render graphics
	maze,err := loadMaze()
	if err != nil{
		log.Printf("Error rendering graphics: %v", err)
	}

	// 3 - Game loop
	for{
		
		printScreen(maze)


		break
	}

}