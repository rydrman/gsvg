package svg

import (
  "fmt"
  "strconv"
)

//Dom represents an entire svg document
type Dom struct {
  Element
  Version float32 `xml:"version,attr"`
  ViewBox string  `xml:"viewBox,attr"`
}

//Elementer represents an svg
//element that has a named type
type Elementer interface {
  ToElement() *Element
}

//Element represents a generic SVG element
type Element struct {
  Name     string `xml:"data-name,attr"`
  ID       string `xml:"id,attr"`
  Style    Style  `xml:"style,attr"`
  Children []Elementer
}

//NewElement creates a new generic element
//from the given attributes
func NewElement(attrs *Attributes) (*Element, error) {
  id := (*attrs)["id"]
  style := new(Style)
  style.UnmarshalAttribute((*attrs)["style"])
  dataName := (*attrs)["data-name"]

  return &Element{dataName, id, *style, make([]Elementer, 0)}, nil
}

//ToElement satisfies the Elementer interface
func (elem *Element) ToElement() *Element {
  return elem
}

//Rect represents an svg rectangle element
type Rect struct {
  *Element
  X      float32
  Y      float32
  Width  float32
  Height float32
}

//NewRect creates a new rectangle element
//from the given attributes (assumes it's valid rect attrs)
func NewRect(attrs *Attributes) (*Rect, error) {
  elem, err := NewElement(attrs)
  if nil != err {
    return nil, err
  }
  x, err := strconv.ParseFloat((*attrs)["x"], 32)
  if nil != err {
    x = 0
  }
  y, err := strconv.ParseFloat((*attrs)["y"], 32)
  if nil != err {
    y = 0
  }
  w, err := strconv.ParseFloat((*attrs)["width"], 32)
  if nil != err {
    return nil, err
  }
  h, err := strconv.ParseFloat((*attrs)["height"], 32)
  if nil != err {
    return nil, err
  }

  return &Rect{
      elem,
      float32(x), float32(y),
      float32(w), float32(h)},
    nil
}

//ToElement satisfies the Elementer interface
func (rect *Rect) ToElement() *Element {
  return rect.Element
}

//String print a nice string of this rectangle
func (rect Rect) String() string {
  return fmt.Sprintf(
    "xml:rect{x:%f, y:%f, w:%f, h:%f}",
    rect.X, rect.Y, rect.Width, rect.Height)
}

//PolyLine is just a polygon but rendered differently
type PolyLine struct {
  *Polygon
}

// NewPolyLine creates a new polygon element
//from the given attributes (assumes it's valid polygon attrs)
func NewPolyLine(attrs *Attributes) (*PolyLine, error) {
  poly, err := NewPolygon(attrs)
  return &PolyLine{poly}, err
}

//Polygon represents an svg polygon element
type Polygon struct {
  *Element
  Points *Points
}

//NewPolygon creates a new polygon element
//from the given attributes (assumes it's valid polygon attrs)
func NewPolygon(attrs *Attributes) (*Polygon, error) {
  elem, err := NewElement(attrs)
  if nil != err {
    return nil, err
  }
  points := new(Points)
  points.UnmarshalAttribute((*attrs)["points"])
  if nil != err {
    return nil, err
  }

  return &Polygon{elem, points}, nil
}

//print a nice string of this polygon
func (poly Polygon) String() string {
  return fmt.Sprintf(
    "xml:polygon{%s}", poly.Points)
}

//ToElement satisfies the Elementer interface
func (poly *Polygon) ToElement() *Element {
  return poly.Element
}

//Line represents an svg line element
type Line struct {
  *Element
  X1 float32
  Y1 float32
  X2 float32
  Y2 float32
}

//NewLine creates a new line element from the
//given attributes (assumes it's valid line attrs)
func NewLine(attrs *Attributes) (*Line, error) {
  elem, err := NewElement(attrs)
  if nil != err {
    return nil, err
  }
  x1, err := strconv.ParseFloat((*attrs)["x1"], 32)
  if nil != err {
    return nil, err
  }
  y1, err := strconv.ParseFloat((*attrs)["y1"], 32)
  if nil != err {
    return nil, err
  }
  x2, err := strconv.ParseFloat((*attrs)["x2"], 32)
  if nil != err {
    return nil, err
  }
  y2, err := strconv.ParseFloat((*attrs)["y2"], 32)
  if nil != err {
    return nil, err
  }

  return &Line{
      elem,
      float32(x1), float32(y1),
      float32(x2), float32(y2)},
    nil
}

//ToElement satisfies the Elementer interface
func (line *Line) ToElement() *Element {
  return line.Element
}

//print a nice string of this polygon
func (line Line) String() string {
  return fmt.Sprintf(
    "xml:line{x1:%f, y1:%f, x2:%f, y2:%f}",
    line.X1, line.Y1, line.X2, line.Y2)
}
