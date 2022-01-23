# frozen_string_literal: true

require 'date'

def main
  today = Date.today
  birth_date = Date.parse(ARGV[0])
  count = (ARGV[1] || 10).to_i

  print_header
  print_date(birth_date.year, birth_date)
  0.upto(count - 1) { |i| print_date(today.year + i, birth_date) }
end

def print_date(year, date)
  d = Date.new(year, date.month, date.day)
  puts "#{d}\t#{year - date.year}\t#{d.strftime('%A')}"
end

def print_header
  puts 'Ruby'
  puts "Date      \tAge\tDay"
  puts "----------\t---\t----------"
end

main
