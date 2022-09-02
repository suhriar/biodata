package seeder

import "biodata/models"

func AddStudent(id, name, address, job, reason string) models.Students {
	return models.Students{
		Id:      id,
		Name:    name,
		Address: address,
		Job:     job,
		Reason:  reason,
	}
}

func Biodata() []models.Students {
	var biodatas []models.Students

	bio := AddStudent("1", "Student 1", "Jalan Lurus", "Mahasiswa", "Tertarik belajar Golang")
	biodatas = append(biodatas, bio)

	bio = AddStudent("2", "Student 2", "Jalan Soekarno", "Mahasiswa", "Menambah Insight Bahas Pemrograman")
	biodatas = append(biodatas, bio)

	bio = AddStudent("3", "Student 3", "Perumnas", "Mahasiswa", "Menguasai tech Golang")
	biodatas = append(biodatas, bio)

	bio = AddStudent("4", "Student 4", "Sepanjang jalan", "Mahasiswa", "menambah tech stack")
	biodatas = append(biodatas, bio)

	bio = AddStudent("5", "Student 5", "Jalan Kenangan", "Mahasiswa", "Ingin Jadi Backend engineer")
	biodatas = append(biodatas, bio)

	return biodatas

}
