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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class MPropWriteCon extends CEMI implements Message {

  // Accessors for discriminator values.
  public Short getMessageCode() {
    return (short) 0xF5;
  }

  // Arguments.
  protected final Integer size;

  public MPropWriteCon(Integer size) {
    super(size);
    this.size = size;
  }

  @Override
  protected void serializeCEMIChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MPropWriteCon");

    writeBuffer.popContext("MPropWriteCon");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MPropWriteCon _value = this;

    return lengthInBits;
  }

  public static MPropWriteConBuilder staticParseBuilder(ReadBuffer readBuffer, Integer size)
      throws ParseException {
    readBuffer.pullContext("MPropWriteCon");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    readBuffer.closeContext("MPropWriteCon");
    // Create the instance
    return new MPropWriteConBuilder(size);
  }

  public static class MPropWriteConBuilder implements CEMI.CEMIBuilder {
    private final Integer size;

    public MPropWriteConBuilder(Integer size) {

      this.size = size;
    }

    public MPropWriteCon build(Integer size) {

      MPropWriteCon mPropWriteCon = new MPropWriteCon(size);

      return mPropWriteCon;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MPropWriteCon)) {
      return false;
    }
    MPropWriteCon that = (MPropWriteCon) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
