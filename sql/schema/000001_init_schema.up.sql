CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,       -- ID duy nhất cho người dùng
  "username" varchar NOT NULL,           -- Tên người dùng (đăng nhập)
  "full_name" varchar NOT NULL,          -- Tên đầy đủ của người dùng
  "gender" varchar(10),                  -- Giới tính (Male, Female, Other)
  "email" varchar UNIQUE NOT NULL,       -- Địa chỉ email duy nhất
  "phone_number" varchar(15),            -- Số điện thoại
  "date_of_birth" date,                  -- Ngày sinh
  "address" text,                        -- Địa chỉ đầy đủ
  "created_at" timestamptz NOT NULL DEFAULT (now()), -- Thời gian tạo
  "updated_at" timestamptz DEFAULT NULL, -- Thời gian cập nhật cuối cùng
  "is_active" boolean DEFAULT true,      -- Trạng thái hoạt động
  "role" varchar(50) DEFAULT 'user'      -- Vai trò (admin, user, etc.)
);

CREATE INDEX ON "users" ("username");
CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("phone_number");
