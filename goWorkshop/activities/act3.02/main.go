package main

goodCreditScore := 450

interestRateGoodScore := 0.15
interestRateLessScore := 0.20

mpLessScore := 0.10
mpGoodScore := 0.20


loanAmonunt (total emprestado)
monthlyPayment (parcela)
monthlyIncome (salario)
loanTerm(tempo)
interestPAyment = loanAmount * interestRate * loanTerm

var (
	err_loanTermTime = errors.New("Loan term is not anual")
	err_creditScore = errors.New("Credit score can't be zero ou negative")
	err_monIn = errors.New("Monthly income can't be zero or negative")
	err_loanAmount = errors.New("Loan amount can't be zero or negative")
	err_LoanTerm = errors.New("Loan term can't be zero or negative")
)


if credtScore >= 450 {
	interestRate = 0.15
}else {
	interestRate = 0.20
}



func main() {
	
}