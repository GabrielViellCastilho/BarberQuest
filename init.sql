
-- Conectando ao banco de dados
\c barberquest_database;

CREATE TABLE users (
     id SERIAL PRIMARY KEY,
     name VARCHAR(100) NOT NULL,
     email VARCHAR(150) UNIQUE NOT NULL,
     password VARCHAR(255) NOT NULL,
    role VARCHAR(20) CHECK (role IN ('admin', 'barber','user')) NOT NULL,
    cellphone VARCHAR(11) CHECK (cellphone ~ '^[1-9]{2}9?[0-9]{8}$')  NOT NULL,
    date_of_birth DATE

);

CREATE TABLE services (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          price NUMERIC(10, 2) NOT NULL,
                          duration_minutes INT NOT NULL, -- Duração do serviço em minutos
                          available BOOL
);

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
     client_name VARCHAR(100) NOT NULL,
     client_contact VARCHAR(11),
     appointment_date TIMESTAMP NOT NULL,
    barber_id INT REFERENCES users(id) ON DELETE SET NULL, -- Relacionamento com o barbeiro
    service_id INT REFERENCES services(id) ON DELETE CASCADE,
    completed BOOL
);


CREATE TABLE barber_availability (
                                     id SERIAL PRIMARY KEY,
                                     barber_id INT REFERENCES users(id) ON DELETE CASCADE,
                                     day_of_week INT CHECK (day_of_week BETWEEN 1 AND 6),
                                     start_time TIME NOT NULL,
                                     end_time TIME NOT NULL,
                                     break_start_time TIME,  -- ⬅️ Início do intervalo
                                     break_end_time TIME,    -- ⬅️ Fim do intervalo
                                     CONSTRAINT valid_time_range CHECK (start_time < end_time),
                                     CONSTRAINT valid_break_time CHECK (
                                         break_start_time IS NULL OR break_end_time IS NULL OR
                                         (break_start_time >= start_time AND break_end_time <= end_time AND break_start_time < break_end_time)
                                         ),
                                     CONSTRAINT unique_barber_day UNIQUE (barber_id, day_of_week)
);



CREATE TABLE special_schedule (
    id SERIAL PRIMARY KEY,
    barber_id INT REFERENCES users(id),
    date DATE NOT NULL,
    opening_time TIME,
    closing_time TIME,
    break_start_time TIME,  -- ⬅️ Início do intervalo
    break_end_time TIME,    -- ⬅️ Fim do intervalo
    CONSTRAINT unique_barber_day_special UNIQUE (barber_id, date)
    CONSTRAINT valid_break_time_special CHECK (
        break_start_time IS NULL OR break_end_time IS NULL OR
        break_start_time = '00:00:00' OR break_end_time = '00:00:00' OR
        (break_start_time >= opening_time AND break_end_time <= closing_time AND break_start_time < break_end_time)
        ),
);



