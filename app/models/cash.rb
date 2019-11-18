# frozen_string_literal: true

class Cash < ActiveRecord::Base
  include ActsAsHasFormula

  validates :name, :formula, presence: true

  validates :name, uniqueness: { case_sensitive: false, scope: :currency }

  enum currency: CURRENCIES

  acts_as_paranoid
end
