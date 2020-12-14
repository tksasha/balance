# frozen_string_literal: true

class Cash < ApplicationRecord
  include ActsAsHasFormula

  include ActsAsParanoid

  validates :name, :formula, :currency, presence: true

  validates :name, uniqueness: { case_sensitive: false, scope: :currency }

  enum currency: CURRENCIES

  has_paper_trail
end
