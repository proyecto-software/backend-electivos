PGDMP     1                
    y           administracion_electivos    13.3    13.3 R               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                        0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            !           1262    18656    administracion_electivos    DATABASE     t   CREATE DATABASE administracion_electivos WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Spanish_Spain.1252';
 (   DROP DATABASE administracion_electivos;
                postgres    false            �            1255    18802    formulario_admin()    FUNCTION     +  CREATE FUNCTION public.formulario_admin() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
declare
	indicador integer;
	sem_inc integer;
begin
	--cálculo indicador:
	select semestre_incompleto into sem_inc from alumno a where a.id = new.id_alumno;
	indicador = 9 - sem_inc;
	insert into Registro_Postulacion(rut,nombre,carrera,indicador,electivo) values ((select rut from alumno a where a.id = new.id_alumno),(select nombre from alumno a where a.id = new.id_alumno),(select nombre from carrera c where c.id = (select id_carrera from alumno a where a.id = new.id_alumno)),indicador,(select nombre from electivo e where e.id = (select id_electivo from postulacion p where p.id = new.id_postulacion_1)));
	insert into Registro_Postulacion(rut,nombre,carrera,indicador,electivo) values ((select rut from alumno a where a.id = new.id_alumno),(select nombre from alumno a where a.id = new.id_alumno),(select nombre from carrera c where c.id = (select id_carrera from alumno a where a.id = new.id_alumno)),indicador,(select nombre from electivo e where e.id = (select id_electivo from postulacion p where p.id = new.id_postulacion_2)));
	insert into Registro_Postulacion(rut,nombre,carrera,indicador,electivo) values ((select rut from alumno a where a.id = new.id_alumno),(select nombre from alumno a where a.id = new.id_alumno),(select nombre from carrera c where c.id = (select id_carrera from alumno a where a.id = new.id_alumno)),indicador,(select nombre from electivo e where e.id = (select id_electivo from postulacion p where p.id = new.id_postulacion_3)));
	return new;
end;
$$;
 )   DROP FUNCTION public.formulario_admin();
       public          postgres    false            �            1259    18657    administrador    TABLE     z   CREATE TABLE public.administrador (
    id integer NOT NULL,
    correo text,
    rut text NOT NULL,
    password text
);
 !   DROP TABLE public.administrador;
       public         heap    postgres    false            �            1259    18663    Administrador_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Administrador_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public."Administrador_id_seq";
       public          postgres    false    200            "           0    0    Administrador_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public."Administrador_id_seq" OWNED BY public.administrador.id;
          public          postgres    false    201            �            1259    18665    alumno    TABLE     �   CREATE TABLE public.alumno (
    id integer NOT NULL,
    rut text NOT NULL,
    nombre text,
    correo text,
    id_carrera integer NOT NULL,
    semestre_incompleto integer,
    cantidad_ramos integer
);
    DROP TABLE public.alumno;
       public         heap    postgres    false            �            1259    18671    Alumno_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Alumno_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public."Alumno_id_seq";
       public          postgres    false    202            #           0    0    Alumno_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public."Alumno_id_seq" OWNED BY public.alumno.id;
          public          postgres    false    203            �            1259    18673    carrera    TABLE     J   CREATE TABLE public.carrera (
    id integer NOT NULL,
    nombre text
);
    DROP TABLE public.carrera;
       public         heap    postgres    false            �            1259    18679    Carrera_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Carrera_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public."Carrera_id_seq";
       public          postgres    false    204            $           0    0    Carrera_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public."Carrera_id_seq" OWNED BY public.carrera.id;
          public          postgres    false    205            �            1259    18681    profesor    TABLE     \   CREATE TABLE public.profesor (
    id integer NOT NULL,
    nombre text,
    correo text
);
    DROP TABLE public.profesor;
       public         heap    postgres    false            �            1259    18687    Profesor_id_seq    SEQUENCE     �   CREATE SEQUENCE public."Profesor_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public."Profesor_id_seq";
       public          postgres    false    206            %           0    0    Profesor_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public."Profesor_id_seq" OWNED BY public.profesor.id;
          public          postgres    false    207            �            1259    18689    electivo    TABLE     t   CREATE TABLE public.electivo (
    id integer NOT NULL,
    nombre text,
    "NRC" text,
    id_profesor integer
);
    DROP TABLE public.electivo;
       public         heap    postgres    false            �            1259    18695    electivo_id_seq    SEQUENCE     �   CREATE SEQUENCE public.electivo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.electivo_id_seq;
       public          postgres    false    208            &           0    0    electivo_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.electivo_id_seq OWNED BY public.electivo.id;
          public          postgres    false    209            �            1259    18697    postulacion    TABLE     �   CREATE TABLE public.postulacion (
    id integer NOT NULL,
    cantidad integer,
    id_electivo integer,
    aprobado boolean
);
    DROP TABLE public.postulacion;
       public         heap    postgres    false            �            1259    18764    ramo    TABLE     k   CREATE TABLE public.ramo (
    id integer NOT NULL,
    nrc text,
    nombre text,
    creditos integer
);
    DROP TABLE public.ramo;
       public         heap    postgres    false            �            1259    18775    ramo_alumno    TABLE     �   CREATE TABLE public.ramo_alumno (
    id_ramo integer NOT NULL,
    id_alumno integer NOT NULL,
    nota double precision,
    oportunidad integer
);
    DROP TABLE public.ramo_alumno;
       public         heap    postgres    false            �            1259    18773    ramo_alumno_id_ramo_seq    SEQUENCE     �   CREATE SEQUENCE public.ramo_alumno_id_ramo_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.ramo_alumno_id_ramo_seq;
       public          postgres    false    216            '           0    0    ramo_alumno_id_ramo_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.ramo_alumno_id_ramo_seq OWNED BY public.ramo_alumno.id_ramo;
          public          postgres    false    215            �            1259    18762    ramo_id_seq    SEQUENCE     �   CREATE SEQUENCE public.ramo_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.ramo_id_seq;
       public          postgres    false    214            (           0    0    ramo_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.ramo_id_seq OWNED BY public.ramo.id;
          public          postgres    false    213            �            1259    18793    registro_postulacion    TABLE     �   CREATE TABLE public.registro_postulacion (
    id integer NOT NULL,
    rut text,
    nombre text,
    carrera text,
    indicador text,
    electivo text
);
 (   DROP TABLE public.registro_postulacion;
       public         heap    postgres    false            �            1259    18791    registro_postulacion_id_seq    SEQUENCE     �   CREATE SEQUENCE public.registro_postulacion_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 2   DROP SEQUENCE public.registro_postulacion_id_seq;
       public          postgres    false    218            )           0    0    registro_postulacion_id_seq    SEQUENCE OWNED BY     [   ALTER SEQUENCE public.registro_postulacion_id_seq OWNED BY public.registro_postulacion.id;
          public          postgres    false    217            �            1259    18700 	   solicitud    TABLE     �   CREATE TABLE public.solicitud (
    id integer NOT NULL,
    id_alumno integer NOT NULL,
    id_postulacion_1 integer NOT NULL,
    id_postulacion_2 integer,
    id_postulacion_3 integer
);
    DROP TABLE public.solicitud;
       public         heap    postgres    false            �            1259    18703    solicitud_id_seq    SEQUENCE     �   CREATE SEQUENCE public.solicitud_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.solicitud_id_seq;
       public          postgres    false    211            *           0    0    solicitud_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.solicitud_id_seq OWNED BY public.solicitud.id;
          public          postgres    false    212            ^           2604    18705    administrador id    DEFAULT     v   ALTER TABLE ONLY public.administrador ALTER COLUMN id SET DEFAULT nextval('public."Administrador_id_seq"'::regclass);
 ?   ALTER TABLE public.administrador ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    201    200            _           2604    18706 	   alumno id    DEFAULT     h   ALTER TABLE ONLY public.alumno ALTER COLUMN id SET DEFAULT nextval('public."Alumno_id_seq"'::regclass);
 8   ALTER TABLE public.alumno ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202            `           2604    18707 
   carrera id    DEFAULT     j   ALTER TABLE ONLY public.carrera ALTER COLUMN id SET DEFAULT nextval('public."Carrera_id_seq"'::regclass);
 9   ALTER TABLE public.carrera ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    205    204            b           2604    18708    electivo id    DEFAULT     j   ALTER TABLE ONLY public.electivo ALTER COLUMN id SET DEFAULT nextval('public.electivo_id_seq'::regclass);
 :   ALTER TABLE public.electivo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    209    208            a           2604    18709    profesor id    DEFAULT     l   ALTER TABLE ONLY public.profesor ALTER COLUMN id SET DEFAULT nextval('public."Profesor_id_seq"'::regclass);
 :   ALTER TABLE public.profesor ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    207    206            d           2604    18767    ramo id    DEFAULT     b   ALTER TABLE ONLY public.ramo ALTER COLUMN id SET DEFAULT nextval('public.ramo_id_seq'::regclass);
 6   ALTER TABLE public.ramo ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    213    214            e           2604    18778    ramo_alumno id_ramo    DEFAULT     z   ALTER TABLE ONLY public.ramo_alumno ALTER COLUMN id_ramo SET DEFAULT nextval('public.ramo_alumno_id_ramo_seq'::regclass);
 B   ALTER TABLE public.ramo_alumno ALTER COLUMN id_ramo DROP DEFAULT;
       public          postgres    false    216    215    216            f           2604    18796    registro_postulacion id    DEFAULT     �   ALTER TABLE ONLY public.registro_postulacion ALTER COLUMN id SET DEFAULT nextval('public.registro_postulacion_id_seq'::regclass);
 F   ALTER TABLE public.registro_postulacion ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            c           2604    18710    solicitud id    DEFAULT     l   ALTER TABLE ONLY public.solicitud ALTER COLUMN id SET DEFAULT nextval('public.solicitud_id_seq'::regclass);
 ;   ALTER TABLE public.solicitud ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    212    211            	          0    18657    administrador 
   TABLE DATA           B   COPY public.administrador (id, correo, rut, password) FROM stdin;
    public          postgres    false    200   rb                 0    18665    alumno 
   TABLE DATA           j   COPY public.alumno (id, rut, nombre, correo, id_carrera, semestre_incompleto, cantidad_ramos) FROM stdin;
    public          postgres    false    202   �b                 0    18673    carrera 
   TABLE DATA           -   COPY public.carrera (id, nombre) FROM stdin;
    public          postgres    false    204   %c                 0    18689    electivo 
   TABLE DATA           B   COPY public.electivo (id, nombre, "NRC", id_profesor) FROM stdin;
    public          postgres    false    208   Sc                 0    18697    postulacion 
   TABLE DATA           J   COPY public.postulacion (id, cantidad, id_electivo, aprobado) FROM stdin;
    public          postgres    false    210   �c                 0    18681    profesor 
   TABLE DATA           6   COPY public.profesor (id, nombre, correo) FROM stdin;
    public          postgres    false    206   �c                 0    18764    ramo 
   TABLE DATA           9   COPY public.ramo (id, nrc, nombre, creditos) FROM stdin;
    public          postgres    false    214   !d                 0    18775    ramo_alumno 
   TABLE DATA           L   COPY public.ramo_alumno (id_ramo, id_alumno, nota, oportunidad) FROM stdin;
    public          postgres    false    216   >d                 0    18793    registro_postulacion 
   TABLE DATA           ]   COPY public.registro_postulacion (id, rut, nombre, carrera, indicador, electivo) FROM stdin;
    public          postgres    false    218   [d                 0    18700 	   solicitud 
   TABLE DATA           h   COPY public.solicitud (id, id_alumno, id_postulacion_1, id_postulacion_2, id_postulacion_3) FROM stdin;
    public          postgres    false    211   �d       +           0    0    Administrador_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public."Administrador_id_seq"', 1, false);
          public          postgres    false    201            ,           0    0    Alumno_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public."Alumno_id_seq"', 1, true);
          public          postgres    false    203            -           0    0    Carrera_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public."Carrera_id_seq"', 1, false);
          public          postgres    false    205            .           0    0    Profesor_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public."Profesor_id_seq"', 1, false);
          public          postgres    false    207            /           0    0    electivo_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.electivo_id_seq', 1, false);
          public          postgres    false    209            0           0    0    ramo_alumno_id_ramo_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.ramo_alumno_id_ramo_seq', 1, false);
          public          postgres    false    215            1           0    0    ramo_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.ramo_id_seq', 1, false);
          public          postgres    false    213            2           0    0    registro_postulacion_id_seq    SEQUENCE SET     I   SELECT pg_catalog.setval('public.registro_postulacion_id_seq', 3, true);
          public          postgres    false    217            3           0    0    solicitud_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.solicitud_id_seq', 1, false);
          public          postgres    false    212            h           2606    18712     administrador Administrador_pkey 
   CONSTRAINT     e   ALTER TABLE ONLY public.administrador
    ADD CONSTRAINT "Administrador_pkey" PRIMARY KEY (id, rut);
 L   ALTER TABLE ONLY public.administrador DROP CONSTRAINT "Administrador_pkey";
       public            postgres    false    200    200            j           2606    18714    alumno Alumno_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_pkey" PRIMARY KEY (id, rut);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_pkey";
       public            postgres    false    202    202            n           2606    18716    carrera Carrera_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.carrera
    ADD CONSTRAINT "Carrera_pkey" PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.carrera DROP CONSTRAINT "Carrera_pkey";
       public            postgres    false    204            p           2606    18718    profesor Profesor_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.profesor
    ADD CONSTRAINT "Profesor_pkey" PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.profesor DROP CONSTRAINT "Profesor_pkey";
       public            postgres    false    206            l           2606    18720    alumno alumno_id_key 
   CONSTRAINT     M   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT alumno_id_key UNIQUE (id);
 >   ALTER TABLE ONLY public.alumno DROP CONSTRAINT alumno_id_key;
       public            postgres    false    202            r           2606    18722    electivo electivo_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_pkey;
       public            postgres    false    208            t           2606    18724    postulacion postulacion_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_pkey;
       public            postgres    false    210            z           2606    18780    ramo_alumno ramo_alumno_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_pkey PRIMARY KEY (id_ramo, id_alumno);
 F   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_pkey;
       public            postgres    false    216    216            x           2606    18772    ramo ramo_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.ramo
    ADD CONSTRAINT ramo_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.ramo DROP CONSTRAINT ramo_pkey;
       public            postgres    false    214            |           2606    18801 .   registro_postulacion registro_postulacion_pkey 
   CONSTRAINT     l   ALTER TABLE ONLY public.registro_postulacion
    ADD CONSTRAINT registro_postulacion_pkey PRIMARY KEY (id);
 X   ALTER TABLE ONLY public.registro_postulacion DROP CONSTRAINT registro_postulacion_pkey;
       public            postgres    false    218            v           2606    18726    solicitud solicitud_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_pkey;
       public            postgres    false    211            �           2620    18803 "   solicitud trigger_formulario_admin    TRIGGER     �   CREATE TRIGGER trigger_formulario_admin BEFORE INSERT ON public.solicitud FOR EACH ROW EXECUTE FUNCTION public.formulario_admin();
 ;   DROP TRIGGER trigger_formulario_admin ON public.solicitud;
       public          postgres    false    211    219            }           2606    18727    alumno Alumno_id_carrera_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.alumno
    ADD CONSTRAINT "Alumno_id_carrera_fkey" FOREIGN KEY (id_carrera) REFERENCES public.carrera(id);
 I   ALTER TABLE ONLY public.alumno DROP CONSTRAINT "Alumno_id_carrera_fkey";
       public          postgres    false    204    2926    202            ~           2606    18732 "   electivo electivo_id_profesor_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.electivo
    ADD CONSTRAINT electivo_id_profesor_fkey FOREIGN KEY (id_profesor) REFERENCES public.profesor(id);
 L   ALTER TABLE ONLY public.electivo DROP CONSTRAINT electivo_id_profesor_fkey;
       public          postgres    false    206    2928    208                       2606    18737 (   postulacion postulacion_id_electivo_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.postulacion
    ADD CONSTRAINT postulacion_id_electivo_fkey FOREIGN KEY (id_electivo) REFERENCES public.electivo(id);
 R   ALTER TABLE ONLY public.postulacion DROP CONSTRAINT postulacion_id_electivo_fkey;
       public          postgres    false    208    2930    210            �           2606    18786 &   ramo_alumno ramo_alumno_id_alumno_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 P   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_id_alumno_fkey;
       public          postgres    false    2924    216    202            �           2606    18781 $   ramo_alumno ramo_alumno_id_ramo_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.ramo_alumno
    ADD CONSTRAINT ramo_alumno_id_ramo_fkey FOREIGN KEY (id_ramo) REFERENCES public.ramo(id);
 N   ALTER TABLE ONLY public.ramo_alumno DROP CONSTRAINT ramo_alumno_id_ramo_fkey;
       public          postgres    false    2936    216    214            �           2606    18742 "   solicitud solicitud_id_alumno_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_alumno_fkey FOREIGN KEY (id_alumno) REFERENCES public.alumno(id);
 L   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_alumno_fkey;
       public          postgres    false    202    2924    211            �           2606    18747 )   solicitud solicitud_id_postulacion_1_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_1_fkey FOREIGN KEY (id_postulacion_1) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_1_fkey;
       public          postgres    false    2932    210    211            �           2606    18752 )   solicitud solicitud_id_postulacion_2_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_2_fkey FOREIGN KEY (id_postulacion_2) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_2_fkey;
       public          postgres    false    2932    211    210            �           2606    18757 )   solicitud solicitud_id_postulacion_3_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.solicitud
    ADD CONSTRAINT solicitud_id_postulacion_3_fkey FOREIGN KEY (id_postulacion_3) REFERENCES public.postulacion(id);
 S   ALTER TABLE ONLY public.solicitud DROP CONSTRAINT solicitud_id_postulacion_3_fkey;
       public          postgres    false    210    211    2932            	      x������ � �         �   x�e�A
�0EדS�!�0�Z"z�nBaʘ�N�<�3Hw�_|�� ����H*)��,��U��֢vԭeb�;�dߎ�۳��6W73�n=&vu�}W_D?o{m����O����u�h&g����3i            x�3��tv��2����2r<�b���� 9�K         (   x�3�t5�44b.#NW# ���9]��LcN#�=... l�c         2   x�346��4�4.C#(���� �4612!L3 ��4�*����� �H&         D   x�3��M,J�/Vp�H�,��MQ��yz�9\F�ΉE9@I�ļ�ļ|�
 �I�$2�JP�b���� c��            x������ � �            x������ � �         @   x�3�420437303���/>�R��4�X�-'�(�����ٓӘ�Րˈ(uF\�D�3����� V�!1         *   x��02�4�42�0@\�FƜF�F�&� �D�s��qqq �Sj     