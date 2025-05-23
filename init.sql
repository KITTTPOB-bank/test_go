CREATE TABLE IF NOT EXISTS patient (
    id SERIAL PRIMARY KEY,
    first_name_th TEXT NOT NULL,
    middle_name_th TEXT,
    last_name_th TEXT NOT NULL,
    first_name_en TEXT,
    middle_name_en TEXT,
    last_name_en TEXT,
    date_of_birth TEXT,
    patient_hn TEXT,
    national_id VARCHAR(13) UNIQUE,
    passport_id VARCHAR(20) UNIQUE,
    phone_number VARCHAR(20),
    email TEXT,
    gender CHAR(1) CHECK (gender IN ('M', 'F')),
    staff_id INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    password TEXT,
    hospital TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
