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

// SALCommandTypeContainer is an enum
type SALCommandTypeContainer uint8

type ISALCommandTypeContainer interface {
	CommandType() SALCommandType
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	SALCommandTypeContainer_SALCommandOff                       SALCommandTypeContainer = 0x01
	SALCommandTypeContainer_SALCommandOn                        SALCommandTypeContainer = 0x79
	SALCommandTypeContainer_SALCommandRampToLevel_Instantaneous SALCommandTypeContainer = 0x02
	SALCommandTypeContainer_SALCommandRampToLevel_4Second       SALCommandTypeContainer = 0x0A
	SALCommandTypeContainer_SALCommandRampToLevel_8Second       SALCommandTypeContainer = 0x12
	SALCommandTypeContainer_SALCommandRampToLevel_12Second      SALCommandTypeContainer = 0x1A
	SALCommandTypeContainer_SALCommandRampToLevel_20Second      SALCommandTypeContainer = 0x22
	SALCommandTypeContainer_SALCommandRampToLevel_30Second      SALCommandTypeContainer = 0x2A
	SALCommandTypeContainer_SALCommandRampToLevel_40Second      SALCommandTypeContainer = 0x32
	SALCommandTypeContainer_SALCommandRampToLevel_60Second      SALCommandTypeContainer = 0x3A
	SALCommandTypeContainer_SALCommandRampToLevel_90Second      SALCommandTypeContainer = 0x42
	SALCommandTypeContainer_SALCommandRampToLevel_120Second     SALCommandTypeContainer = 0x4A
	SALCommandTypeContainer_SALCommandRampToLevel_180Second     SALCommandTypeContainer = 0x52
	SALCommandTypeContainer_SALCommandRampToLevel_300Second     SALCommandTypeContainer = 0x5A
	SALCommandTypeContainer_SALCommandRampToLevel_420Second     SALCommandTypeContainer = 0x62
	SALCommandTypeContainer_SALCommandRampToLevel_600Second     SALCommandTypeContainer = 0x6A
	SALCommandTypeContainer_SALCommandRampToLevel_900Second     SALCommandTypeContainer = 0x72
	SALCommandTypeContainer_SALCommandRampToLevel_1020Second    SALCommandTypeContainer = 0x7A
	SALCommandTypeContainer_SALCommandTerminateRamp             SALCommandTypeContainer = 0x09
)

var SALCommandTypeContainerValues []SALCommandTypeContainer

func init() {
	_ = errors.New
	SALCommandTypeContainerValues = []SALCommandTypeContainer{
		SALCommandTypeContainer_SALCommandOff,
		SALCommandTypeContainer_SALCommandOn,
		SALCommandTypeContainer_SALCommandRampToLevel_Instantaneous,
		SALCommandTypeContainer_SALCommandRampToLevel_4Second,
		SALCommandTypeContainer_SALCommandRampToLevel_8Second,
		SALCommandTypeContainer_SALCommandRampToLevel_12Second,
		SALCommandTypeContainer_SALCommandRampToLevel_20Second,
		SALCommandTypeContainer_SALCommandRampToLevel_30Second,
		SALCommandTypeContainer_SALCommandRampToLevel_40Second,
		SALCommandTypeContainer_SALCommandRampToLevel_60Second,
		SALCommandTypeContainer_SALCommandRampToLevel_90Second,
		SALCommandTypeContainer_SALCommandRampToLevel_120Second,
		SALCommandTypeContainer_SALCommandRampToLevel_180Second,
		SALCommandTypeContainer_SALCommandRampToLevel_300Second,
		SALCommandTypeContainer_SALCommandRampToLevel_420Second,
		SALCommandTypeContainer_SALCommandRampToLevel_600Second,
		SALCommandTypeContainer_SALCommandRampToLevel_900Second,
		SALCommandTypeContainer_SALCommandRampToLevel_1020Second,
		SALCommandTypeContainer_SALCommandTerminateRamp,
	}
}

func (e SALCommandTypeContainer) CommandType() SALCommandType {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return SALCommandType_OFF
		}
	case 0x02:
		{ /* '0x02' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x09:
		{ /* '0x09' */
			return SALCommandType_TERMINATE_RAMP
		}
	case 0x0A:
		{ /* '0x0A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x12:
		{ /* '0x12' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x1A:
		{ /* '0x1A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x22:
		{ /* '0x22' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x2A:
		{ /* '0x2A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x32:
		{ /* '0x32' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x3A:
		{ /* '0x3A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x42:
		{ /* '0x42' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x4A:
		{ /* '0x4A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x52:
		{ /* '0x52' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x5A:
		{ /* '0x5A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x62:
		{ /* '0x62' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x6A:
		{ /* '0x6A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x72:
		{ /* '0x72' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	case 0x79:
		{ /* '0x79' */
			return SALCommandType_ON
		}
	case 0x7A:
		{ /* '0x7A' */
			return SALCommandType_RAMP_TO_LEVEL
		}
	default:
		{
			return 0
		}
	}
}

func SALCommandTypeContainerFirstEnumForFieldCommandType(value SALCommandType) (SALCommandTypeContainer, error) {
	for _, sizeValue := range SALCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func SALCommandTypeContainerByValue(value uint8) SALCommandTypeContainer {
	switch value {
	case 0x01:
		return SALCommandTypeContainer_SALCommandOff
	case 0x02:
		return SALCommandTypeContainer_SALCommandRampToLevel_Instantaneous
	case 0x09:
		return SALCommandTypeContainer_SALCommandTerminateRamp
	case 0x0A:
		return SALCommandTypeContainer_SALCommandRampToLevel_4Second
	case 0x12:
		return SALCommandTypeContainer_SALCommandRampToLevel_8Second
	case 0x1A:
		return SALCommandTypeContainer_SALCommandRampToLevel_12Second
	case 0x22:
		return SALCommandTypeContainer_SALCommandRampToLevel_20Second
	case 0x2A:
		return SALCommandTypeContainer_SALCommandRampToLevel_30Second
	case 0x32:
		return SALCommandTypeContainer_SALCommandRampToLevel_40Second
	case 0x3A:
		return SALCommandTypeContainer_SALCommandRampToLevel_60Second
	case 0x42:
		return SALCommandTypeContainer_SALCommandRampToLevel_90Second
	case 0x4A:
		return SALCommandTypeContainer_SALCommandRampToLevel_120Second
	case 0x52:
		return SALCommandTypeContainer_SALCommandRampToLevel_180Second
	case 0x5A:
		return SALCommandTypeContainer_SALCommandRampToLevel_300Second
	case 0x62:
		return SALCommandTypeContainer_SALCommandRampToLevel_420Second
	case 0x6A:
		return SALCommandTypeContainer_SALCommandRampToLevel_600Second
	case 0x72:
		return SALCommandTypeContainer_SALCommandRampToLevel_900Second
	case 0x79:
		return SALCommandTypeContainer_SALCommandOn
	case 0x7A:
		return SALCommandTypeContainer_SALCommandRampToLevel_1020Second
	}
	return 0
}

func SALCommandTypeContainerByName(value string) (enum SALCommandTypeContainer, ok bool) {
	ok = true
	switch value {
	case "SALCommandOff":
		enum = SALCommandTypeContainer_SALCommandOff
	case "SALCommandRampToLevel_Instantaneous":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_Instantaneous
	case "SALCommandTerminateRamp":
		enum = SALCommandTypeContainer_SALCommandTerminateRamp
	case "SALCommandRampToLevel_4Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_4Second
	case "SALCommandRampToLevel_8Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_8Second
	case "SALCommandRampToLevel_12Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_12Second
	case "SALCommandRampToLevel_20Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_20Second
	case "SALCommandRampToLevel_30Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_30Second
	case "SALCommandRampToLevel_40Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_40Second
	case "SALCommandRampToLevel_60Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_60Second
	case "SALCommandRampToLevel_90Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_90Second
	case "SALCommandRampToLevel_120Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_120Second
	case "SALCommandRampToLevel_180Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_180Second
	case "SALCommandRampToLevel_300Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_300Second
	case "SALCommandRampToLevel_420Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_420Second
	case "SALCommandRampToLevel_600Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_600Second
	case "SALCommandRampToLevel_900Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_900Second
	case "SALCommandOn":
		enum = SALCommandTypeContainer_SALCommandOn
	case "SALCommandRampToLevel_1020Second":
		enum = SALCommandTypeContainer_SALCommandRampToLevel_1020Second
	default:
		enum = 0
		ok = false
	}
	return
}

func SALCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range SALCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSALCommandTypeContainer(structType interface{}) SALCommandTypeContainer {
	castFunc := func(typ interface{}) SALCommandTypeContainer {
		if sSALCommandTypeContainer, ok := typ.(SALCommandTypeContainer); ok {
			return sSALCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m SALCommandTypeContainer) GetLengthInBits() uint16 {
	return 8
}

func (m SALCommandTypeContainer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALCommandTypeContainerParse(readBuffer utils.ReadBuffer) (SALCommandTypeContainer, error) {
	val, err := readBuffer.ReadUint8("SALCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SALCommandTypeContainer")
	}
	return SALCommandTypeContainerByValue(val), nil
}

func (e SALCommandTypeContainer) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("SALCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SALCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case SALCommandTypeContainer_SALCommandOff:
		return "SALCommandOff"
	case SALCommandTypeContainer_SALCommandRampToLevel_Instantaneous:
		return "SALCommandRampToLevel_Instantaneous"
	case SALCommandTypeContainer_SALCommandTerminateRamp:
		return "SALCommandTerminateRamp"
	case SALCommandTypeContainer_SALCommandRampToLevel_4Second:
		return "SALCommandRampToLevel_4Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_8Second:
		return "SALCommandRampToLevel_8Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_12Second:
		return "SALCommandRampToLevel_12Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_20Second:
		return "SALCommandRampToLevel_20Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_30Second:
		return "SALCommandRampToLevel_30Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_40Second:
		return "SALCommandRampToLevel_40Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_60Second:
		return "SALCommandRampToLevel_60Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_90Second:
		return "SALCommandRampToLevel_90Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_120Second:
		return "SALCommandRampToLevel_120Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_180Second:
		return "SALCommandRampToLevel_180Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_300Second:
		return "SALCommandRampToLevel_300Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_420Second:
		return "SALCommandRampToLevel_420Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_600Second:
		return "SALCommandRampToLevel_600Second"
	case SALCommandTypeContainer_SALCommandRampToLevel_900Second:
		return "SALCommandRampToLevel_900Second"
	case SALCommandTypeContainer_SALCommandOn:
		return "SALCommandOn"
	case SALCommandTypeContainer_SALCommandRampToLevel_1020Second:
		return "SALCommandRampToLevel_1020Second"
	}
	return ""
}

func (e SALCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
