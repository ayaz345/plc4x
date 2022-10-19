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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Attribute is an enum
type Attribute uint8

type IAttribute interface {
	BytesReturned() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	Attribute_Manufacturer              Attribute = 0x00
	Attribute_Type                      Attribute = 0x01
	Attribute_FirmwareVersion           Attribute = 0x02
	Attribute_Summary                   Attribute = 0x03
	Attribute_ExtendedDiagnosticSummary Attribute = 0x04
	Attribute_NetworkTerminalLevels     Attribute = 0x05
	Attribute_TerminalLevel             Attribute = 0x06
	Attribute_NetworkVoltage            Attribute = 0x07
	Attribute_GAVValuesCurrent          Attribute = 0x08
	Attribute_GAVValuesStored           Attribute = 0x09
	Attribute_GAVPhysicalAddresses      Attribute = 0x0A
	Attribute_LogicalAssignment         Attribute = 0x0B
	Attribute_Delays                    Attribute = 0x0C
	Attribute_MinimumLevels             Attribute = 0x0D
	Attribute_MaximumLevels             Attribute = 0x0E
	Attribute_CurrentSenseLevels        Attribute = 0x0F
	Attribute_OutputUnitSummary         Attribute = 0x10
	Attribute_DSIStatus                 Attribute = 0x11
)

var AttributeValues []Attribute

func init() {
	_ = errors.New
	AttributeValues = []Attribute{
		Attribute_Manufacturer,
		Attribute_Type,
		Attribute_FirmwareVersion,
		Attribute_Summary,
		Attribute_ExtendedDiagnosticSummary,
		Attribute_NetworkTerminalLevels,
		Attribute_TerminalLevel,
		Attribute_NetworkVoltage,
		Attribute_GAVValuesCurrent,
		Attribute_GAVValuesStored,
		Attribute_GAVPhysicalAddresses,
		Attribute_LogicalAssignment,
		Attribute_Delays,
		Attribute_MinimumLevels,
		Attribute_MaximumLevels,
		Attribute_CurrentSenseLevels,
		Attribute_OutputUnitSummary,
		Attribute_DSIStatus,
	}
}

func (e Attribute) BytesReturned() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 8
		}
	case 0x01:
		{ /* '0x01' */
			return 8
		}
	case 0x02:
		{ /* '0x02' */
			return 8
		}
	case 0x03:
		{ /* '0x03' */
			return 9
		}
	case 0x04:
		{ /* '0x04' */
			return 13
		}
	case 0x05:
		{ /* '0x05' */
			return 13
		}
	case 0x06:
		{ /* '0x06' */
			return 13
		}
	case 0x07:
		{ /* '0x07' */
			return 5
		}
	case 0x08:
		{ /* '0x08' */
			return 16
		}
	case 0x09:
		{ /* '0x09' */
			return 16
		}
	case 0x0A:
		{ /* '0x0A' */
			return 16
		}
	case 0x0B:
		{ /* '0x0B' */
			return 13
		}
	case 0x0C:
		{ /* '0x0C' */
			return 14
		}
	case 0x0D:
		{ /* '0x0D' */
			return 13
		}
	case 0x0E:
		{ /* '0x0E' */
			return 13
		}
	case 0x0F:
		{ /* '0x0F' */
			return 8
		}
	case 0x10:
		{ /* '0x10' */
			return 4
		}
	case 0x11:
		{ /* '0x11' */
			return 10
		}
	default:
		{
			return 0
		}
	}
}

func AttributeFirstEnumForFieldBytesReturned(value uint8) (Attribute, error) {
	for _, sizeValue := range AttributeValues {
		if sizeValue.BytesReturned() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing BytesReturned not found", value)
}
func AttributeByValue(value uint8) (enum Attribute, ok bool) {
	switch value {
	case 0x00:
		return Attribute_Manufacturer, true
	case 0x01:
		return Attribute_Type, true
	case 0x02:
		return Attribute_FirmwareVersion, true
	case 0x03:
		return Attribute_Summary, true
	case 0x04:
		return Attribute_ExtendedDiagnosticSummary, true
	case 0x05:
		return Attribute_NetworkTerminalLevels, true
	case 0x06:
		return Attribute_TerminalLevel, true
	case 0x07:
		return Attribute_NetworkVoltage, true
	case 0x08:
		return Attribute_GAVValuesCurrent, true
	case 0x09:
		return Attribute_GAVValuesStored, true
	case 0x0A:
		return Attribute_GAVPhysicalAddresses, true
	case 0x0B:
		return Attribute_LogicalAssignment, true
	case 0x0C:
		return Attribute_Delays, true
	case 0x0D:
		return Attribute_MinimumLevels, true
	case 0x0E:
		return Attribute_MaximumLevels, true
	case 0x0F:
		return Attribute_CurrentSenseLevels, true
	case 0x10:
		return Attribute_OutputUnitSummary, true
	case 0x11:
		return Attribute_DSIStatus, true
	}
	return 0, false
}

func AttributeByName(value string) (enum Attribute, ok bool) {
	switch value {
	case "Manufacturer":
		return Attribute_Manufacturer, true
	case "Type":
		return Attribute_Type, true
	case "FirmwareVersion":
		return Attribute_FirmwareVersion, true
	case "Summary":
		return Attribute_Summary, true
	case "ExtendedDiagnosticSummary":
		return Attribute_ExtendedDiagnosticSummary, true
	case "NetworkTerminalLevels":
		return Attribute_NetworkTerminalLevels, true
	case "TerminalLevel":
		return Attribute_TerminalLevel, true
	case "NetworkVoltage":
		return Attribute_NetworkVoltage, true
	case "GAVValuesCurrent":
		return Attribute_GAVValuesCurrent, true
	case "GAVValuesStored":
		return Attribute_GAVValuesStored, true
	case "GAVPhysicalAddresses":
		return Attribute_GAVPhysicalAddresses, true
	case "LogicalAssignment":
		return Attribute_LogicalAssignment, true
	case "Delays":
		return Attribute_Delays, true
	case "MinimumLevels":
		return Attribute_MinimumLevels, true
	case "MaximumLevels":
		return Attribute_MaximumLevels, true
	case "CurrentSenseLevels":
		return Attribute_CurrentSenseLevels, true
	case "OutputUnitSummary":
		return Attribute_OutputUnitSummary, true
	case "DSIStatus":
		return Attribute_DSIStatus, true
	}
	return 0, false
}

func AttributeKnows(value uint8) bool {
	for _, typeValue := range AttributeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAttribute(structType interface{}) Attribute {
	castFunc := func(typ interface{}) Attribute {
		if sAttribute, ok := typ.(Attribute); ok {
			return sAttribute
		}
		return 0
	}
	return castFunc(structType)
}

func (m Attribute) GetLengthInBits() uint16 {
	return 8
}

func (m Attribute) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AttributeParse(readBuffer utils.ReadBuffer) (Attribute, error) {
	val, err := readBuffer.ReadUint8("Attribute", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading Attribute")
	}
	if enum, ok := AttributeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return Attribute(val), nil
	} else {
		return enum, nil
	}
}

func (e Attribute) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("Attribute", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e Attribute) PLC4XEnumName() string {
	switch e {
	case Attribute_Manufacturer:
		return "Manufacturer"
	case Attribute_Type:
		return "Type"
	case Attribute_FirmwareVersion:
		return "FirmwareVersion"
	case Attribute_Summary:
		return "Summary"
	case Attribute_ExtendedDiagnosticSummary:
		return "ExtendedDiagnosticSummary"
	case Attribute_NetworkTerminalLevels:
		return "NetworkTerminalLevels"
	case Attribute_TerminalLevel:
		return "TerminalLevel"
	case Attribute_NetworkVoltage:
		return "NetworkVoltage"
	case Attribute_GAVValuesCurrent:
		return "GAVValuesCurrent"
	case Attribute_GAVValuesStored:
		return "GAVValuesStored"
	case Attribute_GAVPhysicalAddresses:
		return "GAVPhysicalAddresses"
	case Attribute_LogicalAssignment:
		return "LogicalAssignment"
	case Attribute_Delays:
		return "Delays"
	case Attribute_MinimumLevels:
		return "MinimumLevels"
	case Attribute_MaximumLevels:
		return "MaximumLevels"
	case Attribute_CurrentSenseLevels:
		return "CurrentSenseLevels"
	case Attribute_OutputUnitSummary:
		return "OutputUnitSummary"
	case Attribute_DSIStatus:
		return "DSIStatus"
	}
	return ""
}

func (e Attribute) String() string {
	return e.PLC4XEnumName()
}