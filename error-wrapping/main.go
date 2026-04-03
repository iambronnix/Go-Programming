
func readConfig() error {
	_, err := readFile("file.txt")
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	return nil
	//alternatively the wrapped error return can be rep by :
	//return fmt.Errorf("failed to read config fil: %w", err)
}