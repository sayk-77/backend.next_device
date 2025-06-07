--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: order_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.order_status AS ENUM (
    'pending',
    'paid',
    'shipped',
    'delivered',
    'canceled',
    'returned'
);


ALTER TYPE public.order_status OWNER TO postgres;

--
-- Name: payment_method; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.payment_method AS ENUM (
    'credit_card',
    'paypal',
    'bank_transfer'
);


ALTER TYPE public.payment_method OWNER TO postgres;

--
-- Name: payment_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.payment_status AS ENUM (
    'pending',
    'completed',
    'failed'
);


ALTER TYPE public.payment_status OWNER TO postgres;

--
-- Name: role; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.role AS ENUM (
    'customer',
    'admin',
    'moderator'
);


ALTER TYPE public.role OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.addresses (
    id bigint NOT NULL,
    user_id bigint,
    country text NOT NULL,
    city text NOT NULL,
    street text NOT NULL,
    postal_code text NOT NULL
);


ALTER TABLE public.addresses OWNER TO postgres;

--
-- Name: addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.addresses_id_seq OWNER TO postgres;

--
-- Name: addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


--
-- Name: brand_banners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brand_banners (
    id bigint NOT NULL,
    brand_id bigint NOT NULL,
    title text,
    image_url text
);


ALTER TABLE public.brand_banners OWNER TO postgres;

--
-- Name: brand_banners_brand_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.brand_banners_brand_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.brand_banners_brand_id_seq OWNER TO postgres;

--
-- Name: brand_banners_brand_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brand_banners_brand_id_seq OWNED BY public.brand_banners.brand_id;


--
-- Name: brand_banners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.brand_banners_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.brand_banners_id_seq OWNER TO postgres;

--
-- Name: brand_banners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brand_banners_id_seq OWNED BY public.brand_banners.id;


--
-- Name: brands; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brands (
    id bigint NOT NULL,
    name text NOT NULL,
    image_url text
);


ALTER TABLE public.brands OWNER TO postgres;

--
-- Name: brands_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.brands_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.brands_id_seq OWNER TO postgres;

--
-- Name: brands_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brands_id_seq OWNED BY public.brands.id;


--
-- Name: cart_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cart_items (
    id bigint NOT NULL,
    cart_id bigint,
    quantity bigint,
    product_id bigint
);


ALTER TABLE public.cart_items OWNER TO postgres;

--
-- Name: cart_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cart_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cart_items_id_seq OWNER TO postgres;

--
-- Name: cart_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cart_items_id_seq OWNED BY public.cart_items.id;


--
-- Name: carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.carts (
    id bigint NOT NULL,
    user_id bigint
);


ALTER TABLE public.carts OWNER TO postgres;

--
-- Name: carts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.carts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.carts_id_seq OWNER TO postgres;

--
-- Name: carts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.carts_id_seq OWNED BY public.carts.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    name text NOT NULL,
    parent_category_id bigint,
    title text,
    category_image text
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: laptop_filters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.laptop_filters (
    id bigint NOT NULL,
    laptop_id bigint,
    processor_brand text NOT NULL,
    processor_name text NOT NULL,
    graphics_card_type text NOT NULL,
    graphics_card_name text NOT NULL,
    ram bigint NOT NULL,
    storage bigint NOT NULL,
    battery_capacity bigint NOT NULL,
    operating_system text NOT NULL,
    screen_size numeric NOT NULL,
    screen_refresh_rate bigint NOT NULL,
    screen_type text NOT NULL,
    body_material text NOT NULL,
    width numeric NOT NULL,
    height numeric NOT NULL,
    depth numeric NOT NULL,
    weight numeric NOT NULL
);


ALTER TABLE public.laptop_filters OWNER TO postgres;

--
-- Name: laptop_filters_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.laptop_filters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.laptop_filters_id_seq OWNER TO postgres;

--
-- Name: laptop_filters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.laptop_filters_id_seq OWNED BY public.laptop_filters.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    id bigint NOT NULL,
    order_id bigint,
    product_id bigint,
    quantity bigint NOT NULL,
    price numeric(10,2) NOT NULL
);


ALTER TABLE public.order_items OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    user_id bigint,
    total_price numeric(10,2) NOT NULL,
    status public.order_status DEFAULT 'pending'::public.order_status NOT NULL,
    address_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id bigint NOT NULL,
    order_id bigint,
    payment_id text,
    payment_method public.payment_method NOT NULL,
    payment_status public.payment_status DEFAULT 'pending'::public.payment_status NOT NULL,
    amount numeric(10,2) NOT NULL,
    payment_date timestamp with time zone
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: product_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_details (
    id bigint NOT NULL,
    product_id bigint,
    processor text,
    ram text,
    storage text,
    display text,
    camera text,
    battery text,
    os text,
    dimensions text,
    weight text,
    graphics_card text
);


ALTER TABLE public.product_details OWNER TO postgres;

--
-- Name: product_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_details_id_seq OWNER TO postgres;

--
-- Name: product_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_details_id_seq OWNED BY public.product_details.id;


--
-- Name: product_filters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_filters (
    id bigint NOT NULL,
    product_id bigint NOT NULL,
    display_size numeric(10,2),
    ram bigint NOT NULL,
    storage bigint NOT NULL,
    camera_quality bigint NOT NULL,
    processor character varying(100),
    battery bigint NOT NULL,
    os character varying(50),
    width numeric(10,2),
    height numeric(10,2),
    length numeric(10,2),
    weight numeric(10,2)
);


ALTER TABLE public.product_filters OWNER TO postgres;

--
-- Name: product_filters_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_filters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_filters_id_seq OWNER TO postgres;

--
-- Name: product_filters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_filters_id_seq OWNED BY public.product_filters.id;


--
-- Name: product_images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_images (
    id bigint NOT NULL,
    product_id bigint,
    image_url text NOT NULL,
    is_main boolean
);


ALTER TABLE public.product_images OWNER TO postgres;

--
-- Name: product_images_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_images_id_seq OWNER TO postgres;

--
-- Name: product_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_images_id_seq OWNED BY public.product_images.id;


--
-- Name: product_variants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_variants (
    id bigint NOT NULL,
    product_id bigint,
    variant_name text NOT NULL,
    variant_value text NOT NULL
);


ALTER TABLE public.product_variants OWNER TO postgres;

--
-- Name: product_variants_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_variants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_variants_id_seq OWNER TO postgres;

--
-- Name: product_variants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_variants_id_seq OWNED BY public.product_variants.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    sku text NOT NULL,
    name text NOT NULL,
    description text,
    price numeric(10,2) NOT NULL,
    discount_price numeric(10,2) DEFAULT 0,
    stock bigint NOT NULL,
    category_id bigint,
    brand_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    search_name text
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: review_images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.review_images (
    id bigint NOT NULL,
    review_id bigint,
    image_url character varying(255)
);


ALTER TABLE public.review_images OWNER TO postgres;

--
-- Name: review_images_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.review_images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.review_images_id_seq OWNER TO postgres;

--
-- Name: review_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.review_images_id_seq OWNED BY public.review_images.id;


--
-- Name: review_images_review_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.review_images_review_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.review_images_review_id_seq OWNER TO postgres;

--
-- Name: review_images_review_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.review_images_review_id_seq OWNED BY public.review_images.review_id;


--
-- Name: reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reviews (
    id bigint NOT NULL,
    user_id bigint,
    product_id bigint,
    rating bigint NOT NULL,
    created_at timestamp with time zone,
    pros text,
    cons text,
    comment text,
    is_moder boolean DEFAULT false,
    CONSTRAINT chk_reviews_rating CHECK (((rating >= 1) AND (rating <= 5)))
);


ALTER TABLE public.reviews OWNER TO postgres;

--
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reviews_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reviews_id_seq OWNER TO postgres;

--
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    email text NOT NULL,
    password_hash text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    role public.role DEFAULT 'customer'::public.role NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: addresses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


--
-- Name: brand_banners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand_banners ALTER COLUMN id SET DEFAULT nextval('public.brand_banners_id_seq'::regclass);


--
-- Name: brand_banners brand_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand_banners ALTER COLUMN brand_id SET DEFAULT nextval('public.brand_banners_brand_id_seq'::regclass);


--
-- Name: brands id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brands ALTER COLUMN id SET DEFAULT nextval('public.brands_id_seq'::regclass);


--
-- Name: cart_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items ALTER COLUMN id SET DEFAULT nextval('public.cart_items_id_seq'::regclass);


--
-- Name: carts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts ALTER COLUMN id SET DEFAULT nextval('public.carts_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: laptop_filters id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.laptop_filters ALTER COLUMN id SET DEFAULT nextval('public.laptop_filters_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Name: product_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details ALTER COLUMN id SET DEFAULT nextval('public.product_details_id_seq'::regclass);


--
-- Name: product_filters id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_filters ALTER COLUMN id SET DEFAULT nextval('public.product_filters_id_seq'::regclass);


--
-- Name: product_images id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_images ALTER COLUMN id SET DEFAULT nextval('public.product_images_id_seq'::regclass);


--
-- Name: product_variants id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_variants ALTER COLUMN id SET DEFAULT nextval('public.product_variants_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: review_images id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review_images ALTER COLUMN id SET DEFAULT nextval('public.review_images_id_seq'::regclass);


--
-- Name: review_images review_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review_images ALTER COLUMN review_id SET DEFAULT nextval('public.review_images_review_id_seq'::regclass);


--
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.addresses (id, user_id, country, city, street, postal_code) FROM stdin;
6	4	Россия	Псков	Псковская	125152
\.


--
-- Data for Name: brand_banners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brand_banners (id, brand_id, title, image_url) FROM stdin;
1	1	Xiaomi 11T Pro	ximi_banner_3.webp
2	1	Xiaomi Mi 11	ximi_banner_1.jpg
3	1	Xiaomi Redmi Note 9	ximi_banner_2.webp
9	17	Samsung S 22	samsung_banner_1.jpg
8	17	Samsung	samsung_banner_3.png
7	17	Samsung S10+	samsung_banner_2.jfif
10	12	Iphone 16	apple_banner_1.jpg
11	12	Mackbook Pro	apple_banner_2.webp
12	12	Mackbook Air M3	apple_banner_3.jpg
13	19	Acer	acer_banner_1.jpg
14	19	Acer	acer_banner_3.jpg
15	19	Acer	acer_banner_2.jpg
16	18	Honor	honor_banner_1.jpg
17	18	Honor	honor_banner_3.jpg
18	18	Honor	honor_banner_2.webp
\.


--
-- Data for Name: brands; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brands (id, name, image_url) FROM stdin;
1	Xiaomi	brand_xiaomi.webp
12	Apple	brand_apple.webp
17	Samsung	brand_samsung.webp
18	Honor	brand_honor.webp
19	Acer	brand_acer.webp
\.


--
-- Data for Name: cart_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cart_items (id, cart_id, quantity, product_id) FROM stdin;
11	1	1	9
12	1	1	64
13	1	1	66
14	1	1	67
15	1	1	67
16	1	1	67
17	1	1	67
18	1	1	67
19	1	1	13
20	1	1	64
21	1	1	13
22	1	1	13
23	1	1	13
24	1	1	9
25	1	1	67
26	1	1	71
27	1	1	64
28	1	1	66
29	1	1	13
30	1	1	12
31	1	1	11
32	1	1	66
33	1	1	65
34	1	1	68
35	1	1	67
36	1	1	9
\.


--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.carts (id, user_id) FROM stdin;
1	4
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, parent_category_id, title, category_image) FROM stdin;
1	mobile	\N	Смартфоны	mobile.png
4	laptop	\N	Ноутбуки	laptop.png
5	tablet	\N	Планшеты	tablet.png
6	smart-watch	\N	Умные часы	watch.png
7	pc	\N	Компьютеры	pc.png
8	accessories	\N	Аксесуары	accessories.png
\.


--
-- Data for Name: laptop_filters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.laptop_filters (id, laptop_id, processor_brand, processor_name, graphics_card_type, graphics_card_name, ram, storage, battery_capacity, operating_system, screen_size, screen_refresh_rate, screen_type, body_material, width, height, depth, weight) FROM stdin;
3	70	Intel	Core i7-1260P	integrated	Intel Iris Xe	16	512	6800	Windows 11 Home	13.3	60	AMOLED	aluminum	305.4	202.9	11.2	1.04
4	66	Intel	Core i5-1135G7	integrated	Intel Iris Xe	8	512	4800	Windows 11 Home	14.0	60	IPS	aluminum	318.1	208.6	15.9	1.2
5	64	Intel	Core i9-12900HX	Dedicated	NVIDIA GeForce RTX 3080	32	1000	7500	Windows 11	18.0	165	IPS	Aluminum	412.5	280.0	22.5	3.5
6	65	Intel	Core i5-1135G7	Integrated	Intel Iris Xe	8	512	4800	Windows 10 Home	15.6	60	IPS	Plastic	363.4	238.4	19.9	1.8
7	66	AMD	Ryzen 7 5800U	Integrated	AMD Radeon Vega 8	16	512	4800	Windows 11	14.0	60	IPS	Aluminum	314.0	214.0	16.5	1.2
8	67	AMD	Ryzen 5 4600U	Integrated	AMD Radeon Vega 6	8	512	4600	Windows 10 Home	14.0	60	IPS	Aluminum	323.0	215.0	16.9	1.38
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_items (id, order_id, product_id, quantity, price) FROM stdin;
1	1	13	1	120990.00
2	1	9	1	87990.00
3	1	67	1	40000.00
4	1	71	1	25000.00
21	4	66	2	65000.00
22	4	13	1	120990.00
23	4	12	1	80990.00
24	4	11	1	119990.00
25	4	65	1	42000.00
26	4	68	1	70000.00
27	4	67	1	38000.00
28	5	9	1	87990.00
29	6	9	1	87990.00
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, user_id, total_price, status, address_id, created_at, updated_at) FROM stdin;
4	4	601970.00	delivered	6	2024-11-07 19:02:52.325432+03	2024-11-20 22:41:06.113385+03
1	4	273980.00	delivered	6	2024-11-07 18:09:24.732416+03	2024-11-20 22:53:56.066676+03
5	4	87990.00	delivered	6	2024-11-20 22:54:48.395205+03	2024-11-20 22:56:14.90023+03
6	4	87990.00	pending	6	2024-11-20 23:04:28.215035+03	2024-11-20 23:04:28.215035+03
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payments (id, order_id, payment_id, payment_method, payment_status, amount, payment_date) FROM stdin;
1	1	cs_test_a1LhT5LxscogMMG26jcuqzfsqQXWNOnmnKxdfQj5kxzlwS7YcyXCJ67zTf	credit_card	completed	27398000.00	0001-01-01 02:30:17+02:30:17
2	1	cs_test_a1LhT5LxscogMMG26jcuqzfsqQXWNOnmnKxdfQj5kxzlwS7YcyXCJ67zTf	credit_card	completed	27398000.00	0001-01-01 02:30:17+02:30:17
3	4	cs_test_a1D3LWWFTs3AbUkrjCrULobrqhiOupudIJOhld8MgcU9ogB1paSe38Essn	credit_card	completed	60197000.00	0001-01-01 02:30:17+02:30:17
4	5	cs_test_a1YXvwFdhUZAHLNidSGcPG637orudNtJX6q9ar5E3oK983vMaF2O05OSpZ	credit_card	completed	8799000.00	0001-01-01 02:30:17+02:30:17
5	6	cs_test_a1S4NDcnf01X94ZM6Wjys7bXFjfrQESZyFdVtiozEwlMsBop3dguOxHmCJ	credit_card	completed	8799000.00	0001-01-01 02:30:17+02:30:17
\.


--
-- Data for Name: product_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_details (id, product_id, processor, ram, storage, display, camera, battery, os, dimensions, weight, graphics_card) FROM stdin;
13	1	Snapdragon 8 gen 3, 3ггц	12 ГБ LPDDR5X	512 ГБ UFS 4.0	6.73" AMOLED, 3200x1440, 120 Гц	Sony IMX989, 50 Мп	5000 мА*ч, поддержка быстрой зарядки 120 Вт	Android 14 с оболочкой HyperOS	163.2x74.6x9.1 мм	234 г	
14	10	Какой то	много	много	Крутой	крутая	5000	Андроед	компактный	тяжелый	
15	64	Intel Core i7-13700HX	16 GB DDR5	1 TB SSD	18" QHD 2560x1600, 165Hz, IPS	HD 720p	90 Wh, до 6 часов работы	Windows 11 Home	399.2 x 299.3 x 23.5 mm	2.9 kg	
17	8	Qualcomm Snapdragon 870	6 GB LPDDR4X	256 GB UFS 3.1	11" LCD, 2560x1600, 120Hz	13 MP rear, 8 MP front	8840 mAh, 33W fast charging	Android 12, MIUI for Pad	254.7 x 166.3 x 6.9 mm	515 g	Adreno 650
18	9	Qualcomm Snapdragon 8 Gen 2	12 GB LPDDR5X	512 GB UFS 3.1	6.81" OLED, 3200x1440, 120Hz	50+50+50 MP, 3D ToF, OIS	4600 mAh, 100W fast charging	MagicOS 7.0 (Android 13)	163.3 x 74.1 x 8.9 mm	209 g	Adreno 740
19	10	Qualcomm Snapdragon 7 Gen 1	8 GB LPDDR4X	128 GB UFS 2.1	6.7" AMOLED, 2664x1200, 120Hz	200 MP + 12 MP + 2 MP, OIS	5000 mAh, 66W fast charging	MagicOS 7.0 (Android 13)	162.9 x 76.1 x 8.1 mm	183 g	Adreno 662
20	11	Apple A17 Pro	8 GB	128 GB / 256 GB / 512 GB / 1 TB	6.1" OLED, 2532x1170, 120Hz	48 MP main, 12 MP telephoto, 12 MP ultra-wide	3200 mAh, 20W fast charging	iOS 17	146.7 x 70.9 x 8.25 mm	187 g	Apple GPU (6-core)
21	12	Apple A15 Bionic	6 GB	128 GB / 256 GB / 512 GB	6.1" OLED, 2532x1170, 60Hz	12 MP main, 12 MP ultra-wide	3279 mAh, 20W fast charging	iOS 16	146.7 x 71.5 x 7.8 mm	172 g	Apple GPU (5-core)
22	13	Qualcomm Snapdragon 8 Gen 2	12 GB LPDDR5X	256 GB / 512 GB / 1 TB UFS 4.0	6.8" Dynamic AMOLED 2X, 3088x1440, 120Hz	200 MP main, 12 MP ultra-wide, 10 MP telephoto, 10 MP periscope	5000 mAh, 45W fast charging	Android 13, One UI 5.1	163.4 x 78.1 x 8.9 mm	234 g	Adreno 740
23	64	Intel Core i7-13700HX	16 GB DDR5	1 TB SSD	18" QHD 2560x1600, 165Hz, IPS	N/A	5000 mAh, 80W fast charging	Windows 11	411.2 x 286.6 x 22.9 mm	2.75 kg	NVIDIA GeForce RTX 4060
24	65	AMD Ryzen 5 5500U	8 GB LPDDR4	512 GB SSD	15.6" FHD, 1920x1080, 60Hz	N/A	48Wh, 65W fast charging	Windows 11	359.7 x 236.5 x 19.9 mm	1.8 kg	Integrated AMD Radeon
25	66	Intel Core i5-1135G7	8 GB LPDDR4X	512 GB SSD	14" FHD, 1920x1080, 60Hz	N/A	50Wh, 65W fast charging	Windows 10	319.9 x 210.5 x 17.95 mm	1.2 kg	Intel Iris Xe
26	67	Intel Core i7-1165G7	16 GB LPDDR4X	512 GB SSD	14" FHD, 1920x1080, 60Hz	N/A	56Wh, 65W fast charging	Windows 10	324.2 x 229.9 x 15.9 mm	1.38 kg	Intel Iris Xe
27	68	Intel Core i7-11800H	16 GB DDR4	512 GB SSD	15.6" FHD, 1920x1080, 144Hz	N/A	56Wh, 135W fast charging	Windows 11	359.6 x 258.8 x 20.9 mm	2.4 kg	NVIDIA GeForce RTX 3060
28	69	Intel Core i5-1235U	8 GB LPDDR4X	256 GB SSD	15.6" FHD, 1920x1080, 60Hz	N/A	42Wh, 65W fast charging	Windows 11	362.3 x 252.4 x 15.9 mm	1.58 kg	Intel Iris Xe
29	70	Intel Core i7-1265U	16 GB LPDDR4X	512 GB SSD	15.6" FHD, 1920x1080, 60Hz	N/A	54Wh, 65W fast charging	Windows 11	357.6 x 230.3 x 15.9 mm	1.7 kg	Intel Iris Xe
30	71	Qualcomm Snapdragon 870	6 GB LPDDR4X	128 GB UFS 3.1	10.1" LCD, 2560x1600, 120Hz	13 MP rear, 8 MP front	8600 mAh, 33W fast charging	Android 11	254.7 x 166.3 x 6.9 mm	511 g	Adreno 650
31	72	Qualcomm Snapdragon 865	6 GB LPDDR4X	128 GB UFS 3.1	12.4" AMOLED, 2800x1752, 120Hz	13 MP rear, 8 MP front	10090 mAh, 45W fast charging	Android 11	285.0 x 185.0 x 6.3 mm	572 g	Adreno 650
32	73	Apple A14 Bionic	4 GB	64 GB / 256 GB	10.9" Liquid Retina, 2360x1640	12 MP rear, 7 MP front	7606 mAh, 20W fast charging	iPadOS 14	247.6 x 178.5 x 6.1 mm	458 g	Apple GPU (4-core)
33	74	Samsung Exynos W920	1 GB	16 GB	1.4" Super AMOLED, 450x450	N/A	247 mAh	WearOS 3.0	44.4 x 44.4 x 9.8 mm	30 g	N/A
34	75	Samsung Exynos W320	1 GB	4 GB	1.4" Super AMOLED, 360x360	N/A	247 mAh	Tizen OS	40 x 40 x 10.5 mm	25 g	N/A
35	76	Apple S7	1 GB	32 GB	1.9" Retina OLED, 396x484	N/A	309 mAh	watchOS 8	44 x 38 x 10.7 mm	32 g	N/A
36	77	Huawei Kirin A1	1 GB	4 GB	1.39" AMOLED, 454x454	N/A	455 mAh	LiteOS	46.7 x 46.7 x 10.7 mm	41 g	N/A
37	78	Honor	1 GB	4 GB	1.39" AMOLED, 454x454	N/A	455 mAh	LiteOS	48 x 48 x 11 mm	42 g	N/A
38	67	AMD Ryzen 5 5500U	16 GB LPDDR4X	512 GB SSD	14" FHD, 1920x1080, 60Hz	N/A	56Wh, 65W fast charging	Windows 10	324.2 x 229.9 x 15.9 mm	1.38 kg	Intel Iris Xe
39	65	Intel Core i5 10300H	8 GB LPDDR4	512 GB SSD	15.6" FHD, 1920x1080, 60Hz	N/A	48Wh, 65W fast charging	Windows 11	359.7 x 236.5 x 19.9 mm	1.8 kg	Integrated AMD Radeon
40	66	AMD Ryzen 7 5700U	8 GB LPDDR4X	512 GB SSD	14" FHD, 1920x1080, 60Hz	N/A	50Wh, 65W fast charging	Windows 10	319.9 x 210.5 x 17.95 mm	1.2 kg	Intel Iris Xe
42	83	123	123	123	123	123	123	123	123	123	
43	64	Intel Core i7-13700HX	16 GB DDR5	1 TB SSD	18" QHD 2560x1600, 165Hz, IPS	N/A	5000 mAh, 80W fast charging	Windows 11	411.2 x 286.6 x 22.9 mm	2.75 kg	NVIDIA GeForce RTX 4060
44	83	123	123	123	123	123	123	123	123	123	
\.


--
-- Data for Name: product_filters; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_filters (id, product_id, display_size, ram, storage, camera_quality, processor, battery, os, width, height, length, weight) FROM stdin;
1	1	6.73	12	512	50	Snapdragon 8 Gen 3	5000	Android	163.20	74.60	9.10	234.00
2	11	6.10	8	128	48	A17 Pro	3279	iOS	146.60	70.60	7.80	187.00
5	72	12.40	8	256	13	Snapdragon 865+	10090	Android	285.00	185.00	6.30	575.00
6	73	10.90	8	64	12	Apple M2	7566	iOS	247.60	178.50	6.10	461.00
8	8	11.00	8	256	13	Qualcomm Snapdragon 870	8840	Android	254.72	166.00	6.90	510.00
9	9	6.76	12	512	108	Kirin 9000	4600	Android	160.50	74.20	8.90	210.00
10	10	6.70	8	256	50	Snapdragon 7 Gen 1	4000	Android	161.00	73.30	7.90	183.00
12	12	6.10	6	128	12	A15 Bionic	3279	iOS	71.50	146.70	7.80	174.00
13	13	6.80	12	512	200	Snapdragon 8 Gen 2	5000	Android	163.40	78.10	8.90	234.00
14	71	11.00	6	128	13	Snapdragon 860	8720	Android	254.72	166.00	6.90	511.00
\.


--
-- Data for Name: product_images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_images (id, product_id, image_url, is_main) FROM stdin;
77	68	honor_hunter_700_main.webp	t
78	68	honor_hunter_700_1.webp	f
79	68	honor_hunter_700_2.webp	f
80	68	honor_hunter_700_3.webp	f
19	8	ximi_pad_main.webp	t
20	8	ximi_pad_2.webp	f
21	8	ximi_pad_7.webp	f
22	8	ximi_pad_6.webp	f
23	8	ximi_pad_5.webp	f
24	8	ximi_pad_4.webp	f
25	8	ximi_pad_3.webp	f
36	11	apple_15_pro_main.webp	t
37	11	apple_15_pro_1.webp	f
38	11	apple_15_pro_2.webp	f
39	11	apple_15_pro_3.webp	f
40	11	apple_15_pro_4.webp	f
41	12	apple_14_main.webp	t
42	12	apple_14_1.webp	f
43	12	apple_14_2.webp	f
44	12	apple_14_3.webp	f
45	12	apple_14_4.webp	f
46	13	samsung_s23_ultra_main.webp	t
47	13	samsung_s23_ultra_1.webp	f
48	13	samsung_s23_ultra_2.webp	f
49	13	samsung_s23_ultra_3.webp	f
50	13	samsung_s23_ultra_4.webp	f
26	9	honor_magic6_pro_main.webp	t
27	9	honor_magic6_pro_1.webp	f
28	9	honor_magic6_pro_2.webp	f
29	9	honor_magic6_pro_3.webp	f
30	9	honor_magic6_pro_4.webp	f
51	9	honor_magic6_pro_5.webp	f
54	11	apple_15_pro_5.webp	f
55	12	apple_14_5.webp	f
56	13	samsung_s23_ultra_5.webp	f
82	69	samsung_galaxy_book_3_main.webp	t
83	69	samsung_galaxy_book_3_1.webp	f
84	69	samsung_galaxy_book_3_2.webp	f
85	69	samsung_galaxy_book_3_3.webp	f
86	69	samsung_galaxy_book_3_4.webp	f
87	70	samsung_galaxy_book_4_pro_main.webp	t
88	70	samsung_galaxy_book_4_pro_1.webp	f
89	70	samsung_galaxy_book_4_pro_2.webp	f
90	70	samsung_galaxy_book_4_pro_3.webp	f
91	70	samsung_galaxy_book_4_pro_4.webp	f
92	71	xiaomi_pad_5_main.webp	t
93	71	xiaomi_pad_5_1.webp	f
94	71	xiaomi_pad_5_2.webp	f
95	71	xiaomi_pad_5_3.webp	f
96	71	xiaomi_pad_5_4.webp	f
102	73	ipad_air_m2_main.webp	t
103	73	ipad_air_m2_1.webp	f
104	73	ipad_air_m2_2.webp	f
105	73	ipad_air_m2_3.webp	f
107	74	samsung_watch_4_main.webp	t
108	74	samsung_watch_4_1.webp	f
109	74	samsung_watch_4_2.webp	f
110	74	samsung_watch_4_3.webp	f
111	74	samsung_watch_4_4.webp	f
112	75	samsung_watch_active_2_main.webp	t
113	75	samsung_watch_active_2_1.webp	f
114	75	samsung_watch_active_2_2.webp	f
81	68	honor_hunter_700_4.webp	f
7	1	ximi14u_main.webp	t
31	10	honor_90_main.webp	t
32	10	honor_90_1.webp	f
97	72	samsung_tab_7_plus_main.webp	t
98	72	samsung_tab_7_plus_1.webp	f
99	72	samsung_tab_7_plus_2.webp	f
100	72	samsung_tab_7_plus_3.webp	f
101	72	samsung_tab_7_plus_4.webp	f
33	10	honor_90_2.webp	f
34	10	honor_90_3.webp	f
35	10	honor_90_4.webp	f
53	10	honor_90_5.webp	f
72	67	honor_magic_14_main.webp	t
73	67	honor_magic_14_1.webp	f
74	67	honor_magic_14_2.webp	f
75	67	honor_magic_14_3.webp	f
76	67	honor_magic_14_4.webp	f
62	65	acer_aspire_5_main.webp	t
63	65	acer_aspire_5_1.webp	f
64	65	acer_aspire_5_2.webp	f
65	65	acer_aspire_5_3.webp	f
66	65	acer_aspire_5_4.webp	f
67	66	acer_swift_3_main.webp	t
68	66	acer_swift_3_1.webp	f
69	66	acer_swift_3_2.webp	f
70	66	acer_swift_3_3.webp	f
71	66	acer_swift_3_4.webp	f
57	64	acer_predator_helios_18_main.webp	t
58	64	acer_predator_helios_18_1.webp	f
59	64	acer_predator_helios_18_2.webp	f
60	64	acer_predator_helios_18_3.webp	f
61	64	acer_predator_helios_18_4.webp	f
115	75	samsung_watch_active_2_3.webp	f
116	75	samsung_watch_active_2_4.webp	f
117	76	apple_watch_series_7_main.webp	t
118	76	apple_watch_series_7_1.webp	f
119	76	apple_watch_series_7_2.webp	f
120	76	apple_watch_series_7_3.webp	f
121	76	apple_watch_series_7_4.webp	f
122	77	honor_magic_watch_2_main.webp	t
123	77	honor_magic_watch_2_1.webp	f
124	77	honor_magic_watch_2_2.webp	f
125	77	honor_magic_watch_2_3.webp	f
126	77	honor_magic_watch_2_4.webp	f
127	78	honor_watch_gs_pro_main.webp	t
128	78	honor_watch_gs_pro_1.webp	f
129	78	honor_watch_gs_pro_2.webp	f
130	78	honor_watch_gs_pro_3.webp	f
131	78	honor_watch_gs_pro_4.webp	f
132	72	samsung_tab_7_plus_5.webp	f
137	68	honor_hunter_700_5.webp	f
138	69	samsung_galaxy_book_3_5.webp	f
139	70	samsung_galaxy_book_4_pro_5.webp	f
140	71	xiaomi_pad_5_5.webp	f
141	74	samsung_watch_4_5.webp	f
142	75	samsung_watch_active_2_5.webp	f
143	76	apple_watch_series_7_5.webp	f
144	77	honor_magic_watch_2_5.webp	f
145	78	honor_watch_gs_pro_5.webp	f
146	1	ximi14u_2.webp	f
147	1	ximi14u_5.webp	f
148	1	ximi14u_4.webp	f
149	1	ximi14u_3.webp	f
136	67	honor_magic_14_5.webp	f
134	65	acer_aspire_5_5.webp	f
135	66	acer_swift_3_5.webp	f
133	64	acer_predator_helios_18_5.webp	f
151	83	image_83_honor_90_1.webp	f
152	83	image_83_honor_90_2.webp	f
153	83	image_83_honor_90_3.webp	f
150	83	image_83_apple_watch_series_7_main.webp	t
\.


--
-- Data for Name: product_variants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_variants (id, product_id, variant_name, variant_value) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, sku, name, description, price, discount_price, stock, category_id, brand_id, created_at, updated_at, search_name) FROM stdin;
11	SKU10003	Смартфон Apple iPhone 15 Pro	Apple iPhone 15 Pro — это один из самых мощных и технологичных смартфонов на рынке, который объединяет высокую производительность, уникальный дизайн и инновационные функции. Оснащен процессором A17 Bionic, который позволяет работать с любыми приложениями и играми без задержек и с минимальным энергопотреблением. 6.1-дюймовый Super Retina XDR дисплей с яркостью до 2000 нит и HDR10 делает просмотр видео и игр на устройстве незабываемым. Камера на 48 Мп с поддержкой 3x оптического зума позволяет делать потрясающие фотографии в любых условиях. Смартфон поддерживает 5G для сверхбыстрой связи и оснащен аккумулятором на 3200 мАч с поддержкой быстрой зарядки. Это идеальный выбор для тех, кто ищет флагманский смартфон с превосходными характеристиками и современными технологиями.	124990.00	5000.00	20	1	12	2024-10-15 23:07:08.948306+03	2024-10-15 23:07:08.948306+03	apple-iphone-15-pro
12	SKU10004	6.1" Смартфон Apple iPhone 14	Apple iPhone 14 с экраном 6.1 дюймов и технологией Super Retina XDR — это флагманский смартфон с улучшенными характеристиками и современными функциями. Внутри устройства установлен процессор A15 Bionic, который обеспечивает высокую производительность и энергоэффективность. Камера на 12 Мп с улучшенной оптической стабилизацией позволяет делать четкие и яркие фотографии даже при низком освещении. Устройство также поддерживает записи видео в 4K с улучшенной стабилизацией. iPhone 14 обладает аккумулятором на 3279 мАч, что обеспечит до 20 часов работы в режиме воспроизведения видео. Быстрая зарядка и поддержка беспроводной зарядки MagSafe делают использование смартфона более удобным. Кроме того, устройство поддерживает 5G-сети для максимально быстрой передачи данных. iPhone 14 сочетает в себе мощность и элегантность, предлагая идеальный баланс между производительностью и функциональностью.	89990.00	9000.00	25	1	12	2024-10-15 23:07:08.948306+03	2024-10-15 23:07:08.948306+03	apple-iphone-14
1	fawf41fw	Xiaomi 14 Ultra White	Xiaomi 14 Ultra White — флагманский смартфон с ультрасовременными характеристиками и элегантным дизайном. Оснащен мощным процессором Snapdragon 8 Gen 3, который обеспечивает безупречную производительность и скорость работы в любых приложениях и играх. 6.73-дюймовый AMOLED экран с разрешением 3200x1440 пикселей и частотой обновления 120 Гц позволяет наслаждаться ярким и плавным изображением в любых условиях. Основная камера с четырьмя сенсорами на 50 Мп + 50 Мп + 50 Мп + 50 Мп поддерживает записи видео в 8K и позволяет делать фотографии с высокой детализацией, даже при слабом освещении. Мощный аккумулятор на 5000 мАч с поддержкой быстрой зарядки на 120 Вт позволяет быстро восстановить заряд и использовать устройство без перебоев в течение всего дня. Поддержка 5G и NFC, а также стильный белый корпус делают этот смартфон отличным выбором для ценителей технологий и стиля.	134999.00	0.00	13	1	1	2024-10-10 16:30:56.023118+03	2024-11-20 17:41:36.50837+03	xiaomi-14-ultra
8	awf1241	Планшет Xiaomi Pad 6 Wi-Fi 256 ГБ черный	Xiaomi Pad 6 Wi-Fi 256 ГБ черный — это стильный и мощный планшет, который сочетает в себе отличные характеристики и современный дизайн. 11-дюймовый экран с разрешением 2560x1600 пикселей и частотой обновления 120 Гц обеспечивает яркое и четкое изображение, идеально подходящее для просмотра видео, игр и работы с графикой. Под капотом планшета скрывается процессор Snapdragon 870, который обеспечивает отличную производительность в многозадачности и играх. 256 ГБ встроенной памяти позволяют хранить огромное количество данных, включая фотографии, видео и приложения. Камера на 13 Мп и фронтальная камера на 8 Мп позволяют снимать качественные фотографии и проводить видеозвонки в отличном качестве. Планшет поддерживает Wi-Fi 6 для сверхбыстрого интернета и оснащен аккумулятором на 8600 мАч, который обеспечит длительное время работы без подзарядки. Это идеальный выбор для работы и развлечений в любом месте.	40999.00	12000.00	54	5	1	2024-10-10 16:30:56.023118+03	2024-10-10 16:30:56.023118+03	xiaomi-pad-6
13	SKU10005	6.8" Смартфон Samsung Galaxy S23 Ultra	Samsung Galaxy S23 Ultra — это флагманский смартфон с огромным 6.8-дюймовым Dynamic AMOLED 2X экраном, который предлагает невероятно четкое и яркое изображение с разрешением 3088x1440 пикселей и частотой обновления 120 Гц. Этот смартфон оснащен самым современным процессором Snapdragon 8 Gen 2, что гарантирует отличную производительность в любых приложениях и играх. Основная камера на 200 Мп позволяет создавать потрясающие фотографии с высокой детализацией, а система из четырех камер также включает телеобъектив с 10-кратным оптическим зумом, что позволяет делать великолепные снимки на дальнем расстоянии. Аккумулятор на 5000 мАч с поддержкой быстрой зарядки на 45 Вт обеспечит длительное использование устройства и быструю зарядку. Поддержка 5G и Wi-Fi 6E позволит вам всегда оставаться на связи, а встроенная поддержка S Pen добавит удобства в повседневном использовании. Samsung Galaxy S23 Ultra — это идеальный выбор для тех, кто ищет сочетание мощности, камеры и дизайна.	139990.00	19000.00	15	1	17	2024-10-15 23:07:08.948306+03	2024-10-15 23:07:08.948306+03	samsung-galaxy-s23-ultra
68	SKU10010	Honor Hunter V700	Honor Hunter V700 — это мощный игровой ноутбук, предназначенный для самых требовательных пользователей. С 16.1-дюймовым экраном Full HD и частотой обновления 144 Гц, он предлагает отличную плавность изображения и яркие, насыщенные цвета, что делает его идеальным выбором для геймеров. Оснащенный процессором Intel Core i7 11-го поколения и видеокартой NVIDIA GeForce RTX 3060, этот ноутбук справляется с любыми современными играми на высоких настройках. 16 ГБ оперативной памяти и SSD на 512 ГБ обеспечат молниеносную скорость загрузки и плавную работу системы. Honor Hunter V700 также оснащен системой охлаждения с тремя вентиляторами и несколькими теплотрубами, что позволяет поддерживать низкие температуры даже при долгих игровых сессиях. Время автономной работы до 8 часов делает его удобным для использования вне дома. Этот ноутбук сочетает в себе отличную производительность и стильный внешний вид, что делает его одним из лучших вариантов для геймеров.	80000.00	10000.00	10	4	18	\N	\N	honor-hunter-v700
66	SKU10008	Acer Swift 3	Acer Swift 3 — это стильный и легкий ультрабук, который сочетает в себе высокую производительность и компактные размеры. Оснащенный 14-дюймовым экраном Full HD с яркими цветами и тонкими рамками, он идеален для работы в дороге. Внутри устройства установлен процессор AMD Ryzen 7 5700U, который позволяет легко справляться с многозадачностью и запуском различных приложений. Ноутбук обладает 16 ГБ оперативной памяти и SSD на 1 ТБ, что обеспечит молниеносную скорость работы и быструю загрузку приложений и файлов. Время автономной работы до 12 часов делает Acer Swift 3 отличным выбором для людей, которые часто находятся в движении и ценят продолжительное время без подзарядки. Дополнительные преимущества включают поддержку Wi-Fi 6 и наличие портов для подключения множества внешних устройств.	65000.00	4000.00	8	4	19	0001-01-01 02:30:17+02:30:17	2024-11-20 19:39:03.451796+03	acer-swift-3
83	123	123	123	123.00	12.00	123	1	1	2024-11-20 21:37:29.02417+03	2024-11-20 22:03:38.293162+03	123-123
64	SKU10006	Acer Predator Helios 18	Acer Predator Helios 18 — это мощный игровой ноутбук с экраном 18 дюймов, предназначенный для самых требовательных пользователей. Он оснащен процессором Intel Core i9 13-го поколения и видеокартой NVIDIA GeForce RTX 4080, что обеспечивает невероятную производительность в самых современных играх и графических приложениях. Разрешение экрана 2560x1600 пикселей и частота обновления 165 Гц позволяют наслаждаться плавным и четким изображением, даже в самых динамичных сценах. Ноутбук также имеет 32 ГБ оперативной памяти и 1 ТБ SSD, что обеспечивает быстрый доступ к данным и эффективную работу с большими объемами информации. Встроенная система охлаждения с несколькими вентиляторами и тепловыми трубками гарантирует стабильную работу даже при длительных игровых сессиях. С максимальным набором портов и поддержкой Wi-Fi 6, этот ноутбук идеально подходит для профессионалов и геймеров.	450000.00	1500.00	10	4	19	0001-01-01 02:30:17+02:30:17	2024-11-20 22:03:14.057982+03	acer-predator-helios-18
71	SKU10013	Xiaomi Pad 5	Xiaomi Pad 5 — это стильный и мощный планшет с 11-дюймовым дисплеем 2.5K, который идеально подходит для мультимедийных развлечений и продуктивной работы. Экран с частотой обновления 120 Гц и высоким разрешением предлагает отличное качество изображения, а яркие и насыщенные цвета делают его идеальным для просмотра фильмов и игр. Внутри устройства установлен процессор Qualcomm Snapdragon 860, который обеспечивает отличную производительность и позволяет запускать любые приложения и игры без задержек. Xiaomi Pad 5 также оснащен 6 ГБ оперативной памяти и 128 ГБ встроенной памяти для хранения файлов. Планшет поддерживает быструю зарядку, и его батареи хватает на длительное использование. Помимо этого, Xiaomi Pad 5 поддерживает поддержку стилуса, что делает его удобным для рисования и работы с заметками. Это идеальный планшет для людей, которым нужно сочетание производительности, качества и доступной цены.	25000.00	2000.00	20	5	1	\N	\N	xiaomi-pad-5
76	SKU10018	Apple Watch Series 7	Apple Watch Series 7 — это последние умные часы от Apple, которые обладают лучшими характеристиками в своей линейке. Экран стал на 20% больше по сравнению с предыдущими моделями, а также более прочным благодаря стеклу, устойчивому к ударам. Часы оснащены чипом S7, который улучшает производительность и скорость отклика на действия пользователя. Apple Watch Series 7 включает множество функций для здоровья, таких как измерение уровня кислорода в крови, отслеживание сердечного ритма и анализ сна. Они также оснащены функцией ЭКГ, которая позволяет следить за сердечным здоровьем. С поддержкой постоянного экрана, улучшенной батареей и возможностью быстрой зарядки, эти часы идеально подходят для повседневного использования и интенсивных тренировок. Водонепроницаемость до 50 метров и возможность настройки множества циферблатов позволяют адаптировать часы под любой стиль жизни.	40000.00	3000.00	15	6	12	\N	\N	apple-watch-series-7
9	SKU10001	Смартфон Honor Magic6 Pro, 12/512, черный	Смартфон Honor Magic6 Pro в черном цвете — это мощное устройство для тех, кто ценит производительность, стиль и инновационные технологии. С 12 ГБ оперативной памяти и 512 ГБ встроенной памяти, он обеспечивает отличную многозадачность и достаточный объем для хранения всех данных. 6.76-дюймовый AMOLED дисплей с разрешением 3200x1440 пикселей и частотой обновления 120 Гц подарит вам яркое и четкое изображение, которое идеально подойдет как для игр, так и для мультимедийных приложений. В качестве процессора используется флагманский чипset, что гарантирует быструю работу устройства и отличную производительность. Основная камера с разрешением 50 Мп и дополнительными сенсорами позволяет делать фотографии высокого качества в любых условиях, а поддержка записи видео в 4K и AI-режимы помогут вам раскрыть потенциал камеры. Аккумулятор на 4600 мАч с быстрой зарядкой на 66 Вт обеспечит долгую работу смартфона и быструю зарядку. Honor Magic6 Pro — это идеальный выбор для тех, кто ищет сочетание мощности и функционала.	99990.00	12000.00	30	1	18	2024-10-15 23:07:08.948306+03	2024-10-15 23:07:08.948306+03	honor-magic5-pro
10	SKU10002	Смартфон Honor 90	Honor 90 — это современный и стильный смартфон, который предлагает отличное сочетание производительности и инновационных технологий. С 12 ГБ оперативной и 256 ГБ встроенной памяти он обеспечит быструю работу и достаточно места для хранения ваших данных. Смартфон оснащен 6.7-дюймовым AMOLED экраном с разрешением 2400x1080 пикселей и частотой обновления 120 Гц, что позволяет наслаждаться ярким и плавным изображением. Основная камера с 200 Мп сенсором делает потрясающие снимки с отличной детализацией, а фронтальная камера на 50 Мп позволяет делать яркие селфи. Honor 90 поддерживает быструю зарядку на 66 Вт, что позволяет быстро восстановить заряд и продолжить пользоваться устройством в любое время. Аккумулятор на 5000 мАч обеспечит долгую работу смартфона, а процессор Qualcomm Snapdragon 8+ Gen 1 гарантирует отличную производительность в играх и приложениях. Это идеальный выбор для любителей технологий и стильных устройств.	49990.00	3000.00	40	1	18	2024-10-15 23:07:08.948306+03	2024-11-20 17:47:33.625567+03	honor-90
69	SKU10011	Samsung Galaxy Book 3	Samsung Galaxy Book 3 — это ультратонкий и легкий ноутбук с экраном 15.6 дюймов Full HD, который подходит как для работы, так и для развлечений. Он оснащен процессором Intel Core i5 12-го поколения и 8 ГБ оперативной памяти, что обеспечивает отличную производительность для офисных задач и мультимедиа. Благодаря SSD на 512 ГБ, Galaxy Book 3 быстро загружает операционную систему и приложения, а также предлагает много места для хранения данных. Элегантный металлический корпус и длительное время работы от аккумулятора (до 12 часов) делают этот ноутбук идеальным для тех, кто часто находится в движении. Galaxy Book 3 также поддерживает технологию быстрой зарядки, что позволяет быстро восполнить заряд устройства. Это универсальный ноутбук с хорошим балансом между производительностью и мобильностью.	50000.00	4000.00	5	4	17	\N	\N	samsung-galaxy-book-3
70	SKU10012	Samsung Galaxy Book 4 Pro	Samsung Galaxy Book 4 Pro — это премиум-ноутбук, предназначенный для пользователей, которым необходимы высокие производительность и мобильность. С 14-дюймовым экраном AMOLED, поддерживающим разрешение Full HD, он идеально подходит для работы с графикой, просмотра видео и работы в интернете. Внутри устройства установлен процессор Intel Core i7 12-го поколения и 16 ГБ оперативной памяти, что гарантирует отличную производительность при работе с несколькими приложениями и файлами. SSD на 512 ГБ обеспечит молниеносную скорость работы и быстрый доступ к данным. Ноутбук также оснащен батареей с длительным временем работы и функцией быстрой зарядки, что позволяет пользоваться устройством на протяжении всего дня. С тонким и стильным металлическим корпусом, Galaxy Book 4 Pro сочетает в себе все необходимые качества для профессионалов и студентов, работающих в любых условиях.	70000.00	5000.00	7	4	17	\N	\N	samsung-galaxybook-4-pro
72	SKU10014	Samsung Galaxy Tab S7+	Samsung Galaxy Tab S7+ — это премиум-планшет с 12.4-дюймовым AMOLED экраном, который выводит качество изображения на новый уровень. С разрешением 2800x1752 пикселей и частотой обновления 120 Гц, этот планшет идеально подходит для просмотра видео, игры и работы с графикой. Оснащенный процессором Qualcomm Snapdragon 865+ и 6 ГБ оперативной памяти, Samsung Galaxy Tab S7+ обеспечивает отличную производительность даже при запуске самых требовательных приложений. Планшет поддерживает S Pen, который входит в комплект и позволяет рисовать и записывать заметки с высокой точностью. Батарея на 10 090 мАч обеспечит долгую работу без подзарядки. Благодаря поддержке 5G, Wi-Fi 6 и многозадачности, Galaxy Tab S7+ является отличным выбором для профессионалов и пользователей, которым нужен мощный и универсальный планшет для работы и развлечений.	40000.00	2500.00	15	5	17	\N	\N	samsung-galaxy-tab-s7-plus
73	SKU10015	Apple iPad Air	Apple iPad Air — это универсальный планшет с 10.9-дюймовым экраном Liquid Retina, который поддерживает цветовую гамму P3 и True Tone, что делает его идеальным для творчества и просмотра мультимедиа. Оснащенный процессором A14 Bionic, iPad Air предлагает отличную производительность и плавность работы, даже при запуске самых требовательных приложений. Планшет поддерживает Apple Pencil (2-го поколения) и Magic Keyboard, что позволяет использовать его как для работы, так и для рисования и создания контента. iPad Air имеет 64 ГБ или 256 ГБ памяти, а также 10-часовой срок службы батареи, что делает его идеальным выбором для длительных рабочих сессий. Он также поддерживает Wi-Fi 6 и опционально сотовую связь, что позволяет оставаться на связи в любой ситуации. Этот планшет — отличный выбор для профессионалов, студентов и пользователей, которые ценят качество и мобильность.	50000.00	3000.00	10	5	12	\N	\N	apple-ipad-air
74	SKU10016	Samsung Galaxy Watch 4	Samsung Galaxy Watch 4 — это умные часы, которые предлагают не только стильный и элегантный дизайн, но и множество функций для здоровья и фитнеса. Оснащенные AMOLED экраном с яркими цветами и четкостью, они идеально подойдут для повседневного использования и спорта. Встроенные датчики измеряют уровень кислорода в крови, пульс, а также отслеживают уровень стресса и сон. Galaxy Watch 4 оснащены процессором Exynos W920, что обеспечивает отличную производительность и плавность работы. Часы работают на новой операционной системе Wear OS, что дает доступ к множеству приложений через Google Play. Они поддерживают функцию мониторинга активности, умный помощник Bixby и уведомления с вашего смартфона. С водонепроницаемостью IP68 и батареей, которая может работать до 40 часов, Galaxy Watch 4 — это идеальный выбор для людей, следящих за своим здоровьем и активностью, а также для тех, кто хочет иметь все функции умных часов в одном устройстве.	25000.00	2000.00	25	6	17	\N	\N	samsung-galaxy-watch-4
75	SKU10017	Samsung Galaxy Watch Active 2	Samsung Galaxy Watch Active 2 — это компактные и легкие умные часы с сенсорным экраном Super AMOLED, которые идеально подходят для тех, кто ведет активный образ жизни. Они оснащены функциями отслеживания сердечного ритма, стресса, сна и уровня кислорода в крови, что позволяет следить за состоянием здоровья в реальном времени. В часах имеется встроенный GPS для точного отслеживания маршрутов, а также различные режимы тренировок для более эффективных тренировок. Galaxy Watch Active 2 могут синхронизироваться с вашим смартфоном для получения уведомлений, звонков и сообщений прямо на экран. Часы обладают водонепроницаемостью до 50 метров, что делает их идеальными для плавания. С батареей, которая работает до 2 суток, и обновленной операционной системой Tizen, они обеспечивают надежную и длительную работу, а также стильный и современный дизайн, который подойдет для любой ситуации.	22000.00	1500.00	20	6	17	\N	\N	samsung-galaxy-watch-active-2
77	SKU10019	Honor Magic Watch 2	Honor Magic Watch 2 — это умные часы с элегантным дизайном и мощной функциональностью для тех, кто ценит как стиль, так и здоровье. Часы оснащены AMOLED экраном для ярких и четких изображений, а также множества спортивных и фитнес-режимов для отслеживания различных тренировок. Встроенный датчик сердечного ритма, мониторинг уровня кислорода в крови и отслеживание качества сна позволяют пользователю контролировать свое здоровье и физическую форму на протяжении дня. Honor Magic Watch 2 может работать без подзарядки до 14 дней благодаря мощной батарее, а водонепроницаемость до 50 метров делает их идеальными для плавания. Эти часы также поддерживают уведомления с вашего смартфона, а встроенный GPS точно отслеживает маршруты во время прогулок или пробежек. Honor Magic Watch 2 — это отличный выбор для людей, которые хотят следить за своим здоровьем и оставаться на связи с миром, не снимая часы.	3000.00	1000.00	30	6	18	\N	\N	honor-magic-watch-2
78	SKU10020	Honor Watch GS Pro	Honor Watch GS Pro — это умные часы, предназначенные для любителей активного отдыха и экстремальных условий. Вдохновленные концепцией "Explorer", эти часы имеют сверхпрочную конструкцию с военным стандартом MIL-STD-810G, что делает их идеальными для путешествий и приключений. Часы оснащены большим 1.39-дюймовым AMOLED экраном, который отлично читается при ярком солнечном свете. Honor Watch GS Pro поддерживает более 100 различных спортивных режимов и может отслеживать множество показателей здоровья, включая пульс, уровень кислорода в крови и качество сна. Встроенный GPS и поддержка навигации позволяют уверенно ориентироваться в любых условиях, а батарея, которая работает до 25 дней на одном заряде, делает часы идеальными для длительных путешествий. Водонепроницаемость до 50 метров и наличие множества полезных функций для активного образа жизни делают Honor Watch GS Pro отличным выбором для людей, которые любят проводить время на свежем воздухе.	30000.00	1500.00	25	6	18	\N	\N	honor-watch-gs-pro
67	SKU10009	Honor MagicBook 14	Honor MagicBook 14 — это тонкий и легкий ноутбук с экраном 14 дюймов Full HD, который идеально подходит для работы, учебы и развлечений. Благодаря процессору AMD Ryzen 5 5500U и видеокарте AMD Radeon, он предоставляет отличную производительность в многозадачных приложениях, а также обеспечит плавную работу в популярных офисных и графических программах. 16 ГБ оперативной памяти и 512 ГБ SSD обеспечат быструю загрузку приложений и файлов, а батарея с продолжительностью работы до 10 часов позволит вам работать без подзарядки в течение всего дня. Honor MagicBook 14 оснащен современной системой охлаждения, которая поддерживает оптимальную температуру устройства даже при длительных нагрузках. Этот ноутбук — отличный выбор для пользователей, которым нужно сочетание мощности и портативности.	40000.00	2000.00	12	4	18	0001-01-01 02:30:17+02:30:17	2024-11-20 19:37:24.221402+03	honor-magicbook-14
65	SKU10007	Acer Aspire 5	Acer Aspire 5 — это универсальный ноутбук, предназначенный для работы и развлечений. Он оснащен 15.6-дюймовым экраном с разрешением Full HD (1920x1080) и антибликовым покрытием, что позволяет работать с комфортом даже при ярком освещении. Ноутбук оснащен процессором Intel Core i5 12-го поколения и интегрированной графикой Intel Iris Xe, что делает его отличным выбором для повседневных задач, включая работу с офисными приложениями, серфинг в интернете и просмотр видео. 8 ГБ оперативной памяти и SSD на 512 ГБ обеспечивают быструю работу системы и быстрый доступ к файлам. Батарея с длительным временем работы позволяет оставаться продуктивным в течение всего дня без необходимости частой подзарядки. Acer Aspire 5 — это идеальное решение для студентов и профессионалов, которые ищут баланс между производительностью, удобством и ценой.	45000.00	3000.00	15	4	19	0001-01-01 02:30:17+02:30:17	2024-11-20 19:38:00.218323+03	acer-aspire-5
\.


--
-- Data for Name: review_images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.review_images (id, review_id, image_url) FROM stdin;
14	19	review_eee0c67b267a463380c2dd72f42de663.jpg
15	20	review_42ac05e954e94d3cb917e716c5962bcf.png
16	21	review_583f56533fe8479f9f7faa5a697bc81b.png
17	22	review_6f229a96708743c9a6a5298201c20e47.png
\.


--
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.reviews (id, user_id, product_id, rating, created_at, pros, cons, comment, is_moder) FROM stdin;
19	4	1	2	2024-11-11 00:04:48.798987+03	фц	пц	фцп	t
20	4	1	2	2024-11-11 00:05:04.524457+03	пап	цп	цп	t
21	4	11	1	2024-11-20 13:17:00.938592+03	1	1	1	t
22	4	13	5	2024-11-20 17:53:46.30949+03	крутой телефон	нет	пуыпуыупыпу	t
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password_hash, first_name, last_name, role, created_at, updated_at) FROM stdin;
7	yawaihv10@gmail.com	$2a$10$2P5o5TcQY6q.a1hAPK8vb.51uEx8H0TYZkAKPt66FjWt615o6z0dm	Nikita	Swrawr	customer	2024-11-15 23:26:21.340869+03	2024-11-15 23:26:29.301772+03
4	yawaihv4@gmail.com	$2a$10$d3Tme.gAZbN.JUHl5k1q6OJRVe0r553FJjtkW5GWc9vH.0vJlNGjO	Dimka	Kruck	admin	2024-11-03 17:26:41.181195+03	2024-11-05 15:28:09.758159+03
\.


--
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.addresses_id_seq', 6, true);


--
-- Name: brand_banners_brand_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brand_banners_brand_id_seq', 1, false);


--
-- Name: brand_banners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brand_banners_id_seq', 18, true);


--
-- Name: brands_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brands_id_seq', 19, true);


--
-- Name: cart_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cart_items_id_seq', 36, true);


--
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.carts_id_seq', 1, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 8, true);


--
-- Name: laptop_filters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.laptop_filters_id_seq', 11, true);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_items_id_seq', 29, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 6, true);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 5, true);


--
-- Name: product_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_details_id_seq', 44, true);


--
-- Name: product_filters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_filters_id_seq', 14, true);


--
-- Name: product_images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_images_id_seq', 153, true);


--
-- Name: product_variants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_variants_id_seq', 1, false);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 83, true);


--
-- Name: review_images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.review_images_id_seq', 17, true);


--
-- Name: review_images_review_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.review_images_review_id_seq', 1, false);


--
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reviews_id_seq', 22, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 7, true);


--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- Name: brand_banners brand_banners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand_banners
    ADD CONSTRAINT brand_banners_pkey PRIMARY KEY (id, brand_id);


--
-- Name: brands brands_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);


--
-- Name: cart_items cart_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_pkey PRIMARY KEY (id);


--
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: laptop_filters laptop_filters_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.laptop_filters
    ADD CONSTRAINT laptop_filters_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: product_details product_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_pkey PRIMARY KEY (id);


--
-- Name: product_filters product_filters_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_filters
    ADD CONSTRAINT product_filters_pkey PRIMARY KEY (id);


--
-- Name: product_images product_images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT product_images_pkey PRIMARY KEY (id);


--
-- Name: product_variants product_variants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_variants
    ADD CONSTRAINT product_variants_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: review_images review_images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review_images
    ADD CONSTRAINT review_images_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_products_sku; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_products_sku ON public.products USING btree (sku);


--
-- Name: idx_review_images_review_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_review_images_review_id ON public.review_images USING btree (review_id);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: trgm_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX trgm_idx ON public.products USING gin (name public.gin_trgm_ops);


--
-- Name: trgm_idx_brand; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX trgm_idx_brand ON public.brands USING gin (name public.gin_trgm_ops);


--
-- Name: trgm_idx_category; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX trgm_idx_category ON public.categories USING gin (name public.gin_trgm_ops);


--
-- Name: trgm_idx_title; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX trgm_idx_title ON public.categories USING gin (title public.gin_trgm_ops);


--
-- Name: brand_banners fk_brands_banners; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand_banners
    ADD CONSTRAINT fk_brands_banners FOREIGN KEY (brand_id) REFERENCES public.brands(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: cart_items fk_carts_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT fk_carts_items FOREIGN KEY (cart_id) REFERENCES public.carts(id);


--
-- Name: categories fk_categories_sub_categories; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT fk_categories_sub_categories FOREIGN KEY (parent_category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: order_items fk_order_items_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_products FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: orders fk_orders_address; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_orders_address FOREIGN KEY (address_id) REFERENCES public.addresses(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: order_items fk_orders_order_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_orders_order_items FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: payments fk_orders_payment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT fk_orders_payment FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: products fk_products_brand; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_brand FOREIGN KEY (brand_id) REFERENCES public.brands(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: products fk_products_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: product_details fk_products_details; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT fk_products_details FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: product_filters fk_products_filters; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_filters
    ADD CONSTRAINT fk_products_filters FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: product_images fk_products_images; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT fk_products_images FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: reviews fk_products_reviews; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_products_reviews FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: product_variants fk_products_variants; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_variants
    ADD CONSTRAINT fk_products_variants FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: review_images fk_reviews_images; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.review_images
    ADD CONSTRAINT fk_reviews_images FOREIGN KEY (review_id) REFERENCES public.reviews(id);


--
-- Name: addresses fk_users_addresses; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT fk_users_addresses FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: carts fk_users_cart_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT fk_users_cart_items FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: orders fk_users_orders; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_users_orders FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: reviews fk_users_reviews; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_users_reviews FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

