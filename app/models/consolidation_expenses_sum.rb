class ConsolidationExpensesSum < Struct.new(:name, :sum, :slug, :year, :month)
  @@sum = 0

  def name
    I18n.translate 'consolidation.sum'
  end

  def sum
    @@sum
  end

  class << self
    def sum= sum
      @@sum = sum
    end
  end
end
