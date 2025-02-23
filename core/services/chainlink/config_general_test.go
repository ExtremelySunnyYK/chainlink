//go:build !dev

package chainlink

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestTOMLGeneralConfig_Defaults(t *testing.T) {
	config, err := GeneralConfigOpts{}.New(logger.TestLogger(t))
	require.NoError(t, err)
	assert.Equal(t, (*url.URL)(nil), config.BridgeResponseURL())
	assert.Nil(t, config.DefaultChainID())
	assert.False(t, config.EVMRPCEnabled())
	assert.False(t, config.EVMEnabled())
	assert.False(t, config.CosmosEnabled())
	assert.False(t, config.SolanaEnabled())
	assert.False(t, config.StarkNetEnabled())
	assert.Equal(t, false, config.FeatureExternalInitiators())
	assert.Equal(t, 15*time.Minute, config.SessionTimeout().Duration())
}

func TestTOMLGeneralConfig_InsecureConfig(t *testing.T) {
	t.Parallel()

	t.Run("all insecure configs are false by default", func(t *testing.T) {
		config, err := GeneralConfigOpts{}.New(logger.TestLogger(t))
		require.NoError(t, err)

		assert.False(t, config.DevWebServer())
		assert.False(t, config.DisableRateLimiting())
		assert.False(t, config.InfiniteDepthQueries())
		assert.False(t, config.OCRDevelopmentMode())
	})

	t.Run("insecure config ignore override on non-dev builds", func(t *testing.T) {
		config, err := GeneralConfigOpts{
			OverrideFn: func(c *Config, s *Secrets) {
				*c.Insecure.DevWebServer = true
				*c.Insecure.DisableRateLimiting = true
				*c.Insecure.InfiniteDepthQueries = true
				*c.Insecure.OCRDevelopmentMode = true
				*c.AuditLogger.Enabled = true
			}}.New(logger.TestLogger(t))
		require.NoError(t, err)

		// Just asserting that override logic work on a safe config
		assert.True(t, config.AuditLoggerEnabled())

		assert.False(t, config.DevWebServer())
		assert.False(t, config.DisableRateLimiting())
		assert.False(t, config.InfiniteDepthQueries())
		assert.False(t, config.OCRDevelopmentMode())
	})

	t.Run("ValidateConfig fails if insecure config is set on non-dev builds", func(t *testing.T) {
		opts := GeneralConfigOpts{}
		err := opts.ParseConfig(`
		  [insecure]
		  DevWebServer = true
		  DisableRateLimiting = false
		  InfiniteDepthQueries = false
		  OCRDevelopmentMode = false
		`)
		require.NoError(t, err)
		cfg, err := opts.init()
		require.NoError(t, err)
		err = cfg.Validate()
		require.Contains(t, err.Error(), "invalid configuration: Insecure.DevWebServer: invalid value (true): insecure configs are not allowed on secure builds")
	})
}
