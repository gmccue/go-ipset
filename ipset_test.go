package ipset

import (
	"flag"
	"log"
	"os"
	"testing"
)

var (
	ipset       IPSet
	testSet     = "test"
	testIP      = "192.168.1.100"
	testOutfile = "./ipset.txt"
)

// Perform test setup and teardown functions.
func TestMain(m *testing.M) {
	flag.Parse()

	set, err := New()
	if err != nil {
		log.Fatal("Could not setup the tests. Is ipset installed? Message:", err)
	}

	ipset = *set

	os.Exit(m.Run())
}

func TestNewIPSet(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Error("Could not construct a new set. Error was:", err)
	}
}

func TestCreate(t *testing.T) {
	err := ipset.Create(testSet, "hash:ip", "timeout", "300")
	if err != nil {
		t.Error("Could not create a new set. Error was:", err)
	}
}

func TestAdd(t *testing.T) {
	err := ipset.Add(testSet, testIP)
	if err != nil {
		t.Error("Could not add an entry to the test set. Error was:", err)
	}
}

func TestAddUnique(t *testing.T) {
	err := ipset.AddUnique(testSet, testIP)
	if err != nil {
		t.Error("Could not add a unique entry to the test set. Error was:", err)
	}
}

func TestTest(t *testing.T) {
	err := ipset.Test(testSet, testIP)
	if err != nil {
		t.Error("Error while testing for a set entry. Error was:", err)
	}

	// The following IP should not exist in the set, so we expect an error.
	nonexistErr := ipset.Test(testSet, "192.168.100.1")
	if nonexistErr == nil {
		t.Error("Error while testing for a non-existent set entry. Error was:", err)
	}
}

func TestDelete(t *testing.T) {
	err := ipset.Delete(testSet, testIP)
	if err != nil {
		t.Error("Could not delete an entry from the test set. Error was:", err)
	}
}

func TestSave(t *testing.T) {
	err := ipset.Save(testSet, testOutfile)
	if err != nil {
		t.Error("Could not save the test set. Error was:", err)
	}
}

func TestRestore(t *testing.T) {
	destroyErr := ipset.Destroy(testSet)
	if destroyErr != nil {
		t.Error("Could not destroy the set. Error was:", destroyErr)
	}

	err := ipset.Restore(testOutfile)
	if err != nil {
		t.Error("Could not restore the set from a file. Error was:", err)
	}
}

func TestFlush(t *testing.T) {
	err := ipset.Flush(testSet)
	if err != nil {
		t.Error("Could not flush the test set. Error was:", err)
	}
}

func TestRename(t *testing.T) {
	err := ipset.Rename(testSet, "renamedSet")
	if err != nil {
		t.Error("Could not rename the test set. Error was:", err)
	}
}

func TestSwap(t *testing.T) {
	newsetErr := ipset.Create(testSet, "hash:ip")
	if newsetErr != nil {
		t.Error("Could not create a new set for swapping. Error was:", newsetErr)
	}

	err := ipset.Swap("renamedSet", testSet)
	if err != nil {
		t.Error("Could not swap sets. Error was:", err)
	}
}

func TestDestroy(t *testing.T) {
	destErr := ipset.Destroy("renamedSet")
	if destErr != nil {
		t.Error("Could not destroy the copied set. Error was:", destErr)
	}

	err := ipset.Destroy(testSet)
	if err != nil {
		t.Error("Could not destroy the test set. Error was:", err)
	}
}
