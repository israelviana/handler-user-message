package domain

type MetaRegisterNumberBody struct {
	MessagingProduct string `json:"messaging_product"`
}

type MetaRegisterNumberResponse struct {
}

type MetaEditTemplateBody struct {
	Category   *string         `json:"category,omitempty"`
	Components *map[string]any `json:"components,omitempty"`
}

type MetaDeleteTemplateParams struct {
	ID           string `json:"id"`
	TemplateName string `json:"templateName"`
}
type MetaDeleteTemplateOption func(*MetaDeleteTemplateParams)

func WithID(id string) MetaDeleteTemplateOption {
	return func(p *MetaDeleteTemplateParams) {
		p.ID = id
	}
}

type MetaListTemplatesParams struct {
	Fields []string `json:"fields"`
	Limit  int      `json:"limit"`
}
type MetaListTemplatesOption func(*MetaListTemplatesParams)

func WithFields(fields []string) MetaListTemplatesOption {
	return func(p *MetaListTemplatesParams) {
		p.Fields = fields
	}
}

func WithLimit(limit int) MetaListTemplatesOption {
	return func(p *MetaListTemplatesParams) {
		p.Limit = limit
	}
}

type MetaListTemplatesResponse struct {
	Data   []MetaListTemplatesData `json:"data"`
	Paging MetaPagination          `json:"paging"`
}

type MetaListTemplatesData struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Id     string `json:"id"`
}

type MetaPagination struct {
	Cursors struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"cursors"`
	Next string `json:"next"`
}

type MetaCreateTemplateResponse struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	Category string `json:"category"`
}

type Category int

const (
	MARKETING Category = iota
	UTILITY
	AUTHENTICATION
)

func (c Category) String() string {
	switch c {
	case MARKETING:
		return "MARKETING"
	case UTILITY:
		return "UTILITY"
	case AUTHENTICATION:
		return "AUTHENTICATION"
	default:
		return ""
	}
}

func (c Category) IsValid() bool {
	switch c {
	case MARKETING, UTILITY, AUTHENTICATION:
		return true
	default:
		return false
	}
}

type MetaCreateTemplateBody struct {
	Name            string                  `json:"name"`
	Category        string                  `json:"category"`
	ParameterFormat interface{}             `json:"parameter_format"`
	Language        string                  `json:"language"`
	Components      []MetaTemplateComponent `json:"components"`
}

type MetaTemplateComponent struct {
	Type    string                        `json:"type"`             // HEADER, BODY, etc.
	Format  *string                       `json:"format,omitempty"` // HEADER: TEXT, IMAGE, etc.
	Text    *string                       `json:"text,omitempty"`
	Buttons *[]MetaTemplateButton         `json:"buttons,omitempty"`
	Example *MetaTemplateComponentExample `json:"example,omitempty"`
}

type MetaTemplateButton struct {
	Type        string `json:"type"` // QUICK_REPLY, URL, PHONE_NUMBER
	Text        string `json:"text"`
	URL         string `json:"url,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type MetaTemplateComponentExample struct {
	HeaderText []string `json:"header_text,omitempty"`
	BodyText   []string `json:"body_text,omitempty"`
}

type MetaSendWhatsappMessageWithTemplateResponse struct {
	MessagingProduct string        `json:"messaging_product"`
	Contacts         []MetaContact `json:"contacts"`
	Messages         []MetaMessage `json:"messages"`
}

type MetaContact struct {
	Input string `json:"input"`
	WaId  string `json:"wa_id"`
}

type MetaMessage struct {
	Id string `json:"id"`
}

type MetaSendWhatsappMessageBody struct {
	MessagingProduct string        `json:"messaging_product"`
	RecipientType    string        `json:"recipient_type"`
	To               string        `json:"to"`
	Type             string        `json:"type"`
	Template         *MetaTemplate `json:"template,omitempty"` //type: template
}

type MetaTemplate struct {
	NameSpace  string          `json:"namespace"`
	Name       string          `json:"name"`
	Language   MetaLanguage    `json:"language"`
	Components []MetaComponent `json:"components"`
}

type MetaLanguage struct {
	Code string `json:"code"`
}

type MetaComponent struct {
	Type       string          `json:"type"`
	Parameters []MetaParameter `json:"parameters"`
}

type MetaParameter struct {
	Type         string        `json:"type"`
	Text         *string       `json:"text,omitempty"`
	MetaCurrency *MetaCurrency `json:"currency,omitempty"`
	MetaButton   *MetaButton   `json:"button,omitempty"`
}

type MetaButton struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Text    string `json:"text"`
}

type MetaCurrency struct {
	FallbackValue string  `json:"fallback_value"`
	Code          string  `json:"code"`
	Amount1000    float64 `json:"amount_1000"`
}
