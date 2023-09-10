package recdn

import "github.com/evanw/esbuild/pkg/api"

func Min(content string, loader api.Loader) string {
	result := api.Transform(content, api.TransformOptions{
		Loader: loader,
		MinifyWhitespace: true,
		MinifyIdentifiers: true,
		MinifySyntax: true,
	})
	return string(result.Code)
}