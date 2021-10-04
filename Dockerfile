<?xml version="1.0" encoding="UTF-8"?>
<!--Created by TIBCO WSDL-->
<wsdl:definitions xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/" xmlns:tns="http://esb.kaspibank.kz/Bpm_KolesaServices_Ws" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" name="Untitled" targetNamespace="http://esb.kaspibank.kz/Bpm_KolesaServices_Ws">
    <wsdl:types>
        <xs:schema xmlns="http://esb.kaspibank.kz/Bpm_KolesaServices_Ws" xmlns:xs="http://www.w3.org/2001/XMLSchema" targetNamespace="http://esb.kaspibank.kz/Bpm_KolesaServices_Ws" elementFormDefault="qualified" attributeFormDefault="unqualified">
            <xs:element name="pingEsbRequest">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="msg" type="xs:string"/>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="pingEsbResponse">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="result">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="status" type="xs:string"/>
                                    <xs:element name="system" type="xs:string" minOccurs="0"/>
                                    <xs:element name="code" type="xs:string" minOccurs="0"/>
                                    <xs:element name="message" type="xs:string" minOccurs="0"/>
                                    <xs:element name="stackTrace" type="xs:string" minOccurs="0"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                        <xs:element name="data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="msg" type="xs:string" minOccurs="0"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="sendBpmMsgRequest">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="userID" type="xs:string" minOccurs="0"/>
                        <xs:element name="customerID" type="xs:string" minOccurs="0"/>
                        <xs:element name="eventName" type="xs:string" minOccurs="0"/>
                        <xs:element name="priority" type="xs:string" minOccurs="0"/>
                        <xs:element name="requireProcessing" type="xs:string" minOccurs="0"/>
                        <xs:element name="externalMessageID" type="xs:string" minOccurs="0"/>
                        <xs:element name="bpdInstanceId" type="xs:string" minOccurs="0"/>
                        <xs:element name="systemSource" type="xs:string" minOccurs="0"/>
                        <xs:element name="business_data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="ngParField" minOccurs="0" maxOccurs="unbounded">
                                        <xs:complexType>
                                            <xs:sequence>
                                                <xs:element name="codePar" type="xs:string" minOccurs="0"/>
                                                <xs:element name="valuePar" type="xs:string" minOccurs="0"/>
                                            </xs:sequence>
                                        </xs:complexType>
                                    </xs:element>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="sendBpmMsgResponse">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="result">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="status" type="xs:string"/>
                                    <xs:element name="system" type="xs:string" minOccurs="0"/>
                                    <xs:element name="code" type="xs:string" minOccurs="0"/>
                                    <xs:element name="message" type="xs:string" minOccurs="0"/>
                                    <xs:element name="stackTrace" type="xs:string" minOccurs="0"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                        <xs:element name="data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="JmsMessageID" type="xs:string"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="sendKaspiMsgRequest">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="eventName" type="xs:string" minOccurs="0"/>
                        <xs:element name="sourceSystem" type="xs:string" minOccurs="0"/>
						<xs:element name="targetSystem" type="xs:string" minOccurs="0"/>
                        <xs:element name="business_data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="ngParField" minOccurs="0" maxOccurs="unbounded">
                                        <xs:complexType>
                                            <xs:sequence>
                                                <xs:element name="codePar" type="xs:string" minOccurs="0"/>
                                                <xs:element name="valuePar" type="xs:string" minOccurs="0"/>
                                            </xs:sequence>
                                        </xs:complexType>
                                    </xs:element>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="sendKaspiMsgResponse">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="result">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="status" type="xs:string"/>
                                    <xs:element name="system" type="xs:string" minOccurs="0"/>
                                    <xs:element name="code" type="xs:string" minOccurs="0"/>
                                    <xs:element name="message" type="xs:string" minOccurs="0"/>
                                    <xs:element name="stackTrace" type="xs:string" minOccurs="0"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                        <xs:element name="data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="JmsMessageID" type="xs:string"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="sendCamundaMsgRequest">
                <xs:complexType>
                    <xs:all>
                        <xs:element name="bpdInstanceId" type="xs:string" minOccurs="0"/>
                        <xs:element name="business_data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="ngParField" minOccurs="0" maxOccurs="unbounded">
                                        <xs:complexType>
                                            <xs:sequence>
                                                <xs:element name="codePar" type="xs:string" minOccurs="0"/>
                                                <xs:element name="valuePar" type="xs:string" minOccurs="0"/>
                                            </xs:sequence>
                                        </xs:complexType>
                                    </xs:element>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                        <xs:element name="customerIIN" type="xs:string" minOccurs="0"/>
                        <xs:element name="eventGroup" type="xs:string" minOccurs="0"/>
                        <xs:element name="eventName" type="xs:string" minOccurs="0"/>
                        <xs:element name="externalMessageID" type="xs:string" minOccurs="0"/>
                        <xs:element name="sourceSystem" type="xs:string" minOccurs="0"/>
                        <xs:element name="targetSystem" type="xs:string" minOccurs="0"/>
                    </xs:all>
                </xs:complexType>
            </xs:element>
            <xs:element name="sendCamundaMsgResponse">
                <xs:complexType>
                    <xs:sequence>
                        	<xs:element name="result">
				<xs:annotation>
					<xs:documentation>Блок содержит информацию о результате выполнения операции.</xs:documentation>
				</xs:annotation>
				<xs:complexType>
					<xs:sequence>
						<xs:element name="status" type="xs:string">
							<xs:annotation>
								<xs:documentation>Результат работы операции. Возращается:
								OK - вызов прошёл успешно, ошибок нет
								FAILED - при обращении в другую систему, получили бизнес ошибку. Нет смысла делать повторный вызов операции. Проблема с данными.
								ERROR - произошла ошибка при обращении в друю систему или внутри сервиса ESB. Нужно повторить вызов ещё раз через некоторый промежуток времени.
								</xs:documentation>
							</xs:annotation>
						</xs:element>
						<xs:element name="system" type="xs:string" minOccurs="0">
							<xs:annotation>
								<xs:documentation>Код системы от которой получили ответ или ошибку</xs:documentation>
							</xs:annotation>
						</xs:element>
						<xs:element name="code" type="xs:string">
							<xs:annotation>
								<xs:documentation>Код ответа полученный от вызываемой системы.
								Если "status" равен "ERROR", то коды могут быть:
								INTERNAL_EXCEPTION - произошла ошибка при обращении к стороннему сервису или БД. Код системы см. в теге "system".
								TIMEOUT_EXCEPTION - истекло время ожидания ответа от системы, к которой было обращение. Код системы см. в теге "system".
								LOGIN_TIMEOUT_EXCEPTION - истекло время ожидания подключения к БД. Код системы см. в теге "system".
								RUNTIME_EXCEPTION - ошибка некорректного обращеия к сервису ESB или необработанная внутренняя ошикба.
								PARSE_EXCEPTION - ошибка при парсинге XML или JSON сообщения.
								</xs:documentation>
							</xs:annotation>
						</xs:element>
						<xs:element name="message" type="xs:string" minOccurs="0">
							<xs:annotation>
								<xs:documentation>сообщение об ошибке</xs:documentation>
							</xs:annotation>
						</xs:element>
						<xs:element name="stacktrace" type="xs:string" minOccurs="0">
							<xs:annotation>
								<xs:documentation>Полная расшифровка ошибки </xs:documentation>
							</xs:annotation>
						</xs:element>
					</xs:sequence>
				</xs:complexType>
			</xs:element>
                        <xs:element name="data" minOccurs="0">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="JmsMessageID" type="xs:string"/>
                                </xs:sequence>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
        </xs:schema>
    </wsdl:types>
    <wsdl:service name="Bpm_KolesaServices_Ws">
        <wsdl:port name="Bpm_KolesaServices_Ws_PortTypeEndpoint0" binding="tns:Bpm_KolesaServices_Ws_PortTypeEndpoint0Binding">
            <soap:address location="http://localhost:25048/Bpm_KolesaServices_Ws"/>
        </wsdl:port>
    </wsdl:service>
    <wsdl:portType name="Bpm_KolesaServices_Ws_PortType">
        <wsdl:operation name="sendBpmMsg">
            <wsdl:input message="tns:sendBpmMsgRequest"/>
            <wsdl:output message="tns:sendBpmMsgResponse"/>
        </wsdl:operation>
        <wsdl:operation name="sendCamundaMsg">
            <wsdl:input message="tns:sendCamundaMsgRequest"/>
            <wsdl:output message="tns:sendCamundaMsgResponse"/>
        </wsdl:operation>
		<wsdl:operation name="sendKaspiMsg">
            <wsdl:input message="tns:sendKaspiMsgRequest"/>
            <wsdl:output message="tns:sendKaspiMsgResponse"/>
        </wsdl:operation>
        <wsdl:operation name="pingEsb">
            <wsdl:input message="tns:pingEsbRequest"/>
            <wsdl:output message="tns:pingEsbResponse"/>
        </wsdl:operation>
    </wsdl:portType>
    <wsdl:binding name="Bpm_KolesaServices_Ws_PortTypeEndpoint0Binding" type="tns:Bpm_KolesaServices_Ws_PortType">
        <soap:binding style="document" transport="http://schemas.xmlsoap.org/soap/http"/>
        <wsdl:operation name="sendBpmMsg">
            <soap:operation style="document" soapAction="/sendBpmMsg"/>
            <wsdl:input>
                <soap:body use="literal" parts="request"/>
            </wsdl:input>
            <wsdl:output>
                <soap:body use="literal" parts="response"/>
            </wsdl:output>
        </wsdl:operation>
        <wsdl:operation name="sendCamundaMsg">
            <soap:operation style="document" soapAction="/sendCamundaMsg"/>
            <wsdl:input>
                <soap:body use="literal" parts="sendCamundaMsgRequest"/>
            </wsdl:input>
            <wsdl:output>
                <soap:body use="literal" parts="sendCamundaMsgResponse"/>
            </wsdl:output>
        </wsdl:operation>
        <wsdl:operation name="sendKaspiMsg">
            <soap:operation style="document" soapAction="/sendKaspiMsg"/>
            <wsdl:input>
                <soap:body use="literal" parts="sendKaspiMsgRequest"/>
            </wsdl:input>
            <wsdl:output>
                <soap:body use="literal" parts="sendKaspiMsgResponse"/>
            </wsdl:output>
        </wsdl:operation>
        <wsdl:operation name="pingEsb">
            <soap:operation style="document" soapAction="/pingEsb"/>
            <wsdl:input>
                <soap:body use="literal" parts="request"/>
            </wsdl:input>
            <wsdl:output>
                <soap:body use="literal" parts="response"/>
            </wsdl:output>
        </wsdl:operation>
    </wsdl:binding>
    <wsdl:message name="sendBpmMsgRequest">
        <wsdl:part name="request" element="tns:sendBpmMsgRequest"/>
    </wsdl:message>
    <wsdl:message name="sendBpmMsgResponse">
        <wsdl:part name="response" element="tns:sendBpmMsgResponse"/>
    </wsdl:message>
    <wsdl:message name="sendCamundaMsgRequest">
        <wsdl:part name="sendCamundaMsgRequest" element="tns:sendCamundaMsgRequest"/>
    </wsdl:message>
    <wsdl:message name="sendCamundaMsgResponse">
        <wsdl:part name="sendCamundaMsgResponse" element="tns:sendCamundaMsgResponse"/>
    </wsdl:message>
    <wsdl:message name="sendKaspiMsgRequest">
        <wsdl:part name="sendKaspiMsgRequest" element="tns:sendKaspiMsgRequest"/>
    </wsdl:message>
    <wsdl:message name="sendKaspiMsgResponse">
        <wsdl:part name="sendKaspiMsgResponse" element="tns:sendKaspiMsgResponse"/>
    </wsdl:message>
    <wsdl:message name="pingEsbRequest">
        <wsdl:part name="request" element="tns:pingEsbRequest"/>
    </wsdl:message>
    <wsdl:message name="pingEsbResponse">
        <wsdl:part name="response" element="tns:pingEsbResponse"/>
    </wsdl:message>
</wsdl:definitions>