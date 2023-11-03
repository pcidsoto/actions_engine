package engine

type ParametroDeEntrada struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type ObtenerDato struct {
	Accion  string                 `json:"accion"`
	Param   ParametroDeEntrada     `json:"param"`
	Retorno map[string]interface{} `json:"retorno"`
}

type ActivarFiltro struct {
	Accion string             `json:"accion"`
	Param  ParametroDeEntrada `json:"param"`
}

type Step struct {
	ObtenerDato   ObtenerDato     `json:"ObtenerDato"`
	ActivarFiltro []ActivarFiltro `json:"ActivarFiltro"`
}

type Config struct {
	ParametroDeEntrada ParametroDeEntrada `json:"Parametro_de_entrada"`
	Steps              []Step             `json:"steps"`
}
