package vo

type ThemeOptions struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
type ThemePages struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Screenshot  string `json:"screenshot"`
	File        string `json:"file"`
}
type ThemeSchema struct {
	Kit     string         `json:"kit"`
	Check   string         `json:"check"`
	Name    string         `json:"name"`
	Label   string         `json:"label"`
	Value   string         `json:"value"`
	Help    string         `json:"help"`
	Options []ThemeOptions `json:"options,omitempty"`
	Group   []ThemeSchema  `json:"group,omitempty"`
}
type ThemeForms struct {
	Group  string        `json:"group"`
	Label  string        `json:"label"`
	Schema []ThemeSchema `json:"schema"`
}
type ThemeSettingConfig struct {
	Name        string        `json:"name"`
	Kind        string        `json:"kind"`
	Version     string        `json:"version"`
	Author      string        `json:"author"`
	Website     string        `json:"website"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Repo        string        `json:"repo"`
	Pages       []ThemePages  `json:"pages"`
	Forms       []ThemeForms  `json:"forms"`
	Posts       []ThemeSchema `json:"posts,omitempty"`
}
