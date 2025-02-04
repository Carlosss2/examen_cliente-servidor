package servidorprincipal


import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Producto struct{
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Cantidad string `json:"cantidad"`
	CodigoBarra string `json:"cantidad"`
}

type Cambio struct{
	Accion string `json:"accion"`
	Producto   Producto   `json:"producto"`
}

var (
	bd      []Producto
	cambios []Cambio
)

func sendToReplicationServer(product Producto, accion string) {
	url := fmt.Sprintf("http://localhost:5000/replication?product_id=%d&name=%s&product=%s&accion=%s",
		product.ID,
		url.QueryEscape(product.Name),
		url.QueryEscape(product.Cantidad),
		url.QueryEscape(product.CodigoBarra),
		url.QueryEscape(accion),
	)

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error al enviar al servidor de replicación:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error replicando el producto:", resp.Status)
	}
}