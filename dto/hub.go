package dto

import "github.com/google/uuid"

type Hub struct {
	// untuk bikin list channel, ketika user stream ke system kita, ia akan membuat channel, channel ini akan di buat oleh user
	// jadi nanti ketika ada notif yang berhubungan dengan user tsb, kita tinggal ngirim ke channel user nya
	NotificationChannel map[uuid.UUID]chan NotificationData // int64 ini berisi user_id, user_id daripada usernya
}
