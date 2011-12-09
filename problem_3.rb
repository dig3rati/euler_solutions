def is_prime?(num, trivial = false)
  end_val = trivial ? 5 : Math.sqrt(num).ceil
  (2..end_val).each do |n|
    return false if num % n == 0
  end
  return true
end

def get_prime_factors(num)
  # Uncomment if we need all prime factors
  # factors = []
  num.downto(2) do |n|
    # factors << n if num % n == 0 && is_prime?(n)
    return n if num % n == 0 && is_prime?(n)
  end
  # factors
end

puts get_prime_factors(600851475143)