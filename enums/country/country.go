package country

type country struct {
	Code   string
	CnName string
	EnName string
}

var (
	NZL                                        = &country{Code: "NZL", CnName: "新西兰", EnName: "New Zealand"}
	FJI                                        = &country{Code: "FJI", CnName: "斐济", EnName: "Fiji"}
	PNG                                        = &country{Code: "PNG", CnName: "巴布亚新几内亚", EnName: "Papua New Guinea"}
	GLP                                        = &country{Code: "GLP", CnName: "瓜德罗普岛", EnName: "Guadeloupe"}
	STP                                        = &country{Code: "STP", CnName: "圣多美和普林西比", EnName: "Sao Tome and Principe"}
	MHL                                        = &country{Code: "MHL", CnName: "马绍尔群岛", EnName: "Marshall Islands (the)"}
	CUB                                        = &country{Code: "CUB", CnName: "古巴", EnName: "Cuba"}
	SDN                                        = &country{Code: "SDN", CnName: "苏丹", EnName: "Sudan (the)"}
	GMB                                        = &country{Code: "GMB", CnName: "冈比亚", EnName: "Gambia (the)"}
	CUW                                        = &country{Code: "CUW", CnName: "库拉索", EnName: "Curaçao"}
	MYS                                        = &country{Code: "MYS", CnName: "马来西亚", EnName: "Malaysia"}
	MYT                                        = &country{Code: "MYT", CnName: "马约特", EnName: "Mayotte"}
	POL                                        = &country{Code: "POL", CnName: "波兰", EnName: "Poland"}
	OMN                                        = &country{Code: "OMN", CnName: "阿曼", EnName: "Oman"}
	SUR                                        = &country{Code: "SUR", CnName: "苏里南", EnName: "Suriname"}
	ARE                                        = &country{Code: "ARE", CnName: "阿拉伯联合酋长国", EnName: "United Arab Emirates (the)"}
	KEN                                        = &country{Code: "KEN", CnName: "肯尼亚", EnName: "Kenya"}
	ARG                                        = &country{Code: "ARG", CnName: "阿根廷", EnName: "Argentina"}
	GNB                                        = &country{Code: "GNB", CnName: "几内亚比绍", EnName: "Guinea-Bissau"}
	ARM                                        = &country{Code: "ARM", CnName: "亚美尼亚", EnName: "Armenia"}
	UZB                                        = &country{Code: "UZB", CnName: "乌兹别克斯坦", EnName: "Uzbekistan"}
	BTN                                        = &country{Code: "BTN", CnName: "不丹", EnName: "Bhutan"}
	SEN                                        = &country{Code: "SEN", CnName: "塞内加尔", EnName: "Senegal"}
	TGO                                        = &country{Code: "TGO", CnName: "多哥", EnName: "Togo"}
	IRL                                        = &country{Code: "IRL", CnName: "爱尔兰", EnName: "Ireland"}
	FLK                                        = &country{Code: "FLK", CnName: "福克兰群岛（马尔维纳斯群岛）", EnName: "Falkland Islands (the) [Malvinas]"}
	IRN                                        = &country{Code: "IRN", CnName: "伊朗（伊斯兰共和国）", EnName: "Iran (Islamic Republic of)"}
	QAT                                        = &country{Code: "QAT", CnName: "卡塔尔", EnName: "Qatar"}
	BDI                                        = &country{Code: "BDI", CnName: "布隆迪", EnName: "Burundi"}
	NLD                                        = &country{Code: "NLD", CnName: "荷兰", EnName: "Netherlands (the)"}
	IRQ                                        = &country{Code: "IRQ", CnName: "伊拉克", EnName: "Iraq"}
	SVK                                        = &country{Code: "SVK", CnName: "斯洛伐克", EnName: "Slovakia"}
	SVN                                        = &country{Code: "SVN", CnName: "斯洛文尼亚", EnName: "Slovenia"}
	GNQ                                        = &country{Code: "GNQ", CnName: "赤道几内亚", EnName: "Equatorial Guinea"}
	THA                                        = &country{Code: "THA", CnName: "泰国", EnName: "Thailand"}
	ABW                                        = &country{Code: "ABW", CnName: "阿鲁巴", EnName: "Aruba"}
	ASM                                        = &country{Code: "ASM", CnName: "美属萨摩亚", EnName: "American Samoa"}
	SWE                                        = &country{Code: "SWE", CnName: "瑞典", EnName: "Sweden"}
	ISL                                        = &country{Code: "ISL", CnName: "冰岛", EnName: "Iceland"}
	BEL                                        = &country{Code: "BEL", CnName: "比利时", EnName: "Belgium"}
	ISR                                        = &country{Code: "ISR", CnName: "以色列", EnName: "Israel"}
	KWT                                        = &country{Code: "KWT", CnName: "科威特", EnName: "Kuwait"}
	LIE                                        = &country{Code: "LIE", CnName: "列支敦士登", EnName: "Liechtenstein"}
	DZA                                        = &country{Code: "DZA", CnName: "阿尔及利亚", EnName: "Algeria"}
	BEN                                        = &country{Code: "BEN", CnName: "贝宁", EnName: "Benin"}
	ATA                                        = &country{Code: "ATA", CnName: "南极洲", EnName: "Antarctica"}
	RUS                                        = &country{Code: "RUS", CnName: "俄罗斯联邦", EnName: "Russian Federation (the)"}
	ATF                                        = &country{Code: "ATF", CnName: "法属南部领土", EnName: "French Southern Territories (the)"}
	ATG                                        = &country{Code: "ATG", CnName: "安提瓜和巴布达", EnName: "Antigua and Barbuda"}
	SWZ                                        = &country{Code: "SWZ", CnName: "斯威士兰", EnName: "Eswatini"}
	ITA                                        = &country{Code: "ITA", CnName: "意大利", EnName: "Italy"}
	TZA                                        = &country{Code: "TZA", CnName: "坦桑尼亚联合共和国", EnName: "Tanzania, the United Republic of"}
	PAK                                        = &country{Code: "PAK", CnName: "巴基斯坦", EnName: "Pakistan"}
	BFA                                        = &country{Code: "BFA", CnName: "布基纳法索", EnName: "Burkina Faso"}
	CXR                                        = &country{Code: "CXR", CnName: "圣诞岛", EnName: "Christmas Island"}
	PAN                                        = &country{Code: "PAN", CnName: "巴拿马", EnName: "Panama"}
	SGP                                        = &country{Code: "SGP", CnName: "新加坡", EnName: "Singapore"}
	UKR                                        = &country{Code: "UKR", CnName: "乌克兰", EnName: "Ukraine"}
	SGS                                        = &country{Code: "SGS", CnName: "南乔治亚和南桑威奇群岛", EnName: "South Georgia and the South Sandwich Islands"}
	KGZ                                        = &country{Code: "KGZ", CnName: "吉尔吉斯斯坦", EnName: "Kyrgyzstan"}
	BVT                                        = &country{Code: "BVT", CnName: "布韦岛", EnName: "Bouvet Island"}
	CHE                                        = &country{Code: "CHE", CnName: "瑞士", EnName: "Switzerland"}
	DJI                                        = &country{Code: "DJI", CnName: "吉布提", EnName: "Djibouti"}
	REU                                        = &country{Code: "REU", CnName: "留尼汪", EnName: "Réunion"}
	CHL                                        = &country{Code: "CHL", CnName: "智利", EnName: "Chile"}
	PRI                                        = &country{Code: "PRI", CnName: "波多黎各", EnName: "Puerto Rico"}
	CHN                                        = &country{Code: "CHN", CnName: "中国", EnName: "China"}
	PRK                                        = &country{Code: "PRK", CnName: "朝鲜民主主义人民共和国", EnName: "Korea (the Democratic People's Republic of)"}
	MLI                                        = &country{Code: "MLI", CnName: "马里", EnName: "Mali"}
	BWA                                        = &country{Code: "BWA", CnName: "博茨瓦纳", EnName: "Botswana"}
	HRV                                        = &country{Code: "HRV", CnName: "克罗地亚", EnName: "Croatia"}
	KHM                                        = &country{Code: "KHM", CnName: "柬埔寨", EnName: "Cambodia"}
	IDN                                        = &country{Code: "IDN", CnName: "印度尼西亚", EnName: "Indonesia"}
	PRT                                        = &country{Code: "PRT", CnName: "葡萄牙", EnName: "Portugal"}
	MLT                                        = &country{Code: "MLT", CnName: "马耳他", EnName: "Malta"}
	TJK                                        = &country{Code: "TJK", CnName: "塔吉克斯坦", EnName: "Tajikistan"}
	VNM                                        = &country{Code: "VNM", CnName: "越南", EnName: "Viet Nam"}
	CYM                                        = &country{Code: "CYM", CnName: "开曼群岛", EnName: "Cayman Islands (the)"}
	PRY                                        = &country{Code: "PRY", CnName: "巴拉圭", EnName: "Paraguay"}
	SHN                                        = &country{Code: "SHN", CnName: "圣赫勒拿、阿森松和特里斯坦达库尼亚", EnName: "Saint Helena, Ascension and Tristan da Cunha"}
	CYP                                        = &country{Code: "CYP", CnName: "塞浦路斯", EnName: "Cyprus"}
	SYC                                        = &country{Code: "SYC", CnName: "塞舌尔", EnName: "Seychelles"}
	RWA                                        = &country{Code: "RWA", CnName: "卢旺达", EnName: "Rwanda"}
	BGD                                        = &country{Code: "BGD", CnName: "孟加拉国", EnName: "Bangladesh"}
	AUS                                        = &country{Code: "AUS", CnName: "澳大利亚", EnName: "Australia"}
	AUT                                        = &country{Code: "AUT", CnName: "奥地利", EnName: "Austria"}
	PSE                                        = &country{Code: "PSE", CnName: "巴勒斯坦国", EnName: "Palestine, State of"}
	LKA                                        = &country{Code: "LKA", CnName: "斯里兰卡", EnName: "Sri Lanka"}
	GAB                                        = &country{Code: "GAB", CnName: "加蓬", EnName: "Gabon"}
	ZWE                                        = &country{Code: "ZWE", CnName: "津巴布韦", EnName: "Zimbabwe"}
	BGR                                        = &country{Code: "BGR", CnName: "保加利亚", EnName: "Bulgaria"}
	NOR                                        = &country{Code: "NOR", CnName: "挪威", EnName: "Norway"}
	CIV                                        = &country{Code: "CIV", CnName: "科特迪瓦", EnName: "Côte d'Ivoire"}
	MMR                                        = &country{Code: "MMR", CnName: "缅甸", EnName: "Myanmar"}
	TKL                                        = &country{Code: "TKL", CnName: "托克劳", EnName: "Tokelau"}
	KIR                                        = &country{Code: "KIR", CnName: "基里巴斯", EnName: "Kiribati"}
	TKM                                        = &country{Code: "TKM", CnName: "土库曼斯坦", EnName: "Turkmenistan"}
	GRD                                        = &country{Code: "GRD", CnName: "格林纳达", EnName: "Grenada"}
	GRC                                        = &country{Code: "GRC", CnName: "希腊", EnName: "Greece"}
	PCN                                        = &country{Code: "PCN", CnName: "皮特凯恩", EnName: "Pitcairn"}
	HTI                                        = &country{Code: "HTI", CnName: "海地", EnName: "Haiti"}
	GRL                                        = &country{Code: "GRL", CnName: "格陵兰岛", EnName: "Greenland"}
	YEM                                        = &country{Code: "YEM", CnName: "也门", EnName: "Yemen"}
	AFG                                        = &country{Code: "AFG", CnName: "阿富汗", EnName: "Afghanistan"}
	MNE                                        = &country{Code: "MNE", CnName: "黑山", EnName: "Montenegro"}
	MNG                                        = &country{Code: "MNG", CnName: "蒙古", EnName: "Mongolia"}
	NPL                                        = &country{Code: "NPL", CnName: "尼泊尔", EnName: "Nepal"}
	BHS                                        = &country{Code: "BHS", CnName: "巴哈马", EnName: "Bahamas (the)"}
	BHR                                        = &country{Code: "BHR", CnName: "巴林", EnName: "Bahrain"}
	MNP                                        = &country{Code: "MNP", CnName: "北马里亚纳群岛", EnName: "Northern Mariana Islands (the)"}
	GBR                                        = &country{Code: "GBR", CnName: "大不列颠及北爱尔兰联合王国", EnName: "United Kingdom of Great Britain and Northern Ireland (the)"}
	DMA                                        = &country{Code: "DMA", CnName: "多米尼克", EnName: "Dominica"}
	TLS                                        = &country{Code: "TLS", CnName: "东帝汶", EnName: "Timor-Leste"}
	HUN                                        = &country{Code: "HUN", CnName: "匈牙利", EnName: "Hungary"}
	AGO                                        = &country{Code: "AGO", CnName: "安哥拉", EnName: "Angola"}
	WSM                                        = &country{Code: "WSM", CnName: "萨摩亚", EnName: "Samoa"}
	FRA                                        = &country{Code: "FRA", CnName: "法国", EnName: "France"}
	MOZ                                        = &country{Code: "MOZ", CnName: "莫桑比克", EnName: "Mozambique"}
	NAM                                        = &country{Code: "NAM", CnName: "纳米比亚", EnName: "Namibia"}
	PER                                        = &country{Code: "PER", CnName: "秘鲁", EnName: "Peru"}
	DNK                                        = &country{Code: "DNK", CnName: "丹麦", EnName: "Denmark"}
	GTM                                        = &country{Code: "GTM", CnName: "危地马拉", EnName: "Guatemala"}
	FRO                                        = &country{Code: "FRO", CnName: "法罗群岛", EnName: "Faroe Islands (the)"}
	SLB                                        = &country{Code: "SLB", CnName: "所罗门群岛", EnName: "Solomon Islands"}
	SLE                                        = &country{Code: "SLE", CnName: "塞拉利昂", EnName: "Sierra Leone"}
	NRU                                        = &country{Code: "NRU", CnName: "瑙鲁", EnName: "Nauru"}
	AIA                                        = &country{Code: "AIA", CnName: "安圭拉", EnName: "Anguilla"}
	GUF                                        = &country{Code: "GUF", CnName: "法属圭亚那", EnName: "French Guiana"}
	SLV                                        = &country{Code: "SLV", CnName: "萨尔瓦多", EnName: "El Salvador"}
	GUM                                        = &country{Code: "GUM", CnName: "关岛", EnName: "Guam"}
	FSM                                        = &country{Code: "FSM", CnName: "密克罗尼西亚联邦", EnName: "Micronesia (Federated States of)"}
	DOM                                        = &country{Code: "DOM", CnName: "多米尼加共和国", EnName: "Dominican Republic (the)"}
	CMR                                        = &country{Code: "CMR", CnName: "喀麦隆", EnName: "Cameroon"}
	GUY                                        = &country{Code: "GUY", CnName: "圭亚那", EnName: "Guyana"}
	AZE                                        = &country{Code: "AZE", CnName: "阿塞拜疆", EnName: "Azerbaijan"}
	MAC                                        = &country{Code: "MAC", CnName: "澳门", EnName: "Macao"}
	TON                                        = &country{Code: "TON", CnName: "汤加", EnName: "Tonga"}
	MAF                                        = &country{Code: "MAF", CnName: "圣马丁（法语部分）", EnName: "Saint Martin (French part)"}
	NCL                                        = &country{Code: "NCL", CnName: "新喀里多尼亚", EnName: "New Caledonia"}
	SMR                                        = &country{Code: "SMR", CnName: "圣马力诺", EnName: "San Marino"}
	ERI                                        = &country{Code: "ERI", CnName: "厄立特里亚", EnName: "Eritrea"}
	MAR                                        = &country{Code: "MAR", CnName: "摩洛哥", EnName: "Morocco"}
	KNA                                        = &country{Code: "KNA", CnName: "圣基茨和尼维斯", EnName: "Saint Kitts and Nevis"}
	BLM                                        = &country{Code: "BLM", CnName: "圣巴勒米", EnName: "Saint Barthélemy"}
	VCT                                        = &country{Code: "VCT", CnName: "圣文森特和格林纳丁斯", EnName: "Saint Vincent and the Grenadines"}
	BLR                                        = &country{Code: "BLR", CnName: "白俄罗斯", EnName: "Belarus"}
	MRT                                        = &country{Code: "MRT", CnName: "毛里塔尼亚", EnName: "Mauritania"}
	BLZ                                        = &country{Code: "BLZ", CnName: "伯利兹", EnName: "Belize"}
	PHL                                        = &country{Code: "PHL", CnName: "菲律宾", EnName: "Philippines (the)"}
	COD                                        = &country{Code: "COD", CnName: "刚果（民主共和国）", EnName: "Congo (the Democratic Republic of the)"}
	COG                                        = &country{Code: "COG", CnName: "刚果", EnName: "Congo (the)"}
	ESH                                        = &country{Code: "ESH", CnName: "西撒哈拉*", EnName: "Western Sahara*"}
	PYF                                        = &country{Code: "PYF", CnName: "法属波利尼西亚", EnName: "French Polynesia"}
	URY                                        = &country{Code: "URY", CnName: "乌拉圭", EnName: "Uruguay"}
	COK                                        = &country{Code: "COK", CnName: "库克群岛", EnName: "Cook Islands (the)"}
	COM                                        = &country{Code: "COM", CnName: "科摩罗", EnName: "Comoros (the)"}
	COL                                        = &country{Code: "COL", CnName: "哥伦比亚", EnName: "Colombia"}
	USA                                        = &country{Code: "USA", CnName: "美利坚合众国", EnName: "United States of America (the)"}
	ESP                                        = &country{Code: "ESP", CnName: "西班牙", EnName: "Spain"}
	EST                                        = &country{Code: "EST", CnName: "爱沙尼亚", EnName: "Estonia"}
	BMU                                        = &country{Code: "BMU", CnName: "百慕大群岛", EnName: "Bermuda"}
	MSR                                        = &country{Code: "MSR", CnName: "蒙特塞拉特", EnName: "Montserrat"}
	ZMB                                        = &country{Code: "ZMB", CnName: "赞比亚", EnName: "Zambia"}
	KOR                                        = &country{Code: "KOR", CnName: "韩国（共和国）", EnName: "Korea (the Republic of)"}
	SOM                                        = &country{Code: "SOM", CnName: "索马里", EnName: "Somalia"}
	VUT                                        = &country{Code: "VUT", CnName: "瓦努阿图", EnName: "Vanuatu"}
	ECU                                        = &country{Code: "ECU", CnName: "厄瓜多尔", EnName: "Ecuador"}
	ALA                                        = &country{Code: "ALA", CnName: "奥兰群岛", EnName: "Åland Islands"}
	ALB                                        = &country{Code: "ALB", CnName: "阿尔巴尼亚", EnName: "Albania"}
	ETH                                        = &country{Code: "ETH", CnName: "埃塞俄比亚", EnName: "Ethiopia"}
	GGY                                        = &country{Code: "GGY", CnName: "根西岛", EnName: "Guernsey"}
	MCO                                        = &country{Code: "MCO", CnName: "摩纳哥", EnName: "Monaco"}
	NER                                        = &country{Code: "NER", CnName: "尼日尔", EnName: "Niger (the)"}
	LAO                                        = &country{Code: "LAO", CnName: "老挝人民民主共和国", EnName: "Lao People's Democratic Republic (the)"}
	VEN                                        = &country{Code: "VEN", CnName: "委内瑞拉玻利瓦尔共和国", EnName: "Venezuela (Bolivarian Republic of)"}
	GHA                                        = &country{Code: "GHA", CnName: "加纳", EnName: "Ghana"}
	CPV                                        = &country{Code: "CPV", CnName: "佛得角", EnName: "Cabo Verde"}
	MTQ                                        = &country{Code: "MTQ", CnName: "马提尼克岛", EnName: "Martinique"}
	MDA                                        = &country{Code: "MDA", CnName: "摩尔多瓦共和国", EnName: "Moldova (the Republic of)"}
	MDG                                        = &country{Code: "MDG", CnName: "马达加斯加", EnName: "Madagascar"}
	SPM                                        = &country{Code: "SPM", CnName: "圣皮埃尔和密克隆", EnName: "Saint Pierre and Miquelon"}
	NFK                                        = &country{Code: "NFK", CnName: "诺福克岛", EnName: "Norfolk Island"}
	LBN                                        = &country{Code: "LBN", CnName: "黎巴嫩", EnName: "Lebanon"}
	LBR                                        = &country{Code: "LBR", CnName: "利比里亚", EnName: "Liberia"}
	BOL                                        = &country{Code: "BOL", CnName: "玻利维亚（多民族国）", EnName: "Bolivia (Plurinational State of)"}
	MDV                                        = &country{Code: "MDV", CnName: "马尔代夫", EnName: "Maldives"}
	GIB                                        = &country{Code: "GIB", CnName: "直布罗陀", EnName: "Gibraltar"}
	LBY                                        = &country{Code: "LBY", CnName: "利比亚", EnName: "Libya"}
	HKG                                        = &country{Code: "HKG", CnName: "香港", EnName: "Hong Kong"}
	CAF                                        = &country{Code: "CAF", CnName: "中非共和国", EnName: "Central African Republic (the)"}
	LSO                                        = &country{Code: "LSO", CnName: "莱索托", EnName: "Lesotho"}
	NGA                                        = &country{Code: "NGA", CnName: "尼日利亚", EnName: "Nigeria"}
	MUS                                        = &country{Code: "MUS", CnName: "毛里求斯", EnName: "Mauritius"}
	IMN                                        = &country{Code: "IMN", CnName: "马恩岛", EnName: "Isle of Man"}
	LCA                                        = &country{Code: "LCA", CnName: "圣卢西亚", EnName: "Saint Lucia"}
	VGB                                        = &country{Code: "VGB", CnName: "维尔京群岛（英属）", EnName: "Virgin Islands (British)"}
	CAN                                        = &country{Code: "CAN", CnName: "加拿大", EnName: "Canada"}
	TCA                                        = &country{Code: "TCA", CnName: "特克斯和凯科斯群岛", EnName: "Turks and Caicos Islands (the)"}
	TCD                                        = &country{Code: "TCD", CnName: "乍得", EnName: "Chad"}
	AND                                        = &country{Code: "AND", CnName: "安道尔", EnName: "Andorra"}
	TCH                                        = &country{Code: "TCH", CnName: "捷克斯洛伐克", EnName: "Czechia"}
	ROU                                        = &country{Code: "ROU", CnName: "罗马尼亚", EnName: "Romania"}
	CRI                                        = &country{Code: "CRI", CnName: "哥斯达黎加", EnName: "Costa Rica"}
	IND                                        = &country{Code: "IND", CnName: "印度", EnName: "India"}
	MEX                                        = &country{Code: "MEX", CnName: "墨西哥", EnName: "Mexico"}
	SRB                                        = &country{Code: "SRB", CnName: "塞尔维亚", EnName: "Serbia"}
	KAZ                                        = &country{Code: "KAZ", CnName: "哈萨克斯坦", EnName: "Kazakhstan"}
	SAU                                        = &country{Code: "SAU", CnName: "沙特阿拉伯", EnName: "Saudi Arabia"}
	JPN                                        = &country{Code: "JPN", CnName: "日本", EnName: "Japan"}
	LTU                                        = &country{Code: "LTU", CnName: "立陶宛", EnName: "Lithuania"}
	TTO                                        = &country{Code: "TTO", CnName: "特立尼达和多巴哥", EnName: "Trinidad and Tobago"}
	PLW                                        = &country{Code: "PLW", CnName: "帕劳", EnName: "Palau"}
	HMD                                        = &country{Code: "HMD", CnName: "赫德岛和麦克唐纳群岛", EnName: "Heard Island and McDonald Islands"}
	MWI                                        = &country{Code: "MWI", CnName: "马拉维", EnName: "Malawi"}
	FYR                                        = &country{Code: "FYR", CnName: "北马其顿", EnName: "North Macedonia"}
	SSD                                        = &country{Code: "SSD", CnName: "南苏丹", EnName: "South Sudan"}
	NIC                                        = &country{Code: "NIC", CnName: "尼加拉瓜", EnName: "Nicaragua"}
	CCK                                        = &country{Code: "CCK", CnName: "科科斯（基林）群岛", EnName: "Cocos (Keeling) Islands (the)"}
	FIN                                        = &country{Code: "FIN", CnName: "芬兰", EnName: "Finland"}
	TUN                                        = &country{Code: "TUN", CnName: "突尼斯", EnName: "Tunisia"}
	LUX                                        = &country{Code: "LUX", CnName: "卢森堡", EnName: "Luxembourg"}
	UGA                                        = &country{Code: "UGA", CnName: "乌干达", EnName: "Uganda"}
	IOT                                        = &country{Code: "IOT", CnName: "英属印度洋领地", EnName: "British Indian Ocean Territory (the)"}
	BRA                                        = &country{Code: "BRA", CnName: "巴西", EnName: "Brazil"}
	TUR                                        = &country{Code: "TUR", CnName: "土耳其", EnName: "Türkiye"}
	BRB                                        = &country{Code: "BRB", CnName: "巴巴多斯", EnName: "Barbados"}
	TUV                                        = &country{Code: "TUV", CnName: "图瓦卢", EnName: "Tuvalu"}
	DEU                                        = &country{Code: "DEU", CnName: "德国", EnName: "Germany"}
	EGY                                        = &country{Code: "EGY", CnName: "埃及", EnName: "Egypt"}
	LVA                                        = &country{Code: "LVA", CnName: "拉脱维亚", EnName: "Latvia"}
	JAM                                        = &country{Code: "JAM", CnName: "牙买加", EnName: "Jamaica"}
	NIU                                        = &country{Code: "NIU", CnName: "纽埃", EnName: "Niue"}
	ZAF                                        = &country{Code: "ZAF", CnName: "南非", EnName: "South Africa"}
	VIR                                        = &country{Code: "VIR", CnName: "维尔京群岛（美国）", EnName: "Virgin Islands (U.S.)"}
	BRN                                        = &country{Code: "BRN", CnName: "文莱达鲁萨兰国", EnName: "Brunei Darussalam"}
	HND                                        = &country{Code: "HND", CnName: "洪都拉斯", EnName: "Honduras"}
	TWN                                        = &country{Code: "TWN", CnName: "中国台湾", EnName: "Taiwan (Province of China)"}
	GEO                                        = &country{Code: "GEO", CnName: "格鲁吉亚", EnName: "Georgia"}
	GIN                                        = &country{Code: "GIN", CnName: "几内亚", EnName: "Guinea"}
	BONAIRE_SINT_EUSTATIUS_AND_SABA            = &country{Code: "Bonaire, Sint Eustatius and Saba", CnName: "博奈尔、圣尤斯特修斯和萨巴", EnName: "Bonaire, Sint Eustatius and Saba"}
	BOSNIA_AND_HERZEGOVINA                     = &country{Code: "Bosnia and Herzegovina", CnName: "波斯尼亚和黑塞哥维那", EnName: "Bosnia and Herzegovina"}
	GEORGIA                                    = &country{Code: "Georgia", CnName: "佐治亚州", EnName: "Georgia"}
	GUINEA                                     = &country{Code: "Guinea", CnName: "几尼", EnName: "Guinea"}
	HOLY_SEE__THE_                             = &country{Code: "Holy See (the)", CnName: "罗马教廷", EnName: "Holy See (the)"}
	JERSEY                                     = &country{Code: "Jersey", CnName: "泽西岛", EnName: "Jersey"}
	JORDAN                                     = &country{Code: "Jordan", CnName: "约旦", EnName: "Jordan"}
	SINT_MAARTEN__DUTCH_PART_                  = &country{Code: "Sint Maarten (Dutch part)", CnName: "圣马丁岛（荷兰语部分）", EnName: "Sint Maarten (Dutch part)"}
	SVALBARD_AND_JAN_MAYEN                     = &country{Code: "Svalbard and Jan Mayen", CnName: "斯瓦尔巴和扬马延", EnName: "Svalbard and Jan Mayen"}
	SYRIAN_ARAB_REPUBLIC__THE_                 = &country{Code: "Syrian Arab Republic (the)", CnName: "阿拉伯叙利亚共和国", EnName: "Syrian Arab Republic (the)"}
	UNITED_STATES_MINOR_OUTLYING_ISLANDS__THE_ = &country{Code: "United States Minor Outlying Islands (the)", CnName: "美国小离岛", EnName: "United States Minor Outlying Islands (the)"}
	WALLIS_AND_FUTUNA                          = &country{Code: "Wallis and Futuna", CnName: "沃利斯和富图纳", EnName: "Wallis and Futuna"}
)

var ALL_COUNTRY_MAP = map[string]*country{
	"NZL":                             NZL,
	"FJI":                             FJI,
	"PNG":                             PNG,
	"GLP":                             GLP,
	"STP":                             STP,
	"MHL":                             MHL,
	"CUB":                             CUB,
	"SDN":                             SDN,
	"GMB":                             GMB,
	"CUW":                             CUW,
	"MYS":                             MYS,
	"MYT":                             MYT,
	"POL":                             POL,
	"OMN":                             OMN,
	"SUR":                             SUR,
	"ARE":                             ARE,
	"KEN":                             KEN,
	"ARG":                             ARG,
	"GNB":                             GNB,
	"ARM":                             ARM,
	"UZB":                             UZB,
	"BTN":                             BTN,
	"SEN":                             SEN,
	"TGO":                             TGO,
	"IRL":                             IRL,
	"FLK":                             FLK,
	"IRN":                             IRN,
	"QAT":                             QAT,
	"BDI":                             BDI,
	"NLD":                             NLD,
	"IRQ":                             IRQ,
	"SVK":                             SVK,
	"SVN":                             SVN,
	"GNQ":                             GNQ,
	"THA":                             THA,
	"ABW":                             ABW,
	"ASM":                             ASM,
	"SWE":                             SWE,
	"ISL":                             ISL,
	"BEL":                             BEL,
	"ISR":                             ISR,
	"KWT":                             KWT,
	"LIE":                             LIE,
	"DZA":                             DZA,
	"BEN":                             BEN,
	"ATA":                             ATA,
	"RUS":                             RUS,
	"ATF":                             ATF,
	"ATG":                             ATG,
	"SWZ":                             SWZ,
	"ITA":                             ITA,
	"TZA":                             TZA,
	"PAK":                             PAK,
	"BFA":                             BFA,
	"CXR":                             CXR,
	"PAN":                             PAN,
	"SGP":                             SGP,
	"UKR":                             UKR,
	"SGS":                             SGS,
	"KGZ":                             KGZ,
	"BVT":                             BVT,
	"CHE":                             CHE,
	"DJI":                             DJI,
	"REU":                             REU,
	"CHL":                             CHL,
	"PRI":                             PRI,
	"CHN":                             CHN,
	"PRK":                             PRK,
	"MLI":                             MLI,
	"BWA":                             BWA,
	"HRV":                             HRV,
	"KHM":                             KHM,
	"IDN":                             IDN,
	"PRT":                             PRT,
	"MLT":                             MLT,
	"TJK":                             TJK,
	"VNM":                             VNM,
	"CYM":                             CYM,
	"PRY":                             PRY,
	"SHN":                             SHN,
	"CYP":                             CYP,
	"SYC":                             SYC,
	"RWA":                             RWA,
	"BGD":                             BGD,
	"AUS":                             AUS,
	"AUT":                             AUT,
	"PSE":                             PSE,
	"LKA":                             LKA,
	"GAB":                             GAB,
	"ZWE":                             ZWE,
	"BGR":                             BGR,
	"NOR":                             NOR,
	"CIV":                             CIV,
	"MMR":                             MMR,
	"TKL":                             TKL,
	"KIR":                             KIR,
	"TKM":                             TKM,
	"GRD":                             GRD,
	"GRC":                             GRC,
	"PCN":                             PCN,
	"HTI":                             HTI,
	"GRL":                             GRL,
	"YEM":                             YEM,
	"AFG":                             AFG,
	"MNE":                             MNE,
	"MNG":                             MNG,
	"NPL":                             NPL,
	"BHS":                             BHS,
	"BHR":                             BHR,
	"MNP":                             MNP,
	"GBR":                             GBR,
	"DMA":                             DMA,
	"TLS":                             TLS,
	"HUN":                             HUN,
	"AGO":                             AGO,
	"WSM":                             WSM,
	"FRA":                             FRA,
	"MOZ":                             MOZ,
	"NAM":                             NAM,
	"PER":                             PER,
	"DNK":                             DNK,
	"GTM":                             GTM,
	"FRO":                             FRO,
	"SLB":                             SLB,
	"SLE":                             SLE,
	"NRU":                             NRU,
	"AIA":                             AIA,
	"GUF":                             GUF,
	"SLV":                             SLV,
	"GUM":                             GUM,
	"FSM":                             FSM,
	"DOM":                             DOM,
	"CMR":                             CMR,
	"GUY":                             GUY,
	"AZE":                             AZE,
	"MAC":                             MAC,
	"TON":                             TON,
	"MAF":                             MAF,
	"NCL":                             NCL,
	"SMR":                             SMR,
	"ERI":                             ERI,
	"MAR":                             MAR,
	"KNA":                             KNA,
	"BLM":                             BLM,
	"VCT":                             VCT,
	"BLR":                             BLR,
	"MRT":                             MRT,
	"BLZ":                             BLZ,
	"PHL":                             PHL,
	"COD":                             COD,
	"COG":                             COG,
	"ESH":                             ESH,
	"PYF":                             PYF,
	"URY":                             URY,
	"COK":                             COK,
	"COM":                             COM,
	"COL":                             COL,
	"USA":                             USA,
	"ESP":                             ESP,
	"EST":                             EST,
	"BMU":                             BMU,
	"MSR":                             MSR,
	"ZMB":                             ZMB,
	"KOR":                             KOR,
	"SOM":                             SOM,
	"VUT":                             VUT,
	"ECU":                             ECU,
	"ALA":                             ALA,
	"ALB":                             ALB,
	"ETH":                             ETH,
	"GGY":                             GGY,
	"MCO":                             MCO,
	"NER":                             NER,
	"LAO":                             LAO,
	"VEN":                             VEN,
	"GHA":                             GHA,
	"CPV":                             CPV,
	"MTQ":                             MTQ,
	"MDA":                             MDA,
	"MDG":                             MDG,
	"SPM":                             SPM,
	"NFK":                             NFK,
	"LBN":                             LBN,
	"LBR":                             LBR,
	"BOL":                             BOL,
	"MDV":                             MDV,
	"GIB":                             GIB,
	"LBY":                             LBY,
	"HKG":                             HKG,
	"CAF":                             CAF,
	"LSO":                             LSO,
	"NGA":                             NGA,
	"MUS":                             MUS,
	"IMN":                             IMN,
	"LCA":                             LCA,
	"VGB":                             VGB,
	"CAN":                             CAN,
	"TCA":                             TCA,
	"TCD":                             TCD,
	"AND":                             AND,
	"TCH":                             TCH,
	"ROU":                             ROU,
	"CRI":                             CRI,
	"IND":                             IND,
	"MEX":                             MEX,
	"SRB":                             SRB,
	"KAZ":                             KAZ,
	"SAU":                             SAU,
	"JPN":                             JPN,
	"LTU":                             LTU,
	"TTO":                             TTO,
	"PLW":                             PLW,
	"HMD":                             HMD,
	"MWI":                             MWI,
	"FYR":                             FYR,
	"SSD":                             SSD,
	"NIC":                             NIC,
	"CCK":                             CCK,
	"FIN":                             FIN,
	"TUN":                             TUN,
	"LUX":                             LUX,
	"UGA":                             UGA,
	"IOT":                             IOT,
	"BRA":                             BRA,
	"TUR":                             TUR,
	"BRB":                             BRB,
	"TUV":                             TUV,
	"DEU":                             DEU,
	"EGY":                             EGY,
	"LVA":                             LVA,
	"JAM":                             JAM,
	"NIU":                             NIU,
	"ZAF":                             ZAF,
	"VIR":                             VIR,
	"BRN":                             BRN,
	"HND":                             HND,
	"TWN":                             TWN,
	"GEO":                             GEO,
	"GIN":                             GIN,
	"BONAIRE_SINT_EUSTATIUS_AND_SABA": BONAIRE_SINT_EUSTATIUS_AND_SABA,
	"BOSNIA_AND_HERZEGOVINA":          BOSNIA_AND_HERZEGOVINA,
	"GEORGIA":                         GEORGIA,
	"GUINEA":                          GUINEA,
	"HOLY_SEE__THE_":                  HOLY_SEE__THE_,
	"JERSEY":                          JERSEY,
	"JORDAN":                          JORDAN,
	"SINT_MAARTEN__DUTCH_PART_":       SINT_MAARTEN__DUTCH_PART_,
	"SVALBARD_AND_JAN_MAYEN":          SVALBARD_AND_JAN_MAYEN,
	"SYRIAN_ARAB_REPUBLIC__THE_":      SYRIAN_ARAB_REPUBLIC__THE_,
	"UNITED_STATES_MINOR_OUTLYING_ISLANDS__THE_": UNITED_STATES_MINOR_OUTLYING_ISLANDS__THE_,
	"WALLIS_AND_FUTUNA":                          WALLIS_AND_FUTUNA,
}

func GetByCode(code string) *country {
	return ALL_COUNTRY_MAP[code]
}

func GetCnNameByCode(code string) string {
	cty := ALL_COUNTRY_MAP[code]
	if cty != nil {
		return cty.CnName
	}
	return ""
}

func GetEnNameByCode(code string) string {
	cty := ALL_COUNTRY_MAP[code]
	if cty != nil {
		return cty.EnName
	}
	return ""
}
