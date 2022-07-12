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

// KnxDatapointMainType is an enum
type KnxDatapointMainType uint16

type IKnxDatapointMainType interface {
	Number() uint16
	Name() string
	SizeInBits() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	KnxDatapointMainType_DPT_UNKNOWN                                       KnxDatapointMainType = 0
	KnxDatapointMainType_DPT_64_BIT_SET                                    KnxDatapointMainType = 1
	KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE                         KnxDatapointMainType = 2
	KnxDatapointMainType_DPT_8_BYTE_SIGNED_VALUE                           KnxDatapointMainType = 3
	KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE                          KnxDatapointMainType = 4
	KnxDatapointMainType_DPT_8_BYTE_FLOAT_VALUE                            KnxDatapointMainType = 5
	KnxDatapointMainType_DPT_1_BIT                                         KnxDatapointMainType = 6
	KnxDatapointMainType_DPT_1_BIT_CONTROLLED                              KnxDatapointMainType = 7
	KnxDatapointMainType_DPT_3_BIT_CONTROLLED                              KnxDatapointMainType = 8
	KnxDatapointMainType_DPT_CHARACTER                                     KnxDatapointMainType = 9
	KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE                          KnxDatapointMainType = 10
	KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE                            KnxDatapointMainType = 11
	KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE                         KnxDatapointMainType = 12
	KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE                           KnxDatapointMainType = 13
	KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE                            KnxDatapointMainType = 14
	KnxDatapointMainType_DPT_TIME                                          KnxDatapointMainType = 15
	KnxDatapointMainType_DPT_DATE                                          KnxDatapointMainType = 16
	KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE                         KnxDatapointMainType = 17
	KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE                           KnxDatapointMainType = 18
	KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE                            KnxDatapointMainType = 19
	KnxDatapointMainType_DPT_ENTRANCE_ACCESS                               KnxDatapointMainType = 20
	KnxDatapointMainType_DPT_CHARACTER_STRING                              KnxDatapointMainType = 21
	KnxDatapointMainType_DPT_SCENE_NUMBER                                  KnxDatapointMainType = 22
	KnxDatapointMainType_DPT_SCENE_CONTROL                                 KnxDatapointMainType = 23
	KnxDatapointMainType_DPT_DATE_TIME                                     KnxDatapointMainType = 24
	KnxDatapointMainType_DPT_1_BYTE                                        KnxDatapointMainType = 25
	KnxDatapointMainType_DPT_8_BIT_SET                                     KnxDatapointMainType = 26
	KnxDatapointMainType_DPT_16_BIT_SET                                    KnxDatapointMainType = 27
	KnxDatapointMainType_DPT_2_BIT_SET                                     KnxDatapointMainType = 28
	KnxDatapointMainType_DPT_2_NIBBLE_SET                                  KnxDatapointMainType = 29
	KnxDatapointMainType_DPT_8_BIT_SET_2                                   KnxDatapointMainType = 30
	KnxDatapointMainType_DPT_32_BIT_SET                                    KnxDatapointMainType = 31
	KnxDatapointMainType_DPT_ELECTRICAL_ENERGY                             KnxDatapointMainType = 32
	KnxDatapointMainType_DPT_24_TIMES_CHANNEL_ACTIVATION                   KnxDatapointMainType = 33
	KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM          KnxDatapointMainType = 34
	KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM           KnxDatapointMainType = 35
	KnxDatapointMainType_DPT_DATAPOINT_TYPE_VERSION                        KnxDatapointMainType = 36
	KnxDatapointMainType_DPT_ALARM_INFO                                    KnxDatapointMainType = 37
	KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE                         KnxDatapointMainType = 38
	KnxDatapointMainType_DPT_SCALING_SPEED                                 KnxDatapointMainType = 39
	KnxDatapointMainType_DPT_4_1_1_BYTE_COMBINED_INFORMATION               KnxDatapointMainType = 40
	KnxDatapointMainType_DPT_MBUS_ADDRESS                                  KnxDatapointMainType = 41
	KnxDatapointMainType_DPT_3_BYTE_COLOUR_RGB                             KnxDatapointMainType = 42
	KnxDatapointMainType_DPT_LANGUAGE_CODE_ISO_639_1                       KnxDatapointMainType = 43
	KnxDatapointMainType_DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY KnxDatapointMainType = 44
	KnxDatapointMainType_DPT_PRIORITISED_MODE_CONTROL                      KnxDatapointMainType = 45
	KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_16_BIT              KnxDatapointMainType = 46
	KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_8_BIT               KnxDatapointMainType = 47
	KnxDatapointMainType_DPT_POSITIONS                                     KnxDatapointMainType = 48
	KnxDatapointMainType_DPT_STATUS_32_BIT                                 KnxDatapointMainType = 49
	KnxDatapointMainType_DPT_STATUS_48_BIT                                 KnxDatapointMainType = 50
	KnxDatapointMainType_DPT_CONVERTER_STATUS                              KnxDatapointMainType = 51
	KnxDatapointMainType_DPT_CONVERTER_TEST_RESULT                         KnxDatapointMainType = 52
	KnxDatapointMainType_DPT_BATTERY_INFORMATION                           KnxDatapointMainType = 53
	KnxDatapointMainType_DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION      KnxDatapointMainType = 54
	KnxDatapointMainType_DPT_STATUS_24_BIT                                 KnxDatapointMainType = 55
	KnxDatapointMainType_DPT_COLOUR_RGBW                                   KnxDatapointMainType = 56
	KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGBW                         KnxDatapointMainType = 57
	KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGB                          KnxDatapointMainType = 58
	KnxDatapointMainType_DPT_F32F32                                        KnxDatapointMainType = 59
	KnxDatapointMainType_DPT_F16F16F16F16                                  KnxDatapointMainType = 60
)

var KnxDatapointMainTypeValues []KnxDatapointMainType

func init() {
	_ = errors.New
	KnxDatapointMainTypeValues = []KnxDatapointMainType{
		KnxDatapointMainType_DPT_UNKNOWN,
		KnxDatapointMainType_DPT_64_BIT_SET,
		KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE,
		KnxDatapointMainType_DPT_8_BYTE_SIGNED_VALUE,
		KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE,
		KnxDatapointMainType_DPT_8_BYTE_FLOAT_VALUE,
		KnxDatapointMainType_DPT_1_BIT,
		KnxDatapointMainType_DPT_1_BIT_CONTROLLED,
		KnxDatapointMainType_DPT_3_BIT_CONTROLLED,
		KnxDatapointMainType_DPT_CHARACTER,
		KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE,
		KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE,
		KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE,
		KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE,
		KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE,
		KnxDatapointMainType_DPT_TIME,
		KnxDatapointMainType_DPT_DATE,
		KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE,
		KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE,
		KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE,
		KnxDatapointMainType_DPT_ENTRANCE_ACCESS,
		KnxDatapointMainType_DPT_CHARACTER_STRING,
		KnxDatapointMainType_DPT_SCENE_NUMBER,
		KnxDatapointMainType_DPT_SCENE_CONTROL,
		KnxDatapointMainType_DPT_DATE_TIME,
		KnxDatapointMainType_DPT_1_BYTE,
		KnxDatapointMainType_DPT_8_BIT_SET,
		KnxDatapointMainType_DPT_16_BIT_SET,
		KnxDatapointMainType_DPT_2_BIT_SET,
		KnxDatapointMainType_DPT_2_NIBBLE_SET,
		KnxDatapointMainType_DPT_8_BIT_SET_2,
		KnxDatapointMainType_DPT_32_BIT_SET,
		KnxDatapointMainType_DPT_ELECTRICAL_ENERGY,
		KnxDatapointMainType_DPT_24_TIMES_CHANNEL_ACTIVATION,
		KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM,
		KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM,
		KnxDatapointMainType_DPT_DATAPOINT_TYPE_VERSION,
		KnxDatapointMainType_DPT_ALARM_INFO,
		KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE,
		KnxDatapointMainType_DPT_SCALING_SPEED,
		KnxDatapointMainType_DPT_4_1_1_BYTE_COMBINED_INFORMATION,
		KnxDatapointMainType_DPT_MBUS_ADDRESS,
		KnxDatapointMainType_DPT_3_BYTE_COLOUR_RGB,
		KnxDatapointMainType_DPT_LANGUAGE_CODE_ISO_639_1,
		KnxDatapointMainType_DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY,
		KnxDatapointMainType_DPT_PRIORITISED_MODE_CONTROL,
		KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_16_BIT,
		KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_8_BIT,
		KnxDatapointMainType_DPT_POSITIONS,
		KnxDatapointMainType_DPT_STATUS_32_BIT,
		KnxDatapointMainType_DPT_STATUS_48_BIT,
		KnxDatapointMainType_DPT_CONVERTER_STATUS,
		KnxDatapointMainType_DPT_CONVERTER_TEST_RESULT,
		KnxDatapointMainType_DPT_BATTERY_INFORMATION,
		KnxDatapointMainType_DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION,
		KnxDatapointMainType_DPT_STATUS_24_BIT,
		KnxDatapointMainType_DPT_COLOUR_RGBW,
		KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGBW,
		KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGB,
		KnxDatapointMainType_DPT_F32F32,
		KnxDatapointMainType_DPT_F16F16F16F16,
	}
}

func (e KnxDatapointMainType) Number() uint16 {
	switch e {
	case 0:
		{ /* '0' */
			return 0
		}
	case 1:
		{ /* '1' */
			return 0
		}
	case 10:
		{ /* '10' */
			return 5
		}
	case 11:
		{ /* '11' */
			return 6
		}
	case 12:
		{ /* '12' */
			return 7
		}
	case 13:
		{ /* '13' */
			return 8
		}
	case 14:
		{ /* '14' */
			return 9
		}
	case 15:
		{ /* '15' */
			return 10
		}
	case 16:
		{ /* '16' */
			return 11
		}
	case 17:
		{ /* '17' */
			return 12
		}
	case 18:
		{ /* '18' */
			return 13
		}
	case 19:
		{ /* '19' */
			return 14
		}
	case 2:
		{ /* '2' */
			return 0
		}
	case 20:
		{ /* '20' */
			return 15
		}
	case 21:
		{ /* '21' */
			return 16
		}
	case 22:
		{ /* '22' */
			return 17
		}
	case 23:
		{ /* '23' */
			return 18
		}
	case 24:
		{ /* '24' */
			return 19
		}
	case 25:
		{ /* '25' */
			return 20
		}
	case 26:
		{ /* '26' */
			return 21
		}
	case 27:
		{ /* '27' */
			return 22
		}
	case 28:
		{ /* '28' */
			return 23
		}
	case 29:
		{ /* '29' */
			return 25
		}
	case 3:
		{ /* '3' */
			return 0
		}
	case 30:
		{ /* '30' */
			return 26
		}
	case 31:
		{ /* '31' */
			return 27
		}
	case 32:
		{ /* '32' */
			return 29
		}
	case 33:
		{ /* '33' */
			return 30
		}
	case 34:
		{ /* '34' */
			return 206
		}
	case 35:
		{ /* '35' */
			return 207
		}
	case 36:
		{ /* '36' */
			return 217
		}
	case 37:
		{ /* '37' */
			return 219
		}
	case 38:
		{ /* '38' */
			return 222
		}
	case 39:
		{ /* '39' */
			return 225
		}
	case 4:
		{ /* '4' */
			return 0
		}
	case 40:
		{ /* '40' */
			return 229
		}
	case 41:
		{ /* '41' */
			return 230
		}
	case 42:
		{ /* '42' */
			return 232
		}
	case 43:
		{ /* '43' */
			return 234
		}
	case 44:
		{ /* '44' */
			return 235
		}
	case 45:
		{ /* '45' */
			return 236
		}
	case 46:
		{ /* '46' */
			return 237
		}
	case 47:
		{ /* '47' */
			return 238
		}
	case 48:
		{ /* '48' */
			return 240
		}
	case 49:
		{ /* '49' */
			return 241
		}
	case 5:
		{ /* '5' */
			return 0
		}
	case 50:
		{ /* '50' */
			return 242
		}
	case 51:
		{ /* '51' */
			return 244
		}
	case 52:
		{ /* '52' */
			return 245
		}
	case 53:
		{ /* '53' */
			return 246
		}
	case 54:
		{ /* '54' */
			return 249
		}
	case 55:
		{ /* '55' */
			return 250
		}
	case 56:
		{ /* '56' */
			return 251
		}
	case 57:
		{ /* '57' */
			return 252
		}
	case 58:
		{ /* '58' */
			return 254
		}
	case 59:
		{ /* '59' */
			return 255
		}
	case 6:
		{ /* '6' */
			return 1
		}
	case 60:
		{ /* '60' */
			return 275
		}
	case 7:
		{ /* '7' */
			return 2
		}
	case 8:
		{ /* '8' */
			return 3
		}
	case 9:
		{ /* '9' */
			return 4
		}
	default:
		{
			return 0
		}
	}
}

func KnxDatapointMainTypeFirstEnumForFieldNumber(value uint16) (KnxDatapointMainType, error) {
	for _, sizeValue := range KnxDatapointMainTypeValues {
		if sizeValue.Number() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing Number not found", value)
}

func (e KnxDatapointMainType) Name() string {
	switch e {
	case 0:
		{ /* '0' */
			return "Unknown Datapoint Type"
		}
	case 1:
		{ /* '1' */
			return "Unknown Datapoint Type"
		}
	case 10:
		{ /* '10' */
			return "8-bit unsigned value"
		}
	case 11:
		{ /* '11' */
			return "8-bit signed value"
		}
	case 12:
		{ /* '12' */
			return "2-byte unsigned value"
		}
	case 13:
		{ /* '13' */
			return "2-byte signed value"
		}
	case 14:
		{ /* '14' */
			return "2-byte float value"
		}
	case 15:
		{ /* '15' */
			return "time"
		}
	case 16:
		{ /* '16' */
			return "date"
		}
	case 17:
		{ /* '17' */
			return "4-byte unsigned value"
		}
	case 18:
		{ /* '18' */
			return "4-byte signed value"
		}
	case 19:
		{ /* '19' */
			return "4-byte float value"
		}
	case 2:
		{ /* '2' */
			return "Unknown Datapoint Type"
		}
	case 20:
		{ /* '20' */
			return "entrance access"
		}
	case 21:
		{ /* '21' */
			return "character string"
		}
	case 22:
		{ /* '22' */
			return "scene number"
		}
	case 23:
		{ /* '23' */
			return "scene control"
		}
	case 24:
		{ /* '24' */
			return "Date Time"
		}
	case 25:
		{ /* '25' */
			return "1-byte"
		}
	case 26:
		{ /* '26' */
			return "8-bit set"
		}
	case 27:
		{ /* '27' */
			return "16-bit set"
		}
	case 28:
		{ /* '28' */
			return "2-bit set"
		}
	case 29:
		{ /* '29' */
			return "2-nibble set"
		}
	case 3:
		{ /* '3' */
			return "Unknown Datapoint Type"
		}
	case 30:
		{ /* '30' */
			return "8-bit set"
		}
	case 31:
		{ /* '31' */
			return "32-bit set"
		}
	case 32:
		{ /* '32' */
			return "electrical energy"
		}
	case 33:
		{ /* '33' */
			return "24 times channel activation"
		}
	case 34:
		{ /* '34' */
			return "16-bit unsigned value & 8-bit enum"
		}
	case 35:
		{ /* '35' */
			return "8-bit unsigned value & 8-bit enum"
		}
	case 36:
		{ /* '36' */
			return "datapoint type version"
		}
	case 37:
		{ /* '37' */
			return "alarm info"
		}
	case 38:
		{ /* '38' */
			return "3x 2-byte float value"
		}
	case 39:
		{ /* '39' */
			return "scaling speed"
		}
	case 4:
		{ /* '4' */
			return "Unknown Datapoint Type"
		}
	case 40:
		{ /* '40' */
			return "4-1-1 byte combined information"
		}
	case 41:
		{ /* '41' */
			return "MBus address"
		}
	case 42:
		{ /* '42' */
			return "3-byte colour RGB"
		}
	case 43:
		{ /* '43' */
			return "language code ISO 639-1"
		}
	case 44:
		{ /* '44' */
			return "Signed value with classification and validity"
		}
	case 45:
		{ /* '45' */
			return "Prioritised Mode Control"
		}
	case 46:
		{ /* '46' */
			return "configuration/ diagnostics"
		}
	case 47:
		{ /* '47' */
			return "configuration/ diagnostics"
		}
	case 48:
		{ /* '48' */
			return "positions"
		}
	case 49:
		{ /* '49' */
			return "status"
		}
	case 5:
		{ /* '5' */
			return "Unknown Datapoint Type"
		}
	case 50:
		{ /* '50' */
			return "status"
		}
	case 51:
		{ /* '51' */
			return "Converter Status"
		}
	case 52:
		{ /* '52' */
			return "Converter test result"
		}
	case 53:
		{ /* '53' */
			return "Battery Information"
		}
	case 54:
		{ /* '54' */
			return "brightness colour temperature transition"
		}
	case 55:
		{ /* '55' */
			return "status"
		}
	case 56:
		{ /* '56' */
			return "Colour RGBW"
		}
	case 57:
		{ /* '57' */
			return "Relative Control RGBW"
		}
	case 58:
		{ /* '58' */
			return "Relative Control RGB"
		}
	case 59:
		{ /* '59' */
			return "F32F32"
		}
	case 6:
		{ /* '6' */
			return "1-bit"
		}
	case 60:
		{ /* '60' */
			return "F16F16F16F16"
		}
	case 7:
		{ /* '7' */
			return "1-bit controlled"
		}
	case 8:
		{ /* '8' */
			return "3-bit controlled"
		}
	case 9:
		{ /* '9' */
			return "character"
		}
	default:
		{
			return ""
		}
	}
}

func KnxDatapointMainTypeFirstEnumForFieldName(value string) (KnxDatapointMainType, error) {
	for _, sizeValue := range KnxDatapointMainTypeValues {
		if sizeValue.Name() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing Name not found", value)
}

func (e KnxDatapointMainType) SizeInBits() uint8 {
	switch e {
	case 0:
		{ /* '0' */
			return 0
		}
	case 1:
		{ /* '1' */
			return 64
		}
	case 10:
		{ /* '10' */
			return 8
		}
	case 11:
		{ /* '11' */
			return 8
		}
	case 12:
		{ /* '12' */
			return 16
		}
	case 13:
		{ /* '13' */
			return 16
		}
	case 14:
		{ /* '14' */
			return 16
		}
	case 15:
		{ /* '15' */
			return 24
		}
	case 16:
		{ /* '16' */
			return 24
		}
	case 17:
		{ /* '17' */
			return 32
		}
	case 18:
		{ /* '18' */
			return 32
		}
	case 19:
		{ /* '19' */
			return 32
		}
	case 2:
		{ /* '2' */
			return 64
		}
	case 20:
		{ /* '20' */
			return 32
		}
	case 21:
		{ /* '21' */
			return 112
		}
	case 22:
		{ /* '22' */
			return 8
		}
	case 23:
		{ /* '23' */
			return 8
		}
	case 24:
		{ /* '24' */
			return 64
		}
	case 25:
		{ /* '25' */
			return 8
		}
	case 26:
		{ /* '26' */
			return 8
		}
	case 27:
		{ /* '27' */
			return 16
		}
	case 28:
		{ /* '28' */
			return 2
		}
	case 29:
		{ /* '29' */
			return 8
		}
	case 3:
		{ /* '3' */
			return 64
		}
	case 30:
		{ /* '30' */
			return 8
		}
	case 31:
		{ /* '31' */
			return 32
		}
	case 32:
		{ /* '32' */
			return 64
		}
	case 33:
		{ /* '33' */
			return 24
		}
	case 34:
		{ /* '34' */
			return 24
		}
	case 35:
		{ /* '35' */
			return 16
		}
	case 36:
		{ /* '36' */
			return 16
		}
	case 37:
		{ /* '37' */
			return 48
		}
	case 38:
		{ /* '38' */
			return 48
		}
	case 39:
		{ /* '39' */
			return 24
		}
	case 4:
		{ /* '4' */
			return 96
		}
	case 40:
		{ /* '40' */
			return 48
		}
	case 41:
		{ /* '41' */
			return 64
		}
	case 42:
		{ /* '42' */
			return 24
		}
	case 43:
		{ /* '43' */
			return 16
		}
	case 44:
		{ /* '44' */
			return 48
		}
	case 45:
		{ /* '45' */
			return 8
		}
	case 46:
		{ /* '46' */
			return 16
		}
	case 47:
		{ /* '47' */
			return 8
		}
	case 48:
		{ /* '48' */
			return 24
		}
	case 49:
		{ /* '49' */
			return 32
		}
	case 5:
		{ /* '5' */
			return 64
		}
	case 50:
		{ /* '50' */
			return 48
		}
	case 51:
		{ /* '51' */
			return 16
		}
	case 52:
		{ /* '52' */
			return 48
		}
	case 53:
		{ /* '53' */
			return 16
		}
	case 54:
		{ /* '54' */
			return 48
		}
	case 55:
		{ /* '55' */
			return 24
		}
	case 56:
		{ /* '56' */
			return 48
		}
	case 57:
		{ /* '57' */
			return 40
		}
	case 58:
		{ /* '58' */
			return 24
		}
	case 59:
		{ /* '59' */
			return 64
		}
	case 6:
		{ /* '6' */
			return 1
		}
	case 60:
		{ /* '60' */
			return 64
		}
	case 7:
		{ /* '7' */
			return 2
		}
	case 8:
		{ /* '8' */
			return 4
		}
	case 9:
		{ /* '9' */
			return 8
		}
	default:
		{
			return 0
		}
	}
}

func KnxDatapointMainTypeFirstEnumForFieldSizeInBits(value uint8) (KnxDatapointMainType, error) {
	for _, sizeValue := range KnxDatapointMainTypeValues {
		if sizeValue.SizeInBits() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing SizeInBits not found", value)
}
func KnxDatapointMainTypeByValue(value uint16) KnxDatapointMainType {
	switch value {
	case 0:
		return KnxDatapointMainType_DPT_UNKNOWN
	case 1:
		return KnxDatapointMainType_DPT_64_BIT_SET
	case 10:
		return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
	case 11:
		return KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE
	case 12:
		return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
	case 13:
		return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
	case 14:
		return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
	case 15:
		return KnxDatapointMainType_DPT_TIME
	case 16:
		return KnxDatapointMainType_DPT_DATE
	case 17:
		return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
	case 18:
		return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
	case 19:
		return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
	case 2:
		return KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE
	case 20:
		return KnxDatapointMainType_DPT_ENTRANCE_ACCESS
	case 21:
		return KnxDatapointMainType_DPT_CHARACTER_STRING
	case 22:
		return KnxDatapointMainType_DPT_SCENE_NUMBER
	case 23:
		return KnxDatapointMainType_DPT_SCENE_CONTROL
	case 24:
		return KnxDatapointMainType_DPT_DATE_TIME
	case 25:
		return KnxDatapointMainType_DPT_1_BYTE
	case 26:
		return KnxDatapointMainType_DPT_8_BIT_SET
	case 27:
		return KnxDatapointMainType_DPT_16_BIT_SET
	case 28:
		return KnxDatapointMainType_DPT_2_BIT_SET
	case 29:
		return KnxDatapointMainType_DPT_2_NIBBLE_SET
	case 3:
		return KnxDatapointMainType_DPT_8_BYTE_SIGNED_VALUE
	case 30:
		return KnxDatapointMainType_DPT_8_BIT_SET_2
	case 31:
		return KnxDatapointMainType_DPT_32_BIT_SET
	case 32:
		return KnxDatapointMainType_DPT_ELECTRICAL_ENERGY
	case 33:
		return KnxDatapointMainType_DPT_24_TIMES_CHANNEL_ACTIVATION
	case 34:
		return KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
	case 35:
		return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
	case 36:
		return KnxDatapointMainType_DPT_DATAPOINT_TYPE_VERSION
	case 37:
		return KnxDatapointMainType_DPT_ALARM_INFO
	case 38:
		return KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE
	case 39:
		return KnxDatapointMainType_DPT_SCALING_SPEED
	case 4:
		return KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE
	case 40:
		return KnxDatapointMainType_DPT_4_1_1_BYTE_COMBINED_INFORMATION
	case 41:
		return KnxDatapointMainType_DPT_MBUS_ADDRESS
	case 42:
		return KnxDatapointMainType_DPT_3_BYTE_COLOUR_RGB
	case 43:
		return KnxDatapointMainType_DPT_LANGUAGE_CODE_ISO_639_1
	case 44:
		return KnxDatapointMainType_DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY
	case 45:
		return KnxDatapointMainType_DPT_PRIORITISED_MODE_CONTROL
	case 46:
		return KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_16_BIT
	case 47:
		return KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_8_BIT
	case 48:
		return KnxDatapointMainType_DPT_POSITIONS
	case 49:
		return KnxDatapointMainType_DPT_STATUS_32_BIT
	case 5:
		return KnxDatapointMainType_DPT_8_BYTE_FLOAT_VALUE
	case 50:
		return KnxDatapointMainType_DPT_STATUS_48_BIT
	case 51:
		return KnxDatapointMainType_DPT_CONVERTER_STATUS
	case 52:
		return KnxDatapointMainType_DPT_CONVERTER_TEST_RESULT
	case 53:
		return KnxDatapointMainType_DPT_BATTERY_INFORMATION
	case 54:
		return KnxDatapointMainType_DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION
	case 55:
		return KnxDatapointMainType_DPT_STATUS_24_BIT
	case 56:
		return KnxDatapointMainType_DPT_COLOUR_RGBW
	case 57:
		return KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGBW
	case 58:
		return KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGB
	case 59:
		return KnxDatapointMainType_DPT_F32F32
	case 6:
		return KnxDatapointMainType_DPT_1_BIT
	case 60:
		return KnxDatapointMainType_DPT_F16F16F16F16
	case 7:
		return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
	case 8:
		return KnxDatapointMainType_DPT_3_BIT_CONTROLLED
	case 9:
		return KnxDatapointMainType_DPT_CHARACTER
	}
	return 0
}

func KnxDatapointMainTypeByName(value string) (enum KnxDatapointMainType, ok bool) {
	ok = true
	switch value {
	case "DPT_UNKNOWN":
		enum = KnxDatapointMainType_DPT_UNKNOWN
	case "DPT_64_BIT_SET":
		enum = KnxDatapointMainType_DPT_64_BIT_SET
	case "DPT_8_BIT_UNSIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
	case "DPT_8_BIT_SIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE
	case "DPT_2_BYTE_UNSIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
	case "DPT_2_BYTE_SIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
	case "DPT_2_BYTE_FLOAT_VALUE":
		enum = KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
	case "DPT_TIME":
		enum = KnxDatapointMainType_DPT_TIME
	case "DPT_DATE":
		enum = KnxDatapointMainType_DPT_DATE
	case "DPT_4_BYTE_UNSIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
	case "DPT_4_BYTE_SIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
	case "DPT_4_BYTE_FLOAT_VALUE":
		enum = KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
	case "DPT_8_BYTE_UNSIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE
	case "DPT_ENTRANCE_ACCESS":
		enum = KnxDatapointMainType_DPT_ENTRANCE_ACCESS
	case "DPT_CHARACTER_STRING":
		enum = KnxDatapointMainType_DPT_CHARACTER_STRING
	case "DPT_SCENE_NUMBER":
		enum = KnxDatapointMainType_DPT_SCENE_NUMBER
	case "DPT_SCENE_CONTROL":
		enum = KnxDatapointMainType_DPT_SCENE_CONTROL
	case "DPT_DATE_TIME":
		enum = KnxDatapointMainType_DPT_DATE_TIME
	case "DPT_1_BYTE":
		enum = KnxDatapointMainType_DPT_1_BYTE
	case "DPT_8_BIT_SET":
		enum = KnxDatapointMainType_DPT_8_BIT_SET
	case "DPT_16_BIT_SET":
		enum = KnxDatapointMainType_DPT_16_BIT_SET
	case "DPT_2_BIT_SET":
		enum = KnxDatapointMainType_DPT_2_BIT_SET
	case "DPT_2_NIBBLE_SET":
		enum = KnxDatapointMainType_DPT_2_NIBBLE_SET
	case "DPT_8_BYTE_SIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_8_BYTE_SIGNED_VALUE
	case "DPT_8_BIT_SET_2":
		enum = KnxDatapointMainType_DPT_8_BIT_SET_2
	case "DPT_32_BIT_SET":
		enum = KnxDatapointMainType_DPT_32_BIT_SET
	case "DPT_ELECTRICAL_ENERGY":
		enum = KnxDatapointMainType_DPT_ELECTRICAL_ENERGY
	case "DPT_24_TIMES_CHANNEL_ACTIVATION":
		enum = KnxDatapointMainType_DPT_24_TIMES_CHANNEL_ACTIVATION
	case "DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM":
		enum = KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
	case "DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM":
		enum = KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
	case "DPT_DATAPOINT_TYPE_VERSION":
		enum = KnxDatapointMainType_DPT_DATAPOINT_TYPE_VERSION
	case "DPT_ALARM_INFO":
		enum = KnxDatapointMainType_DPT_ALARM_INFO
	case "DPT_3X_2_BYTE_FLOAT_VALUE":
		enum = KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE
	case "DPT_SCALING_SPEED":
		enum = KnxDatapointMainType_DPT_SCALING_SPEED
	case "DPT_12_BYTE_SIGNED_VALUE":
		enum = KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE
	case "DPT_4_1_1_BYTE_COMBINED_INFORMATION":
		enum = KnxDatapointMainType_DPT_4_1_1_BYTE_COMBINED_INFORMATION
	case "DPT_MBUS_ADDRESS":
		enum = KnxDatapointMainType_DPT_MBUS_ADDRESS
	case "DPT_3_BYTE_COLOUR_RGB":
		enum = KnxDatapointMainType_DPT_3_BYTE_COLOUR_RGB
	case "DPT_LANGUAGE_CODE_ISO_639_1":
		enum = KnxDatapointMainType_DPT_LANGUAGE_CODE_ISO_639_1
	case "DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY":
		enum = KnxDatapointMainType_DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY
	case "DPT_PRIORITISED_MODE_CONTROL":
		enum = KnxDatapointMainType_DPT_PRIORITISED_MODE_CONTROL
	case "DPT_CONFIGURATION_DIAGNOSTICS_16_BIT":
		enum = KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_16_BIT
	case "DPT_CONFIGURATION_DIAGNOSTICS_8_BIT":
		enum = KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_8_BIT
	case "DPT_POSITIONS":
		enum = KnxDatapointMainType_DPT_POSITIONS
	case "DPT_STATUS_32_BIT":
		enum = KnxDatapointMainType_DPT_STATUS_32_BIT
	case "DPT_8_BYTE_FLOAT_VALUE":
		enum = KnxDatapointMainType_DPT_8_BYTE_FLOAT_VALUE
	case "DPT_STATUS_48_BIT":
		enum = KnxDatapointMainType_DPT_STATUS_48_BIT
	case "DPT_CONVERTER_STATUS":
		enum = KnxDatapointMainType_DPT_CONVERTER_STATUS
	case "DPT_CONVERTER_TEST_RESULT":
		enum = KnxDatapointMainType_DPT_CONVERTER_TEST_RESULT
	case "DPT_BATTERY_INFORMATION":
		enum = KnxDatapointMainType_DPT_BATTERY_INFORMATION
	case "DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION":
		enum = KnxDatapointMainType_DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION
	case "DPT_STATUS_24_BIT":
		enum = KnxDatapointMainType_DPT_STATUS_24_BIT
	case "DPT_COLOUR_RGBW":
		enum = KnxDatapointMainType_DPT_COLOUR_RGBW
	case "DPT_RELATIVE_CONTROL_RGBW":
		enum = KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGBW
	case "DPT_RELATIVE_CONTROL_RGB":
		enum = KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGB
	case "DPT_F32F32":
		enum = KnxDatapointMainType_DPT_F32F32
	case "DPT_1_BIT":
		enum = KnxDatapointMainType_DPT_1_BIT
	case "DPT_F16F16F16F16":
		enum = KnxDatapointMainType_DPT_F16F16F16F16
	case "DPT_1_BIT_CONTROLLED":
		enum = KnxDatapointMainType_DPT_1_BIT_CONTROLLED
	case "DPT_3_BIT_CONTROLLED":
		enum = KnxDatapointMainType_DPT_3_BIT_CONTROLLED
	case "DPT_CHARACTER":
		enum = KnxDatapointMainType_DPT_CHARACTER
	default:
		enum = 0
		ok = false
	}
	return
}

func KnxDatapointMainTypeKnows(value uint16) bool {
	for _, typeValue := range KnxDatapointMainTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastKnxDatapointMainType(structType interface{}) KnxDatapointMainType {
	castFunc := func(typ interface{}) KnxDatapointMainType {
		if sKnxDatapointMainType, ok := typ.(KnxDatapointMainType); ok {
			return sKnxDatapointMainType
		}
		return 0
	}
	return castFunc(structType)
}

func (m KnxDatapointMainType) GetLengthInBits() uint16 {
	return 16
}

func (m KnxDatapointMainType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxDatapointMainTypeParse(readBuffer utils.ReadBuffer) (KnxDatapointMainType, error) {
	val, err := readBuffer.ReadUint16("KnxDatapointMainType", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading KnxDatapointMainType")
	}
	return KnxDatapointMainTypeByValue(val), nil
}

func (e KnxDatapointMainType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("KnxDatapointMainType", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e KnxDatapointMainType) PLC4XEnumName() string {
	switch e {
	case KnxDatapointMainType_DPT_UNKNOWN:
		return "DPT_UNKNOWN"
	case KnxDatapointMainType_DPT_64_BIT_SET:
		return "DPT_64_BIT_SET"
	case KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE:
		return "DPT_8_BIT_UNSIGNED_VALUE"
	case KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE:
		return "DPT_8_BIT_SIGNED_VALUE"
	case KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE:
		return "DPT_2_BYTE_UNSIGNED_VALUE"
	case KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE:
		return "DPT_2_BYTE_SIGNED_VALUE"
	case KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE:
		return "DPT_2_BYTE_FLOAT_VALUE"
	case KnxDatapointMainType_DPT_TIME:
		return "DPT_TIME"
	case KnxDatapointMainType_DPT_DATE:
		return "DPT_DATE"
	case KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE:
		return "DPT_4_BYTE_UNSIGNED_VALUE"
	case KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE:
		return "DPT_4_BYTE_SIGNED_VALUE"
	case KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE:
		return "DPT_4_BYTE_FLOAT_VALUE"
	case KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE:
		return "DPT_8_BYTE_UNSIGNED_VALUE"
	case KnxDatapointMainType_DPT_ENTRANCE_ACCESS:
		return "DPT_ENTRANCE_ACCESS"
	case KnxDatapointMainType_DPT_CHARACTER_STRING:
		return "DPT_CHARACTER_STRING"
	case KnxDatapointMainType_DPT_SCENE_NUMBER:
		return "DPT_SCENE_NUMBER"
	case KnxDatapointMainType_DPT_SCENE_CONTROL:
		return "DPT_SCENE_CONTROL"
	case KnxDatapointMainType_DPT_DATE_TIME:
		return "DPT_DATE_TIME"
	case KnxDatapointMainType_DPT_1_BYTE:
		return "DPT_1_BYTE"
	case KnxDatapointMainType_DPT_8_BIT_SET:
		return "DPT_8_BIT_SET"
	case KnxDatapointMainType_DPT_16_BIT_SET:
		return "DPT_16_BIT_SET"
	case KnxDatapointMainType_DPT_2_BIT_SET:
		return "DPT_2_BIT_SET"
	case KnxDatapointMainType_DPT_2_NIBBLE_SET:
		return "DPT_2_NIBBLE_SET"
	case KnxDatapointMainType_DPT_8_BYTE_SIGNED_VALUE:
		return "DPT_8_BYTE_SIGNED_VALUE"
	case KnxDatapointMainType_DPT_8_BIT_SET_2:
		return "DPT_8_BIT_SET_2"
	case KnxDatapointMainType_DPT_32_BIT_SET:
		return "DPT_32_BIT_SET"
	case KnxDatapointMainType_DPT_ELECTRICAL_ENERGY:
		return "DPT_ELECTRICAL_ENERGY"
	case KnxDatapointMainType_DPT_24_TIMES_CHANNEL_ACTIVATION:
		return "DPT_24_TIMES_CHANNEL_ACTIVATION"
	case KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM:
		return "DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM"
	case KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM:
		return "DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM"
	case KnxDatapointMainType_DPT_DATAPOINT_TYPE_VERSION:
		return "DPT_DATAPOINT_TYPE_VERSION"
	case KnxDatapointMainType_DPT_ALARM_INFO:
		return "DPT_ALARM_INFO"
	case KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE:
		return "DPT_3X_2_BYTE_FLOAT_VALUE"
	case KnxDatapointMainType_DPT_SCALING_SPEED:
		return "DPT_SCALING_SPEED"
	case KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE:
		return "DPT_12_BYTE_SIGNED_VALUE"
	case KnxDatapointMainType_DPT_4_1_1_BYTE_COMBINED_INFORMATION:
		return "DPT_4_1_1_BYTE_COMBINED_INFORMATION"
	case KnxDatapointMainType_DPT_MBUS_ADDRESS:
		return "DPT_MBUS_ADDRESS"
	case KnxDatapointMainType_DPT_3_BYTE_COLOUR_RGB:
		return "DPT_3_BYTE_COLOUR_RGB"
	case KnxDatapointMainType_DPT_LANGUAGE_CODE_ISO_639_1:
		return "DPT_LANGUAGE_CODE_ISO_639_1"
	case KnxDatapointMainType_DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY:
		return "DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY"
	case KnxDatapointMainType_DPT_PRIORITISED_MODE_CONTROL:
		return "DPT_PRIORITISED_MODE_CONTROL"
	case KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_16_BIT:
		return "DPT_CONFIGURATION_DIAGNOSTICS_16_BIT"
	case KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_8_BIT:
		return "DPT_CONFIGURATION_DIAGNOSTICS_8_BIT"
	case KnxDatapointMainType_DPT_POSITIONS:
		return "DPT_POSITIONS"
	case KnxDatapointMainType_DPT_STATUS_32_BIT:
		return "DPT_STATUS_32_BIT"
	case KnxDatapointMainType_DPT_8_BYTE_FLOAT_VALUE:
		return "DPT_8_BYTE_FLOAT_VALUE"
	case KnxDatapointMainType_DPT_STATUS_48_BIT:
		return "DPT_STATUS_48_BIT"
	case KnxDatapointMainType_DPT_CONVERTER_STATUS:
		return "DPT_CONVERTER_STATUS"
	case KnxDatapointMainType_DPT_CONVERTER_TEST_RESULT:
		return "DPT_CONVERTER_TEST_RESULT"
	case KnxDatapointMainType_DPT_BATTERY_INFORMATION:
		return "DPT_BATTERY_INFORMATION"
	case KnxDatapointMainType_DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION:
		return "DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION"
	case KnxDatapointMainType_DPT_STATUS_24_BIT:
		return "DPT_STATUS_24_BIT"
	case KnxDatapointMainType_DPT_COLOUR_RGBW:
		return "DPT_COLOUR_RGBW"
	case KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGBW:
		return "DPT_RELATIVE_CONTROL_RGBW"
	case KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGB:
		return "DPT_RELATIVE_CONTROL_RGB"
	case KnxDatapointMainType_DPT_F32F32:
		return "DPT_F32F32"
	case KnxDatapointMainType_DPT_1_BIT:
		return "DPT_1_BIT"
	case KnxDatapointMainType_DPT_F16F16F16F16:
		return "DPT_F16F16F16F16"
	case KnxDatapointMainType_DPT_1_BIT_CONTROLLED:
		return "DPT_1_BIT_CONTROLLED"
	case KnxDatapointMainType_DPT_3_BIT_CONTROLLED:
		return "DPT_3_BIT_CONTROLLED"
	case KnxDatapointMainType_DPT_CHARACTER:
		return "DPT_CHARACTER"
	}
	return ""
}

func (e KnxDatapointMainType) String() string {
	return e.PLC4XEnumName()
}
