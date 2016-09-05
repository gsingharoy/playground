class Fibonacci
  def initialize(length : Int32)
    @length = length
  end

  def calculate
    result = [0, 1]
    while @length > 1
      result << result[-1] + result[-2]
      @length -= 1
    end
    result[1..-1]
  end
end

puts Fibonacci.new(10).calculate
