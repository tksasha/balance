# frozen_string_literal: true

module ActsAsHasFormula
  extend ActiveSupport::Concern

  def formula=(formula)
    self.sum = FormulaService.calculate formula

    super
  end
end
