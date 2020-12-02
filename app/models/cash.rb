# frozen_string_literal: true

class Cash < ApplicationRecord
  include ActsAsHasFormula

  validates :name, :formula, :currency, presence: true

  validates :name, uniqueness: { case_sensitive: false, scope: :currency }

  enum currency: CURRENCIES

  acts_as_paranoid

  has_paper_trail
end
