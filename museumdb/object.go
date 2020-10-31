package main

type Object struct{
	Id 					int
	AccessionNumber  	string  `json:"accession_number"`
	// ArtChampionsText	___		`json:"art_champions_text"`
	Artist				string	`json:"artist"`
	// Catalogue			string	`json:"catalogue_raissonne`
	// Classification		string	`json:"classification"`
	Continent			string	`json:"continent"`
	Country				string	`json:"country"`
	// CreditLine			string	`json:"creditline"`
	// Culture						`json:"culture"`: null,
	// CuratorApproved		bool	`json:"curator_approved"`
	Dated				string	`json:"dated"`
	Department			string 	`json:"department"`
	Description			string 	`json:"description"`
	Dimension			string	`json:"dimension"`
	IdURL				string	`json:"id"`
	Image				string	`json:"image"`
	// ImageCopyright 		string	`json:"image_copyright"`
	Height				int		`json:"image_height"`
	Width				int		`json:"image_width"`
	// Inscription			string	`json:"inscription"`
	// LifeDate			string	`json:"life_date"`
	// Markings			string	`json:"markings"`
	Medium 				string	`json:"medium"`
	// Nationality			string	`json:"nationality"`
	// ObjectName			string  `json:"object_name"`
	// Portfolio			string	`json:"portfolio"`
	// Provenance 			string	`json:"provenance"`
	// Restricted			bool	`json:"restricted"`
	// RightsType			string	`json:"rights_type"`
	// Role				string	`json:"role"`
	Room				string	`json:"room"`
	// SeeAlso				[]string	`json:"see_also"`
	// Signed				string	`json:"signed"`
	Style 				string	`json:"style"`
	Text				string	`json:"text"`
	Title				string	`json:"title"`
}