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

// AccessLevel is an enum
type AccessLevel uint8

type IAccessLevel interface {
	Purpose() string
	NeedsAuthentication() bool
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	AccessLevel_Level0  AccessLevel = 0x0
	AccessLevel_Level1  AccessLevel = 0x1
	AccessLevel_Level2  AccessLevel = 0x2
	AccessLevel_Level3  AccessLevel = 0x3
	AccessLevel_Level15 AccessLevel = 0xF
)

var AccessLevelValues []AccessLevel

func init() {
	_ = errors.New
	AccessLevelValues = []AccessLevel{
		AccessLevel_Level0,
		AccessLevel_Level1,
		AccessLevel_Level2,
		AccessLevel_Level3,
		AccessLevel_Level15,
	}
}

func (e AccessLevel) Purpose() string {
	switch e {
	case 0x0:
		{ /* '0x0' */
			return "system manufacturer"
		}
	case 0x1:
		{ /* '0x1' */
			return "product manufacturer"
		}
	case 0x2:
		{ /* '0x2' */
			return "configuration"
		}
	case 0x3:
		{ /* '0x3' */
			return "end-user"
		}
	case 0xF:
		{ /* '0xF' */
			return "read access"
		}
	default:
		{
			return ""
		}
	}
}

func AccessLevelFirstEnumForFieldPurpose(value string) (AccessLevel, error) {
	for _, sizeValue := range AccessLevelValues {
		if sizeValue.Purpose() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing Purpose not found", value)
}

func (e AccessLevel) NeedsAuthentication() bool {
	switch e {
	case 0x0:
		{ /* '0x0' */
			return true
		}
	case 0x1:
		{ /* '0x1' */
			return true
		}
	case 0x2:
		{ /* '0x2' */
			return true
		}
	case 0x3:
		{ /* '0x3' */
			return false
		}
	case 0xF:
		{ /* '0xF' */
			return false
		}
	default:
		{
			return false
		}
	}
}

func AccessLevelFirstEnumForFieldNeedsAuthentication(value bool) (AccessLevel, error) {
	for _, sizeValue := range AccessLevelValues {
		if sizeValue.NeedsAuthentication() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NeedsAuthentication not found", value)
}
func AccessLevelByValue(value uint8) AccessLevel {
	switch value {
	case 0x0:
		return AccessLevel_Level0
	case 0x1:
		return AccessLevel_Level1
	case 0x2:
		return AccessLevel_Level2
	case 0x3:
		return AccessLevel_Level3
	case 0xF:
		return AccessLevel_Level15
	}
	return 0
}

func AccessLevelByName(value string) (enum AccessLevel, ok bool) {
	ok = true
	switch value {
	case "Level0":
		enum = AccessLevel_Level0
	case "Level1":
		enum = AccessLevel_Level1
	case "Level2":
		enum = AccessLevel_Level2
	case "Level3":
		enum = AccessLevel_Level3
	case "Level15":
		enum = AccessLevel_Level15
	default:
		enum = 0
		ok = false
	}
	return
}

func AccessLevelKnows(value uint8) bool {
	for _, typeValue := range AccessLevelValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAccessLevel(structType interface{}) AccessLevel {
	castFunc := func(typ interface{}) AccessLevel {
		if sAccessLevel, ok := typ.(AccessLevel); ok {
			return sAccessLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessLevel) GetLengthInBits() uint16 {
	return 4
}

func (m AccessLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AccessLevelParse(readBuffer utils.ReadBuffer) (AccessLevel, error) {
	val, err := readBuffer.ReadUint8("AccessLevel", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessLevel")
	}
	return AccessLevelByValue(val), nil
}

func (e AccessLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AccessLevel", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessLevel) PLC4XEnumName() string {
	switch e {
	case AccessLevel_Level0:
		return "Level0"
	case AccessLevel_Level1:
		return "Level1"
	case AccessLevel_Level2:
		return "Level2"
	case AccessLevel_Level3:
		return "Level3"
	case AccessLevel_Level15:
		return "Level15"
	}
	return ""
}

func (e AccessLevel) String() string {
	return e.PLC4XEnumName()
}
