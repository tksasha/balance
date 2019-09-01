# frozen_string_literal: true

class Cash < ActiveRecord::Base
  include ActsAsHasFormula

  acts_as_paranoid

  validates :name, :formula, presence: true

  validates :name, uniqueness: { case_sensitive: false }

  class << self
    def at_end
      Item.income.sum(:sum) - Item.expense.sum(:sum)
    end

    def balance
      (sum(:sum) - at_end).round(2)
    end
  end
end
