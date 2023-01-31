create table nasabah(
    id int unsigned primary key auto_increment,
    nama_lengkap varchar(300),
    alamat text,
    tempat_lahir varchar(300),
    tanggal_lahir date,
    avg_penghasilan int (255),
    jenis_nasabah tinyint (1),
    id_customer varchar(15) unique
) engine = innodb;