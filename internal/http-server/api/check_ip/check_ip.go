package check_ip

import (
	"context"
	"github.com/ShevelevEvgeniy/app/config"
	ipInfoClient "github.com/ShevelevEvgeniy/app/internal/clients/ip_info_client"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	retryFunc "github.com/ShevelevEvgeniy/app/pkg/retry_func"
	"log/slog"
	"net"
	"net/http"
	"strings"
)

func CheckIp(ctx context.Context, client *ipInfoClient.IpInfoClient, retry *retryFunc.RetryFunc, cfg *config.IpInfo, log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ipStr := getIpStr(r)
			if cfg.LocalMachine && ipStr == "::1" {
				next.ServeHTTP(w, r)
			}

			ip := net.ParseIP(ipStr)

			var country string
			err := retry.Do(ctx, func() error {
				c, err := client.GetIpInfo(ip)
				if err != nil {
					return err
				}

				country = c
				return nil
			})
			if err != nil {
				log.Error("Failed get ip info", sl.Err(err))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if country != cfg.Country {
				log.Info("Blocked ip", slog.String("ip", ipStr))
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getIpStr(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")

		return strings.TrimSpace(ips[0])
	}

	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return xri
	}

	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)

	return remoteIP
}
