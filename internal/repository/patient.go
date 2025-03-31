package repository

type PatientRepository struct {
}

func NewPatientRepository() *PatientRepository {
	return &PatientRepository{}
}

func (pr *PatientRepository) GetPatient() {
	panic("impl me!")
}

func (pr *PatientRepository) GetAllPatients() {
	panic("impl me!")
}

func (pr *PatientRepository) UpdatePatient() {
	panic("impl me!")
}

func (pr *PatientRepository) CreatePatient() {
	panic("impl me!")
}

func (pr *PatientRepository) DeletePatient() {
	panic("impl me!")
}
