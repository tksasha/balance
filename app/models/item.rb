# frozen_string_literal: true

class Item < ApplicationRecord
  include ActsAsHasFormula

  include ActsAsParanoid

  belongs_to :category

  validates :date, :category_id, :formula, :currency, presence: true

  scope :income, -> { joins(:category).merge(Category.income) }

  scope :expense, -> { joins(:category).merge(Category.expense) }

  enum currency: CURRENCIES
end
