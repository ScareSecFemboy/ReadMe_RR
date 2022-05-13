/*

ALL CODE WAS CREATED AND GENERATED BY ARK ANGEL43 YOU MAY NOT TAKE THIS CODE WITHOUT PROPER AUTHOR PREMISSION OR AUTHOR CREDITS

YOU HAVE BEEN WARNED!

Package parses all json data files and tries to find and split apart the first few digit's of the number and tries to compare it to country codes in the json file and willtry
to output a country name or location name if it can find it, not location gathering and addr searching will only be allowed to be used if the number is a United States number
if it is not it can not get the exact addr of the number, however you will be able to get the following from international supported numbers

Mask           : EX -> +973-=####-####
CC             : EX -> BH
English name   : EX -> Bahrain
English desc   : EX -> <nil>???
Russian Name   : EX -> Болгария

if the name or phone number ex: +973-####-#### is a deep equal to a hash or code in the database it will return a true value, and parse the data in the json file


the more the red rabbit project the more organized and structured this database will be and will eventually be turned into a PostGreSQL database or MONGO database
to store all numbers and relative data to those numbers, the plan for this would be to NOT use a website to grab phone information since we can not see
if this JSON data is true to its values, or if it is false information being spread

Dev Note:
	This might be a bit glitchy, however it will be more structured as time goes on and modified multiple times
*/

package Geo_Nunbers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type Numbers struct {
	Numbers []Data `json:"International_masks"`
}

type Data struct {
	Mask                        string `json:"mask"`
	Country_Code                string `json:"cc"`
	Country_English             string `json:"name_en"`
	Country_Russian             string `json:"name_ru"`
	Country_English_Description string `json:"desc_en"`
	Country_Russian_Description string `json:"desc_ru"`
}

// us phone numbers

type Us struct {
	Us []Us_Codes `json:"United_States_Phone_Codes_And_Regions"`
}

type Us_Codes struct {
	Dial_Mask                      string `json:"Phone_Dialcode"`
	Dial_Country_Code              string `json:"Country_Code_I"`
	Dial_Country_Currency          string `json:"Country_Currency"`
	Dial_Country_Language          string `json:"Country_Lang"`
	Dial_Country_Code_FULL         string `json:"Country_Code_Full"`
	Dial_Code_State_Location       string `json:"State_of_Number"`
	Dial_Code_State_Location_short string `json:"State_of_Number_Short"`
	Dial_Code_City_Location        string `json:"City_of_number"`
}

// be phone numbers

type Be struct {
	Be []Be_codes `json:"BE_Phone codes"`
}

type Be_codes struct {
	Dial_Mask    string `json:"mask"`
	Country_Code string `json:"cc"`
	Country_Name string `json:"cd"`
	Country_City string `json:"city"`
}

var Countr_Codes = map[string]string{
	"BD": "880",
	"BE": "32",
	"BF": "226",
	"BG": "359",
	"BA": "387",
	"BB": "+1-246",
	"WF": "681",
	"BL": "590",
	"BM": "+1-441",
	"BN": "673",
	"BO": "591",
	"BH": "973",
	"BI": "257",
	"BJ": "229",
	"BT": "975",
	"JM": "+1-876",
	"BV": "",
	"BW": "267",
	"WS": "685",
	"BQ": "599",
	"BR": "55",
	"BS": "+1-242",
	"JE": "+44-1534",
	"BY": "375",
	"BZ": "501",
	"RU": "7",
	"RW": "250",
	"RS": "381",
	"TL": "670",
	"RE": "262",
	"TM": "993",
	"TJ": "992",
	"RO": "40",
	"TK": "690",
	"GW": "245",
	"GU": "+1-671",
	"GT": "502",
	"GS": "",
	"GR": "30",
	"GQ": "240",
	"GP": "590",
	"JP": "81",
	"GY": "592",
	"GG": "+44-1481",
	"GF": "594",
	"GE": "995",
	"GD": "+1-473",
	"GB": "44",
	"GA": "241",
	"SV": "503",
	"GN": "224",
	"GM": "220",
	"GL": "299",
	"GI": "350",
	"GH": "233",
	"OM": "968",
	"TN": "216",
	"JO": "962",
	"HR": "385",
	"HT": "509",
	"HU": "36",
	"HK": "852",
	"HN": "504",
	"HM": " ",
	"VE": "58",
	"PR": "+1-787 and 1-939",
	"PS": "970",
	"PW": "680",
	"PT": "351",
	"SJ": "47",
	"PY": "595",
	"IQ": "964",
	"PA": "507",
	"PF": "689",
	"PG": "675",
	"PE": "51",
	"PK": "92",
	"PH": "63",
	"PN": "870",
	"PL": "48",
	"PM": "508",
	"ZM": "260",
	"EH": "212",
	"EE": "372",
	"EG": "20",
	"ZA": "27",
	"EC": "593",
	"IT": "39",
	"VN": "84",
	"SB": "677",
	"ET": "251",
	"SO": "252",
	"ZW": "263",
	"SA": "966",
	"ES": "34",
	"ER": "291",
	"ME": "382",
	"MD": "373",
	"MG": "261",
	"MF": "590",
	"MA": "212",
	"MC": "377",
	"UZ": "998",
	"MM": "95",
	"ML": "223",
	"MO": "853",
	"MN": "976",
	"MH": "692",
	"MK": "389",
	"MU": "230",
	"MT": "356",
	"MW": "265",
	"MV": "960",
	"MQ": "596",
	"MP": "+1-670",
	"MS": "+1-664",
	"MR": "222",
	"IM": "+44-1624",
	"UG": "256",
	"TZ": "255",
	"MY": "60",
	"MX": "52",
	"IL": "972",
	"FR": "33",
	"IO": "246",
	"SH": "290",
	"FI": "358",
	"FJ": "679",
	"FK": "500",
	"FM": "691",
	"FO": "298",
	"NI": "505",
	"NL": "31",
	"NO": "47",
	"NA": "264",
	"VU": "678",
	"NC": "687",
	"NE": "227",
	"NF": "672",
	"NG": "234",
	"NZ": "64",
	"NP": "977",
	"NR": "674",
	"NU": "683",
	"CK": "682",
	"XK": "",
	"CI": "225",
	"CH": "41",
	"CO": "57",
	"CN": "86",
	"CM": "237",
	"CL": "56",
	"CC": "61",
	"CA": "1",
	"CG": "242",
	"CF": "236",
	"CD": "243",
	"CZ": "420",
	"CY": "357",
	"CX": "61",
	"CR": "506",
	"CW": "599",
	"CV": "238",
	"CU": "53",
	"SZ": "268",
	"SY": "963",
	"SX": "599",
	"KG": "996",
	"KE": "254",
	"SS": "211",
	"SR": "597",
	"KI": "686",
	"KH": "855",
	"KN": "+1-869",
	"KM": "269",
	"ST": "239",
	"SK": "421",
	"KR": "82",
	"SI": "386",
	"KP": "850",
	"KW": "965",
	"SN": "221",
	"SM": "378",
	"SL": "232",
	"SC": "248",
	"KZ": "7",
	"KY": "+1-345",
	"SG": "65",
	"SE": "46",
	"SD": "249",
	"DO": "+1-809 and 1-829",
	"DM": "+1-767",
	"DJ": "253",
	"DK": "45",
	"VG": "+1-284",
	"DE": "49",
	"YE": "967",
	"DZ": "213",
	"US": "1",
	"UY": "598",
	"YT": "262",
	"UM": "1",
	"LB": "961",
	"LC": "+1-758",
	"LA": "856",
	"TV": "688",
	"TW": "886",
	"TT": "+1-868",
	"TR": "90",
	"LK": "94",
	"LI": "423",
	"LV": "371",
	"TO": "676",
	"LT": "370",
	"LU": "352",
	"LR": "231",
	"LS": "266",
	"TH": "66",
	"TF": "",
	"TG": "228",
	"TD": "235",
	"TC": "+1-649",
	"LY": "218",
	"VA": "379",
	"VC": "+1-784",
	"AE": "971",
	"AD": "376",
	"AG": "+1-268",
	"AF": "93",
	"AI": "+1-264",
	"VI": "+1-340",
	"IS": "354",
	"IR": "98",
	"AM": "374",
	"AL": "355",
	"AO": "244",
	"AQ": "",
	"AS": "+1-684",
	"AR": "54",
	"AU": "61",
	"AT": "43",
	"AW": "297",
	"IN": "91",
	"AX": "+358-18",
	"AZ": "994",
	"IE": "353",
	"ID": "62",
	"UA": "380",
	"QA": "974",
	"MZ": "258",
}

var International_Country_Code_To_Currency = map[string]string{
	"BD": "BDT",
	"BE": "EUR",
	"BF": "XOF",
	"BG": "BGN",
	"BA": "BAM",
	"BB": "BBD",
	"WF": "XPF",
	"BL": "EUR",
	"BM": "BMD",
	"BN": "BND",
	"BO": "BOB",
	"BH": "BHD",
	"BI": "BIF",
	"BJ": "XOF",
	"BT": "BTN",
	"JM": "JMD",
	"BV": "NOK",
	"BW": "BWP",
	"WS": "WST",
	"BQ": "USD",
	"BR": "BRL",
	"BS": "BSD",
	"JE": "GBP",
	"BY": "BYR",
	"BZ": "BZD",
	"RU": "RUB",
	"RW": "RWF",
	"RS": "RSD",
	"TL": "USD",
	"RE": "EUR",
	"TM": "TMT",
	"TJ": "TJS",
	"RO": "RON",
	"TK": "NZD",
	"GW": "XOF",
	"GU": "USD",
	"GT": "GTQ",
	"GS": "GBP",
	"GR": "EUR",
	"GQ": "XAF",
	"GP": "EUR",
	"JP": "JPY",
	"GY": "GYD",
	"GG": "GBP",
	"GF": "EUR",
	"GE": "GEL",
	"GD": "XCD",
	"GB": "GBP",
	"GA": "XAF",
	"SV": "USD",
	"GN": "GNF",
	"GM": "GMD",
	"GL": "DKK",
	"GI": "GIP",
	"GH": "GHS",
	"OM": "OMR",
	"TN": "TND",
	"JO": "JOD",
	"HR": "HRK",
	"HT": "HTG",
	"HU": "HUF",
	"HK": "HKD",
	"HN": "HNL",
	"HM": "AUD",
	"VE": "VEF",
	"PR": "USD",
	"PS": "ILS",
	"PW": "USD",
	"PT": "EUR",
	"SJ": "NOK",
	"PY": "PYG",
	"IQ": "IQD",
	"PA": "PAB",
	"PF": "XPF",
	"PG": "PGK",
	"PE": "PEN",
	"PK": "PKR",
	"PH": "PHP",
	"PN": "NZD",
	"PL": "PLN",
	"PM": "EUR",
	"ZM": "ZMK",
	"EH": "MAD",
	"EE": "EUR",
	"EG": "EGP",
	"ZA": "ZAR",
	"EC": "USD",
	"IT": "EUR",
	"VN": "VND",
	"SB": "SBD",
	"ET": "ETB",
	"SO": "SOS",
	"ZW": "ZWL",
	"SA": "SAR",
	"ES": "EUR",
	"ER": "ERN",
	"ME": "EUR",
	"MD": "MDL",
	"MG": "MGA",
	"MF": "EUR",
	"MA": "MAD",
	"MC": "EUR",
	"UZ": "UZS",
	"MM": "MMK",
	"ML": "XOF",
	"MO": "MOP",
	"MN": "MNT",
	"MH": "USD",
	"MK": "MKD",
	"MU": "MUR",
	"MT": "EUR",
	"MW": "MWK",
	"MV": "MVR",
	"MQ": "EUR",
	"MP": "USD",
	"MS": "XCD",
	"MR": "MRO",
	"IM": "GBP",
	"UG": "UGX",
	"TZ": "TZS",
	"MY": "MYR",
	"MX": "MXN",
	"IL": "ILS",
	"FR": "EUR",
	"IO": "USD",
	"SH": "SHP",
	"FI": "EUR",
	"FJ": "FJD",
	"FK": "FKP",
	"FM": "USD",
	"FO": "DKK",
	"NI": "NIO",
	"NL": "EUR",
	"NO": "NOK",
	"NA": "NAD",
	"VU": "VUV",
	"NC": "XPF",
	"NE": "XOF",
	"NF": "AUD",
	"NG": "NGN",
	"NZ": "NZD",
	"NP": "NPR",
	"NR": "AUD",
	"NU": "NZD",
	"CK": "NZD",
	"XK": "EUR",
	"CI": "XOF",
	"CH": "CHF",
	"CO": "COP",
	"CN": "CNY",
	"CM": "XAF",
	"CL": "CLP",
	"CC": "AUD",
	"CA": "CAD",
	"CG": "XAF",
	"CF": "XAF",
	"CD": "CDF",
	"CZ": "CZK",
	"CY": "EUR",
	"CX": "AUD",
	"CR": "CRC",
	"CW": "ANG",
	"CV": "CVE",
	"CU": "CUP",
	"SZ": "SZL",
	"SY": "SYP",
	"SX": "ANG",
	"KG": "KGS",
	"KE": "KES",
	"SS": "SSP",
	"SR": "SRD",
	"KI": "AUD",
	"KH": "KHR",
	"KN": "XCD",
	"KM": "KMF",
	"ST": "STD",
	"SK": "EUR",
	"KR": "KRW",
	"SI": "EUR",
	"KP": "KPW",
	"KW": "KWD",
	"SN": "XOF",
	"SM": "EUR",
	"SL": "SLL",
	"SC": "SCR",
	"KZ": "KZT",
	"KY": "KYD",
	"SG": "SGD",
	"SE": "SEK",
	"SD": "SDG",
	"DO": "DOP",
	"DM": "XCD",
	"DJ": "DJF",
	"DK": "DKK",
	"VG": "USD",
	"DE": "EUR",
	"YE": "YER",
	"DZ": "DZD",
	"US": "USD",
	"UY": "UYU",
	"YT": "EUR",
	"UM": "USD",
	"LB": "LBP",
	"LC": "XCD",
	"LA": "LAK",
	"TV": "AUD",
	"TW": "TWD",
	"TT": "TTD",
	"TR": "TRY",
	"LK": "LKR",
	"LI": "CHF",
	"LV": "EUR",
	"TO": "TOP",
	"LT": "LTL",
	"LU": "EUR",
	"LR": "LRD",
	"LS": "LSL",
	"TH": "THB",
	"TF": "EUR",
	"TG": "XOF",
	"TD": "XAF",
	"TC": "USD",
	"LY": "LYD",
	"VA": "EUR",
	"VC": "XCD",
	"AE": "AED",
	"AD": "EUR",
	"AG": "XCD",
	"AF": "AFN",
	"AI": "XCD",
	"VI": "USD",
	"IS": "ISK",
	"IR": "IRR",
	"AM": "AMD",
	"AL": "ALL",
	"AO": "AOA",
	"AQ": "",
	"AS": "USD",
	"AR": "ARS",
	"AU": "AUD",
	"AT": "EUR",
	"AW": "AWG",
	"IN": "INR",
	"AX": "EUR",
	"AZ": "AZN",
	"IE": "EUR",
	"ID": "IDR",
	"UA": "UAH",
	"QA": "QAR",
	"MZ": "MZN",
}

var Countr_Code_To_Country_Name = map[string]string{
	"BD": "Bangladesh",
	"BE": "Belgium",
	"BF": "Burkina Faso",
	"BG": "Bulgaria",
	"BA": "Bosnia and Herzegovina",
	"BB": "Barbados",
	"WF": "Wallis and Futuna",
	"BL": "Saint Barthelemy",
	"BM": "Bermuda",
	"BN": "Brunei",
	"BO": "Bolivia",
	"BH": "Bahrain",
	"BI": "Burundi",
	"BJ": "Benin",
	"BT": "Bhutan",
	"JM": "Jamaica",
	"BV": "Bouvet Island",
	"BW": "Botswana",
	"WS": "Samoa",
	"BQ": "Bonaire, Saint Eustatius and Saba ",
	"BR": "Brazil",
	"BS": "Bahamas",
	"JE": "Jersey",
	"BY": "Belarus",
	"BZ": "Belize",
	"RU": "Russia",
	"RW": "Rwanda",
	"RS": "Serbia",
	"TL": "East Timor",
	"RE": "Reunion",
	"TM": "Turkmenistan",
	"TJ": "Tajikistan",
	"RO": "Romania",
	"TK": "Tokelau",
	"GW": "Guinea-Bissau",
	"GU": "Guam",
	"GT": "Guatemala",
	"GS": "South Georgia and the South Sandwich Islands",
	"GR": "Greece",
	"GQ": "Equatorial Guinea",
	"GP": "Guadeloupe",
	"JP": "Japan",
	"GY": "Guyana",
	"GG": "Guernsey",
	"GF": "French Guiana",
	"GE": "Georgia",
	"GD": "Grenada",
	"GB": "United Kingdom",
	"GA": "Gabon",
	"SV": "El Salvador",
	"GN": "Guinea",
	"GM": "Gambia",
	"GL": "Greenland",
	"GI": "Gibraltar",
	"GH": "Ghana",
	"OM": "Oman",
	"TN": "Tunisia",
	"JO": "Jordan",
	"HR": "Croatia",
	"HT": "Haiti",
	"HU": "Hungary",
	"HK": "Hong Kong",
	"HN": "Honduras",
	"HM": "Heard Island and McDonald Islands",
	"VE": "Venezuela",
	"PR": "Puerto Rico",
	"PS": "Palestinian Territory",
	"PW": "Palau",
	"PT": "Portugal",
	"SJ": "Svalbard and Jan Mayen",
	"PY": "Paraguay",
	"IQ": "Iraq",
	"PA": "Panama",
	"PF": "French Polynesia",
	"PG": "Papua New Guinea",
	"PE": "Peru",
	"PK": "Pakistan",
	"PH": "Philippines",
	"PN": "Pitcairn",
	"PL": "Poland",
	"PM": "Saint Pierre and Miquelon",
	"ZM": "Zambia",
	"EH": "Western Sahara",
	"EE": "Estonia",
	"EG": "Egypt",
	"ZA": "South Africa",
	"EC": "Ecuador",
	"IT": "Italy",
	"VN": "Vietnam",
	"SB": "Solomon Islands",
	"ET": "Ethiopia",
	"SO": "Somalia",
	"ZW": "Zimbabwe",
	"SA": "Saudi Arabia",
	"ES": "Spain",
	"ER": "Eritrea",
	"ME": "Montenegro",
	"MD": "Moldova",
	"MG": "Madagascar",
	"MF": "Saint Martin",
	"MA": "Morocco",
	"MC": "Monaco",
	"UZ": "Uzbekistan",
	"MM": "Myanmar",
	"ML": "Mali",
	"MO": "Macao",
	"MN": "Mongolia",
	"MH": "Marshall Islands",
	"MK": "Macedonia",
	"MU": "Mauritius",
	"MT": "Malta",
	"MW": "Malawi",
	"MV": "Maldives",
	"MQ": "Martinique",
	"MP": "Northern Mariana Islands",
	"MS": "Montserrat",
	"MR": "Mauritania",
	"IM": "Isle of Man",
	"UG": "Uganda",
	"TZ": "Tanzania",
	"MY": "Malaysia",
	"MX": "Mexico",
	"IL": "Israel",
	"FR": "France",
	"IO": "British Indian Ocean Territory",
	"SH": "Saint Helena",
	"FI": "Finland",
	"FJ": "Fiji",
	"FK": "Falkland Islands",
	"FM": "Micronesia",
	"FO": "Faroe Islands",
	"NI": "Nicaragua",
	"NL": "Netherlands",
	"NO": "Norway",
	"NA": "Namibia",
	"VU": "Vanuatu",
	"NC": "New Caledonia",
	"NE": "Niger",
	"NF": "Norfolk Island",
	"NG": "Nigeria",
	"NZ": "New Zealand",
	"NP": "Nepal",
	"NR": "Nauru",
	"NU": "Niue",
	"CK": "Cook Islands",
	"XK": "Kosovo",
	"CI": "Ivory Coast",
	"CH": "Switzerland",
	"CO": "Colombia",
	"CN": "China",
	"CM": "Cameroon",
	"CL": "Chile",
	"CC": "Cocos Islands",
	"CA": "Canada",
	"CG": "Republic of the Congo",
	"CF": "Central African Republic",
	"CD": "Democratic Republic of the Congo",
	"CZ": "Czech Republic",
	"CY": "Cyprus",
	"CX": "Christmas Island",
	"CR": "Costa Rica",
	"CW": "Curacao",
	"CV": "Cape Verde",
	"CU": "Cuba",
	"SZ": "Swaziland",
	"SY": "Syria",
	"SX": "Sint Maarten",
	"KG": "Kyrgyzstan",
	"KE": "Kenya",
	"SS": "South Sudan",
	"SR": "Suriname",
	"KI": "Kiribati",
	"KH": "Cambodia",
	"KN": "Saint Kitts and Nevis",
	"KM": "Comoros",
	"ST": "Sao Tome and Principe",
	"SK": "Slovakia",
	"KR": "South Korea",
	"SI": "Slovenia",
	"KP": "North Korea",
	"KW": "Kuwait",
	"SN": "Senegal",
	"SM": "San Marino",
	"SL": "Sierra Leone",
	"SC": "Seychelles",
	"KZ": "Kazakhstan",
	"KY": "Cayman Islands",
	"SG": "Singapore",
	"SE": "Sweden",
	"SD": "Sudan",
	"DO": "Dominican Republic",
	"DM": "Dominica",
	"DJ": "Djibouti",
	"DK": "Denmark",
	"VG": "British Virgin Islands",
	"DE": "Germany",
	"YE": "Yemen",
	"DZ": "Algeria",
	"US": "United States",
	"UY": "Uruguay",
	"YT": "Mayotte",
	"UM": "United States Minor Outlying Islands",
	"LB": "Lebanon",
	"LC": "Saint Lucia",
	"LA": "Laos",
	"TV": "Tuvalu",
	"TW": "Taiwan",
	"TT": "Trinidad and Tobago",
	"TR": "Turkey",
	"LK": "Sri Lanka",
	"LI": "Liechtenstein",
	"LV": "Latvia",
	"TO": "Tonga",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"LR": "Liberia",
	"LS": "Lesotho",
	"TH": "Thailand",
	"TF": "French Southern Territories",
	"TG": "Togo",
	"TD": "Chad",
	"TC": "Turks and Caicos Islands",
	"LY": "Libya",
	"VA": "Vatican",
	"VC": "Saint Vincent and the Grenadines",
	"AE": "United Arab Emirates",
	"AD": "Andorra",
	"AG": "Antigua and Barbuda",
	"AF": "Afghanistan",
	"AI": "Anguilla",
	"VI": "U.S. Virgin Islands",
	"IS": "Iceland",
	"IR": "Iran",
	"AM": "Armenia",
	"AL": "Albania",
	"AO": "Angola",
	"AQ": "Antarctica",
	"AS": "American Samoa",
	"AR": "Argentina",
	"AU": "Australia",
	"AT": "Austria",
	"AW": "Aruba",
	"IN": "India",
	"AX": "Aland Islands",
	"AZ": "Azerbaijan",
	"IE": "Ireland",
	"ID": "Indonesia",
	"UA": "Ukraine",
	"QA": "Qatar",
	"MZ": "Mozambique",
}

var Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3 = map[string]string{
	"BD": "BGD",
	"BE": "BEL",
	"BF": "BFA",
	"BG": "BGR",
	"BA": "BIH",
	"BB": "BRB",
	"WF": "WLF",
	"BL": "BLM",
	"BM": "BMU",
	"BN": "BRN",
	"BO": "BOL",
	"BH": "BHR",
	"BI": "BDI",
	"BJ": "BEN",
	"BT": "BTN",
	"JM": "JAM",
	"BV": "BVT",
	"BW": "BWA",
	"WS": "WSM",
	"BQ": "BES",
	"BR": "BRA",
	"BS": "BHS",
	"JE": "JEY",
	"BY": "BLR",
	"BZ": "BLZ",
	"RU": "RUS",
	"RW": "RWA",
	"RS": "SRB",
	"TL": "TLS",
	"RE": "REU",
	"TM": "TKM",
	"TJ": "TJK",
	"RO": "ROU",
	"TK": "TKL",
	"GW": "GNB",
	"GU": "GUM",
	"GT": "GTM",
	"GS": "SGS",
	"GR": "GRC",
	"GQ": "GNQ",
	"GP": "GLP",
	"JP": "JPN",
	"GY": "GUY",
	"GG": "GGY",
	"GF": "GUF",
	"GE": "GEO",
	"GD": "GRD",
	"GB": "GBR",
	"GA": "GAB",
	"SV": "SLV",
	"GN": "GIN",
	"GM": "GMB",
	"GL": "GRL",
	"GI": "GIB",
	"GH": "GHA",
	"OM": "OMN",
	"TN": "TUN",
	"JO": "JOR",
	"HR": "HRV",
	"HT": "HTI",
	"HU": "HUN",
	"HK": "HKG",
	"HN": "HND",
	"HM": "HMD",
	"VE": "VEN",
	"PR": "PRI",
	"PS": "PSE",
	"PW": "PLW",
	"PT": "PRT",
	"SJ": "SJM",
	"PY": "PRY",
	"IQ": "IRQ",
	"PA": "PAN",
	"PF": "PYF",
	"PG": "PNG",
	"PE": "PER",
	"PK": "PAK",
	"PH": "PHL",
	"PN": "PCN",
	"PL": "POL",
	"PM": "SPM",
	"ZM": "ZMB",
	"EH": "ESH",
	"EE": "EST",
	"EG": "EGY",
	"ZA": "ZAF",
	"EC": "ECU",
	"IT": "ITA",
	"VN": "VNM",
	"SB": "SLB",
	"ET": "ETH",
	"SO": "SOM",
	"ZW": "ZWE",
	"SA": "SAU",
	"ES": "ESP",
	"ER": "ERI",
	"ME": "MNE",
	"MD": "MDA",
	"MG": "MDG",
	"MF": "MAF",
	"MA": "MAR",
	"MC": "MCO",
	"UZ": "UZB",
	"MM": "MMR",
	"ML": "MLI",
	"MO": "MAC",
	"MN": "MNG",
	"MH": "MHL",
	"MK": "MKD",
	"MU": "MUS",
	"MT": "MLT",
	"MW": "MWI",
	"MV": "MDV",
	"MQ": "MTQ",
	"MP": "MNP",
	"MS": "MSR",
	"MR": "MRT",
	"IM": "IMN",
	"UG": "UGA",
	"TZ": "TZA",
	"MY": "MYS",
	"MX": "MEX",
	"IL": "ISR",
	"FR": "FRA",
	"IO": "IOT",
	"SH": "SHN",
	"FI": "FIN",
	"FJ": "FJI",
	"FK": "FLK",
	"FM": "FSM",
	"FO": "FRO",
	"NI": "NIC",
	"NL": "NLD",
	"NO": "NOR",
	"NA": "NAM",
	"VU": "VUT",
	"NC": "NCL",
	"NE": "NER",
	"NF": "NFK",
	"NG": "NGA",
	"NZ": "NZL",
	"NP": "NPL",
	"NR": "NRU",
	"NU": "NIU",
	"CK": "COK",
	"XK": "XKX",
	"CI": "CIV",
	"CH": "CHE",
	"CO": "COL",
	"CN": "CHN",
	"CM": "CMR",
	"CL": "CHL",
	"CC": "CCK",
	"CA": "CAN",
	"CG": "COG",
	"CF": "CAF",
	"CD": "COD",
	"CZ": "CZE",
	"CY": "CYP",
	"CX": "CXR",
	"CR": "CRI",
	"CW": "CUW",
	"CV": "CPV",
	"CU": "CUB",
	"SZ": "SWZ",
	"SY": "SYR",
	"SX": "SXM",
	"KG": "KGZ",
	"KE": "KEN",
	"SS": "SSD",
	"SR": "SUR",
	"KI": "KIR",
	"KH": "KHM",
	"KN": "KNA",
	"KM": "COM",
	"ST": "STP",
	"SK": "SVK",
	"KR": "KOR",
	"SI": "SVN",
	"KP": "PRK",
	"KW": "KWT",
	"SN": "SEN",
	"SM": "SMR",
	"SL": "SLE",
	"SC": "SYC",
	"KZ": "KAZ",
	"KY": "CYM",
	"SG": "SGP",
	"SE": "SWE",
	"SD": "SDN",
	"DO": "DOM",
	"DM": "DMA",
	"DJ": "DJI",
	"DK": "DNK",
	"VG": "VGB",
	"DE": "DEU",
	"YE": "YEM",
	"DZ": "DZA",
	"US": "USA",
	"UY": "URY",
	"YT": "MYT",
	"UM": "UMI",
	"LB": "LBN",
	"LC": "LCA",
	"LA": "LAO",
	"TV": "TUV",
	"TW": "TWN",
	"TT": "TTO",
	"TR": "TUR",
	"LK": "LKA",
	"LI": "LIE",
	"LV": "LVA",
	"TO": "TON",
	"LT": "LTU",
	"LU": "LUX",
	"LR": "LBR",
	"LS": "LSO",
	"TH": "THA",
	"TF": "ATF",
	"TG": "TGO",
	"TD": "TCD",
	"TC": "TCA",
	"LY": "LBY",
	"VA": "VAT",
	"VC": "VCT",
	"AE": "ARE",
	"AD": "AND",
	"AG": "ATG",
	"AF": "AFG",
	"AI": "AIA",
	"VI": "VIR",
	"IS": "ISL",
	"IR": "IRN",
	"AM": "ARM",
	"AL": "ALB",
	"AO": "AGO",
	"AQ": "ATA",
	"AS": "ASM",
	"AR": "ARG",
	"AU": "AUS",
	"AT": "AUT",
	"AW": "ABW",
	"IN": "IND",
	"AX": "ALA",
	"AZ": "AZE",
	"IE": "IRL",
	"ID": "IDN",
	"UA": "UKR",
	"QA": "QAT",
	"MZ": "MOZ",
}

var Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3_To_Country_Capital = map[string]string{
	"BD": "Dhaka",
	"BE": "Brussels",
	"BF": "Ouagadougou",
	"BG": "Sofia",
	"BA": "Sarajevo",
	"BB": "Bridgetown",
	"WF": "Mata Utu",
	"BL": "Gustavia",
	"BM": "Hamilton",
	"BN": "Bandar Seri Begawan",
	"BO": "Sucre",
	"BH": "Manama",
	"BI": "Bujumbura",
	"BJ": "Porto-Novo",
	"BT": "Thimphu",
	"JM": "Kingston",
	"BV": "",
	"BW": "Gaborone",
	"WS": "Apia",
	"BQ": "",
	"BR": "Brasilia",
	"BS": "Nassau",
	"JE": "Saint Helier",
	"BY": "Minsk",
	"BZ": "Belmopan",
	"RU": "Moscow",
	"RW": "Kigali",
	"RS": "Belgrade",
	"TL": "Dili",
	"RE": "Saint-Denis",
	"TM": "Ashgabat",
	"TJ": "Dushanbe",
	"RO": "Bucharest",
	"TK": "",
	"GW": "Bissau",
	"GU": "Hagatna",
	"GT": "Guatemala City",
	"GS": "Grytviken",
	"GR": "Athens",
	"GQ": "Malabo",
	"GP": "Basse-Terre",
	"JP": "Tokyo",
	"GY": "Georgetown",
	"GG": "St Peter Port",
	"GF": "Cayenne",
	"GE": "Tbilisi",
	"GD": "St. George's",
	"GB": "London",
	"GA": "Libreville",
	"SV": "San Salvador",
	"GN": "Conakry",
	"GM": "Banjul",
	"GL": "Nuuk",
	"GI": "Gibraltar",
	"GH": "Accra",
	"OM": "Muscat",
	"TN": "Tunis",
	"JO": "Amman",
	"HR": "Zagreb",
	"HT": "Port-au-Prince",
	"HU": "Budapest",
	"HK": "Hong Kong",
	"HN": "Tegucigalpa",
	"HM": "",
	"VE": "Caracas",
	"PR": "San Juan",
	"PS": "East Jerusalem",
	"PW": "Melekeok",
	"PT": "Lisbon",
	"SJ": "Longyearbyen",
	"PY": "Asuncion",
	"IQ": "Baghdad",
	"PA": "Panama City",
	"PF": "Papeete",
	"PG": "Port Moresby",
	"PE": "Lima",
	"PK": "Islamabad",
	"PH": "Manila",
	"PN": "Adamstown",
	"PL": "Warsaw",
	"PM": "Saint-Pierre",
	"ZM": "Lusaka",
	"EH": "El-Aaiun",
	"EE": "Tallinn",
	"EG": "Cairo",
	"ZA": "Pretoria",
	"EC": "Quito",
	"IT": "Rome",
	"VN": "Hanoi",
	"SB": "Honiara",
	"ET": "Addis Ababa",
	"SO": "Mogadishu",
	"ZW": "Harare",
	"SA": "Riyadh",
	"ES": "Madrid",
	"ER": "Asmara",
	"ME": "Podgorica",
	"MD": "Chisinau",
	"MG": "Antananarivo",
	"MF": "Marigot",
	"MA": "Rabat",
	"MC": "Monaco",
	"UZ": "Tashkent",
	"MM": "Nay Pyi Taw",
	"ML": "Bamako",
	"MO": "Macao",
	"MN": "Ulan Bator",
	"MH": "Majuro",
	"MK": "Skopje",
	"MU": "Port Louis",
	"MT": "Valletta",
	"MW": "Lilongwe",
	"MV": "Male",
	"MQ": "Fort-de-France",
	"MP": "Saipan",
	"MS": "Plymouth",
	"MR": "Nouakchott",
	"IM": "Douglas, Isle of Man",
	"UG": "Kampala",
	"TZ": "Dodoma",
	"MY": "Kuala Lumpur",
	"MX": "Mexico City",
	"IL": "Jerusalem",
	"FR": "Paris",
	"IO": "Diego Garcia",
	"SH": "Jamestown",
	"FI": "Helsinki",
	"FJ": "Suva",
	"FK": "Stanley",
	"FM": "Palikir",
	"FO": "Torshavn",
	"NI": "Managua",
	"NL": "Amsterdam",
	"NO": "Oslo",
	"NA": "Windhoek",
	"VU": "Port Vila",
	"NC": "Noumea",
	"NE": "Niamey",
	"NF": "Kingston",
	"NG": "Abuja",
	"NZ": "Wellington",
	"NP": "Kathmandu",
	"NR": "Yaren",
	"NU": "Alofi",
	"CK": "Avarua",
	"XK": "Pristina",
	"CI": "Yamoussoukro",
	"CH": "Berne",
	"CO": "Bogota",
	"CN": "Beijing",
	"CM": "Yaounde",
	"CL": "Santiago",
	"CC": "West Island",
	"CA": "Ottawa",
	"CG": "Brazzaville",
	"CF": "Bangui",
	"CD": "Kinshasa",
	"CZ": "Prague",
	"CY": "Nicosia",
	"CX": "Flying Fish Cove",
	"CR": "San Jose",
	"CW": " Willemstad",
	"CV": "Praia",
	"CU": "Havana",
	"SZ": "Mbabane",
	"SY": "Damascus",
	"SX": "Philipsburg",
	"KG": "Bishkek",
	"KE": "Nairobi",
	"SS": "Juba",
	"SR": "Paramaribo",
	"KI": "Tarawa",
	"KH": "Phnom Penh",
	"KN": "Basseterre",
	"KM": "Moroni",
	"ST": "Sao Tome",
	"SK": "Bratislava",
	"KR": "Seoul",
	"SI": "Ljubljana",
	"KP": "Pyongyang",
	"KW": "Kuwait City",
	"SN": "Dakar",
	"SM": "San Marino",
	"SL": "Freetown",
	"SC": "Victoria",
	"KZ": "Astana",
	"KY": "George Town",
	"SG": "Singapur",
	"SE": "Stockholm",
	"SD": "Khartoum",
	"DO": "Santo Domingo",
	"DM": "Roseau",
	"DJ": "Djibouti",
	"DK": "Copenhagen",
	"VG": "Road Town",
	"DE": "Berlin",
	"YE": "Sanaa",
	"DZ": "Algiers",
	"US": "Washington",
	"UY": "Montevideo",
	"YT": "Mamoudzou",
	"UM": "",
	"LB": "Beirut",
	"LC": "Castries",
	"LA": "Vientiane",
	"TV": "Funafuti",
	"TW": "Taipei",
	"TT": "Port of Spain",
	"TR": "Ankara",
	"LK": "Colombo",
	"LI": "Vaduz",
	"LV": "Riga",
	"TO": "Nuku'alofa",
	"LT": "Vilnius",
	"LU": "Luxembourg",
	"LR": "Monrovia",
	"LS": "Maseru",
	"TH": "Bangkok",
	"TF": "Port-aux-Francais",
	"TG": "Lome",
	"TD": "N'Djamena",
	"TC": "Cockburn Town",
	"LY": "Tripolis",
	"VA": "Vatican City",
	"VC": "Kingstown",
	"AE": "Abu Dhabi",
	"AD": "Andorra la Vella",
	"AG": "St. John's",
	"AF": "Kabul",
	"AI": "The Valley",
	"VI": "Charlotte Amalie",
	"IS": "Reykjavik",
	"IR": "Tehran",
	"AM": "Yerevan",
	"AL": "Tirana",
	"AO": "Luanda",
	"AQ": "",
	"AS": "Pago Pago",
	"AR": "Buenos Aires",
	"AU": "Canberra",
	"AT": "Vienna",
	"AW": "Oranjestad",
	"IN": "New Delhi",
	"AX": "Mariehamn",
	"AZ": "Baku",
	"IE": "Dublin",
	"ID": "Jakarta",
	"UA": "Kiev",
	"QA": "Doha",
	"MZ": "Maputo",
}

func Parser(phone_num string) {
	db, e := os.Open("json/phone-code.json")
	if e != nil {
		log.Fatal("<RR6> OSINT module: Could not open up json phone coders database file, this is NOT supposed to happen")
	} else {
		defer db.Close()
		valueb, _ := ioutil.ReadAll(db)
		var values Numbers
		json.Unmarshal(valueb, &values)
		for j := 0; j < len(values.Numbers); j++ {
			if reflect.DeepEqual(phone_num, values.Numbers[j].Mask) {
				Watch_Dog_5 := Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3_To_Country_Capital[values.Numbers[j].Country_Code]
				Watch_Dog_4 := Countr_Code_To_FULL_Country_Code_AKA_ISO2_To_ISO3[values.Numbers[j].Country_Code]
				Watch_Dog_3 := Countr_Code_To_Country_Name[values.Numbers[j].Country_Code]
				Watch_Dog_2 := International_Country_Code_To_Currency[values.Numbers[j].Country_Code]
				Watch_Dog_1 := Countr_Codes[values.Numbers[j].Country_Code]

				fmt.Printf("Phone Value | Phone Mask                           | %s\n", values.Numbers[j].Mask)
				fmt.Printf("Phone Value | Location Name English                | %s\n", values.Numbers[j].Country_English)
				fmt.Printf("Phone Value | Location desc English                | %s\n", values.Numbers[j].Country_English_Description)
				fmt.Printf("Phone Value | Location Country code                | %s\n", values.Numbers[j].Country_Code)
				fmt.Printf("Phone Value | Location Russian name                | %s\n", values.Numbers[j].Country_Russian)
				fmt.Printf("Phone Value | Location Russian desc                | %s\n", values.Numbers[j].Country_Russian_Description)
				fmt.Printf("Phone Value | Country code to prefix               | %s\n", Watch_Dog_1)
				fmt.Printf("Phone Value | Country code to currency             | %s\n", Watch_Dog_2)
				fmt.Printf("Phone Value | Country Code to Country name         | %s\n", Watch_Dog_3)
				fmt.Printf("Phone Value | Country ISO2 to Country ISO3         | %s\n", Watch_Dog_4)
				fmt.Printf("Phone Value | Country ISO2 to Country ISO3 Capital | %s\n", Watch_Dog_5)
			}
		}
	}
}

func Parser_US(ph_Type bool, phone_prefix string) {
	udb, e := os.Open("json/US-Codes.json")
	if e != nil {
		log.Fatal("<RR6> OSINT module: Could not open up json phone coders database file, this is NOT supposed to happen")
	} else {
		defer udb.Close()
		valueb, _ := ioutil.ReadAll(udb)
		var values Us
		json.Unmarshal(valueb, &values)
		for k := 0; k < len(values.Us); k++ {
			if reflect.DeepEqual(phone_prefix, values.Us[k].Dial_Mask) {
				fmt.Printf("Phone Value | Phone Mask                           | %s\n", values.Us[k].Dial_Mask)
				fmt.Printf("Phone Value | Phone Dial Country Language          | %s\n", values.Us[k].Dial_Country_Language)
				fmt.Printf("Phone Value | Phone Dial Country Curreny           | %s\n", values.Us[k].Dial_Country_Currency)
				fmt.Printf("Phone Value | Phone Dial Country Code              | %s\n", values.Us[k].Dial_Country_Code)
				fmt.Printf("Phone Value | Phone Dial Country Code Full         | %s\n", values.Us[k].Dial_Country_Code_FULL)
				fmt.Printf("Phone Value | Phone Dial State Location Short      | %s\n", values.Us[k].Dial_Code_State_Location_short)
				fmt.Printf("Phone Value | Phone Dial State of Number           | %s\n", values.Us[k].Dial_Code_State_Location)
				fmt.Printf("Phone Value | Phone Dial City of Number            | %s\n", values.Us[k].Dial_Code_City_Location)
			}
		}
	}
}

func Parser_BE(phone_prefix string) {
	bdb, e := os.Open("json/be.json")
	if e != nil {
		log.Fatal("<RR6> OSINT module: Could not open up json phone coders database file, this is NOT supposed to happen")
	} else {
		defer bdb.Close()
		valueb, _ := ioutil.ReadAll(bdb)
		var values Be
		json.Unmarshal(valueb, &values)
		for k := 0; k < len(values.Be); k++ {
			Watch_Dog_1 := Countr_Code_To_Country_Name[values.Be[k].Country_Code]
			Watch_Dog_2 := International_Country_Code_To_Currency[values.Be[k].Country_Code]
			if reflect.DeepEqual(phone_prefix, values.Be[k].Dial_Mask) {
				fmt.Printf("Phone Value | Phone Mask                           | %s\n", values.Be[k].Dial_Mask)
				fmt.Printf("Phone Value | Phone Dial Country Code              | %s\n", values.Be[k].Country_Code)
				fmt.Printf("Phone Value | Phone Dial Country Name              | %s\n", values.Be[k].Country_Name)
				fmt.Printf("Phone Value | Phone Dial Country code to name      | %s\n", Watch_Dog_1)
				fmt.Printf("Phone Value | Phone Dial Country code to currency  | %s\n", Watch_Dog_2)
				fmt.Printf("Phone Value | Phone Dial Country code to currency  | %s\n", values.Be[k].Country_City)
			}
		}
	}
}