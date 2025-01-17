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

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum TelephonyCommandTypeContainer {
  TelephonyCommandLineOnHook((short) 0x09, (short) 1, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_0Bytes((short) 0xA0, (short) 0, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_1Bytes((short) 0xA1, (short) 1, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_2Bytes((short) 0xA2, (short) 2, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_3Bytes((short) 0xA3, (short) 3, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_4Bytes((short) 0xA4, (short) 4, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_5Bytes((short) 0xA5, (short) 5, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_6Bytes((short) 0xA6, (short) 6, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_7Bytes((short) 0xA7, (short) 7, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_8Bytes((short) 0xA8, (short) 8, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_9Bytes((short) 0xA9, (short) 9, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_10Bytes((short) 0xAA, (short) 10, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_11Bytes((short) 0xAB, (short) 11, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_12Bytes((short) 0xAC, (short) 12, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_13Bytes((short) 0xAD, (short) 13, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_14Bytes((short) 0xAE, (short) 14, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_15Bytes((short) 0xAF, (short) 15, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_16Bytes((short) 0xB0, (short) 16, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_17Bytes((short) 0xB1, (short) 17, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_18Bytes((short) 0xB2, (short) 18, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_19Bytes((short) 0xB3, (short) 19, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_20Bytes((short) 0xB4, (short) 20, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_21Bytes((short) 0xB5, (short) 21, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_22Bytes((short) 0xB6, (short) 22, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_23Bytes((short) 0xB7, (short) 23, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_24Bytes((short) 0xB8, (short) 24, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_25Bytes((short) 0xB9, (short) 25, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_26Bytes((short) 0xBA, (short) 26, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_27Bytes((short) 0xBB, (short) 27, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_28Bytes((short) 0xBC, (short) 28, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_29Bytes((short) 0xBD, (short) 29, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_30Bytes((short) 0xBE, (short) 30, TelephonyCommandType.EVENT),
  TelephonyCommandLineOffHook_31Bytes((short) 0xBF, (short) 31, TelephonyCommandType.EVENT);
  private static final Map<Short, TelephonyCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (TelephonyCommandTypeContainer value : TelephonyCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final short numBytes;
  private final TelephonyCommandType commandType;

  TelephonyCommandTypeContainer(short value, short numBytes, TelephonyCommandType commandType) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
  }

  public short getValue() {
    return value;
  }

  public short getNumBytes() {
    return numBytes;
  }

  public static TelephonyCommandTypeContainer firstEnumForFieldNumBytes(short fieldValue) {
    for (TelephonyCommandTypeContainer _val : TelephonyCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TelephonyCommandTypeContainer> enumsForFieldNumBytes(short fieldValue) {
    List<TelephonyCommandTypeContainer> _values = new ArrayList<>();
    for (TelephonyCommandTypeContainer _val : TelephonyCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public TelephonyCommandType getCommandType() {
    return commandType;
  }

  public static TelephonyCommandTypeContainer firstEnumForFieldCommandType(
      TelephonyCommandType fieldValue) {
    for (TelephonyCommandTypeContainer _val : TelephonyCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<TelephonyCommandTypeContainer> enumsForFieldCommandType(
      TelephonyCommandType fieldValue) {
    List<TelephonyCommandTypeContainer> _values = new ArrayList<>();
    for (TelephonyCommandTypeContainer _val : TelephonyCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static TelephonyCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
