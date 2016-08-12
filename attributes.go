package svg

import (
  "strconv"
  "strings"

  "github.com/rydrman/geo2"
)

//Attributes represents a map of the attributes on an element
type Attributes map[string]string

//Points represents the points attribute on an svg polygon
type Points struct {
  *geo2.Path
}

//UnmarshalAttribute converts the string point
//list into a slice of float32s
func (points *Points) UnmarshalAttribute(value string) error {
  stripped := strings.Trim(value, " ")
  stripped = strings.Replace(stripped, ",", " ", -1)
  stripped = strings.Replace(stripped, "\n", "", -1)
  stripped = strings.Replace(stripped, "\t", "", -1)
  vals := strings.Split(stripped, " ")
  pts := make([]float32, len(vals))
  for i, v := range vals {
    if v == "" {
      continue
    }
    parsed, _ := strconv.ParseFloat(v, 32)
    pts[i] = float32(parsed)
  }
  if pts[0] == pts[len(pts)-2] &&
    pts[1] == pts[len(pts)-1] {
    pts = pts[:len(pts)-2]
  }
  *points = Points{geo2.NewPathFromFloat32Array(pts)}
  return nil
}

//Style represents the style element of an svg element
type Style map[string]string

//UnmarshalAttribute parses the style property
//list of the given element
func (style *Style) UnmarshalAttribute(value string) error {
  *style = make(Style)
  styleList := strings.Split(value, ";")
  for _, s := range styleList {
    pair := strings.Split(s, ":")
    if len(pair) == 2 {
      (*style)[pair[0]] = strings.Trim(pair[1], " ")
    }
  }
  return nil
}
