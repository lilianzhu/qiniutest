package fopimage

import (
	"qiniutest.com/lib/util"
)

type ImageView2Param struct {
	Mode            string
	WidthLongEdge   string
	HeightShortEdge string
	Quality         string
	Interlace       string
	Format          string
	Baseformat      string
}

func (i *ImageView2Param) NewImageView2Param(mode, widthLongEdge, heightShortEdge, quality, interlace, format, baseformat string) {
	i.Mode = mode
	i.WidthLongEdge = widthLongEdge
	i.HeightShortEdge = heightShortEdge
	i.Quality = quality
	i.Interlace = interlace
	i.Format = format
	i.Baseformat = baseformat
}

func (i *ImageView2Param) To_string() (str string) {
	str = "/" + i.Mode + "/" + i.WidthLongEdge + "/" + i.HeightShortEdge + "/" + i.Quality + "/" + i.Format + "/" + i.Interlace + "/" + i.Baseformat
	return
}

func (i *ImageView2Param) To_url() (url string) {
	url = "imageView2/"
	if i.Mode != "" {
		url += i.Mode + "/"
	}
	if i.WidthLongEdge != "" {
		url += "w/" + i.WidthLongEdge + "/"
	}
	if i.HeightShortEdge != "" {
		url += "h/" + i.HeightShortEdge + "/"
	}
	if i.Quality != "" {
		url += "q/" + i.Quality + "/"
	}
	if i.Format != "" {
		url += "format/" + i.Format + "/"
	}
	if i.Interlace != "" {
		url += "interlace/" + i.Interlace + "/"
	}
	url += util.GetRand(8)
	return
}
