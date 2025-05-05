package fsignin_test

import (
	"context"
	"testing"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/features"
	fsignin "github.com/Dionid/go-boiler/features/sign-in"
	inttests "github.com/Dionid/go-boiler/internal/int-tests"
	"github.com/google/uuid"
)

func TestIntSignIn(t *testing.T) {
	t.Run("SignIn 1", func(t *testing.T) {
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
			Logger: testDeps.Logger,
			MainDb: testDeps.MainDbConnection,
			Config: testDeps.FeaturesConfig,
			RmqT:   testDeps.RmqTransport,
		}

		request := &proto.SignInCallRequest{
			Name: "SignIn",
			Id:   uuid.New().String(),
			Params: &proto.SignInCallRequest_Params{
				Email:    seed.User.Email,
				Password: "1234",
			},
		}

		resp, err := fsignin.SignIn(ctx, featureDeps, request)
		if err != nil {
			t.Fatal(err)
		}

		if resp.Result == nil {
			t.Fatal("result is not ok")
		}
	})
}
