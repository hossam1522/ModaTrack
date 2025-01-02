package models

import "encoding/json"

type Ropa struct {
	nombre string
	talla  Talla
}

func (r Ropa) GetTalla() Talla {
	return r.talla
}

func (r Ropa) GetNombre() string {
	return r.nombre
}

func (r Ropa) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Nombre string `json:"nombre"`
		Talla  Talla  `json:"talla"`
	}{
		Nombre: r.nombre,
		Talla:  r.talla,
	})
}

func (r *Ropa) UnmarshalJSON(data []byte) error {
	aux := struct {
		Nombre string `json:"nombre"`
		Talla  Talla  `json:"talla"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.nombre = aux.Nombre
	r.talla = aux.Talla
	return nil
}
