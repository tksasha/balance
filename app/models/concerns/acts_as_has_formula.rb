# frozen_string_literal: true

module ActsAsHasFormula
  extend ActiveSupport::Concern

  included do
    before_validation :calculate_formula
  end

  private

  def calculate_formula
    self.sum = Formula.calculate formula
  end
end
