package secret_reader

func LoadSecrets(acquirers []string) map[string]SecretReader {
	secretAcquirers := make(map[string]SecretReader)

	for _, acquirer := range acquirers {
		switch acquirer {
		case "permata":
			secretAcquirers["permata"] = New("permata")
		case "aspi":
			secretAcquirers["aspi"] = New("aspi")
		}
	}

	return secretAcquirers
}
