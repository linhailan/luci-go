// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package types

import (
	"time"

	"github.com/luci/luci-go/common/tsmon/field"
)

// Cell is the smallest unit of data recorded by tsmon.  Metrics can be
// thought of as multi-dimensional maps (with fields defining the dimensions) -
// a Cell is one value in that map, with information about its fields and its
// type.
type Cell struct {
	MetricInfo
	CellData
}

// MetricInfo contains the definition of a metric.
type MetricInfo struct {
	MetricName string
	Fields     []field.Field
	ValueType  ValueType
}

// CellData contains the value of a single cell.
type CellData struct {
	FieldVals []interface{}
	Target    Target
	ResetTime time.Time
	Value     interface{}
}
