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

// NPDUNetworkPriority is an enum
type NPDUNetworkPriority uint8

type INPDUNetworkPriority interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	NPDUNetworkPriority_LIFE_SAVETY_MESSAGE        NPDUNetworkPriority = 3
	NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE NPDUNetworkPriority = 2
	NPDUNetworkPriority_URGENT_MESSAGE             NPDUNetworkPriority = 1
	NPDUNetworkPriority_NORMAL_MESSAGE             NPDUNetworkPriority = 0
)

var NPDUNetworkPriorityValues []NPDUNetworkPriority

func init() {
	_ = errors.New
	NPDUNetworkPriorityValues = []NPDUNetworkPriority{
		NPDUNetworkPriority_LIFE_SAVETY_MESSAGE,
		NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE,
		NPDUNetworkPriority_URGENT_MESSAGE,
		NPDUNetworkPriority_NORMAL_MESSAGE,
	}
}

func NPDUNetworkPriorityByValue(value uint8) NPDUNetworkPriority {
	switch value {
	case 0:
		return NPDUNetworkPriority_NORMAL_MESSAGE
	case 1:
		return NPDUNetworkPriority_URGENT_MESSAGE
	case 2:
		return NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE
	case 3:
		return NPDUNetworkPriority_LIFE_SAVETY_MESSAGE
	}
	return 0
}

func NPDUNetworkPriorityByName(value string) (enum NPDUNetworkPriority, ok bool) {
	ok = true
	switch value {
	case "NORMAL_MESSAGE":
		enum = NPDUNetworkPriority_NORMAL_MESSAGE
	case "URGENT_MESSAGE":
		enum = NPDUNetworkPriority_URGENT_MESSAGE
	case "CRITICAL_EQUIPMENT_MESSAGE":
		enum = NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE
	case "LIFE_SAVETY_MESSAGE":
		enum = NPDUNetworkPriority_LIFE_SAVETY_MESSAGE
	default:
		enum = 0
		ok = false
	}
	return
}

func NPDUNetworkPriorityKnows(value uint8) bool {
	for _, typeValue := range NPDUNetworkPriorityValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNPDUNetworkPriority(structType interface{}) NPDUNetworkPriority {
	castFunc := func(typ interface{}) NPDUNetworkPriority {
		if sNPDUNetworkPriority, ok := typ.(NPDUNetworkPriority); ok {
			return sNPDUNetworkPriority
		}
		return 0
	}
	return castFunc(structType)
}

func (m NPDUNetworkPriority) GetLengthInBits() uint16 {
	return 2
}

func (m NPDUNetworkPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NPDUNetworkPriorityParse(readBuffer utils.ReadBuffer) (NPDUNetworkPriority, error) {
	val, err := readBuffer.ReadUint8("NPDUNetworkPriority", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NPDUNetworkPriority")
	}
	return NPDUNetworkPriorityByValue(val), nil
}

func (e NPDUNetworkPriority) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("NPDUNetworkPriority", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NPDUNetworkPriority) PLC4XEnumName() string {
	switch e {
	case NPDUNetworkPriority_NORMAL_MESSAGE:
		return "NORMAL_MESSAGE"
	case NPDUNetworkPriority_URGENT_MESSAGE:
		return "URGENT_MESSAGE"
	case NPDUNetworkPriority_CRITICAL_EQUIPMENT_MESSAGE:
		return "CRITICAL_EQUIPMENT_MESSAGE"
	case NPDUNetworkPriority_LIFE_SAVETY_MESSAGE:
		return "LIFE_SAVETY_MESSAGE"
	}
	return ""
}

func (e NPDUNetworkPriority) String() string {
	return e.PLC4XEnumName()
}
