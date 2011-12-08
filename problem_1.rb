sum = 0
(1..1000-1).each do |i|
  sum += i if (i % 3 == 0 || i % 5 == 0)
end
puts sum