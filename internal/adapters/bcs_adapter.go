package adapters

import (
	"AzubiTool/internal/adapters/mocks"
	"AzubiTool/internal/domain"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type BcsAdapter struct {
	testing     bool
	EndpointURL string
}

func NewBcsAdapter(endpointURL string, testing bool) *BcsAdapter {
	return &BcsAdapter{EndpointURL: endpointURL, testing: testing}
}

const payloadTmpl = `<?xml version="1.0" encoding="UTF-8"?>
                            <s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope"
                            xmlns:wsa="http://www.w3.org/2005/08/addressing"
                            xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2"
                            xmlns:pbse="http://www.projektron.de/bcs/system/events"
                            xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
                            xmlns:dot="http://www.dotsource.de"
                            xmlns:data="http://www.projektron.de/bcs/processing/data">
                                <s:Header>
                                    <wsse:Security s:mustUnderstand="true">
                                        <wsse:UsernameToken>
                                            <wsse:Username>%s</wsse:Username>
                                            <wsse:Password>%s</wsse:Password>
                                        </wsse:UsernameToken>
                                    </wsse:Security>
                                </s:Header>
                                <s:Body>
                                    %s
                                </s:Body>
                            </s:Envelope>`

var (
	getUserOidName = "getUserOidUserShortname"
	getUserOidData = `<dot:getUserOidUserShortname>
        <data:Parameters>
            <data:Parameter name="Shortname">%s</data:Parameter>
        </data:Parameters>
    </dot:getUserOidUserShortname>`

	// getEffortsName = "getEffortsByOidAndDateRange"
	// getEffortsData = `<dot:getEffortsByOidAndDateRange>
	//     <data:Parameters>
	//         <data:Parameter name="UserOid">%s</data:Parameter>
	//         <data:Parameter name="smalerDate">%s</data:Parameter>
	//         <data:Parameter name="biggerDate">%s</data:Parameter>
	//     </data:Parameters>
	// </dot:getEffortsByOidAndDateRange>`

	// getAppointmentsName = "getAppointmentsByOidAndDateRange"
	// getAppointmentsData = `<dot:getAppointmentsByOidAndDateRange>
	// 	<data:Parameters>
	// 		<data:Parameter name="UserOid">%s</data:Parameter>
	// 		<data:Parameter name="smalerDate">%s</data:Parameter>
	// 		<data:Parameter name="biggerDate">%s</data:Parameter>
	// 	</data:Parameters>
	// </dot:getAppointmentsByOidAndDateRange>`
)

func (b *BcsAdapter) GetOid(username, password string) (string, error) {
	actionData := fmt.Sprintf(getUserOidData, username)
	soapAction := fmt.Sprintf("http://www.dotsource.de/AzubiBcsToBlokService/%sRequest", getUserOidName)
	payload := fmt.Sprintf(payloadTmpl, username, password, actionData)

	println(payload)

	req, err := http.NewRequest("POST", b.EndpointURL, strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("Content-Type", "application/soap+xml;charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil && !b.testing {
		return "", err
	}
	defer resp.Body.Close()

	body := make([]byte, 0)
	if !b.testing {
		body, _ = io.ReadAll(resp.Body)
	} else {
		body = []byte(mocks.OID_response)
	}

	if resp.StatusCode != 200 {
		print("body: ", string(body))
		return "", parseSoapError(body)
	}
	oid, err := extractValueByAttr(body, "oid")
	if err != nil {
		print("extract value: ", err.Error())
		return "", err
	}
	return oid, nil
}

func (b *BcsAdapter) GetEfforts(oid, username, password string, from, to time.Time) ([]domain.EffortRecord, error) {
	// Implement the logic to get efforts from BCS API
	return nil, nil // Placeholder return
}

func (b *BcsAdapter) GetAppointments(oid, username, password string, from, to time.Time) ([]domain.Appointment, error) {
	// Implement the logic to get appointments from BCS API
	return nil, nil // Placeholder return
}

// ---- Helper ----

func parseSoapError(body []byte) error {
	type SoapFault struct {
		Reason struct {
			Text string `xml:"Text"`
		} `xml:"Body>Fault>Reason"`
		Code struct {
			Value string `xml:"Body>Fault>Code>Value"`
		} `xml:"Body>Fault>Code"`
	}
	var fault SoapFault
	if err := xml.Unmarshal(body, &fault); err != nil {
		return errors.New("unbekannter SOAP-Fehler")
	}
	return fmt.Errorf("%s: %s", fault.Code.Value, fault.Reason.Text)
}

func extractValueByAttr(xmlBody []byte, attrName string) (string, error) {
	type Part struct {
		XMLName xml.Name
		Name    string `xml:"name,attr"`
		Value   string `xml:",chardata"`
	}
	decoder := xml.NewDecoder(bytes.NewReader(xmlBody))
	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "Part" {
				var p Part
				decoder.DecodeElement(&p, &se)
				if p.Name == attrName {
					return strings.TrimSpace(p.Value), nil
				}
			}
		}
	}
	return "", fmt.Errorf("wert mit Attribut %q nicht gefunden", attrName)
}
