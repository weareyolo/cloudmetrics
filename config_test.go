package cloudmetrics

//	Copyright 2020 @weareyolo
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/gojuno/minimock/v3"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weareyolo/cloudmetrics/mock"
)

func TestGetSettings(t *testing.T) {
	t.Run("OK - No options", func(t *testing.T) {
		s := getSettings(nil)
		require.NotNil(t, s)

		assert.Equal(t, &settings{
			Context:           context.Background(),
			Interval:          1 * time.Minute,
			Units:             map[string]string{},
			Dimensions:        map[string]string{},
			Percentiles:       []float64{.5, .75, .95, .99},
			StorageResolution: 60,
		}, s)
	})

	t.Run("OK - With Builder", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		b := mock.NewDatumBuilderMock(mc)

		s := getSettings([]Option{
			WithBuilder(b),
		})
		require.NotNil(t, s)

		assert.Equal(t, &settings{
			Context:           context.Background(),
			Interval:          1 * time.Minute,
			Units:             map[string]string{},
			Dimensions:        map[string]string{},
			Percentiles:       []float64{.5, .75, .95, .99},
			DatumBuilder:      b,
			StorageResolution: 60,
		}, s)
	})

	t.Run("OK - With all config except builder", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		logger, _ := test.NewNullLogger()
		cw := mock.NewCloudWatchMock(mc)

		units := map[string]string{
			"metric": cloudwatch.StandardUnitTerabits,
		}
		dimensions := map[string]string{"k": "v"}
		percentiles := []float64{.2}

		s := getSettings([]Option{
			WithContext(ctx),
			WithClient(cw),
			WithInterval(6 * time.Millisecond),
			WithLogger(logger),
			WithUnits(units),
			WithDimensions(dimensions),
			WithPercentiles(percentiles),
			WithStorageResolution(30),
		})
		require.NotNil(t, s)

		assert.Equal(t, &settings{
			Context:           ctx,
			Client:            cw,
			Interval:          6 * time.Millisecond,
			Logger:            logger,
			Units:             units,
			Dimensions:        dimensions,
			Percentiles:       percentiles,
			StorageResolution: 30,
		}, s)
	})
}
