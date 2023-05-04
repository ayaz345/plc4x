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

public enum MeteringCommandType {
  EVENT((byte) 0x00, (short) 0xFF);
  private static final Map<Byte, MeteringCommandType> map;

  static {
    map = new HashMap<>();
    for (MeteringCommandType value : MeteringCommandType.values()) {
      map.put((byte) value.getValue(), value);
    }
  }

  private final byte value;
  private short numberOfArguments;

  MeteringCommandType(byte value, short numberOfArguments) {
    this.value = value;
    this.numberOfArguments = numberOfArguments;
  }

  public byte getValue() {
    return value;
  }

  public short getNumberOfArguments() {
    return numberOfArguments;
  }

  public static MeteringCommandType firstEnumForFieldNumberOfArguments(short fieldValue) {
    for (MeteringCommandType _val : MeteringCommandType.values()) {
      if (_val.getNumberOfArguments() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<MeteringCommandType> enumsForFieldNumberOfArguments(short fieldValue) {
    List<MeteringCommandType> _values = new ArrayList();
    for (MeteringCommandType _val : MeteringCommandType.values()) {
      if (_val.getNumberOfArguments() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static MeteringCommandType enumForValue(byte value) {
    return map.get(value);
  }

  public static Boolean isDefined(byte value) {
    return map.containsKey(value);
  }
}
