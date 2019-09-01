# frozen_string_literal: true

class Item < ActiveRecord::Base
  include ActsAsHasFormula

  belongs_to :category

  validates :date, :category_id, :formula, presence: true

  scope :income, -> { joins(:category).merge(Category.income) }

  scope :expense, -> { joins(:category).merge(Category.expense) }

  acts_as_paranoid

  class << self
    #
    # TODO: rename it to `search_by` and refactor it
    # TODO: use ItemSearcher instead
    #
    def search(date_range, slug = nil)
      if slug
        where categories: { slug: slug }
      else
        self
      end.includes(:category).where(date: date_range).order('date DESC')
    end
  end
end
