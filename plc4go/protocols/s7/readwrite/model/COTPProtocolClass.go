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

// COTPProtocolClass is an enum
type COTPProtocolClass uint8

type ICOTPProtocolClass interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	COTPProtocolClass_CLASS_0 COTPProtocolClass = 0x00
	COTPProtocolClass_CLASS_1 COTPProtocolClass = 0x10
	COTPProtocolClass_CLASS_2 COTPProtocolClass = 0x20
	COTPProtocolClass_CLASS_3 COTPProtocolClass = 0x30
	COTPProtocolClass_CLASS_4 COTPProtocolClass = 0x40
)

var COTPProtocolClassValues []COTPProtocolClass

func init() {
	_ = errors.New
	COTPProtocolClassValues = []COTPProtocolClass{
		COTPProtocolClass_CLASS_0,
		COTPProtocolClass_CLASS_1,
		COTPProtocolClass_CLASS_2,
		COTPProtocolClass_CLASS_3,
		COTPProtocolClass_CLASS_4,
	}
}

func COTPProtocolClassByValue(value uint8) COTPProtocolClass {
	switch value {
	case 0x00:
		return COTPProtocolClass_CLASS_0
	case 0x10:
		return COTPProtocolClass_CLASS_1
	case 0x20:
		return COTPProtocolClass_CLASS_2
	case 0x30:
		return COTPProtocolClass_CLASS_3
	case 0x40:
		return COTPProtocolClass_CLASS_4
	}
	return 0
}

func COTPProtocolClassByName(value string) (enum COTPProtocolClass, ok bool) {
	ok = true
	switch value {
	case "CLASS_0":
		enum = COTPProtocolClass_CLASS_0
	case "CLASS_1":
		enum = COTPProtocolClass_CLASS_1
	case "CLASS_2":
		enum = COTPProtocolClass_CLASS_2
	case "CLASS_3":
		enum = COTPProtocolClass_CLASS_3
	case "CLASS_4":
		enum = COTPProtocolClass_CLASS_4
	default:
		enum = 0
		ok = false
	}
	return
}

func COTPProtocolClassKnows(value uint8) bool {
	for _, typeValue := range COTPProtocolClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCOTPProtocolClass(structType interface{}) COTPProtocolClass {
	castFunc := func(typ interface{}) COTPProtocolClass {
		if sCOTPProtocolClass, ok := typ.(COTPProtocolClass); ok {
			return sCOTPProtocolClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m COTPProtocolClass) GetLengthInBits() uint16 {
	return 8
}

func (m COTPProtocolClass) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPProtocolClassParse(readBuffer utils.ReadBuffer) (COTPProtocolClass, error) {
	val, err := readBuffer.ReadUint8("COTPProtocolClass", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading COTPProtocolClass")
	}
	return COTPProtocolClassByValue(val), nil
}

func (e COTPProtocolClass) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("COTPProtocolClass", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e COTPProtocolClass) PLC4XEnumName() string {
	switch e {
	case COTPProtocolClass_CLASS_0:
		return "CLASS_0"
	case COTPProtocolClass_CLASS_1:
		return "CLASS_1"
	case COTPProtocolClass_CLASS_2:
		return "CLASS_2"
	case COTPProtocolClass_CLASS_3:
		return "CLASS_3"
	case COTPProtocolClass_CLASS_4:
		return "CLASS_4"
	}
	return ""
}

func (e COTPProtocolClass) String() string {
	return e.PLC4XEnumName()
}
