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
package org.apache.plc4x.java.test.readwrite;

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

public class TypeSwitchInTypeSwitchParentTypeChild0Infant2
    extends TypeSwitchInTypeSwitchParentTypeChild0 implements Message {

  // Accessors for discriminator values.
  public Short getChildNumber() {
    return (short) 0x03;
  }

  // Properties.
  protected final short infantSomeField02;

  public TypeSwitchInTypeSwitchParentTypeChild0Infant2(
      short parentFieldHurz, short childFieldwolf, short infantSomeField02) {
    super(parentFieldHurz, childFieldwolf);
    this.infantSomeField02 = infantSomeField02;
  }

  public short getInfantSomeField02() {
    return infantSomeField02;
  }

  @Override
  protected void serializeTypeSwitchInTypeSwitchParentTypeChild0Child(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TypeSwitchInTypeSwitchParentTypeChild0Infant2");

    // Simple Field (infantSomeField02)
    writeSimpleField("infantSomeField02", infantSomeField02, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("TypeSwitchInTypeSwitchParentTypeChild0Infant2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TypeSwitchInTypeSwitchParentTypeChild0Infant2 _value = this;

    // Simple field (infantSomeField02)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static TypeSwitchInTypeSwitchParentTypeChild0Builder
      staticParseTypeSwitchInTypeSwitchParentTypeChild0Builder(ReadBuffer readBuffer)
          throws ParseException {
    readBuffer.pullContext("TypeSwitchInTypeSwitchParentTypeChild0Infant2");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short infantSomeField02 =
        readSimpleField("infantSomeField02", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("TypeSwitchInTypeSwitchParentTypeChild0Infant2");
    // Create the instance
    return new TypeSwitchInTypeSwitchParentTypeChild0Infant2BuilderImpl(infantSomeField02);
  }

  public static class TypeSwitchInTypeSwitchParentTypeChild0Infant2BuilderImpl
      implements TypeSwitchInTypeSwitchParentTypeChild0
          .TypeSwitchInTypeSwitchParentTypeChild0Builder {
    private final short infantSomeField02;

    public TypeSwitchInTypeSwitchParentTypeChild0Infant2BuilderImpl(short infantSomeField02) {

      this.infantSomeField02 = infantSomeField02;
    }

    public TypeSwitchInTypeSwitchParentTypeChild0Infant2 build(
        short parentFieldHurz, short childFieldwolf) {
      TypeSwitchInTypeSwitchParentTypeChild0Infant2 typeSwitchInTypeSwitchParentTypeChild0Infant2 =
          new TypeSwitchInTypeSwitchParentTypeChild0Infant2(
              parentFieldHurz, childFieldwolf, infantSomeField02);
      return typeSwitchInTypeSwitchParentTypeChild0Infant2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TypeSwitchInTypeSwitchParentTypeChild0Infant2)) {
      return false;
    }
    TypeSwitchInTypeSwitchParentTypeChild0Infant2 that =
        (TypeSwitchInTypeSwitchParentTypeChild0Infant2) o;
    return (getInfantSomeField02() == that.getInfantSomeField02()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getInfantSomeField02());
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
