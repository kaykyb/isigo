package user

import (
	"fmt"
	"isigo/lang"
	"os"
	"os/exec"
)

func Build(prog lang.Program) {
	code, err := prog.Output()

	if err != nil {
		panic(err)
	}

	outDirPath := "./.isi_output"
	outModuleFilePath := "./.isi_output/go.mod"
	outFilePath := "./.isi_output/main.go"

	maybeDeleteOutDir(outDirPath)
	createOutDir(outDirPath)
	writeArtifact(outModuleFilePath, "module isigoprogram\n\ngo 1.22\n")
	writeArtifact(outFilePath, code)

	buildOutProgram(outDirPath)

	fmt.Println("Programa compilado com sucesso.")
}

func maybeDeleteOutDir(outDirPath string) {
	if _, err := os.Stat(outDirPath); !os.IsNotExist(err) {
		err := os.RemoveAll(outDirPath)
		if err != nil {
			fmt.Println("Erro ao remover o diretório:", err)
			panic("A compilação falhou.")
		}
		fmt.Println("Diretório removido:", outDirPath)
	}
}

func createOutDir(outDirPath string) {
	err := os.Mkdir(outDirPath, 0755)
	if err != nil {
		fmt.Println("Erro criando diretório de saída:", err)
		panic("A compilação falhou.")
	}

	fmt.Println("Diretório de saída criado:", outDirPath)
}

func writeArtifact(filePath string, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro criando artefato:", err)
		stop()
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erro ao escrever no artefato:", err)
		stop()
	}
	fmt.Println("Artefato compilado:", filePath)
}

func buildOutProgram(outputDirPath string) {
	// Define the command and the directory
	cmd := exec.Command("go", "build") // Replace "ls -l" with the command you want to run
	cmd.Dir = outputDirPath            // Set the directory where the command will be run

	// Run the command
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao compilar:", err)
		stop()
	}
}

func stop() {
	panic("A compilação falhou.")
}
