package testhelper_test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"golang.org/x/sync/errgroup"
	"math/big"
	"testing"
)

// TestHelperSuite defines the basic test suite.
type TestHelperSuite struct {
	*testsuite.TestSuite
	// testBackends contins a list of all test backends
	testBackends  []backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *TestHelperSuite {
	tb.Helper()
	return &TestHelperSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *TestHelperSuite) SetupSuite() {
	s.TestSuite.SetupSuite()

	s.SetupBackends(s.GetSuiteContext())
}

func (s *TestHelperSuite) SetupTest() {
	s.TestSuite.SetupTest()

	s.deployManager = testutil.NewDeployManager(s.T())
	// deploy the contract to all backends
	s.deployManager.BulkDeploy(s.GetTestContext(), s.testBackends, testutil.TestContractType)
}

// SetupBackends sets up the test backends that are used for the tests. These need to be setup as embedded backends since
// scribe requires rpc addresses, so we employ some paraellism to speed up the test process.
//
// This can either be done per suite or per test. This is done per suite do to the cost of spinning up fake geth nodes.
func (s *TestHelperSuite) SetupBackends(ctx context.Context) {
	// let's create 3 mock chains
	chainIDs := []uint64{1, 2, 3}

	// preallocate a slice for testbackends to the size of chainIDs
	// this way we can avoid non-deterministic order + needing to acquire/release a lock
	s.testBackends = make([]backends.SimulatedTestBackend, len(chainIDs))

	// TODO: can we use a waitgroup here instead?
	g, _ := errgroup.WithContext(ctx)
	for i, chainID := range chainIDs {
		pos := i           // get position of chain id in array
		chainID := chainID // capture func literal
		g.Go(func() error {
			// we need to use the embedded backend here, because the simulated backend doesn't support rpcs required by scribe
			backend := geth.NewEmbeddedBackendForChainID(ctx, s.T(), new(big.Int).SetUint64(chainID))

			// add the backend to the list of backends
			s.testBackends[pos] = backend
			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		s.T().Fatal(err)
	}
}

func TestTestHelperSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
