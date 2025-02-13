package servidorreplica
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/replication", getReplicatedProducts)

	if err := r.Run(":4000"); err != nil {
		fmt.Println("Error: Replication Server hasn't begun")
	}
}