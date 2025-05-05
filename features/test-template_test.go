package features_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/Dionid/go-boiler/features"
	inttests "github.com/Dionid/go-boiler/internal/int-tests"
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

	// # Global WaitGroup
	gwg := &sync.WaitGroup{}

	// # Graceful shutdown emitter
	gse := make(chan string, 1)

	featureDeps := &features.Deps{
		testDeps.Logger,
		gwg,
		gse,
		testDeps.MainDbConnection,
		testDeps.FeaturesConfig,
	}

	// # WRITE YOUR TEST HERE
	// ...

	// # DELETE THIS LINE
	fmt.Printf("%v, %v\n", seed, featureDeps)
}
