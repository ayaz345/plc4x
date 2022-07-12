/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// MaxSegmentsAccepted is an enum
type MaxSegmentsAccepted uint8

type IMaxSegmentsAccepted interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	MaxSegmentsAccepted_UNSPECIFIED           MaxSegmentsAccepted = 0x0
	MaxSegmentsAccepted_NUM_SEGMENTS_02       MaxSegmentsAccepted = 0x1
	MaxSegmentsAccepted_NUM_SEGMENTS_04       MaxSegmentsAccepted = 0x2
	MaxSegmentsAccepted_NUM_SEGMENTS_08       MaxSegmentsAccepted = 0x3
	MaxSegmentsAccepted_NUM_SEGMENTS_16       MaxSegmentsAccepted = 0x4
	MaxSegmentsAccepted_NUM_SEGMENTS_32       MaxSegmentsAccepted = 0x5
	MaxSegmentsAccepted_NUM_SEGMENTS_64       MaxSegmentsAccepted = 0x6
	MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS MaxSegmentsAccepted = 0x7
)

var MaxSegmentsAcceptedValues []MaxSegmentsAccepted

func init() {
	_ = errors.New
	MaxSegmentsAcceptedValues = []MaxSegmentsAccepted{
		MaxSegmentsAccepted_UNSPECIFIED,
		MaxSegmentsAccepted_NUM_SEGMENTS_02,
		MaxSegmentsAccepted_NUM_SEGMENTS_04,
		MaxSegmentsAccepted_NUM_SEGMENTS_08,
		MaxSegmentsAccepted_NUM_SEGMENTS_16,
		MaxSegmentsAccepted_NUM_SEGMENTS_32,
		MaxSegmentsAccepted_NUM_SEGMENTS_64,
		MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS,
	}
}

func MaxSegmentsAcceptedByValue(value uint8) MaxSegmentsAccepted {
	switch value {
	case 0x0:
		return MaxSegmentsAccepted_UNSPECIFIED
	case 0x1:
		return MaxSegmentsAccepted_NUM_SEGMENTS_02
	case 0x2:
		return MaxSegmentsAccepted_NUM_SEGMENTS_04
	case 0x3:
		return MaxSegmentsAccepted_NUM_SEGMENTS_08
	case 0x4:
		return MaxSegmentsAccepted_NUM_SEGMENTS_16
	case 0x5:
		return MaxSegmentsAccepted_NUM_SEGMENTS_32
	case 0x6:
		return MaxSegmentsAccepted_NUM_SEGMENTS_64
	case 0x7:
		return MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS
	}
	return 0
}

func MaxSegmentsAcceptedByName(value string) (enum MaxSegmentsAccepted, ok bool) {
	ok = true
	switch value {
	case "UNSPECIFIED":
		enum = MaxSegmentsAccepted_UNSPECIFIED
	case "NUM_SEGMENTS_02":
		enum = MaxSegmentsAccepted_NUM_SEGMENTS_02
	case "NUM_SEGMENTS_04":
		enum = MaxSegmentsAccepted_NUM_SEGMENTS_04
	case "NUM_SEGMENTS_08":
		enum = MaxSegmentsAccepted_NUM_SEGMENTS_08
	case "NUM_SEGMENTS_16":
		enum = MaxSegmentsAccepted_NUM_SEGMENTS_16
	case "NUM_SEGMENTS_32":
		enum = MaxSegmentsAccepted_NUM_SEGMENTS_32
	case "NUM_SEGMENTS_64":
		enum = MaxSegmentsAccepted_NUM_SEGMENTS_64
	case "MORE_THAN_64_SEGMENTS":
		enum = MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS
	default:
		enum = 0
		ok = false
	}
	return
}

func MaxSegmentsAcceptedKnows(value uint8) bool {
	for _, typeValue := range MaxSegmentsAcceptedValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMaxSegmentsAccepted(structType interface{}) MaxSegmentsAccepted {
	castFunc := func(typ interface{}) MaxSegmentsAccepted {
		if sMaxSegmentsAccepted, ok := typ.(MaxSegmentsAccepted); ok {
			return sMaxSegmentsAccepted
		}
		return 0
	}
	return castFunc(structType)
}

func (m MaxSegmentsAccepted) GetLengthInBits() uint16 {
	return 3
}

func (m MaxSegmentsAccepted) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MaxSegmentsAcceptedParse(readBuffer utils.ReadBuffer) (MaxSegmentsAccepted, error) {
	val, err := readBuffer.ReadUint8("MaxSegmentsAccepted", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MaxSegmentsAccepted")
	}
	return MaxSegmentsAcceptedByValue(val), nil
}

func (e MaxSegmentsAccepted) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MaxSegmentsAccepted", 3, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MaxSegmentsAccepted) PLC4XEnumName() string {
	switch e {
	case MaxSegmentsAccepted_UNSPECIFIED:
		return "UNSPECIFIED"
	case MaxSegmentsAccepted_NUM_SEGMENTS_02:
		return "NUM_SEGMENTS_02"
	case MaxSegmentsAccepted_NUM_SEGMENTS_04:
		return "NUM_SEGMENTS_04"
	case MaxSegmentsAccepted_NUM_SEGMENTS_08:
		return "NUM_SEGMENTS_08"
	case MaxSegmentsAccepted_NUM_SEGMENTS_16:
		return "NUM_SEGMENTS_16"
	case MaxSegmentsAccepted_NUM_SEGMENTS_32:
		return "NUM_SEGMENTS_32"
	case MaxSegmentsAccepted_NUM_SEGMENTS_64:
		return "NUM_SEGMENTS_64"
	case MaxSegmentsAccepted_MORE_THAN_64_SEGMENTS:
		return "MORE_THAN_64_SEGMENTS"
	}
	return ""
}

func (e MaxSegmentsAccepted) String() string {
	return e.PLC4XEnumName()
}
