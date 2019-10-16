# frozen_string_literal: true

class Category < ActiveRecord::Base
  validates :name, presence: true, uniqueness: { case_sensitive: false }

  scope :income, -> { where income: true }

  scope :expense, -> { where income: false }

  scope :visible, -> { where visible: true }

  before_save :assign_slug

  private

  def assign_slug
    self.slug = SlugService.build name
  end

  class << self
    def group_by_income
      [
        ['Видатки', visible.expense.pluck(:name, :id)],
        ['Надходження', visible.income.pluck(:name, :id)]
      ]
    end
  end
end
