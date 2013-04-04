# 2520 is the smallest number that can be divided by each of the numbers
# from 1 to 10 without any remainder.
# What is the smallest positive number that is evenly divisible by all of
# the numbers from 1 to 20?

@start = 1
@end = 20

@nums = []
@product = 1

def is_prime?(number)
  (2..Math.sqrt(number).floor).none? do |x|
    number % x == 0
  end
end

@end.downto(@start) do |n|
  @nums << n if is_prime?(n)
end

@nums.each { |n| @product *= n }

puts @product