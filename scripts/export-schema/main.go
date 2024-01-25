package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Current pwd
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	envFile := ".maindb.env"

	exportTo := fmt.Sprintf("%s/../../dbs/maindb", pwd)

	cmd := exec.Command(
		"docker",
		"run",
		fmt.Sprintf("--env-file=%s/%s", pwd, envFile),
		// fmt.Sprintf("-v%s:/export", exportTo),
		"--rm",
		"postgres",
		"pg_dump",
		"--schema-only",
		"--schema=public",
		"-x",
		"--disable-triggers",
		"--no-table-access-method",
		"--no-owner",
		"--no-publications",
		"--no-comments",
		"--no-security-labels",
		"--no-toast-compression",
		// "-f/export/schema.sql",
	)

	res, err := cmd.CombinedOutput()

	os.WriteFile(fmt.Sprintf("%s/schema.sql", exportTo), res, 0644)

	if err != nil {
		log.Fatalln(err)
	}
}
