# frozen_string_literal: true

module ActsAsHasFormula
  extend ActiveSupport::Concern

  def formula=(formula)
    self.sum = Formula(formula)

    super
  end
end
