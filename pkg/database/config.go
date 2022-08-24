package database

import (
	"github.com/crowdsecurity/crowdsec/pkg/database/ent"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/configitem"
	"github.com/pkg/errors"
)

func (c *Client) GetConfigItem(key string) (*string, error) {
	result, err := c.Ent.ConfigItem.Query().Where(configitem.NameEQ(key)).First(c.CTX)
	if err != nil && !ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrapf(QueryFail, "select config item: %s", err)
	}

	return &result.Value, nil
}

func (c *Client) SetConfigItem(key string, value string) error {
	err := c.Ent.ConfigItem.Create().SetName(key).SetValue(value).Exec(c.CTX)

	if err != nil {
		return errors.Wrapf(QueryFail, "insert config item: %s", err)
	}

	return nil
}