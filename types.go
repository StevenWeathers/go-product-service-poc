package main

type Product struct {
	Inner struct {
		ProductId string `json:"productId"`
		Type string `json:"type"`
		Brand string `json:"brand"`
		Barcode string `json:"barcode"`
		Bullets []string `json:"bullets"`
		Description string `json:"description"`
		ImgId string `json:"imgId"`
		ItemNum string `json:"itemNum"`
		ModelNum string `json:"modelNum"`
		Specs []string `json:"specs"`
		LgImg string `json:"lgImg"`
		SmImg string `json:"smImg"`
		Rating string `json:"rating"`
		NumRatings string `json:"numRatings"`
		NetworkPrice string `json:"networkPrice"`
	} `json:"product"`
}