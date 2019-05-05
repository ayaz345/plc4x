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
package org.apache.plc4x.codegen.util;

import org.apache.plc4x.codegen.ast.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

import static org.apache.plc4x.codegen.util.EnumFactory.getGetterDefinition;
import static org.apache.plc4x.codegen.util.EnumFactory.getSetterDefinition;

public class PojoFactory {

    private TypeDefinition BUFFER_TYPE;

    private Method readByte = new Method(BUFFER_TYPE, "readByte", Primitive.INTEGER, Collections.emptyList(), Collections.emptyList());

    public ClassDeclaration create(PojoDescription desc) {
        // Create all Fields first
        final List<FieldDeclaration> fields = desc.fields.stream()
            .map(field -> new FieldDeclaration(field.getType(), field.getName()))
            .collect(Collectors.toList());

        // Create the Getters
        final List<MethodDefinition> getters = desc.fields.stream()
            .map(field -> getGetterDefinition(field.getName(), field.getType()))
            .collect(Collectors.toList());

        // Create the Setters
        final List<MethodDefinition> setters = desc.fields.stream()
            .map(field -> getSetterDefinition(field.getName(), field.getType()))
            .collect(Collectors.toList());

        final List<MethodDefinition> methods = new ArrayList<>();

        methods.addAll(getters);
        methods.addAll(setters);

        // Encode Method
        methods.add(new MethodDefinition("encode", Primitive.VOID, Collections.singletonList(
            Expressions.parameter("buffer", Expressions.typeOf("org.apache.plc4x.api.Buffer"))
        ), Block.build().toBlock()));

        // Decode Method
        BUFFER_TYPE = Expressions.typeOf("org.apache.plc4x.api.Buffer");
        final ParameterExpression buffer = Expressions.parameter("buffer", BUFFER_TYPE);
        final TypeDefinition clazz = Expressions.typeOf(desc.getName());
        final ParameterExpression instance = Expressions.parameter("instance", clazz);
        methods.add(new MethodDefinition(Collections.singleton(Modifier.STATIC), "decode", clazz, Collections.singletonList(
            buffer
        ), Block.build()
            .add(Expressions.assignment(instance, Expressions.new_(clazz)))
            .add(Expressions.call(buffer, readByte))
            .add(Expressions.return_(instance))
            .toBlock()
        ));


        return new ClassDeclaration("", desc.getName(), fields, Arrays.asList(new ConstructorDeclaration(Collections.emptyList(), Block.EMPTY_BLOCK)), methods, null);
    }

    public static class PojoDescription {

        private final String name;
        private final List<Field> fields;

        public PojoDescription(String name, Field... fields) {
            this.name = name;
            this.fields = Arrays.asList(fields);
        }

        public PojoDescription(String name, List<Field> fields) {
            this.name = name;
            this.fields = fields;
        }

        public String getName() {
            return name;
        }

        public List<Field> getFields() {
            return fields;
        }
    }

    public static class Field {

        private final TypeDefinition type;
        private final String name;

        public Field(TypeDefinition type, String name) {
            this.type = type;
            this.name = name;
        }

        public TypeDefinition getType() {
            return type;
        }

        public String getName() {
            return name;
        }
    }

}
