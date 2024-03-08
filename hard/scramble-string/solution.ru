def is_scramble(s1, s2)
    memo = {}
    memo_is_scramble(s1, s2, memo)
  end
  
  def memo_is_scramble(s1, s2, memo = {})
    return true if s1 == s2
    return false if s1.length != s2.length
    
    return memo[[s1, s2]] if memo.key?([s1, s2])
    
    count = Array.new(26, 0)
    (0...s1.length).each do |i|
      count[s1[i].ord - 'a'.ord] += 1
      count[s2[i].ord - 'a'.ord] -= 1
    end
    return false if count.any? { |c| c != 0 }
    
    (1...s1.length).each do |i|
      if (memo_is_scramble(s1[0...i], s2[0...i], memo) && memo_is_scramble(s1[i..-1], s2[i..-1], memo)) ||
         (memo_is_scramble(s1[0...i], s2[-i..-1], memo) && memo_is_scramble(s1[i..-1], s2[0...s2.length-i], memo))
        memo[[s1, s2]] = true
        return true
      end
    end
  
    memo[[s1, s2]] = false
    false
  end