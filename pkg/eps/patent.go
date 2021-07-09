package eps

import "encoding/xml"

// EpPatentDocument is a auto generated struct based on the example xml files
type EpPatentDocument struct {
	XMLName    xml.Name `xml:"ep-patent-document"`
	Text       string   `xml:",chardata"`
	ID         string   `xml:"id,attr"`
	File       string   `xml:"file,attr"`
	Lang       string   `xml:"lang,attr"`
	Country    string   `xml:"country,attr"`
	DocNumber  string   `xml:"doc-number,attr"`
	Kind       string   `xml:"kind,attr"`
	DatePubl   string   `xml:"date-publ,attr"`
	Status     string   `xml:"status,attr"`
	DtdVersion string   `xml:"dtd-version,attr"`
	SDOBI      struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		B000 struct {
			Text   string `xml:",chardata"`
			Eptags struct {
				Text   string `xml:",chardata"`
				B001EP string `xml:"B001EP"`
				B005EP string `xml:"B005EP"`
				B007EP string `xml:"B007EP"`
			} `xml:"eptags"`
		} `xml:"B000"`
		B100 struct {
			Text string `xml:",chardata"`
			B110 string `xml:"B110"`
			B120 struct {
				Text string `xml:",chardata"`
				B121 string `xml:"B121"`
			} `xml:"B120"`
			B130 string `xml:"B130"`
			B140 struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"B140"`
			B190 string `xml:"B190"`
		} `xml:"B100"`
		B200 struct {
			Text string `xml:",chardata"`
			B210 string `xml:"B210"`
			B220 struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"B220"`
			B240 struct {
				Text string `xml:",chardata"`
				B241 struct {
					Text string `xml:",chardata"`
					Date string `xml:"date"`
				} `xml:"B241"`
			} `xml:"B240"`
			B250   string `xml:"B250"`
			B251EP string `xml:"B251EP"`
			B260   string `xml:"B260"`
		} `xml:"B200"`
		B400 struct {
			Text string `xml:",chardata"`
			B405 struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Bnum string `xml:"bnum"`
			} `xml:"B405"`
			B430 struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Bnum string `xml:"bnum"`
			} `xml:"B430"`
			B450 struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Bnum string `xml:"bnum"`
			} `xml:"B450"`
			B452EP struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"B452EP"`
		} `xml:"B400"`
		B500 struct {
			Text   string `xml:",chardata"`
			B510EP struct {
				Text               string `xml:",chardata"`
				ClassificationIpcr []struct {
					Chardata string `xml:",chardata"`
					Sequence string `xml:"sequence,attr"`
					Text     string `xml:"text"`
				} `xml:"classification-ipcr"`
			} `xml:"B510EP"`
			B540 struct {
				Text string   `xml:",chardata"`
				B541 []string `xml:"B541"`
				B542 []string `xml:"B542"`
			} `xml:"B540"`
			B560 struct {
				Text string `xml:",chardata"`
				B561 []struct {
					Chardata string `xml:",chardata"`
					Text     string `xml:"text"`
				} `xml:"B561"`
				B562 struct {
					Chardata string `xml:",chardata"`
					Text     string `xml:"text"`
				} `xml:"B562"`
				B565EP struct {
					Text string `xml:",chardata"`
					Date string `xml:"date"`
				} `xml:"B565EP"`
			} `xml:"B560"`
		} `xml:"B500"`
		B700 struct {
			Text string `xml:",chardata"`
			B720 struct {
				Text string `xml:",chardata"`
				B721 struct {
					Text string `xml:",chardata"`
					Snm  string `xml:"snm"`
					Adr  struct {
						Text string `xml:",chardata"`
						Str  string `xml:"str"`
						City string `xml:"city"`
						Ctry string `xml:"ctry"`
					} `xml:"adr"`
				} `xml:"B721"`
			} `xml:"B720"`
			B730 struct {
				Text string `xml:",chardata"`
				B731 struct {
					Text string `xml:",chardata"`
					Snm  string `xml:"snm"`
					Iid  string `xml:"iid"`
					Irf  string `xml:"irf"`
					Adr  struct {
						Text string `xml:",chardata"`
						Str  string `xml:"str"`
						City string `xml:"city"`
						Ctry string `xml:"ctry"`
					} `xml:"adr"`
				} `xml:"B731"`
			} `xml:"B730"`
			B740 struct {
				Text string `xml:",chardata"`
				B741 struct {
					Text string `xml:",chardata"`
					Snm  string `xml:"snm"`
					Iid  string `xml:"iid"`
					Adr  struct {
						Text string `xml:",chardata"`
						Str  string `xml:"str"`
						City string `xml:"city"`
						Ctry string `xml:"ctry"`
					} `xml:"adr"`
				} `xml:"B741"`
			} `xml:"B740"`
		} `xml:"B700"`
		B800 struct {
			Text string `xml:",chardata"`
			B840 struct {
				Text string   `xml:",chardata"`
				Ctry []string `xml:"ctry"`
			} `xml:"B840"`
			B860 struct {
				Text string `xml:",chardata"`
				B861 struct {
					Text string `xml:",chardata"`
					Dnum struct {
						Text string `xml:",chardata"`
						Anum string `xml:"anum"`
					} `xml:"dnum"`
					Date string `xml:"date"`
				} `xml:"B861"`
				B862 string `xml:"B862"`
			} `xml:"B860"`
			B870 struct {
				Text string `xml:",chardata"`
				B871 struct {
					Text string `xml:",chardata"`
					Dnum struct {
						Text string `xml:",chardata"`
						Pnum string `xml:"pnum"`
					} `xml:"dnum"`
					Date string `xml:"date"`
					Bnum string `xml:"bnum"`
				} `xml:"B871"`
			} `xml:"B870"`
		} `xml:"B800"`
	} `xml:"SDOBI"`
	Description struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Lang    string `xml:"lang,attr"`
		Heading []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"heading"`
		P []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Num  string `xml:"num,attr"`
			Ul   struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				ListStyle string `xml:"list-style,attr"`
				Compact   string `xml:"compact,attr"`
				Li        []struct {
					Text   string `xml:",chardata"`
					Patcit struct {
						Chardata string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Dnum     string `xml:"dnum,attr"`
						Text     string `xml:"text"`
					} `xml:"patcit"`
					Figref struct {
						Text  string `xml:",chardata"`
						Idref string `xml:"idref,attr"`
					} `xml:"figref"`
				} `xml:"li"`
			} `xml:"ul"`
			Figref struct {
				Text  string `xml:",chardata"`
				Idref string `xml:"idref,attr"`
			} `xml:"figref"`
			Sub   []string `xml:"sub"`
			Br    string   `xml:"br"`
			Maths struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Num  string `xml:"num,attr"`
				Math struct {
					Text    string   `xml:",chardata"`
					Display string   `xml:"display,attr"`
					Mi      []string `xml:"mi"`
					Mo      []string `xml:"mo"`
					Mfrac   struct {
						Text string `xml:",chardata"`
						Mrow []struct {
							Text string   `xml:",chardata"`
							Mi   string   `xml:"mi"`
							Mo   []string `xml:"mo"`
							Msub struct {
								Text string `xml:",chardata"`
								Mi   string `xml:"mi"`
								Mn   string `xml:"mn"`
							} `xml:"msub"`
							Mn string `xml:"mn"`
						} `xml:"mrow"`
						Mn   string `xml:"mn"`
						Msub struct {
							Text string `xml:",chardata"`
							Mi   []struct {
								Text        string `xml:",chardata"`
								Mathvariant string `xml:"mathvariant,attr"`
							} `xml:"mi"`
						} `xml:"msub"`
					} `xml:"mfrac"`
					Msub struct {
						Text string `xml:",chardata"`
						Mi   string `xml:"mi"`
						Mn   string `xml:"mn"`
					} `xml:"msub"`
					Mn []string `xml:"mn"`
				} `xml:"math"`
				Img struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id,attr"`
					File       string `xml:"file,attr"`
					Wi         string `xml:"wi,attr"`
					He         string `xml:"he,attr"`
					ImgContent string `xml:"img-content,attr"`
					ImgFormat  string `xml:"img-format,attr"`
				} `xml:"img"`
			} `xml:"maths"`
		} `xml:"p"`
	} `xml:"description"`
	Claims []struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Lang  string `xml:"lang,attr"`
		Claim []struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"id,attr"`
			Num       string `xml:"num,attr"`
			ClaimText struct {
				Text  string   `xml:",chardata"`
				Br    []string `xml:"br"`
				Sub   []string `xml:"sub"`
				B     string   `xml:"b"`
				Maths struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Num  string `xml:"num,attr"`
					Math struct {
						Text    string   `xml:",chardata"`
						Display string   `xml:"display,attr"`
						Mi      []string `xml:"mi"`
						Mo      []string `xml:"mo"`
						Mn      string   `xml:"mn"`
						Mfrac   struct {
							Text string `xml:",chardata"`
							Mn   string `xml:"mn"`
							Msub struct {
								Text string `xml:",chardata"`
								Mi   []struct {
									Text        string `xml:",chardata"`
									Mathvariant string `xml:"mathvariant,attr"`
								} `xml:"mi"`
							} `xml:"msub"`
						} `xml:"mfrac"`
					} `xml:"math"`
					Img struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"id,attr"`
						File       string `xml:"file,attr"`
						Wi         string `xml:"wi,attr"`
						He         string `xml:"he,attr"`
						ImgContent string `xml:"img-content,attr"`
						ImgFormat  string `xml:"img-format,attr"`
					} `xml:"img"`
				} `xml:"maths"`
				ClaimText []struct {
					Text  string `xml:",chardata"`
					Maths struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Num  string `xml:"num,attr"`
						Math struct {
							Text    string   `xml:",chardata"`
							Display string   `xml:"display,attr"`
							Mi      []string `xml:"mi"`
							Mo      []string `xml:"mo"`
							Mfrac   struct {
								Text string `xml:",chardata"`
								Mrow []struct {
									Text string   `xml:",chardata"`
									Mi   []string `xml:"mi"`
									Mo   []string `xml:"mo"`
									Mn   string   `xml:"mn"`
								} `xml:"mrow"`
								Mn   string `xml:"mn"`
								Msub struct {
									Text string `xml:",chardata"`
									Mi   []struct {
										Text        string `xml:",chardata"`
										Mathvariant string `xml:"mathvariant,attr"`
									} `xml:"mi"`
								} `xml:"msub"`
							} `xml:"mfrac"`
							Mn string `xml:"mn"`
						} `xml:"math"`
						Img struct {
							Text       string `xml:",chardata"`
							ID         string `xml:"id,attr"`
							File       string `xml:"file,attr"`
							Wi         string `xml:"wi,attr"`
							He         string `xml:"he,attr"`
							ImgContent string `xml:"img-content,attr"`
							ImgFormat  string `xml:"img-format,attr"`
						} `xml:"img"`
					} `xml:"maths"`
					Sub       []string `xml:"sub"`
					ClaimText []string `xml:"claim-text"`
					B         string   `xml:"b"`
				} `xml:"claim-text"`
			} `xml:"claim-text"`
		} `xml:"claim"`
	} `xml:"claims"`
	Drawings struct {
		Text   string `xml:",chardata"`
		ID     string `xml:"id,attr"`
		Lang   string `xml:"lang,attr"`
		Figure []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Num  string `xml:"num,attr"`
			Img  struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"id,attr"`
				File       string `xml:"file,attr"`
				Wi         string `xml:"wi,attr"`
				He         string `xml:"he,attr"`
				ImgContent string `xml:"img-content,attr"`
				ImgFormat  string `xml:"img-format,attr"`
			} `xml:"img"`
		} `xml:"figure"`
	} `xml:"drawings"`
	EpReferenceList struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Heading []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			B    string `xml:"b"`
		} `xml:"heading"`
		P []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Num  string `xml:"num,attr"`
			I    string `xml:"i"`
			Ul   struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				ListStyle string `xml:"list-style,attr"`
				Li        []struct {
					Text   string `xml:",chardata"`
					Patcit struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"id,attr"`
						Dnum       string `xml:"dnum,attr"`
						DocumentID struct {
							Text      string `xml:",chardata"`
							Country   string `xml:"country"`
							DocNumber string `xml:"doc-number"`
							Kind      string `xml:"kind"`
						} `xml:"document-id"`
					} `xml:"patcit"`
					Crossref struct {
						Text  string `xml:",chardata"`
						Idref string `xml:"idref,attr"`
					} `xml:"crossref"`
				} `xml:"li"`
			} `xml:"ul"`
		} `xml:"p"`
	} `xml:"ep-reference-list"`
}
