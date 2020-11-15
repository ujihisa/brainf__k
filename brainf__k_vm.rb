module Brainf__kVm
  Node = Struct.new(:car, :cdr) do
    def initialize(*)
      super
      freeze
    end
  end

  def self.parse(code)
    case code[0]
    when '+', '-', '.', '>', '<'
      (ast, rest) = parse(code[1..])
      [Node.new(code[0], ast), rest]
    when '['
      (ast1, rest1) = parse(code[1..])
      (ast2, rest2) = parse(rest1)
      [Node.new(ast1, ast2), rest2]
    when ']', nil
      [nil, code[1..]]
    end
  end

  def self.execute(ast, tape = {}, di = 0)
    case ast
    in ['+', cdr]
      tape[di] ||= 0
      tape[di] += 1
      execute(cdr, tape, di)
    in ['-', cdr]
      tape[di] ||= 0
      tape[di] -= 1
      execute(cdr, tape, di)
    in ['>', cdr]
      di += 1
      execute(cdr, tape, di)
    in ['<', cdr]
      di -= 1
      execute(cdr, tape, di)
    in ['.', cdr]
      putc tape[di]
      execute(cdr, tape, di)
    in [car, cdr]
      while tape[di] != 0
        (tape, di) = execute(car, tape, di)
      end
      execute(cdr, tape, di)
    in nil
      [tape, di]
    else
      raise "Must not happen! ast: #{ast.inspect}"
    end
  end
end

case Brainf__kVm.parse('++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.')
in [ast, nil]
  Brainf__kVm.execute(ast)
end
# Tested with ruby 3.0.0dev (2020-10-28T00:47:46Z master 8f9c113f35) [x86_64-linux]
