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
package org.apache.plc4x.java.cbus.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum HVACHumidityError {
  NO_ERROR((short) 0x00),
  HUMIDIFIER_TOTAL_FAILURE((short) 0x01),
  DEHUMIDIFIER_TOTAL_FAILURE((short) 0x02),
  FAN_TOTAL_FAILURE((short) 0x03),
  HUMIDITY_SENSOR_FAILURE((short) 0x04),
  HUMIDIFIER_TEMPORARY_PROBLEM((short) 0x05),
  DEHUMIDIFIER_TEMPORARY_PROBLEM((short) 0x06),
  FAN_TEMPORARY_PROBLEM((short) 0x07),
  HUMIDIFIER_SERVICE_REQUIRED((short) 0x08),
  DEHUMIDIFIER_SERVICE_REQUIRED((short) 0x09),
  FAN_SERVICE_REQUIRED((short) 0x0A),
  FILTER_REPLACEMENT_REQUIRED((short) 0x0B),
  CUSTOM_ERROR_0((short) 0x80),
  CUSTOM_ERROR_1((short) 0x81),
  CUSTOM_ERROR_2((short) 0x82),
  CUSTOM_ERROR_3((short) 0x83),
  CUSTOM_ERROR_4((short) 0x84),
  CUSTOM_ERROR_5((short) 0x85),
  CUSTOM_ERROR_6((short) 0x86),
  CUSTOM_ERROR_7((short) 0x87),
  CUSTOM_ERROR_8((short) 0x88),
  CUSTOM_ERROR_9((short) 0x89),
  CUSTOM_ERROR_10((short) 0x8A),
  CUSTOM_ERROR_11((short) 0x8B),
  CUSTOM_ERROR_12((short) 0x8C),
  CUSTOM_ERROR_13((short) 0x8D),
  CUSTOM_ERROR_14((short) 0x8E),
  CUSTOM_ERROR_15((short) 0x8F),
  CUSTOM_ERROR_16((short) 0x90),
  CUSTOM_ERROR_17((short) 0x91),
  CUSTOM_ERROR_18((short) 0x92),
  CUSTOM_ERROR_19((short) 0x93),
  CUSTOM_ERROR_20((short) 0x94),
  CUSTOM_ERROR_21((short) 0x95),
  CUSTOM_ERROR_22((short) 0x96),
  CUSTOM_ERROR_23((short) 0x97),
  CUSTOM_ERROR_24((short) 0x98),
  CUSTOM_ERROR_25((short) 0x99),
  CUSTOM_ERROR_26((short) 0x9A),
  CUSTOM_ERROR_27((short) 0x9B),
  CUSTOM_ERROR_28((short) 0x9C),
  CUSTOM_ERROR_29((short) 0x9D),
  CUSTOM_ERROR_30((short) 0x9E),
  CUSTOM_ERROR_31((short) 0x9F),
  CUSTOM_ERROR_32((short) 0xA0),
  CUSTOM_ERROR_33((short) 0xA1),
  CUSTOM_ERROR_34((short) 0xA2),
  CUSTOM_ERROR_35((short) 0xA3),
  CUSTOM_ERROR_36((short) 0xA4),
  CUSTOM_ERROR_37((short) 0xA5),
  CUSTOM_ERROR_38((short) 0xA6),
  CUSTOM_ERROR_39((short) 0xA7),
  CUSTOM_ERROR_40((short) 0xA8),
  CUSTOM_ERROR_41((short) 0xA9),
  CUSTOM_ERROR_42((short) 0xAA),
  CUSTOM_ERROR_43((short) 0xAB),
  CUSTOM_ERROR_44((short) 0xAC),
  CUSTOM_ERROR_45((short) 0xAD),
  CUSTOM_ERROR_46((short) 0xAE),
  CUSTOM_ERROR_47((short) 0xAF),
  CUSTOM_ERROR_48((short) 0xB0),
  CUSTOM_ERROR_49((short) 0xB1),
  CUSTOM_ERROR_50((short) 0xB2),
  CUSTOM_ERROR_51((short) 0xB3),
  CUSTOM_ERROR_52((short) 0xB4),
  CUSTOM_ERROR_53((short) 0xB5),
  CUSTOM_ERROR_54((short) 0xB6),
  CUSTOM_ERROR_55((short) 0xB7),
  CUSTOM_ERROR_56((short) 0xB8),
  CUSTOM_ERROR_57((short) 0xB9),
  CUSTOM_ERROR_58((short) 0xBA),
  CUSTOM_ERROR_59((short) 0xBB),
  CUSTOM_ERROR_60((short) 0xBC),
  CUSTOM_ERROR_61((short) 0xBD),
  CUSTOM_ERROR_62((short) 0xBE),
  CUSTOM_ERROR_63((short) 0xBF),
  CUSTOM_ERROR_64((short) 0xC0),
  CUSTOM_ERROR_65((short) 0xC1),
  CUSTOM_ERROR_66((short) 0xC2),
  CUSTOM_ERROR_67((short) 0xC3),
  CUSTOM_ERROR_68((short) 0xC4),
  CUSTOM_ERROR_69((short) 0xC5),
  CUSTOM_ERROR_70((short) 0xC6),
  CUSTOM_ERROR_71((short) 0xC7),
  CUSTOM_ERROR_72((short) 0xC8),
  CUSTOM_ERROR_73((short) 0xC9),
  CUSTOM_ERROR_74((short) 0xCA),
  CUSTOM_ERROR_75((short) 0xCB),
  CUSTOM_ERROR_76((short) 0xCC),
  CUSTOM_ERROR_77((short) 0xCD),
  CUSTOM_ERROR_78((short) 0xCE),
  CUSTOM_ERROR_79((short) 0xCF),
  CUSTOM_ERROR_80((short) 0xD0),
  CUSTOM_ERROR_81((short) 0xD1),
  CUSTOM_ERROR_82((short) 0xD2),
  CUSTOM_ERROR_83((short) 0xD3),
  CUSTOM_ERROR_84((short) 0xD4),
  CUSTOM_ERROR_85((short) 0xD5),
  CUSTOM_ERROR_86((short) 0xD6),
  CUSTOM_ERROR_87((short) 0xD7),
  CUSTOM_ERROR_88((short) 0xD8),
  CUSTOM_ERROR_89((short) 0xD9),
  CUSTOM_ERROR_90((short) 0xDA),
  CUSTOM_ERROR_91((short) 0xDB),
  CUSTOM_ERROR_92((short) 0xDC),
  CUSTOM_ERROR_93((short) 0xDD),
  CUSTOM_ERROR_94((short) 0xDE),
  CUSTOM_ERROR_95((short) 0xDF),
  CUSTOM_ERROR_96((short) 0xE0),
  CUSTOM_ERROR_97((short) 0xE1),
  CUSTOM_ERROR_98((short) 0xE2),
  CUSTOM_ERROR_99((short) 0xE3),
  CUSTOM_ERROR_100((short) 0xE4),
  CUSTOM_ERROR_101((short) 0xE5),
  CUSTOM_ERROR_102((short) 0xE6),
  CUSTOM_ERROR_103((short) 0xE7),
  CUSTOM_ERROR_104((short) 0xE8),
  CUSTOM_ERROR_105((short) 0xE9),
  CUSTOM_ERROR_106((short) 0xEA),
  CUSTOM_ERROR_107((short) 0xEB),
  CUSTOM_ERROR_108((short) 0xEC),
  CUSTOM_ERROR_109((short) 0xED),
  CUSTOM_ERROR_110((short) 0xEE),
  CUSTOM_ERROR_111((short) 0xEF),
  CUSTOM_ERROR_112((short) 0xF0),
  CUSTOM_ERROR_113((short) 0xF1),
  CUSTOM_ERROR_114((short) 0xF2),
  CUSTOM_ERROR_115((short) 0xF3),
  CUSTOM_ERROR_116((short) 0xF4),
  CUSTOM_ERROR_117((short) 0xF5),
  CUSTOM_ERROR_118((short) 0xF6),
  CUSTOM_ERROR_119((short) 0xF7),
  CUSTOM_ERROR_120((short) 0xF8),
  CUSTOM_ERROR_121((short) 0xF9),
  CUSTOM_ERROR_122((short) 0xFA),
  CUSTOM_ERROR_123((short) 0xFB),
  CUSTOM_ERROR_124((short) 0xFC),
  CUSTOM_ERROR_125((short) 0xFD),
  CUSTOM_ERROR_126((short) 0xFE),
  CUSTOM_ERROR_127((short) 0xFF);
  private static final Map<Short, HVACHumidityError> map;

  static {
    map = new HashMap<>();
    for (HVACHumidityError value : HVACHumidityError.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  HVACHumidityError(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static HVACHumidityError enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
