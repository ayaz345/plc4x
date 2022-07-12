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

// BACnetLightingOperation is an enum
type BACnetLightingOperation uint16

type IBACnetLightingOperation interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLightingOperation_NONE                     BACnetLightingOperation = 0
	BACnetLightingOperation_FADE_TO                  BACnetLightingOperation = 1
	BACnetLightingOperation_RAMP_TO                  BACnetLightingOperation = 2
	BACnetLightingOperation_STEP_UP                  BACnetLightingOperation = 3
	BACnetLightingOperation_STEP_DOWN                BACnetLightingOperation = 4
	BACnetLightingOperation_STEP_ON                  BACnetLightingOperation = 5
	BACnetLightingOperation_STEP_OFF                 BACnetLightingOperation = 6
	BACnetLightingOperation_WARN                     BACnetLightingOperation = 7
	BACnetLightingOperation_WARN_OFF                 BACnetLightingOperation = 8
	BACnetLightingOperation_WARN_RELINQUISH          BACnetLightingOperation = 9
	BACnetLightingOperation_STOP                     BACnetLightingOperation = 10
	BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE BACnetLightingOperation = 0xFFFF
)

var BACnetLightingOperationValues []BACnetLightingOperation

func init() {
	_ = errors.New
	BACnetLightingOperationValues = []BACnetLightingOperation{
		BACnetLightingOperation_NONE,
		BACnetLightingOperation_FADE_TO,
		BACnetLightingOperation_RAMP_TO,
		BACnetLightingOperation_STEP_UP,
		BACnetLightingOperation_STEP_DOWN,
		BACnetLightingOperation_STEP_ON,
		BACnetLightingOperation_STEP_OFF,
		BACnetLightingOperation_WARN,
		BACnetLightingOperation_WARN_OFF,
		BACnetLightingOperation_WARN_RELINQUISH,
		BACnetLightingOperation_STOP,
		BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLightingOperationByValue(value uint16) BACnetLightingOperation {
	switch value {
	case 0:
		return BACnetLightingOperation_NONE
	case 0xFFFF:
		return BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetLightingOperation_FADE_TO
	case 10:
		return BACnetLightingOperation_STOP
	case 2:
		return BACnetLightingOperation_RAMP_TO
	case 3:
		return BACnetLightingOperation_STEP_UP
	case 4:
		return BACnetLightingOperation_STEP_DOWN
	case 5:
		return BACnetLightingOperation_STEP_ON
	case 6:
		return BACnetLightingOperation_STEP_OFF
	case 7:
		return BACnetLightingOperation_WARN
	case 8:
		return BACnetLightingOperation_WARN_OFF
	case 9:
		return BACnetLightingOperation_WARN_RELINQUISH
	}
	return 0
}

func BACnetLightingOperationByName(value string) (enum BACnetLightingOperation, ok bool) {
	ok = true
	switch value {
	case "NONE":
		enum = BACnetLightingOperation_NONE
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE
	case "FADE_TO":
		enum = BACnetLightingOperation_FADE_TO
	case "STOP":
		enum = BACnetLightingOperation_STOP
	case "RAMP_TO":
		enum = BACnetLightingOperation_RAMP_TO
	case "STEP_UP":
		enum = BACnetLightingOperation_STEP_UP
	case "STEP_DOWN":
		enum = BACnetLightingOperation_STEP_DOWN
	case "STEP_ON":
		enum = BACnetLightingOperation_STEP_ON
	case "STEP_OFF":
		enum = BACnetLightingOperation_STEP_OFF
	case "WARN":
		enum = BACnetLightingOperation_WARN
	case "WARN_OFF":
		enum = BACnetLightingOperation_WARN_OFF
	case "WARN_RELINQUISH":
		enum = BACnetLightingOperation_WARN_RELINQUISH
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetLightingOperationKnows(value uint16) bool {
	for _, typeValue := range BACnetLightingOperationValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLightingOperation(structType interface{}) BACnetLightingOperation {
	castFunc := func(typ interface{}) BACnetLightingOperation {
		if sBACnetLightingOperation, ok := typ.(BACnetLightingOperation); ok {
			return sBACnetLightingOperation
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLightingOperation) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLightingOperation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLightingOperationParse(readBuffer utils.ReadBuffer) (BACnetLightingOperation, error) {
	val, err := readBuffer.ReadUint16("BACnetLightingOperation", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLightingOperation")
	}
	return BACnetLightingOperationByValue(val), nil
}

func (e BACnetLightingOperation) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLightingOperation", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLightingOperation) PLC4XEnumName() string {
	switch e {
	case BACnetLightingOperation_NONE:
		return "NONE"
	case BACnetLightingOperation_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLightingOperation_FADE_TO:
		return "FADE_TO"
	case BACnetLightingOperation_STOP:
		return "STOP"
	case BACnetLightingOperation_RAMP_TO:
		return "RAMP_TO"
	case BACnetLightingOperation_STEP_UP:
		return "STEP_UP"
	case BACnetLightingOperation_STEP_DOWN:
		return "STEP_DOWN"
	case BACnetLightingOperation_STEP_ON:
		return "STEP_ON"
	case BACnetLightingOperation_STEP_OFF:
		return "STEP_OFF"
	case BACnetLightingOperation_WARN:
		return "WARN"
	case BACnetLightingOperation_WARN_OFF:
		return "WARN_OFF"
	case BACnetLightingOperation_WARN_RELINQUISH:
		return "WARN_RELINQUISH"
	}
	return ""
}

func (e BACnetLightingOperation) String() string {
	return e.PLC4XEnumName()
}
