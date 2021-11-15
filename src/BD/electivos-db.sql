PGDMP                     
    y           administracion_electivos    13.2    13.2 9    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    57488    administracion_electivos    DATABASE     u   CREATE DATABASE administracion_electivos WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Spanish_Mexico.1252';
 (   DROP DATABASE administracion_electivos;
                postgres    false            �            1259    57491    administrador    TABLE     z   CREATE TABLE public.administrador (
    id integer NOT NULL,
    correo text,
    rut text NOT NULL,
    password text
);
 !   DROP TABLE public.administrador;
       public         heap    postgres    false            �            1259    57489    Administrador_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Administrador_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public."Administrador_id_seq";
       public          postgres    false    201            �           0    0    Administrador_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public."Administrador_id_seq" OWNED BY public.administrador.id;
          public          postgres    false    200            �            1259    57573    alumno    TABLE     �   CREATE TABLE public.alumno (
    id integer NOT NULL,
    rut text NOT NULL,
    nombre text,
    correo text,
    id_carrera integer NOT NULL,
    semestre_incompleto integer,
    cantidad_ramos integer
);
    DROP TABLE public.alumno;
       public         heap    postgres    false            �            1259    57571    Alumno_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Alumno_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public."Alumno_id_seq";
       public          postgres    false    210            �           0    0    Alumno_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public."Alumno_id_seq" OWNED BY public.alumno.id;
          public          postgres    false    209            �            1259    57562    carrera    TABLE     J   CREATE TABLE public.carrera (
    id integer NOT NULL,
    nombre text
);
    DROP TABLE public.carrera;
       public         heap    postgres    false            �            1259    57560    Carrera_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Carrera_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public."Carrera_id_seq";
       public          postgres    false    208            �           0    0    Carrera_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public."Carrera_id_seq" OWNED BY public.carrera.id;
          public          postgres    false    207            �            1259    57502    profesor    TABLE     \   CREATE TABLE public.profesor (
    id integer NOT NULL,
    nombre text,
    correo text
);
    DROP TABLE public.profesor;
       public         heap    postgres    false            �            1259    57500    Profesor_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Profesor_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."Profesor_id_seq";
       public          postgres    false    203            �           0    0    Profesor_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public."Profesor_id_seq" OWNED BY public.profesor.id;
          public          postgres    false    202            �            1259    57513    electivo    TABLE     t   CREATE TABLE public.electivo (
    id integer NOT NULL,
    nombre text,
    "NRC" text,
    id_profesor integer
);
    DROP TABLE public.electivo;
       public         heap    postgres    false            �            1259    57511    electivo_id_seq    SEQUENCE     �   CREATE SEQUENCE public.electivo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.electivo_id_seq;
       public          postgres    false    205            �           0    0    electivo_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.electivo_id_seq OWNED BY public.electivo.id;
          public          postgres    false    204            �            1259    57527    postulacion    TABLE     �   CREATE TABLE public.postulacion (
    id integer NOT NULL,
    cantidad integer,
    id_electivo integer,
    aprobado boolean
);
    DROP TABLE public.postulacion;
       public         heap    postgres    false            �            1259    57695 	   solicitud    TABLE     �   CREATE TABLE public.solicitud (
    id integer NOT NULL,
    id_alumno integer NOT NULL,
    id_postulacion_1 integer NOT NULL,
    id_postulacion_2 integer,
    id_postulacion_3 integer
);
    DROP TABLE public.solicitud;
       public         heap    postgres    false            �            1259    57693    solicitud_id_seq    SEQUENCE     �   CREATE SEQUENCE public.solicitud_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.solicitud_id_seq;
       public          postgres    false    212                        0    0    solicitud_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.solicitud_id_seq OWNED BY public.solicitud.id;
          public          postgres    false    211            I           2604    57494    administrador id    DEFAULT     v   ALTER TABLE ONLY public.administrador ALTER COLUMN id SET DEFAULT nextval('public."Administrador_id_seq"'::regclass);
 ?   ALTER TABLE public.administrador ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    200    201    201            M           2604    57576 	   alumno id    DEFAULT     h   ALTER TABLE ONLY public.alumno ALTER COLUMN id SET DEFAULT nextval('public."Alumno_id_seq"'::regclass);
 8   ALTER TABLE public.alumno ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    210    209    210            L           2604    57565 
   carrera id    DEFAULT     j   ALTER TABLE ONLY public.carrera ALTER COLUMN id SET DEFAULT nextval('public."Carrera_id_seq"'::regclass);
 9   ALTER TABLE public.carrera ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    207    208    208            K           2604    57516    electivo id    DEFAULT     j   ALTER TABLE ONLY public.electivo ALTER COLUMN id SET DEFAULT nextval('public.electivo_id_seq'::regclass);
 :   ALTER TABLE public.electivo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    204    205    205            J           2604    57505    profesor id    DEFAULT     l   ALTER TABLE ONLY public.profesor ALTER COLUMN id SET DEFAULT nextval('public."Profesor_id_seq"'::regclass);
 :   ALTER TABLE public.profesor ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202    203            N           2604    57698    solicitud id    DEFAULT     l   ALTER TABLE ONLY public.solicitud ALTER COLUMN id SET DEFAULT nextval('public.solicitud_id_seq'::regclass);
 ;   ALTER TABLE public.solicitud ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    211    212    212            �          0    57491    administrador 
   TABLE DATA           B   COPY public.administrador (id, correo, rut, password) FROM stdin;
    public          postgres    false    201   �?       �          0    57573    alumno 
   TABLE DATA           j   COPY public.alumno (id, rut, nombre, correo, id_carrera, semestre_incompleto, cantidad_ramos) FROM stdin;
    public          postgres    false    210   �?       �          0    57562    carrera 
   TABLE DATA           -   COPY public.carrera (id, nombre) FROM stdin;
    public          postgres    false    208   Y@       �          0    57513    electivo 
   TABLE DATA           B   COPY public.electivo (id, nombre, "NRC", id_profesor) FROM stdin;
    public          postgres    false    205   �@       �          0    57527    postulacion 
   TABLE DATA           J   COPY public.postulacion (id, cantidad, id_electivo, aprobado) FROM stdin;
    public          postgres    false    206   �@       �          0    57502    profesor 
   TABLE DATA           6   COPY public.profesor (id, nombre, correo) FROM stdin;
    public          postgres    false    203   �@       �          0    57695 	   solicitud 
   TABLE DATA           h   COPY public.solicitud (id, id_alumno, id_postulacion_1, id_postulacion_2, id_postulacion_3) FROM stdin;
    public          postgres    false    212   EA                  0    0    Administrador_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public."Administrador_id_seq"', 1, false);
          public          postgres    false    200                       0    0    Alumno_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public."Alumno_id_seq"', 1, true);
          public          postgres    false    209                       0    0    Carrera_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public."Carrera_id_seq"', 1, false);
          public          postgres    false    207                       0    0    Profesor_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public."Profesor_id_seq"', 1, false);
          public          postgres    false    202                       0    0    electivo_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.electivo_id_seq', 1, false);
          public          postgres    false    204                       0    0    solicitud_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.solicitud_id_seq', 1, false);
          public          postgres    false    211            P           2606    57499     administrador Administrador_pkey 
   CONSTRAINT     e   ALTER TABLE ONLY public.administrador
    ADD CONSTRAINT "Administrador_pkey" PRIMARY KEY (id, rut);
 L   ALTER TABLE ONLY public.administrador DROP CONSTRAINT "Administrador_pkey";
       public            postgres    false    201    201            Z           2606    57581    alumno Alumno_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_pkey" PRIMARY KEY (id, rut);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_pkey";
       public            postgres    false    210    210            X           2606    57570    carrera Carrera_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.carrera
    ADD CONSTRAINT "Carrera_pkey" PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.carrera DROP CONSTRAINT "Carrera_pkey";
       public            postgres    false    208            R           2606    57510    profesor Profesor_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.profesor
    ADD CONSTRAINT "Profesor_pkey" PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.profesor DROP CONSTRAINT "Profesor_pkey";
       public            postgres    false    203            \           2606    57692    alumno alumno_id_key 
   CONSTRAINT     M   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT alumno_id_key UNIQUE (id);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT alumno_id_key;
       public            postgres    false    210            T           2606    57521    electivo electivo_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_pkey;
       public            postgres    false    205            V           2606    57531    postulacion postulacion_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_pkey;
       public            postgres    false    206            ^           2606    57700    solicitud solicitud_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_pkey;
       public            postgres    false    212            a           2606    57582    alumno Alumno_id_carrera_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_id_carrera_fkey" FOREIGN KEY (id_carrera) REFERENCES public.carrera(id);
 I   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_id_carrera_fkey";
       public          postgres    false    208    210    2904            _           2606    57522 "   electivo electivo_id_profesor_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_id_profesor_fkey FOREIGN KEY (id_profesor) REFERENCES public.profesor(id);
 L   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_id_profesor_fkey;
       public          postgres    false    203    205    2898            `           2606    57532 (   postulacion postulacion_id_electivo_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_id_electivo_fkey FOREIGN KEY (id_electivo) REFERENCES public.electivo(id);
 R   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_id_electivo_fkey;
       public          postgres    false    206    205    2900            b           2606    57701 "   solicitud solicitud_id_alumno_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 L   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_alumno_fkey;
       public          postgres    false    210    212    2908            c           2606    57706 )   solicitud solicitud_id_postulacion_1_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_1_fkey FOREIGN KEY (id_postulacion_1) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_1_fkey;
       public          postgres    false    2902    212    206            d           2606    57711 )   solicitud solicitud_id_postulacion_2_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_2_fkey FOREIGN KEY (id_postulacion_2) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_2_fkey;
       public          postgres    false    212    206    2902            e           2606    57716 )   solicitud solicitud_id_postulacion_3_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_3_fkey FOREIGN KEY (id_postulacion_3) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_3_fkey;
       public          postgres    false    2902    212    206            �      x������ � �      �   \   x�3�420263273�t����,��Wp�I�J�K)�W���,K,J-Vp,.)M�����L���ˇ�9$�����&��%�pr��W� �11      �      x�3��tv��2����2r<�b���� 9�K      �   (   x�3�t5�44b.#NW# ���9]��LcN#�=... l�c      �   "   x�346��4�4.C#(���� ʌ���� x7�      �   D   x�3��M,J�/Vp�H�,��MQ��yz�9\F�ΉE9@I�ļ�ļ|�
 �I�$2�JP�b���� c��      �      x��02�4�42�0@\1z\\\ 9;     