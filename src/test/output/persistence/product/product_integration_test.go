package product_test

import (
	"log"
	"os"
	"testing"

	db_test "github.com/ramonmpacheco/poc-hexagonal-go-v2/src/test/config/db"
)

// No, it is not to test the main function on your program.
// Basically the TestMain function provides more control over
// running tests than was available in the prior releases of Go.
// ##
// This function will be called instead of running the tests directly
// The M struct contains methods to access and run the tests.
func TestMain(m *testing.M) {
	log.Default().Println("starting product integration tests")
	db_test.Connect()
	exitVal := m.Run()
	db_test.Close()
	os.Exit(exitVal)
	log.Default().Println("finishing product integration tests")
}
