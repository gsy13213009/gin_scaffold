package dto

import "time"

type AdminInfoOutPut struct {
	ID           int       `json:"id"`
	Name         string    `json:"user_name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json"roles"`
}
