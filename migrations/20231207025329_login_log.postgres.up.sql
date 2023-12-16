    CREATE TABLE IF NOT EXISTS login_log (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        user_id UUID REFERENCES users (id) ON DELETE CASCADE,
        is_authorized boolean not null,
        ip_address varchar(255) not null,
        timezone varchar not null,
        lat numeric not null,
        lon numeric not null,
        access_time timestamp(0) not null
        );

-- is_authorized ini untuk login atempt, kita check kalau authorized maka bisa masuk
-- lat latitude (lintang)
-- lon longitude (bujur)
-- access_time dia kapan mengakses login tsb, supaya kita bisa melihat kapan waktu loginnya