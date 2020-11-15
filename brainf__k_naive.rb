module Brainf__kNaive
  def self.run(code)
    tape = {}
    dp = 0
    ip = 0
    until code.length <= ip
      case code[ip]
      when '+'
        tape[dp] ||= 0
        tape[dp] += 1
      when '-'
        tape[dp] ||= 0
        tape[dp] -= 1
      when '>'
        dp += 1
      when '<'
        dp -= 1
      when '.'
        putc tape[dp].chr
      when ','
        raise ', is not implemented yet'
      when '['
        if tape[dp] == 0
          ip = next_matching_close(code, ip)
        end
      when']'
        if tape[dp] != 0
          ip = prev_matching_open(code, ip)
        end
      else
        # comment
      end

      ip += 1
    end
  end

  private_class_method def self.next_matching_close(code, ip)
    n = 0
    loop do
      ip += 1
      case [code[ip], n]
      in ['[', _]
        n += 1
      in [']', 0]
        return ip
      in [']', _]
        n -= 1
      else
        # pass
      end
    end
    raise 'Missing corresponding next_matching_close'
  end

  private_class_method def self.prev_matching_open(code, ip)
    n = 0
    loop do
      ip -= 1
      case [code[ip], n]
      in [']', _]
        n += 1
      in ['[', 0]
        return ip
      in ['[', _]
        n -= 1
      else
        # pass
      end
    end
    raise 'Missing corresponding prev_matching_open'
  end
end

Brainf__kNaive.run('++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.')
# Hello World!
# Tested with in ruby 3.0.0dev (2020-10-28T00:47:46Z master 8f9c113f35) [x86_64-linux]
