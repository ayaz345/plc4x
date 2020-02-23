/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.abeth;

import org.apache.plc4x.java.abeth.configuration.AbEthConfiguration;
import org.apache.plc4x.java.abeth.field.AbEthFieldHandler;
import org.apache.plc4x.java.abeth.protocol.AbEthProtocolLogic;
import org.apache.plc4x.java.abeth.readwrite.CIPEncapsulationPacket;
import org.apache.plc4x.java.abeth.readwrite.io.CIPEncapsulationPacketIO;
import org.apache.plc4x.java.api.PlcDriver;
import org.apache.plc4x.java.spi.configuration.Configuration;
import org.apache.plc4x.java.spi.connection.GeneratedDriverBase;
import org.apache.plc4x.java.spi.connection.ProtocolStackConfigurer;
import org.apache.plc4x.java.spi.connection.SingleProtocolStackConfigurer;
import org.osgi.service.component.annotations.Component;

@Component(service = PlcDriver.class, immediate = true)
public class AbEthDriver extends GeneratedDriverBase<CIPEncapsulationPacket> {

    public static final int AB_ETH_PORT = 2222;

    @Override
    public String getProtocolCode() {
        return "ab-eth";
    }

    @Override
    public String getProtocolName() {
        return "Allen Bradley ETH";
    }


    @Override
    protected Class<? extends Configuration> getConfigurationType() {
        return AbEthConfiguration.class;
    }

    @Override
    protected String getDefaultTransport() {
        return "raw";
    }

    @Override
    protected AbEthFieldHandler getFieldHandler() {
        return new AbEthFieldHandler();
    }

    @Override
    protected ProtocolStackConfigurer<CIPEncapsulationPacket> getStackConfigurer() {
        return SingleProtocolStackConfigurer.builder(CIPEncapsulationPacket.class, CIPEncapsulationPacketIO.class)
            .withProtocol(AbEthProtocolLogic.class)
            .build();
    }

}