package myservice_test

import (
	"os"
	"testing"
	"time"

	myservice "github.com/dolanor/gomisc/soapchronopost"
	"github.com/hooklift/gowsdl/soap"
)

func TestEtiquette(t *testing.T) {
	sc := soap.NewClient("https://ws.chronopost.fr/shipping-cxf/ShippingServiceWS")
	c := myservice.NewShippingServiceWS(sc)

	smp := myservice.ShippingMultiParcelV4{
		HeaderValue: &myservice.HeaderValue{
			AccountNumber: 255562,
			IdEmit:        "CHRFR",
		},
		ShipperValue: []*myservice.ShipperValueV2{
			&myservice.ShipperValueV2{
				ShipperValue: &myservice.ShipperValue{
					ShipperCivility: "M",
					ShipperAdress1:  "1 rue des ondes",
					ShipperCity:     "Strasbourg",
					ShipperZipCode:  "67000",
					ShipperCountry:  "FR",
					ShipperEmail:    "m@gmail.com",
					ShipperName:     "MonsieurM",
					ShipperPhone:    "0123456789",
				},
			},
		},
		CustomerValue: &myservice.CustomerValue{
			CustomerCivility: "M",
			CustomerAdress1:  "1 rue des ondes",
			CustomerCity:     "Strasbourg",
			CustomerZipCode:  "67000",
			CustomerCountry:  "FR",
			CustomerEmail:    "m@gmail.com",
			CustomerName:     "MonsieurM",
			CustomerPhone:    "0123456789",
		},
		RecipientValue: []*myservice.RecipientValueV2{
			&myservice.RecipientValueV2{
				RecipientValue: &myservice.RecipientValue{
					RecipientAdress1: "1 rue des ondes",
					RecipientCity:    "Strasbourg",
					RecipientZipCode: "67000",
					RecipientCountry: "FR",
					RecipientEmail:   "m@gmail.com",
					RecipientName:    "MonsieurM",
					RecipientPhone:   "0123456789",
				},
			},
		},
		RefValue: []*myservice.RefValueV2{
			&myservice.RefValueV2{
				RefValue: &myservice.RefValue{},
			},
		},
		SkybillValue: []*myservice.SkybillWithDimensionsValueV6{
			&myservice.SkybillWithDimensionsValueV6{
				SkybillWithDimensionsValueV5: &myservice.SkybillWithDimensionsValueV5{
					SkybillWithDimensionsValueV4: &myservice.SkybillWithDimensionsValueV4{
						SkybillWithDimensionsValueV3: &myservice.SkybillWithDimensionsValueV3{
							SkybillWithDimensionsValueV2: &myservice.SkybillWithDimensionsValueV2{
								SkybillWithDimensionsValue: &myservice.SkybillWithDimensionsValue{
									Height: 1,
									Width:  1,
									Length: 1,
									SkybillValue: &myservice.SkybillValue{
										EvtCode:     "DC",
										ObjectType:  "MAR",
										ProductCode: "LO",
										Service:     "0",
										ShipDate:    soap.CreateXsdDateTime(time.Now(), true),
										ShipHour:    10,
										Weight:      1,
										WeightUnit:  "KGM",
									},
								},
							},
						},
					},
				},
			},
		},
		SkybillParamsValue: &myservice.SkybillParamsValueV2{
			SkybillParamsValue: &myservice.SkybillParamsValue{
				Mode: "PDF",
			},
		},
		Password: "255562",

		NumberOfParcel: 1,
		Version:        "2.0",
		ScheduledValue: []*myservice.ScheduledValue{
			&myservice.ScheduledValue{
				AppointmentValue: &myservice.AppointmentValue{
					TimeSlotStartDate: soap.CreateXsdDateTime(time.Now(), true),
					TimeSlotEndDate:   soap.CreateXsdDateTime(time.Now(), true),
				},
			},
		},
		CustomsValue: []*myservice.CustomsValue{
			&myservice.CustomsValue{
				ArticlesValue: []*myservice.ArticleValue{
					&myservice.ArticleValue{},
				},
			},
		},
	}
	resp, err := c.ShippingMultiParcelV4(&smp)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("out.pdf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write(resp.Return_.ResultMultiParcelValue[0].PdfEtiquette)
	if err != nil {
		t.Fatal(err)
	}
}
