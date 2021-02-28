# frozen_string_literal: true

module ActsAsHasFormula
  extend ActiveSupport::Concern

  def formula=(formula)
    self.sum = CalculateFormulaService.call(formula)

    super
  end
end
