package global

func (p *PageInfo) ValidateAndSetDefault() {
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}
