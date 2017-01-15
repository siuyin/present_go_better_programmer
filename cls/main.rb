class Person
  attr_reader :name
  def initialize(name)
    @name = name
  end
end

class Boss < Person
  attr_reader :title
  def initialize(name,title)
    super(name)
    @title = title
  end
end

p = Person.new('SiuYin')
puts p.name

b = Boss.new('IAmBoss','CEO')
puts b.name, b.title
