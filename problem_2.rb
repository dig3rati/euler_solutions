@sum = 0
ENDS_AT = 4000000

def even_sum_fib(f, s)
  @sum += f if f % 2 == 0
  even_sum_fib(s, f+s) if s <= ENDS_AT
end

even_sum_fib(1, 2)
puts @sum