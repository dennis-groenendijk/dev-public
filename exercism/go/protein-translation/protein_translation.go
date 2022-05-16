package protein

import "errors"

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

// FromRNA splits a RNA string into codons and returns the collection of corresponding proteins
func FromRNA(rna string) ([]string, error) {
    // 1. split the rna string into a slice of codons
	var codons []string
    for i := 1; i < len(rna); i++ {
        if i%3 == 0 {
            codons = append(codons, rna[:i])
            rna = rna[i:]
            i = 1
        }
    }
	codons = append(codons, rna)

    // 2. check for and remove doubles
    check := make(map[string]bool)
    newCodons := []string{}
    for _, codon := range codons {
        if _, double := check[codon]; !double {
            check[codon] = true
            newCodons = append(newCodons, codon)
        }
    }

    // 3. convert the codon to a protein
    var protein []string
    for i := 0; i < len(newCodons); i++ {
        convert, err := FromCodon(newCodons[i])
        if err != nil && err == ErrStop {
             return protein, nil
        } else if err == ErrInvalidBase {
        	return nil, err
        }
        protein = append(protein, convert)
    }
	return protein, nil
}

// FromCodon returns the corresponding protein to the provided codon
func FromCodon(codon string) (string, error) {
	switch codon {
        case "AUG":
    	return "Methionine", nil
        case "UUU", "UUC":
    	return "Phenylalanine", nil
        case "UUA", "UUG":
    	return "Leucine", nil
        case "UCU", "UCC", "UCA", "UCG":
    	return "Serine", nil
        case "UAU", "UAC":
    	return "Tyrosine", nil
        case "UGU", "UGC":
        return "Cysteine", nil
        case "UGG":
    	return "Tryptophan", nil
        case "UAA", "UAG", "UGA":
    	return "", ErrStop
        default:
    	return "", ErrInvalidBase
    }
}
