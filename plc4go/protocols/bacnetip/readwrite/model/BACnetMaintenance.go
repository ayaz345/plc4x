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

// BACnetMaintenance is an enum
type BACnetMaintenance uint8

type IBACnetMaintenance interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetMaintenance_NONE                     BACnetMaintenance = 0
	BACnetMaintenance_PERIODIC_TEST            BACnetMaintenance = 1
	BACnetMaintenance_NEED_SERVICE_OPERATIONAL BACnetMaintenance = 2
	BACnetMaintenance_NEED_SERVICE_INOPERATIVE BACnetMaintenance = 3
	BACnetMaintenance_VENDOR_PROPRIETARY_VALUE BACnetMaintenance = 0xFF
)

var BACnetMaintenanceValues []BACnetMaintenance

func init() {
	_ = errors.New
	BACnetMaintenanceValues = []BACnetMaintenance{
		BACnetMaintenance_NONE,
		BACnetMaintenance_PERIODIC_TEST,
		BACnetMaintenance_NEED_SERVICE_OPERATIONAL,
		BACnetMaintenance_NEED_SERVICE_INOPERATIVE,
		BACnetMaintenance_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetMaintenanceByValue(value uint8) BACnetMaintenance {
	switch value {
	case 0:
		return BACnetMaintenance_NONE
	case 0xFF:
		return BACnetMaintenance_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetMaintenance_PERIODIC_TEST
	case 2:
		return BACnetMaintenance_NEED_SERVICE_OPERATIONAL
	case 3:
		return BACnetMaintenance_NEED_SERVICE_INOPERATIVE
	}
	return 0
}

func BACnetMaintenanceByName(value string) (enum BACnetMaintenance, ok bool) {
	ok = true
	switch value {
	case "NONE":
		enum = BACnetMaintenance_NONE
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetMaintenance_VENDOR_PROPRIETARY_VALUE
	case "PERIODIC_TEST":
		enum = BACnetMaintenance_PERIODIC_TEST
	case "NEED_SERVICE_OPERATIONAL":
		enum = BACnetMaintenance_NEED_SERVICE_OPERATIONAL
	case "NEED_SERVICE_INOPERATIVE":
		enum = BACnetMaintenance_NEED_SERVICE_INOPERATIVE
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetMaintenanceKnows(value uint8) bool {
	for _, typeValue := range BACnetMaintenanceValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetMaintenance(structType interface{}) BACnetMaintenance {
	castFunc := func(typ interface{}) BACnetMaintenance {
		if sBACnetMaintenance, ok := typ.(BACnetMaintenance); ok {
			return sBACnetMaintenance
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetMaintenance) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetMaintenance) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetMaintenanceParse(readBuffer utils.ReadBuffer) (BACnetMaintenance, error) {
	val, err := readBuffer.ReadUint8("BACnetMaintenance", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetMaintenance")
	}
	return BACnetMaintenanceByValue(val), nil
}

func (e BACnetMaintenance) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetMaintenance", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetMaintenance) PLC4XEnumName() string {
	switch e {
	case BACnetMaintenance_NONE:
		return "NONE"
	case BACnetMaintenance_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetMaintenance_PERIODIC_TEST:
		return "PERIODIC_TEST"
	case BACnetMaintenance_NEED_SERVICE_OPERATIONAL:
		return "NEED_SERVICE_OPERATIONAL"
	case BACnetMaintenance_NEED_SERVICE_INOPERATIVE:
		return "NEED_SERVICE_INOPERATIVE"
	}
	return ""
}

func (e BACnetMaintenance) String() string {
	return e.PLC4XEnumName()
}
