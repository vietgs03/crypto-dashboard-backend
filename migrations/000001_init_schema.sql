-- auth: Xác thực và phân quyền
-- core: Thông tin cốt lõi của người dùng
-- market: Dữ liệu thị trường
-- portfolio: Quản lý danh mục đầu tư
-- notify: Thông báo và cảnh báo
-- audit: Ghi log hoạt động


-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS timescaledb;

-- Create schemas
CREATE SCHEMA auth;        -- Authentication and authorization
CREATE SCHEMA core;        -- Core business logic
CREATE SCHEMA market;      -- Market data and analysis
CREATE SCHEMA portfolio;   -- Portfolio management
CREATE SCHEMA notify;      -- Notifications and alerts
CREATE SCHEMA audit;       -- Audit logging

-- Common types across schemas
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'suspended');
CREATE TYPE notification_type AS ENUM ('email', 'push', 'sms');
CREATE TYPE transaction_type AS ENUM ('buy', 'sell', 'transfer', 'deposit', 'withdrawal');
CREATE TYPE alert_severity AS ENUM ('info', 'warning', 'critical');
CREATE TYPE subscription_tier AS ENUM ('free', 'basic', 'premium', 'enterprise');

-- Auth Schema (Authentication & Authorization)
CREATE TABLE auth.users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    status user_status DEFAULT 'active',
    subscription_type subscription_tier DEFAULT 'free',
    email_verified BOOLEAN DEFAULT false,
    two_factor_enabled BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    last_login_at TIMESTAMPTZ,
    version INTEGER DEFAULT 1
);

CREATE TABLE auth.api_keys (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    key_name VARCHAR(50) NOT NULL,
    api_key VARCHAR(255) NOT NULL,
    api_secret VARCHAR(255),
    platform VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    last_used_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, platform)
);

-- Core Schema (User Profiles and Settings)
CREATE TABLE core.user_profiles (
    user_id UUID PRIMARY KEY REFERENCES auth.users(id) ON DELETE CASCADE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    country_code CHAR(2),
    timezone VARCHAR(50),
    kyc_verified BOOLEAN DEFAULT false,
    kyc_verified_at TIMESTAMPTZ,
    preferences JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Market Schema (Market Data and Analysis)
CREATE TABLE market.assets (
    symbol VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE market.price_data (
    asset_symbol VARCHAR(20) NOT NULL REFERENCES market.assets(symbol),
    timestamp TIMESTAMPTZ NOT NULL,
    price DECIMAL(28,8) NOT NULL,
    volume_24h DECIMAL(28,8),
    market_cap DECIMAL(28,8),
    price_change_24h DECIMAL(8,2),
    data_source VARCHAR(50) NOT NULL,
    PRIMARY KEY (asset_symbol, timestamp)
);

CREATE TABLE market.whale_alerts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_symbol VARCHAR(20) NOT NULL REFERENCES market.assets(symbol),
    amount DECIMAL(28,8) NOT NULL,
    from_address VARCHAR(255),
    to_address VARCHAR(255),
    tx_hash VARCHAR(255) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    blockchain VARCHAR(50) NOT NULL,
    usd_value DECIMAL(28,2),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Portfolio Schema (Portfolio Management)
CREATE TABLE portfolio.portfolios (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    is_default BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, name)
);

CREATE TABLE portfolio.holdings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    portfolio_id UUID NOT NULL REFERENCES portfolio.portfolios(id) ON DELETE CASCADE,
    asset_symbol VARCHAR(20) NOT NULL REFERENCES market.assets(symbol),
    quantity DECIMAL(28,8) NOT NULL,
    average_buy_price DECIMAL(28,8),
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(portfolio_id, asset_symbol)
);

CREATE TABLE portfolio.transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    portfolio_id UUID NOT NULL REFERENCES portfolio.portfolios(id),
    type transaction_type NOT NULL,
    asset_symbol VARCHAR(20) NOT NULL REFERENCES market.assets(symbol),
    quantity DECIMAL(28,8) NOT NULL,
    price_per_unit DECIMAL(28,8),
    total_amount DECIMAL(28,8),
    fee_amount DECIMAL(28,8),
    fee_currency VARCHAR(20),
    timestamp TIMESTAMPTZ NOT NULL,
    platform VARCHAR(50),
    tx_hash VARCHAR(255),
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Notification Schema (Alerts and Notifications)
CREATE TABLE notify.alert_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    asset_symbol VARCHAR(20) REFERENCES market.assets(symbol),
    alert_type VARCHAR(50) NOT NULL,
    threshold DECIMAL(28,8),
    comparison_operator VARCHAR(10),
    severity alert_severity DEFAULT 'info',
    notification_channels notification_type[] DEFAULT '{email}',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notify.notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    alert_config_id UUID REFERENCES notify.alert_configs(id),
    type notification_type NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    metadata JSONB,
    read_at TIMESTAMPTZ,
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Audit Schema (Audit Logging)
CREATE TABLE audit.user_activity_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES auth.users(id),
    action VARCHAR(50) NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    entity_id UUID,
    old_value JSONB,
    new_value JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Convert market.price_data to TimescaleDB hypertable
SELECT create_hypertable('market.price_data', 'timestamp');

-- Create indexes
CREATE INDEX idx_portfolio_transactions_timestamp ON portfolio.transactions(portfolio_id, timestamp);
CREATE INDEX idx_market_price_data_timestamp ON market.price_data(timestamp DESC);
CREATE INDEX idx_whale_alerts_timestamp ON market.whale_alerts(timestamp DESC);
CREATE INDEX idx_notifications_user_read ON notify.notifications(user_id, read_at NULLS FIRST);

-- Create roles and permissions
CREATE ROLE app_auth;
CREATE ROLE app_core;
CREATE ROLE app_market;
CREATE ROLE app_portfolio;
CREATE ROLE app_notify;
CREATE ROLE app_audit;

-- Grant permissions per schema
GRANT USAGE ON SCHEMA auth TO app_auth;
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA auth TO app_auth;

GRANT USAGE ON SCHEMA core TO app_core;
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA core TO app_core;

GRANT USAGE ON SCHEMA market TO app_market;
GRANT SELECT, INSERT ON ALL TABLES IN SCHEMA market TO app_market;

GRANT USAGE ON SCHEMA portfolio TO app_portfolio;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA portfolio TO app_portfolio;

GRANT USAGE ON SCHEMA notify TO app_notify;
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA notify TO app_notify;

GRANT USAGE ON SCHEMA audit TO app_audit;
GRANT INSERT ON ALL TABLES IN SCHEMA audit TO app_audit;
GRANT SELECT ON ALL TABLES IN SCHEMA audit TO app_audit;


