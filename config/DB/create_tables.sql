--
-- PostgreSQL database dump
--
-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.1
-- Started on 2024-01-28 18:26:15 CET
SET
    statement_timeout = 0;

SET
    lock_timeout = 0;

SET
    idle_in_transaction_session_timeout = 0;

SET
    client_encoding = 'UTF8';

SET
    standard_conforming_strings = on;

SELECT
    pg_catalog.set_config('search_path', '', false);

SET
    check_function_bodies = false;

SET
    xmloption = content;

SET
    client_min_messages = warning;

SET
    row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--
CREATE SCHEMA public;

ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3435 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--
COMMENT ON SCHEMA public IS 'standard public schema';

SET
    default_tablespace = '';

SET
    default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 16398)
-- Name: auth; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.auth (
    auth_id integer NOT NULL,
    user_id integer,
    password_hash character(60) NOT NULL,
    last_login timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    public.auth OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16397)
-- Name: auth_auth_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.auth_auth_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.auth_auth_id_seq OWNER TO postgres;

--
-- TOC entry 3436 (class 0 OID 0)
-- Dependencies: 217
-- Name: auth_auth_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.auth_auth_id_seq OWNED BY public.auth.auth_id;

--
-- TOC entry 228 (class 1259 OID 16464)
-- Name: invoices; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.invoices (
    invoice_id integer NOT NULL,
    user_id integer,
    purchase_id integer,
    total_amount numeric(10, 2) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    public.invoices OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 16463)
-- Name: invoices_invoice_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.invoices_invoice_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.invoices_invoice_id_seq OWNER TO postgres;

--
-- TOC entry 3437 (class 0 OID 0)
-- Dependencies: 227
-- Name: invoices_invoice_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.invoices_invoice_id_seq OWNED BY public.invoices.invoice_id;

--
-- TOC entry 220 (class 1259 OID 16411)
-- Name: product; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.product (
    product_id integer NOT NULL,
    name character varying(50) NOT NULL,
    pricing numeric(10, 2) NOT NULL,
    description character varying(100) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    public.product OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16410)
-- Name: product_product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.product_product_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.product_product_id_seq OWNER TO postgres;

--
-- TOC entry 3438 (class 0 OID 0)
-- Dependencies: 219
-- Name: product_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.product_product_id_seq OWNED BY public.product.product_id;

--
-- TOC entry 222 (class 1259 OID 16421)
-- Name: product_stock; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.product_stock (
    product_stock_id integer NOT NULL,
    product_id integer,
    quantity integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    public.product_stock OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16420)
-- Name: product_stock_product_stock_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.product_stock_product_stock_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.product_stock_product_stock_id_seq OWNER TO postgres;

--
-- TOC entry 3439 (class 0 OID 0)
-- Dependencies: 221
-- Name: product_stock_product_stock_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.product_stock_product_stock_id_seq OWNED BY public.product_stock.product_stock_id;

--
-- TOC entry 226 (class 1259 OID 16447)
-- Name: purchase_items; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.purchase_items (
    item_id integer NOT NULL,
    purchase_id integer,
    product_id integer,
    quantity integer NOT NULL,
    price_per_unit numeric(10, 2) NOT NULL,
    total_price numeric(10, 2) NOT NULL
);

ALTER TABLE
    public.purchase_items OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16446)
-- Name: purchase_items_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.purchase_items_item_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.purchase_items_item_id_seq OWNER TO postgres;

--
-- TOC entry 3440 (class 0 OID 0)
-- Dependencies: 225
-- Name: purchase_items_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.purchase_items_item_id_seq OWNED BY public.purchase_items.item_id;

--
-- TOC entry 224 (class 1259 OID 16434)
-- Name: user_purchases; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.user_purchases (
    purchase_id integer NOT NULL,
    user_id integer,
    purchase_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    public.user_purchases OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16433)
-- Name: user_purchases_purchase_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.user_purchases_purchase_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.user_purchases_purchase_id_seq OWNER TO postgres;

--
-- TOC entry 3441 (class 0 OID 0)
-- Dependencies: 223
-- Name: user_purchases_purchase_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.user_purchases_purchase_id_seq OWNED BY public.user_purchases.purchase_id;

--
-- TOC entry 216 (class 1259 OID 16386)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--
CREATE TABLE public.users (
    user_id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    public.users OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16385)
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--
CREATE SEQUENCE public.users_user_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.users_user_id_seq OWNER TO postgres;

--
-- TOC entry 3442 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--
ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;

--
-- TOC entry 3235 (class 2604 OID 16401)
-- Name: auth auth_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.auth
ALTER COLUMN
    auth_id
SET
    DEFAULT nextval('public.auth_auth_id_seq' :: regclass);

--
-- TOC entry 3244 (class 2604 OID 16467)
-- Name: invoices invoice_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.invoices
ALTER COLUMN
    invoice_id
SET
    DEFAULT nextval('public.invoices_invoice_id_seq' :: regclass);

--
-- TOC entry 3237 (class 2604 OID 16414)
-- Name: product product_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.product
ALTER COLUMN
    product_id
SET
    DEFAULT nextval('public.product_product_id_seq' :: regclass);

--
-- TOC entry 3239 (class 2604 OID 16424)
-- Name: product_stock product_stock_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.product_stock
ALTER COLUMN
    product_stock_id
SET
    DEFAULT nextval(
        'public.product_stock_product_stock_id_seq' :: regclass
    );

--
-- TOC entry 3243 (class 2604 OID 16450)
-- Name: purchase_items item_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.purchase_items
ALTER COLUMN
    item_id
SET
    DEFAULT nextval('public.purchase_items_item_id_seq' :: regclass);

--
-- TOC entry 3241 (class 2604 OID 16437)
-- Name: user_purchases purchase_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.user_purchases
ALTER COLUMN
    purchase_id
SET
    DEFAULT nextval(
        'public.user_purchases_purchase_id_seq' :: regclass
    );

--
-- TOC entry 3233 (class 2604 OID 16389)
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.users
ALTER COLUMN
    user_id
SET
    DEFAULT nextval('public.users_user_id_seq' :: regclass);

--
-- TOC entry 3419 (class 0 OID 16398)
-- Dependencies: 218
-- Data for Name: auth; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO
    public.auth
VALUES
    (
        1,
        1,
        '$2a$10$nQc7l/HKongmYnsXJEYd0u5OOb0ywhGi85weSt5A3PjkCycD0l1JK',
        NULL,
        '2024-01-28 17:11:24.884884+00'
    );

INSERT INTO
    public.auth
VALUES
    (
        2,
        2,
        '$2a$10$bl7x8ZVkrS8CDTVT3PgLHudn9M2NOusP6u8E54ABLGW2BdfQxgQQK',
        NULL,
        '2024-01-28 17:11:45.069208+00'
    );

--
-- TOC entry 3429 (class 0 OID 16464)
-- Dependencies: 228
-- Data for Name: invoices; Type: TABLE DATA; Schema: public; Owner: postgres
--
--
-- TOC entry 3421 (class 0 OID 16411)
-- Dependencies: 220
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO
    public.product
VALUES
    (
        1,
        'Producto 1',
        5.00,
        'test',
        '2024-01-28 17:12:06.766049+00'
    );

INSERT INTO
    public.product
VALUES
    (
        2,
        'Producto 2',
        10.00,
        'test',
        '2024-01-28 17:12:15.280825+00'
    );

INSERT INTO
    public.product
VALUES
    (
        3,
        'Producto 3',
        20.00,
        'test',
        '2024-01-28 17:12:21.28582+00'
    );

--
-- TOC entry 3423 (class 0 OID 16421)
-- Dependencies: 222
-- Data for Name: product_stock; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO
    public.product_stock
VALUES
    (1, 1, 5, '2024-01-28 17:13:16.938477+00');

INSERT INTO
    public.product_stock
VALUES
    (2, 2, 6, '2024-01-28 17:13:23.300338+00');

--
-- TOC entry 3427 (class 0 OID 16447)
-- Dependencies: 226
-- Data for Name: purchase_items; Type: TABLE DATA; Schema: public; Owner: postgres
--
--
-- TOC entry 3425 (class 0 OID 16434)
-- Dependencies: 224
-- Data for Name: user_purchases; Type: TABLE DATA; Schema: public; Owner: postgres
--
--
-- TOC entry 3417 (class 0 OID 16386)
-- Dependencies: 216
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO
    public.users
VALUES
    (
        1,
        'user2',
        'user2@user.com',
        '2024-01-28 17:11:24.880579+00'
    );

INSERT INTO
    public.users
VALUES
    (
        2,
        'user1',
        'user1@user.com',
        '2024-01-28 17:11:45.063935+00'
    );

--
-- TOC entry 3443 (class 0 OID 0)
-- Dependencies: 217
-- Name: auth_auth_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval('public.auth_auth_id_seq', 2, true);

--
-- TOC entry 3444 (class 0 OID 0)
-- Dependencies: 227
-- Name: invoices_invoice_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval('public.invoices_invoice_id_seq', 1, false);

--
-- TOC entry 3445 (class 0 OID 0)
-- Dependencies: 219
-- Name: product_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval('public.product_product_id_seq', 3, true);

--
-- TOC entry 3446 (class 0 OID 0)
-- Dependencies: 221
-- Name: product_stock_product_stock_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval(
        'public.product_stock_product_stock_id_seq',
        2,
        true
    );

--
-- TOC entry 3447 (class 0 OID 0)
-- Dependencies: 225
-- Name: purchase_items_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval('public.purchase_items_item_id_seq', 1, false);

--
-- TOC entry 3448 (class 0 OID 0)
-- Dependencies: 223
-- Name: user_purchases_purchase_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval(
        'public.user_purchases_purchase_id_seq',
        1,
        false
    );

--
-- TOC entry 3449 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--
SELECT
    pg_catalog.setval('public.users_user_id_seq', 2, true);

--
-- TOC entry 3253 (class 2606 OID 16404)
-- Name: auth auth_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.auth
ADD
    CONSTRAINT auth_pkey PRIMARY KEY (auth_id);

--
-- TOC entry 3265 (class 2606 OID 16470)
-- Name: invoices invoices_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.invoices
ADD
    CONSTRAINT invoices_pkey PRIMARY KEY (invoice_id);

--
-- TOC entry 3255 (class 2606 OID 16419)
-- Name: product product_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.product
ADD
    CONSTRAINT product_name_key UNIQUE (name);

--
-- TOC entry 3257 (class 2606 OID 16417)
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.product
ADD
    CONSTRAINT product_pkey PRIMARY KEY (product_id);

--
-- TOC entry 3259 (class 2606 OID 16427)
-- Name: product_stock product_stock_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.product_stock
ADD
    CONSTRAINT product_stock_pkey PRIMARY KEY (product_stock_id);

--
-- TOC entry 3263 (class 2606 OID 16452)
-- Name: purchase_items purchase_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.purchase_items
ADD
    CONSTRAINT purchase_items_pkey PRIMARY KEY (item_id);

--
-- TOC entry 3261 (class 2606 OID 16440)
-- Name: user_purchases user_purchases_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.user_purchases
ADD
    CONSTRAINT user_purchases_pkey PRIMARY KEY (purchase_id);

--
-- TOC entry 3247 (class 2606 OID 16396)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.users
ADD
    CONSTRAINT users_email_key UNIQUE (email);

--
-- TOC entry 3249 (class 2606 OID 16392)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.users
ADD
    CONSTRAINT users_pkey PRIMARY KEY (user_id);

--
-- TOC entry 3251 (class 2606 OID 16394)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.users
ADD
    CONSTRAINT users_username_key UNIQUE (username);

--
-- TOC entry 3266 (class 2606 OID 16405)
-- Name: auth auth_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.auth
ADD
    CONSTRAINT auth_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);

--
-- TOC entry 3271 (class 2606 OID 16476)
-- Name: invoices invoices_purchase_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.invoices
ADD
    CONSTRAINT invoices_purchase_id_fkey FOREIGN KEY (purchase_id) REFERENCES public.user_purchases(purchase_id);

--
-- TOC entry 3272 (class 2606 OID 16471)
-- Name: invoices invoices_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.invoices
ADD
    CONSTRAINT invoices_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);

--
-- TOC entry 3267 (class 2606 OID 16428)
-- Name: product_stock product_stock_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.product_stock
ADD
    CONSTRAINT product_stock_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(product_id);

--
-- TOC entry 3269 (class 2606 OID 16458)
-- Name: purchase_items purchase_items_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.purchase_items
ADD
    CONSTRAINT purchase_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(product_id);

--
-- TOC entry 3270 (class 2606 OID 16453)
-- Name: purchase_items purchase_items_purchase_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.purchase_items
ADD
    CONSTRAINT purchase_items_purchase_id_fkey FOREIGN KEY (purchase_id) REFERENCES public.user_purchases(purchase_id);

--
-- TOC entry 3268 (class 2606 OID 16441)
-- Name: user_purchases user_purchases_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--
ALTER TABLE
    ONLY public.user_purchases
ADD
    CONSTRAINT user_purchases_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);

-- Completed on 2024-01-28 18:26:15 CET
--
-- PostgreSQL database dump complete
--