package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/util"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, service database.Service, gwService *client.Gateway) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	ginServer, err := NewGinServer(config, service, gwService)

	require.NoError(t, err)

	return ginServer
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
