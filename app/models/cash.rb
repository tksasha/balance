# frozen_string_literal: true

class Cash < ActiveRecord::Base
  include ActsAsHasFormula

  acts_as_paranoid

  validates :name, :formula, presence: true

  validates :name, uniqueness: { case_sensitive: false }

  enum currency: CURRENCIES
end
