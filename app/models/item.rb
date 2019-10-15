# frozen_string_literal: true

class Item < ActiveRecord::Base
  include ActsAsHasFormula

  belongs_to :category

  validates :date, :category_id, :formula, presence: true

  scope :income, -> { joins(:category).merge(Category.income) }

  scope :expense, -> { joins(:category).merge(Category.expense) }

  acts_as_paranoid
end
