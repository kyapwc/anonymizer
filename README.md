# Introduction

You are working on auction platform. The service provides to its user ability to submit and search auctions. Company has to implement some privacy policy, for example, some personal data like emails, Skype usernames or phone numbers must be anonymized.

# Task definition

Your task is to implement 3 content anonymizers:

* for emails (anonymize whole username, leave domain)
* for Skype username (anonymize whole username, leave HTML around if given)
* for phone numbers (anonymize last X digits, leave the rest and code)

To complete this task you should:

* implement methods marked with `@TODO` annotation in *anonymizer package
* check if all anonymizers tests are green
* do not change the interface structure of `anonymizer.Anonymizer`

# Input structure

## Task 1 - Email username anonymizer

Simple usage:
```go
ea := NewEmailAnonymizer("***")
ea.Anonymize("my.mail@gmail.com.pl")      //-> ***@gmail.com.pl
```

Example of valid mail usernames:
* `a@a.com`
* `aa@aa.aa.com`
* `aa12@aa12.aa.com`
* `A-A@A-A.com`
* `A.b+A@AA.com`

Matching mail usernames rules:
* characters: `a-z`, `A-Z`, `0-9`, `.`, `_`, `-`, `+`
* first and last character of username/domain must be a-z, A-Z or 0-9 character:
* all mail usernames must be anonymized in the input string

Examples:
```go
ea := NewEmailAnonymizer("...")
ea.Anonymize("-my-mail-second@gmail.com")   //-> -...@gmail.com
ea.Anonymize(".my-mail.here@gmail.com")     //-> ....@gmail.com
ea.Anonymize(".my-mail.here.@gmail.com")    //-> .my-mail.here.@gmail.com
ea.Anonymize("first-mail@gmail.com and second-mail@gmail.com")    //-> ...@gmail.com and ...@gmail.com
```

For simplicity, you don't have to implement RFC standards.

## Task 2 - Skype usernames anonymizer
Simple usage:
```go
NewSkypeAnonymizer("#").Anonymize(`<a href="skype:loren?call"/>`) //-> <a href="skype:#?call"/>
```

Example of valid Skype usernames:
* `skype:username`
* `skype:USERNAME`
* `<a href="skype:USERNAME?call">call me</a>`

Matching skype usernames rules:
* characters: `a-z`, `A-Z`, `0-9`

## Task 3 - Phone numbers anonymizer

For simplicity, all phone numbers are formatted the same way, you may assume that. There are no different numbers in auction content, like credit card numbers.

Simple example:
```go
pa := NewPhoneAnonymizer()
pa.Anonymize("+48 666 777 888") //-> +48 666 777 XXX
pa.SetLastDigits(5)
pa.Anonymize("+48 666 777 888") //-> +48 666 7XX XXX
```

Example of valid Phone numbers:

* `+48 666 666 666`
* `+234 777 888 999`

Matching Phone number rules:

* international prefix code is always available ( for example `+48`)
* phone number contains 9 digits in 3 groups, 3 digits each, separated by spaces
* maximum 9 digits can be anonymized (meaning prefix should never be anonymized)

# Hints

Think about all edge cases that could be the input to all anonymizers.


# Running test locally

```go
go test ./... -v -race
```


 # Linter used for development

`golangci-lint run`

# Insights

Most of the entire exercise relies heavily upon regex and also looping through the original matches to update the original string to get the result. therefore most of the functions have similar start code as my thought process is:

1. find matches in string [][]int
2. loop through all content in the matches
3. write appropriate logic to replace the content in original string
4. return the value

Generally this is my thought process on the entire test when conducting it. Furthermore, the provided examples in the readme is not detailed enough, as the tests itself is quite different from actual provided examples in the readme, so it took some time to review why some regex did not work.

Lastly, apologies for the late submission, as I am currently travelling with my family in Jeju Island in South Korea, I don't have the time to code as much as I am constantly on the move. Thank you for the patience and opportunity.
