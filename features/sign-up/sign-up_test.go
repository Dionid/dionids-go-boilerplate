package fsignup_test

import (
	"context"
	"testing"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/dbs/maindb"
	"github.com/Dionid/go-boiler/features"
	fsignup "github.com/Dionid/go-boiler/features/sign-up"
	inttests "github.com/Dionid/go-boiler/internal/int-tests"
	"github.com/google/uuid"
)

func TestIntSignUp(t *testing.T) {
	t.Run("SignUp 1", func(t *testing.T) {
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

		_, err = inttests.Seed(ctx,
			testDeps.FeaturesConfig, testDeps.MainDbConnection)
		if err != nil {
			t.Fatal(err)
		}

		featureDeps := &features.Deps{
			Logger:        testDeps.Logger,
			MainDb:        testDeps.MainDbConnection,
			MainDbQueries: testDeps.MainDbQueries,
			Config:        testDeps.FeaturesConfig,
			RmqT:          testDeps.RmqTransport,
		}

		request := &proto.SignUpCallRequest{
			Name: "SignIn",
			Id:   uuid.New().String(),
			Params: &proto.SignUpCallRequest_Params{
				Email:    "new@email.com",
				Password: "1234",
			},
		}

		resp, err := fsignup.SignUp(ctx, featureDeps, request)
		if err != nil {
			t.Fatal(err)
		}

		if resp.Result == nil {
			t.Fatal("result is not ok")
		}

		newUser, err := maindb.SelectUserTableByEmail(ctx, testDeps.MainDbConnection, request.Params.Email)
		if err != nil {
			t.Fatal(err)
		}

		if newUser == nil {
			t.Fatal("new user is nil")
		}

		if newUser.Email != request.Params.Email {
			t.Fatal("new user id is not equal to request email")
		}
	})
}
