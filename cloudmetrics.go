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
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/weareyolo/go-metrics"
)

// Publisher handles the publication of metrics data to CloudWatch
type Publisher interface {
	Publish()
}

// DatumBuilder handles the datum generation per metric type
type DatumBuilder interface {
	BuildCounterData(v metrics.Counter, name string) []*cloudwatch.MetricDatum
	BuildGaugeData(v metrics.Gauge, name string) []*cloudwatch.MetricDatum
	BuildGaugeFloat64Data(v metrics.GaugeFloat64, name string) []*cloudwatch.MetricDatum
	BuildMeterData(v metrics.Meter, name string) []*cloudwatch.MetricDatum
	BuildHistogramData(v metrics.Histogram, name string) []*cloudwatch.MetricDatum
	BuildTimerData(v metrics.Timer, name string) []*cloudwatch.MetricDatum
}

// CloudWatch is an interface for *cloudwatch.CloudWatch that clearly identifies the functions
// used by cloudmetrics
type CloudWatch interface {
	PutMetricData(input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
}
