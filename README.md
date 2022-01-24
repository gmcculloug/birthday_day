# birthday_day

Sample app written in Ruby and Go that displays:
* Year, Age, Day of the week.

# Usage 
birthday_day <birth_date> <number_of_years>

where:
* birth_date - (Format: **YYYY-MM-DD**)
* number_of_years - How many years to process from current year (Default: **10**)

### Example Go output:
```
go run birthday_day.go 2000-01-01 5
GoLang
Date      	Age	Day
----------	---	----------
2000-01-01	0	Saturday
2022-01-01	22	Saturday
2023-01-01	23	Sunday
2024-01-01	24	Monday
2025-01-01	25	Wednesday
2026-01-01	26	Thursday
```

### Example Ruby output:
```
ruby birthday_day.rb 2000-01-01 5
Ruby
Date      	Age	Day
----------	---	----------
2000-01-01	0	Saturday
2022-01-01	22	Saturday
2023-01-01	23	Sunday
2024-01-01	24	Monday
2025-01-01	25	Wednesday
2026-01-01	26	Thursday
```
