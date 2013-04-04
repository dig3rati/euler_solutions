# By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can
# see that the 6th prime is 13.
# What is the 10 001st prime number?

@primes = []
@start = 2  # First prime

def is_prime?(number)
  (2..Math.sqrt(number).floor).none? do |x|
    number % x == 0
  end
end

while @primes.count < 10001
  @primes << @start if is_prime?(@start)
  @start += 1
end

puts @primes.last