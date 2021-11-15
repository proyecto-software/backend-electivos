PGDMP     .                
    y           administracion_electivos    13.3    13.3 I               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    18507    administracion_electivos    DATABASE     t   CREATE DATABASE administracion_electivos WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Spanish_Spain.1252';
 (   DROP DATABASE administracion_electivos;
                postgres    false            �            1259    18508    administrador    TABLE     z   CREATE TABLE public.administrador (
    id integer NOT NULL,
    correo text,
    rut text NOT NULL,
    password text
);
 !   DROP TABLE public.administrador;
       public         heap    postgres    false            �            1259    18514    Administrador_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Administrador_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public."Administrador_id_seq";
       public          postgres    false    200                       0    0    Administrador_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public."Administrador_id_seq" OWNED BY public.administrador.id;
          public          postgres    false    201            �            1259    18516    alumno    TABLE     �   CREATE TABLE public.alumno (
    id integer NOT NULL,
    rut text NOT NULL,
    nombre text,
    correo text,
    id_carrera integer NOT NULL,
    semestre_incompleto integer,
    cantidad_ramos integer
);
    DROP TABLE public.alumno;
       public         heap    postgres    false            �            1259    18522    Alumno_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Alumno_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public."Alumno_id_seq";
       public          postgres    false    202                       0    0    Alumno_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public."Alumno_id_seq" OWNED BY public.alumno.id;
          public          postgres    false    203            �            1259    18524    carrera    TABLE     J   CREATE TABLE public.carrera (
    id integer NOT NULL,
    nombre text
);
    DROP TABLE public.carrera;
       public         heap    postgres    false            �            1259    18530    Carrera_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Carrera_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public."Carrera_id_seq";
       public          postgres    false    204                       0    0    Carrera_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public."Carrera_id_seq" OWNED BY public.carrera.id;
          public          postgres    false    205            �            1259    18532    profesor    TABLE     \   CREATE TABLE public.profesor (
    id integer NOT NULL,
    nombre text,
    correo text
);
    DROP TABLE public.profesor;
       public         heap    postgres    false            �            1259    18538    Profesor_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Profesor_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."Profesor_id_seq";
       public          postgres    false    206                       0    0    Profesor_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public."Profesor_id_seq" OWNED BY public.profesor.id;
          public          postgres    false    207            �            1259    18540    electivo    TABLE     t   CREATE TABLE public.electivo (
    id integer NOT NULL,
    nombre text,
    "NRC" text,
    id_profesor integer
);
    DROP TABLE public.electivo;
       public         heap    postgres    false            �            1259    18546    electivo_id_seq    SEQUENCE     �   CREATE SEQUENCE public.electivo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.electivo_id_seq;
       public          postgres    false    208                       0    0    electivo_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.electivo_id_seq OWNED BY public.electivo.id;
          public          postgres    false    209            �            1259    18548    postulacion    TABLE     �   CREATE TABLE public.postulacion (
    id integer NOT NULL,
    cantidad integer,
    id_electivo integer,
    aprobado boolean
);
    DROP TABLE public.postulacion;
       public         heap    postgres    false            �            1259    18615    ramo    TABLE     k   CREATE TABLE public.ramo (
    id integer NOT NULL,
    nrc text,
    nombre text,
    creditos integer
);
    DROP TABLE public.ramo;
       public         heap    postgres    false            �            1259    18628    ramo_alumno    TABLE     �   CREATE TABLE public.ramo_alumno (
    id_ramo integer NOT NULL,
    id_alumno integer NOT NULL,
    nota double precision,
    oportunidad integer
);
    DROP TABLE public.ramo_alumno;
       public         heap    postgres    false            �            1259    18626    ramo_alumno_id_ramo_seq    SEQUENCE     �   CREATE SEQUENCE public.ramo_alumno_id_ramo_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.ramo_alumno_id_ramo_seq;
       public          postgres    false    216                       0    0    ramo_alumno_id_ramo_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.ramo_alumno_id_ramo_seq OWNED BY public.ramo_alumno.id_ramo;
          public          postgres    false    215            �            1259    18613    ramo_id_seq    SEQUENCE     �   CREATE SEQUENCE public.ramo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.ramo_id_seq;
       public          postgres    false    214                       0    0    ramo_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.ramo_id_seq OWNED BY public.ramo.id;
          public          postgres    false    213            �            1259    18551 	   solicitud    TABLE     �   CREATE TABLE public.solicitud (
    id integer NOT NULL,
    id_alumno integer NOT NULL,
    id_postulacion_1 integer NOT NULL,
    id_postulacion_2 integer,
    id_postulacion_3 integer
);
    DROP TABLE public.solicitud;
       public         heap    postgres    false            �            1259    18554    solicitud_id_seq    SEQUENCE     �   CREATE SEQUENCE public.solicitud_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.solicitud_id_seq;
       public          postgres    false    211                       0    0    solicitud_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.solicitud_id_seq OWNED BY public.solicitud.id;
          public          postgres    false    212            V           2604    18556    administrador id    DEFAULT     v   ALTER TABLE ONLY public.administrador ALTER COLUMN id SET DEFAULT nextval('public."Administrador_id_seq"'::regclass);
 ?   ALTER TABLE public.administrador ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    201    200            W           2604    18557 	   alumno id    DEFAULT     h   ALTER TABLE ONLY public.alumno ALTER COLUMN id SET DEFAULT nextval('public."Alumno_id_seq"'::regclass);
 8   ALTER TABLE public.alumno ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202            X           2604    18558 
   carrera id    DEFAULT     j   ALTER TABLE ONLY public.carrera ALTER COLUMN id SET DEFAULT nextval('public."Carrera_id_seq"'::regclass);
 9   ALTER TABLE public.carrera ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    205    204            Z           2604    18559    electivo id    DEFAULT     j   ALTER TABLE ONLY public.electivo ALTER COLUMN id SET DEFAULT nextval('public.electivo_id_seq'::regclass);
 :   ALTER TABLE public.electivo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    209    208            Y           2604    18560    profesor id    DEFAULT     l   ALTER TABLE ONLY public.profesor ALTER COLUMN id SET DEFAULT nextval('public."Profesor_id_seq"'::regclass);
 :   ALTER TABLE public.profesor ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    207    206            \           2604    18618    ramo id    DEFAULT     b   ALTER TABLE ONLY public.ramo ALTER COLUMN id SET DEFAULT nextval('public.ramo_id_seq'::regclass);
 6   ALTER TABLE public.ramo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    213    214            ]           2604    18631    ramo_alumno id_ramo    DEFAULT     z   ALTER TABLE ONLY public.ramo_alumno ALTER COLUMN id_ramo SET DEFAULT nextval('public.ramo_alumno_id_ramo_seq'::regclass);
 B   ALTER TABLE public.ramo_alumno ALTER COLUMN id_ramo DROP DEFAULT;
       public          postgres    false    215    216    216            [           2604    18561    solicitud id    DEFAULT     l   ALTER TABLE ONLY public.solicitud ALTER COLUMN id SET DEFAULT nextval('public.solicitud_id_seq'::regclass);
 ;   ALTER TABLE public.solicitud ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    212    211            �          0    18508    administrador 
   TABLE DATA           B   COPY public.administrador (id, correo, rut, password) FROM stdin;
    public          postgres    false    200   �Q       �          0    18516    alumno 
   TABLE DATA           j   COPY public.alumno (id, rut, nombre, correo, id_carrera, semestre_incompleto, cantidad_ramos) FROM stdin;
    public          postgres    false    202   �Q                 0    18524    carrera 
   TABLE DATA           -   COPY public.carrera (id, nombre) FROM stdin;
    public          postgres    false    204   +R                 0    18540    electivo 
   TABLE DATA           B   COPY public.electivo (id, nombre, "NRC", id_profesor) FROM stdin;
    public          postgres    false    208   OR                 0    18548    postulacion 
   TABLE DATA           J   COPY public.postulacion (id, cantidad, id_electivo, aprobado) FROM stdin;
    public          postgres    false    210   lR                 0    18532    profesor 
   TABLE DATA           6   COPY public.profesor (id, nombre, correo) FROM stdin;
    public          postgres    false    206   �R                 0    18615    ramo 
   TABLE DATA           9   COPY public.ramo (id, nrc, nombre, creditos) FROM stdin;
    public          postgres    false    214   �R                 0    18628    ramo_alumno 
   TABLE DATA           L   COPY public.ramo_alumno (id_ramo, id_alumno, nota, oportunidad) FROM stdin;
    public          postgres    false    216   �R                 0    18551 	   solicitud 
   TABLE DATA           h   COPY public.solicitud (id, id_alumno, id_postulacion_1, id_postulacion_2, id_postulacion_3) FROM stdin;
    public          postgres    false    211   �R                  0    0    Administrador_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public."Administrador_id_seq"', 1, false);
          public          postgres    false    201                       0    0    Alumno_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public."Alumno_id_seq"', 1, true);
          public          postgres    false    203                       0    0    Carrera_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public."Carrera_id_seq"', 1, false);
          public          postgres    false    205                       0    0    Profesor_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public."Profesor_id_seq"', 1, false);
          public          postgres    false    207                        0    0    electivo_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.electivo_id_seq', 1, false);
          public          postgres    false    209            !           0    0    ramo_alumno_id_ramo_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.ramo_alumno_id_ramo_seq', 1, false);
          public          postgres    false    215            "           0    0    ramo_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.ramo_id_seq', 1, false);
          public          postgres    false    213            #           0    0    solicitud_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.solicitud_id_seq', 1, false);
          public          postgres    false    212            _           2606    18563     administrador Administrador_pkey 
   CONSTRAINT     e   ALTER TABLE ONLY public.administrador
    ADD CONSTRAINT "Administrador_pkey" PRIMARY KEY (id, rut);
 L   ALTER TABLE ONLY public.administrador DROP CONSTRAINT "Administrador_pkey";
       public            postgres    false    200    200            a           2606    18565    alumno Alumno_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_pkey" PRIMARY KEY (id, rut);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_pkey";
       public            postgres    false    202    202            e           2606    18567    carrera Carrera_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.carrera
    ADD CONSTRAINT "Carrera_pkey" PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.carrera DROP CONSTRAINT "Carrera_pkey";
       public            postgres    false    204            g           2606    18569    profesor Profesor_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.profesor
    ADD CONSTRAINT "Profesor_pkey" PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.profesor DROP CONSTRAINT "Profesor_pkey";
       public            postgres    false    206            c           2606    18571    alumno alumno_id_key 
   CONSTRAINT     M   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT alumno_id_key UNIQUE (id);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT alumno_id_key;
       public            postgres    false    202            i           2606    18573    electivo electivo_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_pkey;
       public            postgres    false    208            k           2606    18575    postulacion postulacion_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_pkey;
       public            postgres    false    210            q           2606    18633    ramo_alumno ramo_alumno_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_pkey PRIMARY KEY (id_ramo, id_alumno);
 F   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_pkey;
       public            postgres    false    216    216            o           2606    18623    ramo ramo_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.ramo
    ADD CONSTRAINT ramo_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.ramo DROP CONSTRAINT ramo_pkey;
       public            postgres    false    214            m           2606    18577    solicitud solicitud_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_pkey;
       public            postgres    false    211            r           2606    18578    alumno Alumno_id_carrera_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_id_carrera_fkey" FOREIGN KEY (id_carrera) REFERENCES public.carrera(id);
 I   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_id_carrera_fkey";
       public          postgres    false    202    204    2917            s           2606    18583 "   electivo electivo_id_profesor_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_id_profesor_fkey FOREIGN KEY (id_profesor) REFERENCES public.profesor(id);
 L   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_id_profesor_fkey;
       public          postgres    false    208    2919    206            t           2606    18588 (   postulacion postulacion_id_electivo_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_id_electivo_fkey FOREIGN KEY (id_electivo) REFERENCES public.electivo(id);
 R   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_id_electivo_fkey;
       public          postgres    false    210    2921    208            z           2606    18639 &   ramo_alumno ramo_alumno_id_alumno_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 P   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_id_alumno_fkey;
       public          postgres    false    202    2915    216            y           2606    18634 $   ramo_alumno ramo_alumno_id_ramo_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_id_ramo_fkey FOREIGN KEY (id_ramo) REFERENCES public.ramo(id);
 N   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_id_ramo_fkey;
       public          postgres    false    214    2927    216            u           2606    18593 "   solicitud solicitud_id_alumno_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 L   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_alumno_fkey;
       public          postgres    false    2915    202    211            v           2606    18598 )   solicitud solicitud_id_postulacion_1_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_1_fkey FOREIGN KEY (id_postulacion_1) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_1_fkey;
       public          postgres    false    211    2923    210            w           2606    18603 )   solicitud solicitud_id_postulacion_2_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_2_fkey FOREIGN KEY (id_postulacion_2) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_2_fkey;
       public          postgres    false    210    2923    211            x           2606    18608 )   solicitud solicitud_id_postulacion_3_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_3_fkey FOREIGN KEY (id_postulacion_3) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_3_fkey;
       public          postgres    false    210    2923    211            �      x������ � �      �   \   x�3�420263273�t����,��Wp�I�J�K)�W���,K,J-Vp,.)M�����L���ˇ�9$�����&��%�pr��W� �11            x�3��tv������             x������ � �            x������ � �            x������ � �            x������ � �            x������ � �            x������ � �     