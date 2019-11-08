package client

import "errors"

type Manifest map[string]interface{}

func (m Manifest) Kind() string {
	if _, ok := m["kind"]; !ok {
		return ""
	}
	return m["kind"].(string)
}

func (m Manifest) APIVersion() string {
	if _, ok := m["apiVersion"]; !ok {
		return ""
	}
	return m["apiVersion"].(string)
}

func (m Manifest) Metadata() Metadata {
	if _, ok := m["metadata"]; !ok {
		return nil
	}
	return Metadata(m["metadata"].(map[string]interface{}))
}

func (m Manifest) VerifyLax() error {
	if m.Kind() == "" {
		return errors.New("kind missing")
	}
	if m.APIVersion() == "" {
		return errors.New("apiVersion missing")
	}
	if m.Metadata() == nil {
		return errors.New("metadata missing")
	}
	return nil
}

func (m Manifest) Verify() error {
	if err := m.VerifyLax(); err != nil {
		return err
	}

	if m.Metadata().Name() == "" {
		return errors.New("name missing")
	}

	if m.Metadata().Labels() == nil {
		return errors.New("labels missing")
	}
	return nil
}

type Metadata map[string]interface{}

func (m Metadata) Name() string {
	if _, ok := m["name"]; !ok {
		return ""
	}
	return m["name"].(string)
}

func (m Metadata) Labels() map[string]interface{} {
	if _, ok := m["labels"]; !ok {
		return nil
	}
	return m["labels"].(map[string]interface{})
}
