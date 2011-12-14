@start = 100
@end = 999

@palindromes = {}

def is_palindrome?(number)
  number.to_s == number.to_s.reverse
end

def get_palindrome_in_range
  numbers = @end.downto(@start).to_a
  for i in @end.downto(@start)
    for j in @end.downto(@start)
      @palindromes[i * j] = [i, j] if is_palindrome?(i * j)
    end
  end
end

get_palindrome_in_range
puts @palindromes.keys.max