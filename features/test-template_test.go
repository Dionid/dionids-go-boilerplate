package features_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Dionid/go-boiler/features"
	inttests "github.com/Dionid/go-boiler/int-tests"
)

func TestIntTemplate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testDeps, err := inttests.InitTestDeps(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err := testDeps.Cleanup()
		if err != nil {
			t.Fatal(err)
		}
	})

	seed, err := inttests.Seed(ctx,
		testDeps.FeaturesConfig, testDeps.MainDbConnection)
	if err != nil {
		t.Fatal(err)
	}

	featureDeps := &features.Deps{
		testDeps.Logger,
		testDeps.MainDbConnection,
		testDeps.MainDbQueries,
		testDeps.FeaturesConfig,
		testDeps.RmqTransport,
	}

	// # WRITE YOUR TEST HERE
	// ...

	// # DELETE THIS LINE
	fmt.Printf("%v, %v\n", seed, featureDeps)
}
