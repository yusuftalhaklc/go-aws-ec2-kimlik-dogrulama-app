package client

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	XSI     string   `xml:"xmlns:xsi,attr"`
	XSD     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    TCKimlikNoDogrula
}

type TCKimlik struct {
	TCKimlikNo int64  `xml:"TCKimlikNo" json:"TCKimlikNo"`
	Ad         string `xml:"Ad" json:"Ad"`
	Soyad      string `xml:"Soyad" json:"Soyad"`
	DogumYili  int    `xml:"DogumYili" json:"DogumYili"`
}

type TCKimlikNoDogrula struct {
	XMLName  xml.Name `xml:"soap12:Body"`
	TCKimlik TCKimlik `xml:"http://tckimlik.nvi.gov.tr/WS TCKimlikNoDogrula"`
}

type TCKimlikNoDogrulaResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		Result  struct {
			XMLName xml.Name `xml:"TCKimlikNoDogrulaResponse"`
			Value   bool     `xml:"TCKimlikNoDogrulaResult"`
		} `xml:"TCKimlikNoDogrulaResponse"`
	} `xml:"Body"`
}
