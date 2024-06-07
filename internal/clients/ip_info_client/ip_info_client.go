package ip_info_client

import (
	"github.com/ShevelevEvgeniy/app/config"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/pkg/errors"
	"net"
)

type IpInfoClient struct {
	cfg    *config.IpInfo
	Client *ipinfo.Client
}

func NewIpInfoClient(cfg *config.IpInfo) *IpInfoClient {
	return &IpInfoClient{
		cfg:    cfg,
		Client: ipinfo.NewClient(nil, nil, cfg.IpInfoToken),
	}
}

func (c *IpInfoClient) GetIpInfo(ip net.IP) (string, error) {
	ipInfo, err := c.Client.GetIPCountry(ip)
	if err != nil {
		return "", errors.Wrap(err, "Failed get ip info")
	}

	return ipInfo, nil
}
