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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetConfirmedServiceRequestLifeSafetyOperation extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.LIFE_SAFETY_OPERATION;
  }

  // Properties.
  protected final BACnetContextTagUnsignedInteger requestingProcessIdentifier;
  protected final BACnetContextTagCharacterString requestingSource;
  protected final BACnetLifeSafetyOperationTagged request;
  protected final BACnetContextTagObjectIdentifier objectIdentifier;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestLifeSafetyOperation(
      BACnetContextTagUnsignedInteger requestingProcessIdentifier,
      BACnetContextTagCharacterString requestingSource,
      BACnetLifeSafetyOperationTagged request,
      BACnetContextTagObjectIdentifier objectIdentifier,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.requestingProcessIdentifier = requestingProcessIdentifier;
    this.requestingSource = requestingSource;
    this.request = request;
    this.objectIdentifier = objectIdentifier;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagUnsignedInteger getRequestingProcessIdentifier() {
    return requestingProcessIdentifier;
  }

  public BACnetContextTagCharacterString getRequestingSource() {
    return requestingSource;
  }

  public BACnetLifeSafetyOperationTagged getRequest() {
    return request;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestLifeSafetyOperation");

    // Simple Field (requestingProcessIdentifier)
    writeSimpleField(
        "requestingProcessIdentifier",
        requestingProcessIdentifier,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (requestingSource)
    writeSimpleField(
        "requestingSource", requestingSource, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (request)
    writeSimpleField("request", request, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (objectIdentifier) (Can be skipped, if the value is null)
    writeOptionalField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestLifeSafetyOperation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestLifeSafetyOperation _value = this;

    // Simple field (requestingProcessIdentifier)
    lengthInBits += requestingProcessIdentifier.getLengthInBits();

    // Simple field (requestingSource)
    lengthInBits += requestingSource.getLengthInBits();

    // Simple field (request)
    lengthInBits += request.getLengthInBits();

    // Optional Field (objectIdentifier)
    if (objectIdentifier != null) {
      lengthInBits += objectIdentifier.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestLifeSafetyOperationBuilder staticParseBuilder(
      ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestLifeSafetyOperation");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetContextTagUnsignedInteger requestingProcessIdentifier =
        readSimpleField(
            "requestingProcessIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagCharacterString requestingSource =
        readSimpleField(
            "requestingSource",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    BACnetLifeSafetyOperationTagged request =
        readSimpleField(
            "request",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetLifeSafetyOperationTagged.staticParse(
                        readBuffer, (short) (2), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagObjectIdentifier objectIdentifier =
        readOptionalField(
            "objectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestLifeSafetyOperation");
    // Create the instance
    return new BACnetConfirmedServiceRequestLifeSafetyOperationBuilder(
        requestingProcessIdentifier,
        requestingSource,
        request,
        objectIdentifier,
        serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestLifeSafetyOperationBuilder
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagUnsignedInteger requestingProcessIdentifier;
    private final BACnetContextTagCharacterString requestingSource;
    private final BACnetLifeSafetyOperationTagged request;
    private final BACnetContextTagObjectIdentifier objectIdentifier;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestLifeSafetyOperationBuilder(
        BACnetContextTagUnsignedInteger requestingProcessIdentifier,
        BACnetContextTagCharacterString requestingSource,
        BACnetLifeSafetyOperationTagged request,
        BACnetContextTagObjectIdentifier objectIdentifier,
        Long serviceRequestLength) {

      this.requestingProcessIdentifier = requestingProcessIdentifier;
      this.requestingSource = requestingSource;
      this.request = request;
      this.objectIdentifier = objectIdentifier;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestLifeSafetyOperation build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestLifeSafetyOperation
          bACnetConfirmedServiceRequestLifeSafetyOperation =
              new BACnetConfirmedServiceRequestLifeSafetyOperation(
                  requestingProcessIdentifier,
                  requestingSource,
                  request,
                  objectIdentifier,
                  serviceRequestLength);
      return bACnetConfirmedServiceRequestLifeSafetyOperation;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestLifeSafetyOperation)) {
      return false;
    }
    BACnetConfirmedServiceRequestLifeSafetyOperation that =
        (BACnetConfirmedServiceRequestLifeSafetyOperation) o;
    return (getRequestingProcessIdentifier() == that.getRequestingProcessIdentifier())
        && (getRequestingSource() == that.getRequestingSource())
        && (getRequest() == that.getRequest())
        && (getObjectIdentifier() == that.getObjectIdentifier())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestingProcessIdentifier(),
        getRequestingSource(),
        getRequest(),
        getObjectIdentifier());
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
