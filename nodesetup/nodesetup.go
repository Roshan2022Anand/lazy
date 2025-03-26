package nodesetup

import (
	"os"
	"os/exec"
	"sync"
)

// function to create node setup
func InitNode(wg *sync.WaitGroup) {
	defer wg.Done()
	initCmd := exec.Command("npm", "init", "-y")
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr
	initCmd.Run()
}

// function to create folder stucture
func CreateFolderStruct(wg *sync.WaitGroup) {
	defer wg.Done()

	//create folders
	os.Mkdir("src", 0755)
	os.Mkdir("src/controllers", 0755)
	os.Mkdir("src/models", 0755)
	os.Mkdir("src/routes", 0755)
	os.Mkdir("src/config", 0755)
	os.Mkdir("src/middlewares", 0755)
	os.Mkdir("src/utils", 0755)
	os.Mkdir("src/public", 0755)

	//create files
	os.Create(".env")
	os.Create(".gitignore")

}

// function to create simple node server code with http
func CtreateNodeServer(wg *sync.WaitGroup) {
	defer wg.Done()

	//basic server code
	serverCode := `
	const http = require('http');
	const port = process.env.PORT || 3000;
	
	const server = http.createServer((req, res) => {
	  res.statusCode = 200;
	  res.setHeader('Content-Type', 'text/plain');
	  res.end('Hello World');
	  });
	
	server.listen(port, () => {
	  console.log('Server is running on port ' + port);
	  });
	  `

	os.Create("src/server.js")
	os.WriteFile("src/server.js", []byte(serverCode), 0644)
}
