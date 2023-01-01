package config_reader

type Provider interface {
	Unmarshal([]byte, any) error
}

func LoadConfig(filePath string, dest any, provider Provider, helper IoHelper) error {
	content, err := getFileContents(filePath, helper)
	if err != nil {
		return err
	}

	return provider.Unmarshal(content, dest)
}

func getFileContents(filePath string, helper IoHelper) ([]byte, error) {
	content, err := helper.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
