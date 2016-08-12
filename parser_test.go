package svg

import (
  "reflect"
  "testing"
)

func TestSVGParsing(t *testing.T) {
  svgData := []byte(`
  <?xml version="1.0" encoding="utf-8"?>
  <!-- Generator: Adobe Illustrator 19.2.1, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
  <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
     viewBox="0 0 216 215.75" style="enable-background:new 0 0 216 215.75;" xml:space="preserve">
    <rect x="72" y="72" style="fill:#606060;" width="72" height="72"/>
    <polygon style="fill:none;stroke:#DDB34A;stroke-width:0.1;stroke-miterlimit:10;" points="108.5,132.4289 87.777,120.4644 84.5711,108.5 87.777,96.5356 108.5,84.5711 129.223,96.5356 132.4289,108.5 120.4644,129.223 "/>
  </svg>`)

  data, err := Unmarshal(svgData)
  if nil != err {
    t.Error(err)
  }
  if reflect.TypeOf((*Rect)(nil)) != reflect.TypeOf(data.Children[0]) {
    t.Error("first svg element should be a rect")
  }
  if reflect.TypeOf((*Polygon)(nil)) != reflect.TypeOf(data.Children[1]) {
    t.Error("second svg element should be a rect")
  }
}
