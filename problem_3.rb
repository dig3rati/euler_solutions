# The prime factors of 13195 are 5, 7, 13 and 29.
# What is the largest prime factor of the number 600851475143 ?

def is_prime?(number)
  (2..Math.sqrt(number).floor).none? do |x|
    number % x == 0
  end
end

def first_prime_factor(number)
  (2..number).each do |x|
    return x if number % x == 0 && is_prime?(x)
  end
end

def prime_factors(factors)
  factors = [1, factors] if factors.class == Fixnum || factors.class == Bignum
  return factors if is_prime?(factors.last)
  n = factors.pop
  f = first_prime_factor(n)
  factors.push f
  factors.push n/f
  return prime_factors(factors)
end

puts prime_factors(600851475143).last