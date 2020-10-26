package view

import (
	"strconv"

	"github.com/kataras/iris/v12"
)

//SendResponse envia respuesta a cliente
func SendResponse(ctx iris.Context, status int, data []byte) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("status", strconv.Itoa(status))
	ctx.Write(data)
}

//SendErr envia erro a cliente
func SendErr(ctx iris.Context, status int) {
	data := []byte(`{}`)
	ctx.Header("Content-Type", "application/json")
	ctx.Header("status", strconv.Itoa(status))
	ctx.Write(data)
}

func NewProductProblem(productName, detail string) iris.Problem {
    return iris.NewProblem().
        // The type URI, if relative it automatically convert to absolute.
        Type("/product-error"). 
        // The title, if empty then it gets it from the status code.
        Title("Product validation problem").
        // Any optional details.
        Detail(detail).
        // The status error code, required.
        Status(iris.StatusBadRequest).
        // Any custom key-value pair.
        Key("productName", productName)
        // Optional cause of the problem, chain of Problems.
        // .Cause(other iris.Problem)
}

func FireProblem(ctx iris.Context) {
    // Response like JSON but with indent of "  " and
    // content type of "application/problem+json"
    ctx.Problem(NewProductProblem("product name", "problem details"), iris.ProblemOptions{
        // Optional JSON renderer settings.
        JSON: iris.JSON{
            Indent: "  ",
        },
        // OR
        // Render as XML:
        // RenderXML: true,
        // XML:       iris.XML{Indent: "  "},
        // Sets the "Retry-After" response header.
        //
        // Can accept:
        // time.Time for HTTP-Date,
        // time.Duration, int64, float64, int for seconds
        // or string for date or duration.
        // Examples:
        // time.Now().Add(5 * time.Minute),
        // 300 * time.Second,
        // "5m",
        //
        RetryAfter: 300,
        // A function that, if specified, can dynamically set
        // retry-after based on the request. Useful for ProblemOptions reusability.
        // Overrides the RetryAfter field.
        //
        // RetryAfterFunc: func(iris.Context) interface{} { [...] }
    })
}