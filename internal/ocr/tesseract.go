package ocr

import "github.com/otiai10/gosseract"

func ExtractOCRText(filepath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(filepath)

	text, err := client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
