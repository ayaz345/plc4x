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
package org.apache.plc4x.java.opcua.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class TimeZoneDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "8914";
  }

  // Properties.
  protected final short offset;
  protected final boolean daylightSavingInOffset;
  // Reserved Fields
  private Short reservedField0;

  public TimeZoneDataType(short offset, boolean daylightSavingInOffset) {
    super();
    this.offset = offset;
    this.daylightSavingInOffset = daylightSavingInOffset;
  }

  public short getOffset() {
    return offset;
  }

  public boolean getDaylightSavingInOffset() {
    return daylightSavingInOffset;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TimeZoneDataType");

    // Simple Field (offset)
    writeSimpleField("offset", offset, writeSignedShort(writeBuffer, 16));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 7));

    // Simple Field (daylightSavingInOffset)
    writeSimpleField("daylightSavingInOffset", daylightSavingInOffset, writeBoolean(writeBuffer));

    writeBuffer.popContext("TimeZoneDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TimeZoneDataType _value = this;

    // Simple field (offset)
    lengthInBits += 16;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (daylightSavingInOffset)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static TimeZoneDataTypeBuilder staticParseBuilder(ReadBuffer readBuffer, String identifier)
      throws ParseException {
    readBuffer.pullContext("TimeZoneDataType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short offset = readSimpleField("offset", readSignedShort(readBuffer, 16));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 7), (short) 0x00);

    boolean daylightSavingInOffset =
        readSimpleField("daylightSavingInOffset", readBoolean(readBuffer));

    readBuffer.closeContext("TimeZoneDataType");
    // Create the instance
    return new TimeZoneDataTypeBuilder(offset, daylightSavingInOffset, reservedField0);
  }

  public static class TimeZoneDataTypeBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final short offset;
    private final boolean daylightSavingInOffset;
    private final Short reservedField0;

    public TimeZoneDataTypeBuilder(
        short offset, boolean daylightSavingInOffset, Short reservedField0) {
      this.offset = offset;
      this.daylightSavingInOffset = daylightSavingInOffset;
      this.reservedField0 = reservedField0;
    }

    public TimeZoneDataType build() {
      TimeZoneDataType timeZoneDataType = new TimeZoneDataType(offset, daylightSavingInOffset);
      timeZoneDataType.reservedField0 = reservedField0;
      return timeZoneDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TimeZoneDataType)) {
      return false;
    }
    TimeZoneDataType that = (TimeZoneDataType) o;
    return (getOffset() == that.getOffset())
        && (getDaylightSavingInOffset() == that.getDaylightSavingInOffset())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOffset(), getDaylightSavingInOffset());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
