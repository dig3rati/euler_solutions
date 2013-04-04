# The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
# 
# Find the sum of all the primes below two million.

def is_prime?(number)
  (2..Math.sqrt(number).floor).none? do |x|
    number % x == 0
  end
end

LIMIT = 2000000
@sum = 0

for n in 2.upto(LIMIT) do
  @sum += n if is_prime?(n)
end

puts @sum