package integration_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

// Global or suite-level variables
var (
	cmd              *exec.Cmd // The external process for the recipe-manager
	lastResponse     *http.Response
	lastResponseBody []byte
	port             int
)

// Build & run the server (binary) once for the entire test suite.
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		port = 9090
		cmd = exec.Command("./run_server.sh")
		if err := cmd.Start(); err != nil {
			panic(fmt.Sprintf("failed to start recipe-manager binary: %v", err))
		}

		// 4. Wait for the server to become ready (simple approach: retry a few times)
		if err := waitForServer(port, 5*time.Second); err != nil {
			// If it can't start in time, kill the process
			_ = cmd.Process.Kill()
			panic(fmt.Sprintf("server did not become ready: %v", err))
		}
	})

	ctx.AfterSuite(func() {
		// Stop the binary after tests
		if cmd != nil && cmd.Process != nil {
			_ = cmd.Process.Kill() // or cmd.Process.Signal(os.Interrupt)
		}
	})
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^the application is running on port (\d+)$`, theApplicationIsRunningOnPort)
	s.Step(`^I send a GET request to "([^"]*)"$`, iSendAGETRequestTo)
	s.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	s.Step(`^the response should contain "([^"]*)"$`, theResponseShouldContain)
	s.Step(`^the response should not contain "([^"]*)"$`, theResponseShouldNotContain)
	s.Step(`^the server is running$`, theServerIsRunning)
	s.Step(`^the server is running on port (\d+)$`, theServerIsRunningOnPort)
}

// Wait for the server to become ready by pinging a known endpoint
func waitForServer(port int, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	url := fmt.Sprintf("http://localhost:%d/recipes", port)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return nil
		}
		time.Sleep(250 * time.Millisecond)
	}
	return fmt.Errorf("server did not respond within %v", timeout)
}

// Steps

func theApplicationIsRunningOnPort(p int) error {
	// We already started the server with a known port (8080).
	// If you want to do a dynamic port, you'd parse p and run the binary with that port.
	if port != p {
		return fmt.Errorf("expected port %d but actual port is %d", p, port)
	}
	return nil
}

func iSendAGETRequestTo(path string) error {
	url := fmt.Sprintf("http://localhost:%d%s", port, path)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	lastResponse = resp

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	lastResponseBody = body
	return nil
}

func theResponseCodeShouldBe(code int) error {
	if lastResponse == nil {
		return fmt.Errorf("no response recorded")
	}
	if lastResponse.StatusCode != code {
		return fmt.Errorf("expected status code %d, got %d", code, lastResponse.StatusCode)
	}
	return nil
}

func theResponseShouldContain(substring string) error {
	if !bytes.Contains(lastResponseBody, []byte(substring)) {
		return fmt.Errorf("expected body to contain %q but it did not.\nBody: %s",
			substring, string(lastResponseBody))
	}
	return nil
}

func theResponseShouldNotContain(substring string) error {
	if bytes.Contains(lastResponseBody, []byte(substring)) {
		return fmt.Errorf("expected body to not contain %q but it did.\nBody: %s",
			substring, string(lastResponseBody))
	}
	return nil
}

func theServerIsRunning() error {
	return theServerIsRunningOnPort(9090)
}

func theServerIsRunningOnPort(port int) error {
	if err := waitForServer(port, 5*time.Second); err != nil {
		return fmt.Errorf("server did not become ready: %v", err)
	}
	return nil
}

// Godog test runner
func TestGodog(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "api-bdd",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  FeatureContext,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"./features"},
		},
	}

	if suite.Run() != 0 {
		t.Fail()
	}
}
