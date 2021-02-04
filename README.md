# debt_project
TrueAccord Take Home

# Software Installations 
Install Golang

https://golang.org/doc/install

# Instructions to build and run the project

In the src folder (/TrueAccord/go/src) run the following commands

git clone git@github.com:vontelsunita/debt_project.git

cd debt_project 

go build -o TrueAccordDebtApplication

./TrueAccordDebtApplication

You should see output as below

INFO[0000] All Debts cleared for DebtID 0               

INFO[0000] {"id":0,"amount":123.46,"is_in_payment_plan":true,"remaining_amount":0,"next_payment_due_date":null} 

INFO[0000] {"id":1,"amount":100,"is_in_payment_plan":true,"remaining_amount":50,"next_payment_due_date":"2020-08-08T00:00:00Z"} 

INFO[0000] {"id":2,"amount":4920.34,"is_in_payment_plan":true,"remaining_amount":607.6699,"next_payment_due_date":"2020-01-15T00:00:00Z"} 

INFO[0000] {"id":3,"amount":12938,"is_in_payment_plan":true,"remaining_amount":622.41504,"next_payment_due_date":"2020-08-08T00:00:00Z"} 

INFO[0000] {"id":4,"amount":9238.02,"is_in_payment_plan":false,"remaining_amount":0,"next_payment_due_date":null} 

# Unit Tests

go test ./... 

There are 2 unit tests in the package tests. Those should execute and yield these results

?       debt_project    [no test files]

?       debt_project/db [no test files]

?       debt_project/models     [no test files]

ok      debt_project/tests      0.938s

