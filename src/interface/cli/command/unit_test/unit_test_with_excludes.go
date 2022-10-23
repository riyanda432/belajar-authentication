package command_unit_test

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// folder name than want to exclude from test
var excludes = []string{
	"mocks",
	"unit_test",
}

func UTExclude() error {
	// go clean -testcache
	cleanTestCache()

	// go list ./... > package_list
	packageList()

	packages := getPackageList()

	if err := runUnitTest(packages); err != nil {
		return err
	}

	// // // mkdir -p cov-report
	// // // go tool cover -html=coverage -o ./cov-report/report.html
	// // // go tool cover -func coverage > test-report
	if err := generateReport(); err != nil {
		return err
	}

	if err := Validate(); err != nil {
		return err
	}

	return nil
}

func generateReport() error {
	os.MkdirAll("./cov-report", 0700)
	cmd := exec.Command("go", "tool", "cover", "-html=coverage", "-o", "./cov-report/report.html")
	if err := cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()

	outfile, err := os.Create("./test-report")
	if err != nil {
		return err
	}
	defer outfile.Close()
	cmd = exec.Command("go", "tool", "cover", "-func", "coverage")
	cmd.Stdout = outfile
	if err := cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}

func cleanTestCache() {
	cmd := exec.Command("go", "clean", "-testcache")
	cmd.Wait()
}

func packageList() error {
	outfile, _ := os.Create("./package_list")

	defer outfile.Close()
	cmd := exec.Command("go", "list", "./...")
	cmd.Stdout = outfile

	cmd.Start()
	cmd.Wait()
	return nil
}

func getPackageList() []string {
	file, _ := osOpen("package_list")

	var packages []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ignore := false
		for _, exclude := range excludes {
			if strings.Contains(scanner.Text(), exclude) {
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}
		packages = append(packages, scanner.Text())
	}
	return packages
}

func runUnitTest(packages []string) error {
	args := []string{"test"}
	args = append(args, "-coverpkg="+`"`+strings.Join(packages, ",")+`"`)
	args = append(args, packages...)
	args = append(args, "-coverprofile=coverage")
	args = append(args, "./...")

	cmd := exec.Command("go", args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println("Please run 'go test ./...' to validate")
		return err
	}

	return nil
}
