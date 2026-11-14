package components

// PostPlace - A place a post names, as the platform shows it.
type PostPlace struct {
	Name string `json:"name"`
	// The platform's own id of the place.
	ID        *string  `json:"id,omitzero"`
	Address   *string  `json:"address,omitzero"`
	City      *string  `json:"city,omitzero"`
	Latitude  *float64 `json:"latitude,omitzero"`
	Longitude *float64 `json:"longitude,omitzero"`
}

func (p *PostPlace) GetName() string {
	if p == nil {
		return ""
	}
	return p.Name
}

func (p *PostPlace) GetID() *string {
	if p == nil {
		return nil
	}
	return p.ID
}

func (p *PostPlace) GetAddress() *string {
	if p == nil {
		return nil
	}
	return p.Address
}

func (p *PostPlace) GetCity() *string {
	if p == nil {
		return nil
	}
	return p.City
}

func (p *PostPlace) GetLatitude() *float64 {
	if p == nil {
		return nil
	}
	return p.Latitude
}

func (p *PostPlace) GetLongitude() *float64 {
	if p == nil {
		return nil
	}
	return p.Longitude
}
