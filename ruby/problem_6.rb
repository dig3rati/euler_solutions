# The sum of the squares of the first ten natural numbers is,
# 1^2 + 2^2 + ... + 10^2 = 385
# The square of the sum of the first ten natural numbers is,
# (1 + 2 + ... + 10)^2 = 55^2 = 3025
# Hence the difference between the sum of the squares of the first ten
# natural numbers and the square of the sum is 3025  385 = 2640.
# Find the difference between the sum of the squares of the first one
# hundred natural numbers and the square of the sum.

@till = 100

def sum_to_n(number)
  @sum ||= (number * (number + 1))/2
end

def sum_of_squares_to_n(number)
  @sum2 ||= (sum_to_n(number) * (2*number + 1))/3
end

puts sum_to_n(@till) * sum_to_n(@till) - sum_of_squares_to_n(@till)