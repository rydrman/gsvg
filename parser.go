package svg

import (
  "bytes"
  "encoding/xml"
  "errors"
  "io"
  "io/ioutil"
  "log"
)

// UnmarshalFile reads the given svg file and unmarshals it
func UnmarshalFile(svgFile string) (*Element, error) {
  bytes, err := ioutil.ReadFile(svgFile)
  if nil != err {
    return nil, err
  }
  return Unmarshal(bytes)
}

//Unmarshal decodes the given xvg content
func Unmarshal(rawSVG []byte) (*Element, error) {

  reader := bytes.NewReader(rawSVG)
  decoder := xml.NewDecoder(reader)

  for t, e := decoder.Token(); e != io.EOF; t, e = decoder.Token() {
    if nil != e {
      return nil, e
    }

    switch t.(type) {
    case xml.StartElement:
      root, err := UnmarshalElement(t.(xml.StartElement), decoder)
      if nil != err {
        return nil, err
      }
      return root.(*Element), nil
    }
  }
  return nil, errors.New("no xml start element found")
}

//UnmarshalElement recursively unmarshalls the element
//starting at the given xml.StartElement
func UnmarshalElement(start xml.StartElement, decoder *xml.Decoder) (Elementer, error) {

  //get attributes
  attrs := ParseAttributes(start)

  //start by creating the right shape from the element
  var elem Elementer
  var err error
  switch start.Name.Local {
  case "svg":
    elem, err = NewElement(attrs)
  case "rect":
    elem, err = NewRect(attrs)
  case "polygon":
    elem, err = NewPolygon(attrs)
  case "polyline":
    elem, err = NewPolyLine(attrs)
  case "line":
    elem, err = NewLine(attrs)
  default:
    return nil, errors.New("element type not supported " + start.Name.Local)
  }
  if nil != err {
    return nil, err
  }

  for t, e := decoder.Token(); e != io.EOF; t, e = decoder.Token() {
    switch t.(type) {

    case xml.StartElement:
      e, err := UnmarshalElement(t.(xml.StartElement), decoder)
      if err != nil {
        log.Print("skip " + t.(xml.StartElement).Name.Local)
        log.Print(err)
        decoder.Skip()
      } else {
        element := elem.(*Element)
        element.Children = append(element.Children, e)
      }

    case xml.EndElement:
      return elem, nil

    case xml.CharData:
    case xml.Comment:
    case xml.ProcInst:
    case xml.Directive:
    default:
      //ignore everything else
      break
    }
  }
  return nil, errors.New("xml element never closed")
}

//ParseAttributes parses the attributes on a StartElement
//into an Attributes map
func ParseAttributes(start xml.StartElement) *Attributes {
  attrs := Attributes{}
  for _, attr := range start.Attr {
    switch attr.Name {
    default:
      attrs[attr.Name.Local] = attr.Value
    }
  }
  return &attrs
}
