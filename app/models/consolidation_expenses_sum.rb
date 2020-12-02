# frozen_string_literal: true

ConsolidationExpensesSum = Struct.new(:name, :slug, :year, :month) do
  delegate :sum, to: :class

  def name
    I18n.t('consolidation.sum')
  end

  def income?
    false
  end

  class << self
    attr_writer :sum

    def sum
      @sum || 0
    end
  end
end
