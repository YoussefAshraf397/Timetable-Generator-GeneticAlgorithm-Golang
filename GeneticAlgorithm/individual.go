package GeneticAlgorithm

// Individual for this genetic algorithm
type Individual struct {
	Chromosome []int
	Fitness    float64
}

func createIndividual(timeTable TimeTable) (individual Individual) {
	numClasses := timeTable.GetNumSessions()
	chromosomeLength := numClasses * 3
	newChromosome := make([]int, chromosomeLength)
	chromosomeIndex := 0

	for _, classroom := range timeTable.GetClassrooms() {

		for _, subjectID := range classroom.GetSubjectIDs() {
			timeSlotID := timeTable.GetRandomTimeSlot().TimeSlotID
			newChromosome[chromosomeIndex] = timeSlotID
			chromosomeIndex++

			// Add random professor
			subject := timeTable.GetSubject(subjectID)
			newChromosome[chromosomeIndex] = subject.GetRandomTeacherID()
			chromosomeIndex++
		}
	}

	individual = Individual{
		Chromosome: newChromosome,
		Fitness:    0,
	}
	return
}
