-- Função para atualizar updated_at automaticamente
CREATE OR REPLACE FUNCTION update_updated_at_column () RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at
= NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Tabela users
CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     name VARCHAR(100) NOT NULL,
     email VARCHAR(100) UNIQUE NOT NULL,
     age INTEGER NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_users_updated_at BEFORE
    UPDATE ON users FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

-- Tabela products
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_products_updated_at BEFORE
    UPDATE ON products FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

-- Tabela orders
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),
    total NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_orders_updated_at BEFORE
    UPDATE ON orders FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

-- Tabela categories
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_categories_updated_at BEFORE
    UPDATE ON categories FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

-- Tabela reviews
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products (id),
    rating INTEGER NOT NULL CHECK ( rating >= 1 AND rating <= 5 ),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_reviews_updated_at BEFORE
    UPDATE ON reviews FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column ();

-- Cada um pode escolher uma dessas entidades para implementar o CRUD na API.
-- Isso ajuda a praticar a arquitetura hexagonal com diferentes entidades.