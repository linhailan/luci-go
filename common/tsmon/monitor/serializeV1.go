// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package monitor

import (
	"math"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/luci/luci-go/common/tsmon/distribution"
	"github.com/luci/luci-go/common/tsmon/field"
	"github.com/luci/luci-go/common/tsmon/types"

	pb "github.com/luci/luci-go/common/tsmon/ts_mon_proto_v1"
)

// SerializeCellsV1 creates a MetricsCollection message from a slice of cells.
func SerializeCellsV1(cells []types.Cell) *pb.MetricsCollection {
	collection := pb.MetricsCollection{
		Data: make([]*pb.MetricsData, len(cells)),
	}

	for i, cell := range cells {
		collection.Data[i] = SerializeCellV1(cell)
	}

	return &collection
}

// SerializeCellV1 creates one MetricsData message from a cell.
func SerializeCellV1(c types.Cell) *pb.MetricsData {
	d := pb.MetricsData{}
	d.Name = proto.String(c.Name)
	d.Description = proto.String(c.Description)
	d.MetricNamePrefix = proto.String(metricNamePrefix)
	d.Fields = field.SerializeV1(c.Fields, c.FieldVals)
	if c.ValueType.IsCumulative() {
		d.StartTimestampUs = proto.Uint64(uint64(c.ResetTime.UnixNano() / int64(time.Microsecond)))
	}
	c.Target.PopulateProtoV1(&d)
	if c.Units.IsSpecified() {
		d.Units = c.Units.AsProto()
	}

	SerializeValueV1(c.ValueType, c.Value, &d)
	return &d
}

// SerializeValueV1 writes one metric's value into the MetricsData message.
func SerializeValueV1(typ types.ValueType, value interface{}, d *pb.MetricsData) {
	switch typ {
	case types.NonCumulativeIntType:
		d.Gauge = proto.Int64(value.(int64))
	case types.CumulativeIntType:
		d.Counter = proto.Int64(value.(int64))
	case types.NonCumulativeFloatType:
		d.NoncumulativeDoubleValue = proto.Float64(value.(float64))
	case types.CumulativeFloatType:
		d.CumulativeDoubleValue = proto.Float64(value.(float64))
	case types.StringType:
		d.StringValue = proto.String(value.(string))
	case types.BoolType:
		d.BooleanValue = proto.Bool(value.(bool))
	case types.CumulativeDistributionType:
		d.Distribution = serializeDistributionV1(value.(*distribution.Distribution))
		d.Distribution.IsCumulative = proto.Bool(true)
	case types.NonCumulativeDistributionType:
		d.Distribution = serializeDistributionV1(value.(*distribution.Distribution))
		d.Distribution.IsCumulative = proto.Bool(false)
	}
}

func serializeDistributionV1(d *distribution.Distribution) *pb.PrecomputedDistribution {
	ret := pb.PrecomputedDistribution{}

	// Copy the bucketer params.
	if d.Bucketer().Width() == 0 {
		switch d.Bucketer().GrowthFactor() {
		case 2:
			ret.SpecType = pb.PrecomputedDistribution_CANONICAL_POWERS_OF_2.Enum()
		case math.Pow(10, 0.2):
			ret.SpecType = pb.PrecomputedDistribution_CANONICAL_POWERS_OF_10_P_0_2.Enum()
		case 10:
			ret.SpecType = pb.PrecomputedDistribution_CANONICAL_POWERS_OF_10.Enum()
		}
	}

	if ret.SpecType == nil {
		ret.SpecType = pb.PrecomputedDistribution_CUSTOM_PARAMETERIZED.Enum()
		ret.Width = proto.Float64(d.Bucketer().Width())
		ret.GrowthFactor = proto.Float64(d.Bucketer().GrowthFactor())
		ret.NumBuckets = proto.Int32(int32(d.Bucketer().NumFiniteBuckets()))
	}

	// Copy the distribution bucket values.  Exclude the overflow buckets on each
	// end.
	if len(d.Buckets()) >= 1 {
		if len(d.Buckets()) == d.Bucketer().NumBuckets() {
			ret.Bucket = runningZeroes(d.Buckets()[1 : len(d.Buckets())-1])
		} else {
			ret.Bucket = runningZeroes(d.Buckets()[1:])
		}
	}

	// Add overflow buckets if present.
	if len(d.Buckets()) >= 1 {
		ret.Underflow = proto.Int64(d.Buckets()[0])
	}
	if len(d.Buckets()) == d.Bucketer().NumBuckets() {
		ret.Overflow = proto.Int64(d.Buckets()[d.Bucketer().NumBuckets()-1])
	}

	if d.Count() > 0 {
		ret.Mean = proto.Float64(d.Sum() / float64(d.Count()))
	}

	return &ret
}
